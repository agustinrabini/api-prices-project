// Package docs GENERATED BY SWAG; DO NOT EDIT
// This file was generated by swaggo/swag
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "termsOfService": "http://swagger.io/terms/",
        "contact": {
            "name": "Matias Nuñez",
            "url": "http://www.swagger.io/support",
            "email": "matiasne45@gmail.com"
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
        "/ping": {
            "get": {
                "description": "test if the router works correctly",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "ping"
                ],
                "summary": "Ping",
                "responses": {
                    "200": {
                        "description": ""
                    }
                }
            }
        },
        "/prices": {
            "post": {
                "description": "Create price in db",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Prices"
                ],
                "summary": "Create Price",
                "parameters": [
                    {
                        "description": "Add price",
                        "name": "price",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/domain.Price"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": ""
                    }
                }
            }
        },
        "/prices/item/{id}": {
            "get": {
                "description": "Get details of price by item ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Prices"
                ],
                "summary": "Get price detail",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Item ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/domain.Price"
                        }
                    }
                }
            }
        },
        "/prices/{id}": {
            "put": {
                "description": "Update price in db",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Prices"
                ],
                "summary": "Update Price",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Price ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Add price",
                        "name": "price",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/domain.Price"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": ""
                    }
                }
            },
            "delete": {
                "description": "Delete price in db",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Prices"
                ],
                "summary": "Delete Price",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Price ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": ""
                    }
                }
            }
        }
    },
    "definitions": {
        "domain.Price": {
            "type": "object",
            "required": [
                "amount",
                "currency",
                "item_id"
            ],
            "properties": {
                "amount": {
                    "type": "number",
                    "minimum": 0
                },
                "currency": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "item_id": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "Api Prices Project",
	Description:      "This is a Api Prices Project.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
