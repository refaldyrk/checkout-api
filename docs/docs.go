// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/auth/check-token": {
            "post": {
                "tags": [
                    "Auth"
                ],
                "summary": "check jwt token",
                "parameters": [
                    {
                        "description": "payload",
                        "name": "payload",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/rest.PostCheckTokenReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/rest.BaseJSONResp"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/rest.AuthCheckTokenResp"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/auth/login": {
            "post": {
                "tags": [
                    "Auth"
                ],
                "summary": "generate jwt token (MANDATORY)",
                "parameters": [
                    {
                        "description": "payload",
                        "name": "payload",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/rest.PostLoginReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/rest.BaseJSONResp"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/rest.AuthLoginResp"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/auth/register": {
            "post": {
                "tags": [
                    "Auth"
                ],
                "summary": "Register New User (MANDATORY)",
                "parameters": [
                    {
                        "description": "payload",
                        "name": "payload",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/rest.PostRegisterReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/rest.BaseJSONResp"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/rest.AuthLoginResp"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/cart/checkout": {
            "post": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "tags": [
                    "Cart"
                ],
                "summary": "checkout cart (MANDATORY)",
                "parameters": [
                    {
                        "description": "payload",
                        "name": "payload",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/rest.PostCheckoutCartReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/rest.BaseJSONResp"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/rest.PostCheckoutCartResp"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/cart/orders": {
            "get": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "tags": [
                    "Cart"
                ],
                "summary": "get product list from current cart (MANDATORY)",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/rest.BaseJSONResp"
                        }
                    }
                }
            }
        },
        "/cart/orders/{id}": {
            "delete": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "tags": [
                    "Cart"
                ],
                "summary": "remove item from cart (MANDATORY)",
                "parameters": [
                    {
                        "type": "string",
                        "description": "order id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/rest.BaseJSONResp"
                        }
                    }
                }
            }
        },
        "/products": {
            "get": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "tags": [
                    "Product"
                ],
                "summary": "get product list (MANDATORY)",
                "parameters": [
                    {
                        "type": "string",
                        "name": "category",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "default": 10,
                        "name": "limit",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "default": 1,
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "using regex",
                        "name": "search",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "default": "created_at",
                        "name": "sort_by",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "default": "desc",
                        "name": "sort_order",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/rest.BaseJSONResp"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "type": "array",
                                            "items": {
                                                "$ref": "#/definitions/rest.GetProductListResp"
                                            }
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/products/category-list": {
            "get": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "tags": [
                    "Product"
                ],
                "summary": "get product category list",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/rest.BaseJSONResp"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "type": "array",
                                            "items": {
                                                "type": "string"
                                            }
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/products/{id}/add-to-cart": {
            "post": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "tags": [
                    "Product"
                ],
                "summary": "add item to cart (MANDATORY)",
                "parameters": [
                    {
                        "type": "string",
                        "description": "id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "payload",
                        "name": "payload",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/rest.PostAddItemToCartReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/rest.BaseJSONResp"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "rest.AuthCheckTokenResp": {
            "type": "object",
            "properties": {
                "username": {
                    "type": "string"
                }
            }
        },
        "rest.AuthLoginResp": {
            "type": "object",
            "properties": {
                "access_token": {
                    "type": "string"
                }
            }
        },
        "rest.BaseJSONResp": {
            "type": "object",
            "properties": {
                "data": {},
                "detail": {
                    "type": "string"
                },
                "error": {
                    "type": "boolean"
                },
                "message": {
                    "type": "string"
                }
            }
        },
        "rest.GetProductListResp": {
            "type": "object",
            "properties": {
                "category": {
                    "type": "string"
                },
                "created_at": {
                    "type": "string"
                },
                "deleted_at": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "price": {
                    "type": "number"
                },
                "updated_at": {
                    "type": "string"
                }
            }
        },
        "rest.PostAddItemToCartReq": {
            "type": "object",
            "properties": {
                "quantity": {
                    "type": "integer",
                    "default": 1
                }
            }
        },
        "rest.PostCheckTokenReq": {
            "type": "object",
            "properties": {
                "token": {
                    "type": "string"
                }
            }
        },
        "rest.PostCheckoutCartReq": {
            "type": "object",
            "required": [
                "money_input"
            ],
            "properties": {
                "money_input": {
                    "type": "number"
                }
            }
        },
        "rest.PostCheckoutCartResp": {
            "type": "object",
            "properties": {
                "money_return": {
                    "type": "number"
                }
            }
        },
        "rest.PostLoginReq": {
            "type": "object",
            "properties": {
                "password": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "rest.PostRegisterReq": {
            "type": "object",
            "properties": {
                "confirm_password": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        }
    },
    "securityDefinitions": {
        "BearerAuth": {
            "description": "JWT Authorization header using the Bearer scheme (add 'Bearer ' prefix).",
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "Backend",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
