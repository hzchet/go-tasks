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
        "contact": {
            "name": "API Support",
            "url": "http://www.swagger.io/support",
            "email": "support@swagger.io"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/": {
            "get": {
                "description": "get tasks created by the user",
                "tags": [
                    "handler"
                ],
                "summary": "get all the tasks",
                "operationId": "GetTasks",
                "responses": {
                    "200": {
                        "description": "OK"
                    },
                    "400": {
                        "description": "bad request",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "403": {
                        "description": "forbidden",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "internal error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "post": {
                "description": "create task by the given taskRequest and return taskId",
                "tags": [
                    "handler"
                ],
                "summary": "create task",
                "operationId": "Create",
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "bad request",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "403": {
                        "description": "forbidden",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "internal error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/:taskId": {
            "get": {
                "description": "get task description by the taskId",
                "tags": [
                    "handler"
                ],
                "summary": "get task description by the taskId",
                "operationId": "GetDescription",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "bad request",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "403": {
                        "description": "forbidden",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "not found",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "internal error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "delete": {
                "description": "Delete task by taskId in case it has been assigned to the user",
                "tags": [
                    "handler"
                ],
                "summary": "Delete task by taskId",
                "operationId": "Delete",
                "responses": {
                    "200": {
                        "description": "OK"
                    },
                    "400": {
                        "description": "bad request",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "403": {
                        "description": "forbidden",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "not found",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "internal error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/approve/:taskId": {
            "post": {
                "description": "Approve task by taskId in case it has been assigned to the user",
                "tags": [
                    "handler"
                ],
                "summary": "Approve task by taskId",
                "operationId": "Approve",
                "responses": {
                    "200": {
                        "description": "OK"
                    },
                    "400": {
                        "description": "bad request",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "403": {
                        "description": "forbidden",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "internal error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/decline/:taskId": {
            "post": {
                "description": "Decline task by taskId in case it has been assigned to the user",
                "tags": [
                    "handler"
                ],
                "summary": "Decline task by taskId",
                "operationId": "Decline",
                "responses": {
                    "200": {
                        "description": "OK"
                    },
                    "400": {
                        "description": "bad request",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "403": {
                        "description": "forbidden",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "internal error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    },
    "securityDefinitions": {
        "Auth": {
            "type": "basic"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0.0",
	Host:             "localhost:8080",
	BasePath:         "/tasks/api/v1",
	Schemes:          []string{"http"},
	Title:            "Example for lecture",
	Description:      "Signed token protects our admin endpoints",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}