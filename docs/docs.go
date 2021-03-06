// Package docs GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag
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
        "/cocktail": {
            "put": {
                "security": [
                    {
                        "AuthToken": []
                    }
                ],
                "description": "Add IBA Cocktail",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Add Cocktail",
                "parameters": [
                    {
                        "type": "string",
                        "name": "Garnish",
                        "in": "formData"
                    },
                    {
                        "type": "integer",
                        "name": "Id",
                        "in": "formData"
                    },
                    {
                        "type": "string",
                        "name": "Ingredients",
                        "in": "formData"
                    },
                    {
                        "type": "string",
                        "name": "Method",
                        "in": "formData"
                    },
                    {
                        "type": "string",
                        "name": "Name",
                        "in": "formData"
                    }
                ],
                "responses": {}
            }
        },
        "/cocktail/{id}": {
            "get": {
                "security": [
                    {
                        "AuthToken": []
                    }
                ],
                "description": "get one IBA Cocktail",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "get Cocktail by Id",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Id cocktail",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {}
            },
            "delete": {
                "security": [
                    {
                        "AuthToken": []
                    }
                ],
                "description": "Delete One IBA Cocktail",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Delete Cocktail by Id",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Id cocktail",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {}
            }
        },
        "/cocktails": {
            "get": {
                "security": [
                    {
                        "AuthToken": []
                    }
                ],
                "description": "List IBA Cocktails",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "List Cocktails",
                "responses": {}
            }
        },
        "/login": {
            "post": {
                "description": "Request Login",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Login",
                "parameters": [
                    {
                        "description": "User auth data",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.User"
                        }
                    }
                ],
                "responses": {}
            }
        }
    },
    "definitions": {
        "models.User": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "returnSecureToken": {
                    "type": "string",
                    "default": "true"
                }
            }
        }
    },
    "securityDefinitions": {
        "AuthToken": {
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
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
