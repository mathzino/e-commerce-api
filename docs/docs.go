// Package docs GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag
package docs

import (
	"bytes"
	"encoding/json"
	"strings"
	"text/template"

	"github.com/swaggo/swag"
)

var doc = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "termsOfService": "http://swagger.io/terms/",
        "contact": {
            "name": "API Support",
            "url": "http://www.swagger.io/support",
            "email": "support@swagger.io"
        },
        "license": {
            "name": "Apache 2.0",
            "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/login": {
            "post": {
                "description": "Logging in to get jwt token to access admin or user api by roles.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
                ],
                "summary": "Login as as user.",
                "parameters": [
                    {
                        "description": "the body to login a user",
                        "name": "Body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/controllers.LoginInput"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    }
                }
            }
        },
        "/product": {
            "get": {
                "description": "Get a list of Product.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Product"
                ],
                "summary": "Get all Product.",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.Product"
                            }
                        }
                    }
                }
            },
            "post": {
                "security": [
                    {
                        "BearerToken": []
                    }
                ],
                "description": "create new Product",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Product"
                ],
                "summary": "Create Product",
                "parameters": [
                    {
                        "description": "the body to create new product",
                        "name": "Body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/controllers.productInput"
                        }
                    },
                    {
                        "type": "string",
                        "description": "Authorization. How to input in swagger : 'Bearer \u003cinsert_your_token_here\u003e'",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Product"
                        }
                    }
                }
            }
        },
        "/product-categories": {
            "get": {
                "description": "Get a list of ProductCategory.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "ProductCategory"
                ],
                "summary": "Get all Product Category.",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.ProductCategory"
                            }
                        }
                    }
                }
            },
            "post": {
                "security": [
                    {
                        "BearerToken": []
                    }
                ],
                "description": "create new Product Category",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "ProductCategory"
                ],
                "summary": "Create Product Category",
                "parameters": [
                    {
                        "description": "the body to create new category",
                        "name": "Body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/controllers.productCategoryInput"
                        }
                    },
                    {
                        "type": "string",
                        "description": "Authorization. How to input in swagger : 'Bearer \u003cinsert_your_token_here\u003e'",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.ProductCategory"
                        }
                    }
                }
            }
        },
        "/product-categories/{id}": {
            "get": {
                "description": "Get an ProductCategory by id.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "ProductCategory"
                ],
                "summary": "Get ProductCategory.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "ProductCategory id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.ProductCategory"
                        }
                    }
                }
            },
            "delete": {
                "security": [
                    {
                        "BearerToken": []
                    }
                ],
                "description": "Delete a ProductCategory by id.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "ProductCategory"
                ],
                "summary": "Delete one ProductCategory.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Authorization. How to input in swagger : 'Bearer \u003cinsert_your_token_here\u003e'",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "ProductCategory id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "boolean"
                            }
                        }
                    }
                }
            },
            "patch": {
                "security": [
                    {
                        "BearerToken": []
                    }
                ],
                "description": "Update ProductCategory by id.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "ProductCategory"
                ],
                "summary": "Update ProductCategory.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Authorization. How to input in swagger : 'Bearer \u003cinsert_your_token_here\u003e'",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "ProductCategory id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "the body to update product category",
                        "name": "Body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/controllers.productCategoryInput"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.ProductCategory"
                        }
                    }
                }
            }
        },
        "/product/{id}": {
            "get": {
                "description": "Get an Productby id.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Product"
                ],
                "summary": "Get Product.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Product id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Product"
                        }
                    }
                }
            },
            "delete": {
                "security": [
                    {
                        "BearerToken": []
                    }
                ],
                "description": "Delete a Product by id.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Product"
                ],
                "summary": "Delete one Product.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Authorization. How to input in swagger : 'Bearer \u003cinsert_your_token_here\u003e'",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "ProductCategory id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "boolean"
                            }
                        }
                    }
                }
            },
            "patch": {
                "security": [
                    {
                        "BearerToken": []
                    }
                ],
                "description": "Update Product by id.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Product"
                ],
                "summary": "Update Product.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Authorization. How to input in swagger : 'Bearer \u003cinsert_your_token_here\u003e'",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Product id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "the body to update product",
                        "name": "Body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/controllers.productInput"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Product"
                        }
                    }
                }
            }
        },
        "/register": {
            "post": {
                "description": "registering a user from public access.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
                ],
                "summary": "Register a user.",
                "parameters": [
                    {
                        "description": "the body to register a user",
                        "name": "Body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/controllers.RegisterInput"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    }
                }
            }
        },
        "/seller": {
            "get": {
                "description": "Get a list of Seller.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Product"
                ],
                "summary": "Get all Seller.",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.Seller"
                            }
                        }
                    }
                }
            },
            "post": {
                "security": [
                    {
                        "BearerToken": []
                    }
                ],
                "description": "create new Seller",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Seller"
                ],
                "summary": "Create Seller",
                "parameters": [
                    {
                        "description": "the body to create new seller",
                        "name": "Body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/controllers.sellerInput"
                        }
                    },
                    {
                        "type": "string",
                        "description": "Authorization. How to input in swagger : 'Bearer \u003cinsert_your_token_here\u003e'",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Seller"
                        }
                    }
                }
            }
        },
        "/seller/{id}": {
            "get": {
                "description": "Get an Seller by id.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Product"
                ],
                "summary": "Get Seller.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Seller id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Seller"
                        }
                    }
                }
            },
            "delete": {
                "security": [
                    {
                        "BearerToken": []
                    }
                ],
                "description": "Delete a Seller by id.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Seller"
                ],
                "summary": "Delete one Seller.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Authorization. How to input in swagger : 'Bearer \u003cinsert_your_token_here\u003e'",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Seller id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "boolean"
                            }
                        }
                    }
                }
            },
            "patch": {
                "security": [
                    {
                        "BearerToken": []
                    }
                ],
                "description": "Update Seller by id.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Seller"
                ],
                "summary": "Update Seller.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Authorization. How to input in swagger : 'Bearer \u003cinsert_your_token_here\u003e'",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Seller id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "the body to update seller",
                        "name": "Body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/controllers.sellerInput"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Seller"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "controllers.LoginInput": {
            "type": "object",
            "required": [
                "password",
                "username"
            ],
            "properties": {
                "password": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "controllers.RegisterInput": {
            "type": "object",
            "required": [
                "email",
                "password",
                "username"
            ],
            "properties": {
                "email": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "controllers.productCategoryInput": {
            "type": "object",
            "properties": {
                "name": {
                    "type": "string"
                }
            }
        },
        "controllers.productInput": {
            "type": "object",
            "properties": {
                "description": {
                    "type": "string"
                },
                "image": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "price": {
                    "type": "string"
                },
                "product_category_id": {
                    "type": "integer"
                },
                "seller_id": {
                    "type": "integer"
                }
            }
        },
        "controllers.sellerInput": {
            "type": "object",
            "properties": {
                "store_name": {
                    "type": "string"
                },
                "user_id": {
                    "type": "integer"
                }
            }
        },
        "models.Product": {
            "type": "object",
            "properties": {
                "created_at": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "image": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "price": {
                    "type": "string"
                },
                "productCategoryID": {
                    "type": "integer"
                },
                "sellerID": {
                    "type": "integer"
                },
                "updated_at": {
                    "type": "string"
                }
            }
        },
        "models.ProductCategory": {
            "type": "object",
            "properties": {
                "created_at": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "updated_at": {
                    "type": "string"
                }
            }
        },
        "models.Seller": {
            "type": "object",
            "properties": {
                "created_at": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "store_name": {
                    "type": "string"
                },
                "updated_at": {
                    "type": "string"
                },
                "userID": {
                    "type": "integer"
                }
            }
        }
    }
}`

type swaggerInfo struct {
	Version     string
	Host        string
	BasePath    string
	Schemes     []string
	Title       string
	Description string
}

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = swaggerInfo{
	Version:     "",
	Host:        "",
	BasePath:    "",
	Schemes:     []string{},
	Title:       "",
	Description: "",
}

type s struct{}

func (s *s) ReadDoc() string {
	sInfo := SwaggerInfo
	sInfo.Description = strings.Replace(sInfo.Description, "\n", "\\n", -1)

	t, err := template.New("swagger_info").Funcs(template.FuncMap{
		"marshal": func(v interface{}) string {
			a, _ := json.Marshal(v)
			return string(a)
		},
		"escape": func(v interface{}) string {
			// escape tabs
			str := strings.Replace(v.(string), "\t", "\\t", -1)
			// replace " with \", and if that results in \\", replace that with \\\"
			str = strings.Replace(str, "\"", "\\\"", -1)
			return strings.Replace(str, "\\\\\"", "\\\\\\\"", -1)
		},
	}).Parse(doc)
	if err != nil {
		return doc
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, sInfo); err != nil {
		return doc
	}

	return tpl.String()
}

func init() {
	swag.Register("swagger", &s{})
}
