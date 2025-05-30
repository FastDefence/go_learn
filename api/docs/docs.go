// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
	"schemes": {{ marshal .Schemes }},
	"swagger": "2.0",
	"info": {
		"description": "{{ escape .Description }}",
		"title": "{{ .Title }}",
		"contact": {},
		"version": "{{ .Version }}"
	},
	"host": "{{ .Host }}",
	"basePath": "{{ .BasePath }}",
	"paths": {
		"/users": {
			"post": {
				"description": "write on username",
				"parameters": [
					{
						"name": "name",
						"type": "string",
						"in": "query",
						"description": "ユーザー名",
						"required": true,
					}
				],
				"responses": {
					"200": {
					"description": "Created user",
					}
				}
			},
			"get": {
				"description": "show users",
				"responses": {
					"200": {
					"description": "Created user",
					"schema": {
						"type": "array",
						}
					}
				}
			},
		},
		"/users/{id}": {
			"get": {
				"description": "show users",
				"parameters": [
					{
						"name": "id",
						"type": "integer",
						"in": "path",
						"description": "ユーザーID",
						"required": true,
					}
				],
				"responses": {
					"200": {
					"description": "Created user",
					"schema": {
						"type": "array"
						}
					}
				}
			},
			"put": {
				"description": "change on username",
				"parameters": [
					{
						"name": "id",
						"type": "integer",
						"in": "path",
						"description": "ユーザーID",
						"required": true,
					},
					{
						"name": "name",
						"type": "string",
						"in": "query",
						"description": "ユーザー名",
					}
				],
				"responses": {
					"200": {
					"description": "Updated user",
					}
				}
			},
			"delete": {
				"description": "change on username",
				"parameters": [
					{
						"name": "id",
						"type": "integer",
						"in": "path",
						"description": "ユーザーID",
						"required": true,
					}
				],
				"responses": {
					"200": {
					"description": "Updated user",
					}
				}
			},
		}
	}
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:1323",
	BasePath:         "/",
	Schemes:          []string{"http"},
	Title:            "go_learn API",
	Description:      "go_learn API",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
