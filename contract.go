// --- Enhancement Stubs (GSOC’25) ---
// TODO: Add SchemaMeta struct (version, timestamp, author)
// TODO: Support saving/loading schemas in JSON as well as YAML
// TODO: Implement mergeSchemas for unified contract
// TODO: Implement diffSchemas for schema change tracking
// --- End Enhancement Stubs ---

package main

import (
	"fmt"
)

type HTTPDoc struct {
	Name string
	Spec Spec
}

type Spec struct {
	Request  Request
	Response Response
}

type Request struct {
	Method  string
	URL     string
	Headers map[string]string
	Body    string
}

type Response struct {
	StatusCode int
	Headers    map[string]string
	Body       string
}

type OpenAPI struct {
	Paths map[string]PathItem
}

type PathItem struct {
	Operations map[string]*Operation
}

type Operation struct {
	Responses map[string]ResponseDetail
}

type ResponseDetail struct {
	Description string
	Headers     map[string]string
	Body        string
}

func HTTPDocToOpenAPI(doc HTTPDoc) OpenAPI {
	method := doc.Spec.Request.Method
	if method == "" {
		method = "get"
	}
	return OpenAPI{
		Paths: map[string]PathItem{
			doc.Spec.Request.URL: {
				Operations: map[string]*Operation{
					method: {
						Responses: map[string]ResponseDetail{
							fmt.Sprintf("%d", doc.Spec.Response.StatusCode): {
								Description: "Auto-generated response",
								Headers:     doc.Spec.Response.Headers,
								Body:        doc.Spec.Response.Body,
							},
						},
					},
				},
			},
		},
	}
}

func loadSampleTests() map[string]HTTPDoc {
	return map[string]HTTPDoc{
		"test-get-products": {
			Name: "Get Products",
			Spec: Spec{
				Request: Request{
					Method:  "get",
					URL:     "/api/products",
					Headers: map[string]string{"Content-Type": "application/json"},
					Body:    "",
				},
				Response: Response{
					StatusCode: 200,
					Headers:    map[string]string{"Content-Type": "application/json"},
					Body:       `{"products": [{"id": 1, "name": "Laptop", "price": 999.99}]}`,
				},
			},
		},
		"test-post-order": {
			Name: "Create Order",
			Spec: Spec{
				Request: Request{
					Method:  "post",
					URL:     "/api/orders",
					Headers: map[string]string{"Content-Type": "application/json"},
					Body:    `{"product_id": 1, "quantity": 2}`,
				},
				Response: Response{
					StatusCode: 201,
					Headers:    map[string]string{"Content-Type": "application/json"},
					Body:       `{"order_id": 100, "product_id": 1, "quantity": 2, "total": 1999.98}`,
				},
			},
		},
		"test-get-cart": {
			Name: "Get Cart",
			Spec: Spec{
				Request: Request{
					Method:  "get",
					URL:     "/api/cart",
					Headers: map[string]string{"Content-Type": "application/json"},
					Body:    "",
				},
				Response: Response{
					StatusCode: 200,
					Headers:    map[string]string{"Content-Type": "application/json"},
					Body:       `{"cart": [{"product_id": 1, "quantity": 1}]}`,
				},
			},
		},
	}
}

func loadSampleMocks() map[string]HTTPDoc {
	return map[string]HTTPDoc{
		"mock-get-products": {
			Name: "Mock Get Products",
			Spec: Spec{
				Request: Request{
					Method:  "get",
					URL:     "/api/products",
					Headers: map[string]string{"Content-Type": "application/json"},
					Body:    "",
				},
				Response: Response{
					StatusCode: 200,
					Headers:    map[string]string{"Content-Type": "application/json"},
					Body:       `{"products": [{"id": 1, "name": "Laptop", "price": 999.99}]}`,
				},
			},
		},
		"mock-post-order-mismatch": {
			Name: "Mock Create Order (Mismatch)",
			Spec: Spec{
				Request: Request{
					Method:  "post",
					URL:     "/api/orders",
					Headers: map[string]string{"Content-Type": "application/json"},
					Body:    `{"product_id": 1, "quantity": 2}`,
				},
				Response: Response{
					StatusCode: 200, // Should be 201
					Headers:    map[string]string{"Content-Type": "application/json"},
					Body:       `{"order_id": 100, "total": 1999.98}`, // Missing fields
				},
			},
		},
		"mock-get-cart-extra": {
			Name: "Mock Get Cart (Extra Data)",
			Spec: Spec{
				Request: Request{
					Method:  "get",
					URL:     "/api/cart",
					Headers: map[string]string{"Content-Type": "application/json"},
					Body:    "",
				},
				Response: Response{
					StatusCode: 200,
					Headers:    map[string]string{"Content-Type": "application/json"},
					Body:       `{"cart": [{"product_id": 1, "quantity": 1, "price": 999.99}]}`, // Extra "price"
				},
			},
		},
	}
}