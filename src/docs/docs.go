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
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/api/login": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
                ],
                "summary": "Get access token",
                "operationId": "Login",
                "parameters": [
                    {
                        "description": "User",
                        "name": "Body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/query.UserAuthorizeQuery"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/refresh_token": {
            "get": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
                ],
                "summary": "Refresh access token",
                "operationId": "RefreshToken",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/storage": {
            "get": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Storage"
                ],
                "summary": "Get directory content",
                "operationId": "GetDirectoryContent",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Path",
                        "name": "path",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/entity.DirectoryDataWrapper"
                            }
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "403": {
                        "description": "Forbidden",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/system/cpu-usage": {
            "get": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "System"
                ],
                "summary": "Get CPU usage",
                "operationId": "GetCpuUsage",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/entity.CpuUsage"
                        }
                    }
                }
            }
        },
        "/api/system/disk-usage": {
            "get": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "System"
                ],
                "summary": "Get disk usage",
                "operationId": "GetDiskUsage",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/entity.DiskUsage"
                        }
                    }
                }
            }
        },
        "/api/system/load-avg": {
            "get": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "System"
                ],
                "summary": "Get Get load avg",
                "operationId": "GetLoadAvg",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/entity.LoadAvg"
                        }
                    }
                }
            }
        },
        "/api/system/memory-usage": {
            "get": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "System"
                ],
                "summary": "Get memory usage",
                "operationId": "GetMemoryUsage",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/entity.MemoryUsage"
                        }
                    }
                }
            }
        },
        "/api/system/up-time": {
            "get": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "System"
                ],
                "summary": "Get Get up time",
                "operationId": "GetUpTime",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/entity.UpTime"
                        }
                    }
                }
            }
        },
        "/api/users": {
            "get": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Users"
                ],
                "summary": "Get list of all users",
                "operationId": "GetUsers",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/entity.User"
                            }
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "put": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Users"
                ],
                "summary": "Update user",
                "operationId": "UpdateUser",
                "parameters": [
                    {
                        "description": "User",
                        "name": "Body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/command.UpdateUserCommand"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "integer"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "403": {
                        "description": "Forbidden",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "post": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Users"
                ],
                "summary": "Crete new user",
                "operationId": "CreteUser",
                "parameters": [
                    {
                        "description": "User",
                        "name": "Body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/command.CreateUserCommand"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "integer"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "403": {
                        "description": "Forbidden",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/users/change-password": {
            "put": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Users"
                ],
                "summary": "Change password",
                "operationId": "ChangePassword",
                "parameters": [
                    {
                        "description": "Password",
                        "name": "Body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/command.ChangePasswordCommand"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "integer"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "403": {
                        "description": "Forbidden",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/users/current": {
            "get": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Users"
                ],
                "summary": "Get current",
                "operationId": "CurrentUser",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/entity.User"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/users/{id}": {
            "get": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Users"
                ],
                "summary": "Get user by id",
                "operationId": "GetUser",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "User id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/entity.User"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "delete": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Users"
                ],
                "summary": "Delete user",
                "operationId": "DeleteUser",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "User id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "integer"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "403": {
                        "description": "Forbidden",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "command.ChangePasswordCommand": {
            "type": "object",
            "required": [
                "name",
                "newPassword",
                "previousPassword"
            ],
            "properties": {
                "name": {
                    "type": "string",
                    "example": "adam"
                },
                "newPassword": {
                    "type": "string",
                    "example": "123"
                },
                "previousPassword": {
                    "type": "string",
                    "example": "123"
                }
            }
        },
        "command.CreateUserCommand": {
            "type": "object",
            "required": [
                "name",
                "password"
            ],
            "properties": {
                "isActive": {
                    "type": "boolean",
                    "example": true
                },
                "isAdmin": {
                    "type": "boolean",
                    "example": false
                },
                "name": {
                    "type": "string",
                    "example": "adam"
                },
                "password": {
                    "type": "string",
                    "example": "123"
                }
            }
        },
        "command.UpdateUserCommand": {
            "type": "object",
            "required": [
                "id"
            ],
            "properties": {
                "id": {
                    "type": "integer",
                    "example": 1
                },
                "isActive": {
                    "type": "boolean",
                    "example": false
                },
                "isAdmin": {
                    "type": "boolean",
                    "example": false
                }
            }
        },
        "entity.CpuUsage": {
            "type": "object",
            "properties": {
                "count": {
                    "type": "integer",
                    "example": 1637768672
                },
                "idle": {
                    "type": "number",
                    "example": 1637768672
                },
                "system": {
                    "type": "number",
                    "example": 1637768672
                },
                "total": {
                    "type": "number",
                    "example": 1637768672
                },
                "user": {
                    "type": "number",
                    "example": 1637768672
                }
            }
        },
        "entity.DirectoryData": {
            "type": "object",
            "properties": {
                "name": {
                    "type": "string",
                    "example": "test"
                },
                "size": {
                    "type": "integer",
                    "example": 10
                },
                "type": {
                    "type": "string",
                    "example": "File"
                }
            }
        },
        "entity.DirectoryDataWrapper": {
            "type": "object",
            "properties": {
                "data": {
                    "$ref": "#/definitions/entity.DirectoryData"
                },
                "key": {
                    "type": "string",
                    "example": "test/test2"
                },
                "leaf": {
                    "type": "boolean"
                }
            }
        },
        "entity.DiskUsage": {
            "type": "object",
            "properties": {
                "available": {
                    "type": "integer",
                    "example": 1637768672
                },
                "size": {
                    "type": "integer",
                    "example": 1637768672
                },
                "usage": {
                    "type": "number",
                    "example": 1637768672
                },
                "used": {
                    "type": "integer",
                    "example": 1637768672
                }
            }
        },
        "entity.LoadAvg": {
            "type": "object",
            "properties": {
                "loadavg1": {
                    "type": "number",
                    "example": 1637768672
                },
                "loadavg15": {
                    "type": "number",
                    "example": 1637768672
                },
                "loadavg5": {
                    "type": "number",
                    "example": 1637768672
                }
            }
        },
        "entity.MemoryUsage": {
            "type": "object",
            "properties": {
                "available": {
                    "type": "integer",
                    "example": 1637768672
                },
                "cached": {
                    "type": "integer",
                    "example": 1637768672
                },
                "free": {
                    "type": "integer",
                    "example": 1637768672
                },
                "total": {
                    "type": "integer",
                    "example": 1637768672
                },
                "used": {
                    "type": "integer",
                    "example": 1637768672
                }
            }
        },
        "entity.UpTime": {
            "type": "object",
            "properties": {
                "upTime": {
                    "type": "integer",
                    "example": 1637768672
                }
            }
        },
        "entity.User": {
            "type": "object",
            "properties": {
                "created": {
                    "type": "integer",
                    "example": 1637768672
                },
                "id": {
                    "type": "integer",
                    "example": 1
                },
                "isActive": {
                    "type": "boolean",
                    "example": true
                },
                "isAdmin": {
                    "type": "boolean",
                    "example": true
                },
                "lastLogin": {
                    "type": "integer",
                    "example": 1637768672
                },
                "name": {
                    "type": "string",
                    "example": "Adam"
                }
            }
        },
        "query.UserAuthorizeQuery": {
            "type": "object",
            "required": [
                "password",
                "username"
            ],
            "properties": {
                "password": {
                    "type": "string",
                    "example": "123"
                },
                "username": {
                    "type": "string",
                    "example": "admin"
                }
            }
        }
    },
    "securityDefinitions": {
        "Bearer": {
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
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
