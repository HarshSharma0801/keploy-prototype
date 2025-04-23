// --- Enhancement Stubs (GSOC’25) ---
// TODO: Add dependency graph for advanced schema comparison
// TODO: Add watchSchemas for real-time validation (stub)
// TODO: Enhance compareBodies for partial/deep matching
// TODO: Add contract history tracking for rollback
// --- End Enhancement Stubs ---

package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/fatih/color"
	"github.com/olekukonko/tablewriter"
)

type ValidationResult struct {
	Score      float64
	Pass       bool
	Mismatches []string
}

func match(a, b OpenAPI) ValidationResult {
	result := ValidationResult{Score: 0.0, Pass: true}
	weights := map[string]float64{
		"method":  0.2,
		"status":  0.3,
		"headers": 0.3,
		"body":    0.2,
	}

	if len(a.Paths) != len(b.Paths) {
		result.Mismatches = append(result.Mismatches, "Different number of paths")
		return result
	}

	for path, itemA := range a.Paths {
		itemB, ok := b.Paths[path]
		if !ok {
			result.Mismatches = append(result.Mismatches, fmt.Sprintf("Path %s not found", path))
			return result
		}

		for method, opA := range itemA.Operations {
			opB, exists := itemB.Operations[method]
			if !exists {
				result.Mismatches = append(result.Mismatches, fmt.Sprintf("Method %s for path %s not found", method, path))
				return result
			}
			result.Score += weights["method"]

			for code, respA := range opA.Responses {
				respB, exists := opB.Responses[code]
				if !exists {
					result.Mismatches = append(result.Mismatches, fmt.Sprintf("Status code %s not found for %s %s", code, method, path))
					result.Pass = false
					continue
				}
				result.Score += weights["status"]

				if fmt.Sprint(respA.Headers) != fmt.Sprint(respB.Headers) {
					result.Mismatches = append(result.Mismatches, fmt.Sprintf("Headers mismatch for %s %s %s", method, path, code))
					result.Pass = false
				} else {
					result.Score += weights["headers"]
				}

				if respA.Body != respB.Body {
					result.Mismatches = append(result.Mismatches, fmt.Sprintf("Body mismatch for %s %s %s", method, path, code))
					result.Pass = false
				} else {
					result.Score += weights["body"]
				}
			}
		}
	}

	if !result.Pass {
		result.Score = result.Score / 1.0
	}
	return result
}

func ValidateConsumer(tests, mocks map[string]OpenAPI) {
	table := tablewriter.NewWriter(os.Stdout)
	// Updated headings for clarity
	table.SetHeader([]string{"Consumer Mock", "Provider Test", "Status", "Score", "Mismatches"})
	table.SetAutoMergeCells(true)
	table.SetAlignment(tablewriter.ALIGN_CENTER)
	successColor := color.New(color.FgGreen).SprintFunc()
	failureColor := color.New(color.FgRed).SprintFunc()

	for tName, test := range tests {
		for mName, mock := range mocks {
			result := match(test, mock)
			status := failureColor("Failed")
			if result.Pass {
				status = successColor("Passed")
			}
			mismatchStr := strings.Join(result.Mismatches, "\n")
			if mismatchStr == "" {
				mismatchStr = "-"
			}
			table.Append([]string{mName, tName, status, fmt.Sprintf("%.2f", result.Score), mismatchStr})
		}
	}

	fmt.Println("Validation Results:")
	table.Render()
}

func ValidateProvider(tests, mocks map[string]OpenAPI) {
	table := tablewriter.NewWriter(os.Stdout)
	// Headings remain clear and correct
	table.SetHeader([]string{"Provider Test", "Consumer Mock", "Status", "Score", "Mismatches"})
	table.SetAutoMergeCells(true)
	table.SetAlignment(tablewriter.ALIGN_CENTER)
	successColor := color.New(color.FgGreen).SprintFunc()
	failureColor := color.New(color.FgRed).SprintFunc()

	for tName, test := range tests {
		for mName, mock := range mocks {
			result := match(mock, test)
			status := failureColor("Failed")
			if result.Pass {
				status = successColor("Passed")
			}
			mismatchStr := strings.Join(result.Mismatches, "\n")
			if mismatchStr == "" {
				mismatchStr = "-"
			}
			table.Append([]string{tName, mName, status, fmt.Sprintf("%.2f", result.Score), mismatchStr})
		}
	}

	fmt.Println("Validation Results:")
	table.Render()
}
