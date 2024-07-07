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
        "/task": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Tasks"
                ],
                "summary": "Create a new task",
                "parameters": [
                    {
                        "description": "Task details",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.TaskRequestBody"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Task ID",
                        "schema": {
                            "$ref": "#/definitions/handler.SuccessResponse"
                        }
                    },
                    "400": {
                        "description": "Error message",
                        "schema": {
                            "$ref": "#/definitions/handler.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Error message",
                        "schema": {
                            "$ref": "#/definitions/handler.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/task/{task_id}/start": {
            "post": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Tasks"
                ],
                "summary": "Start a task",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Task ID",
                        "name": "task_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Status of the task",
                        "schema": {
                            "$ref": "#/definitions/handler.SuccessResponse"
                        }
                    },
                    "400": {
                        "description": "Error message",
                        "schema": {
                            "$ref": "#/definitions/handler.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Error message",
                        "schema": {
                            "$ref": "#/definitions/handler.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/task/{task_id}/stop": {
            "post": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Tasks"
                ],
                "summary": "Stop a task",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Task ID",
                        "name": "task_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Status of the task",
                        "schema": {
                            "$ref": "#/definitions/handler.SuccessResponse"
                        }
                    },
                    "400": {
                        "description": "Error message",
                        "schema": {
                            "$ref": "#/definitions/handler.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Error message",
                        "schema": {
                            "$ref": "#/definitions/handler.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/user": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Users"
                ],
                "summary": "Get all users",
                "parameters": [
                    {
                        "type": "string",
                        "name": "address",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "name": "id",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "name": "name",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "default": 1,
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "name": "passport_number",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "name": "passport_serie",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "name": "patronymic",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "default": 10,
                        "name": "per_page",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "name": "surname",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "List of users",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/model.User"
                            }
                        }
                    },
                    "400": {
                        "description": "Error message",
                        "schema": {
                            "$ref": "#/definitions/handler.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Error message",
                        "schema": {
                            "$ref": "#/definitions/handler.ErrorResponse"
                        }
                    }
                }
            },
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Users"
                ],
                "summary": "Create a user",
                "parameters": [
                    {
                        "description": "User details",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.UserRequestBody"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "User ID",
                        "schema": {
                            "$ref": "#/definitions/handler.SuccessResponse"
                        }
                    },
                    "400": {
                        "description": "Error message",
                        "schema": {
                            "$ref": "#/definitions/handler.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Error message",
                        "schema": {
                            "$ref": "#/definitions/handler.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/user/{user_id}": {
            "delete": {
                "tags": [
                    "Users"
                ],
                "summary": "Delete a user",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "User ID",
                        "name": "user_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Message",
                        "schema": {
                            "$ref": "#/definitions/handler.SuccessResponse"
                        }
                    },
                    "400": {
                        "description": "Error message",
                        "schema": {
                            "$ref": "#/definitions/handler.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Error message",
                        "schema": {
                            "$ref": "#/definitions/handler.ErrorResponse"
                        }
                    }
                }
            },
            "patch": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Users"
                ],
                "summary": "Update a user",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "User ID",
                        "name": "user_id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "User details",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.UserUpdateRequestBody"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Message",
                        "schema": {
                            "$ref": "#/definitions/handler.SuccessResponse"
                        }
                    },
                    "400": {
                        "description": "Error message",
                        "schema": {
                            "$ref": "#/definitions/handler.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Error message",
                        "schema": {
                            "$ref": "#/definitions/handler.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/user/{user_id}/time-spent": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Users"
                ],
                "summary": "Get user time spent",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "User ID",
                        "name": "user_id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "example": "\"2023-30-12 00:00:00\"",
                        "description": "Start period",
                        "name": "start_period",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "example": "\"2023-30-12 23:59:59\"",
                        "description": "End period",
                        "name": "end_period",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "List of time spent",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/model.TaskTimeSpent"
                            }
                        }
                    },
                    "400": {
                        "description": "Error message",
                        "schema": {
                            "$ref": "#/definitions/handler.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Error message",
                        "schema": {
                            "$ref": "#/definitions/handler.ErrorResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "handler.ErrorResponse": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "string"
                }
            }
        },
        "handler.SuccessResponse": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                }
            }
        },
        "model.TaskRequestBody": {
            "type": "object",
            "required": [
                "name",
                "user_id"
            ],
            "properties": {
                "description": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "user_id": {
                    "type": "integer"
                }
            }
        },
        "model.TaskTimeSpent": {
            "type": "object",
            "properties": {
                "taskID": {
                    "type": "integer"
                },
                "totalMinutes": {
                    "type": "number"
                }
            }
        },
        "model.User": {
            "type": "object",
            "properties": {
                "address": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "passportNumber": {
                    "type": "integer"
                },
                "passportSerie": {
                    "type": "integer"
                },
                "patronymic": {
                    "type": "string"
                },
                "surname": {
                    "type": "string"
                }
            }
        },
        "model.UserRequestBody": {
            "type": "object",
            "properties": {
                "passportNumber": {
                    "type": "string"
                }
            }
        },
        "model.UserUpdateRequestBody": {
            "type": "object",
            "properties": {
                "address": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "patronymic": {
                    "type": "string"
                },
                "surname": {
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
	BasePath:         "/api",
	Schemes:          []string{},
	Title:            "time-tracker API",
	Description:      "This is the API for the Time Tracker application.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}