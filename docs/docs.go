// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {
            "name": "API Support"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/tasks": {
            "get": {
                "consumes": [
                    "*/*"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "tasks"
                ],
                "summary": "Get all tasks",
                "responses": {
                    "200": {
                        "description": "Succeed gets tasks",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.Task"
                            }
                        }
                    },
                    "500": {
                        "description": "Internal error",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    }
                }
            },
            "post": {
                "consumes": [
                    "*/*"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "tasks"
                ],
                "summary": "Create a new task",
                "parameters": [
                    {
                        "description": "Task to create",
                        "name": "task",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.TaskDto"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Creates",
                        "schema": {
                            "$ref": "#/definitions/models.Task"
                        }
                    },
                    "400": {
                        "description": "Bad request",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal error",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/tasks/{id}": {
            "put": {
                "consumes": [
                    "*/*"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "tasks"
                ],
                "summary": "Update a task",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Task ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Updated task",
                        "name": "task",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.TaskDto"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Task"
                        }
                    },
                    "400": {
                        "description": "Invalid date",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal error",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    }
                }
            },
            "delete": {
                "consumes": [
                    "*/*"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "tasks"
                ],
                "summary": "Delete a task",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Task ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "204": {
                        "description": "Task was deleted"
                    },
                    "500": {
                        "description": "Internal error",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "models.ErrorResponse": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                }
            }
        },
        "models.Task": {
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
                "status": {
                    "type": "string"
                },
                "title": {
                    "type": "string"
                },
                "updated_at": {
                    "type": "string"
                }
            }
        },
        "models.TaskDto": {
            "type": "object",
            "properties": {
                "description": {
                    "type": "string"
                },
                "status": {
                    "type": "string"
                },
                "title": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8080",
	BasePath:         "/",
	Schemes:          []string{"http"},
	Title:            "Fiber Swagger Example API",
	Description:      "This is a sample server...",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
