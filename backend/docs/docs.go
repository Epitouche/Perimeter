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
        "/action/info/{id}": {
            "get": {
                "security": [
                    {
                        "bearerAuth": []
                    }
                ],
                "description": "get action info of service id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Action"
                ],
                "summary": "get action info",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Service ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/schemas.Action"
                            }
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/schemas.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/schemas.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/area/": {
            "get": {
                "security": [
                    {
                        "Bearer": []
                    },
                    {
                        "bearerAuth": []
                    }
                ],
                "description": "get user areas list",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Area"
                ],
                "summary": "get user areas",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/schemas.Area"
                            }
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/schemas.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/schemas.ErrorResponse"
                        }
                    }
                }
            },
            "post": {
                "description": "create area",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Area"
                ],
                "summary": "create area",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/schemas.Response"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/schemas.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/github/auth": {
            "get": {
                "description": "give url to authenticate with github",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Github"
                ],
                "summary": "give url to authenticate with github",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/schemas.AuthenticationUrl"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/schemas.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/github/auth/callback": {
            "post": {
                "description": "give url to authenticate with github",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Github"
                ],
                "summary": "give url to authenticate with github",
                "parameters": [
                    {
                        "description": "Callback Payload",
                        "name": "payload",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/schemas.CodeCredentials"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/schemas.JWT"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/schemas.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/github/info/user": {
            "get": {
                "security": [
                    {
                        "Bearer": []
                    },
                    {
                        "bearerAuth": []
                    }
                ],
                "description": "give user info of github",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Github"
                ],
                "summary": "give user info of github",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/schemas.UserCredentials"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/schemas.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/schemas.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/gmail/auth": {
            "get": {
                "description": "give url to authenticate with gmail",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Gmail"
                ],
                "summary": "give url to authenticate with gmail",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/schemas.AuthenticationUrl"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/schemas.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/gmail/auth/callback": {
            "post": {
                "description": "give url to authenticate with gmail",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Gmail"
                ],
                "summary": "give url to authenticate with gmail",
                "parameters": [
                    {
                        "description": "Callback Payload",
                        "name": "payload",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/schemas.CodeCredentials"
                        }
                    },
                    {
                        "type": "string",
                        "description": "Bearer token",
                        "name": "Authorization",
                        "in": "header"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/schemas.JWT"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/schemas.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/gmail/info/user": {
            "get": {
                "security": [
                    {
                        "Bearer": []
                    },
                    {
                        "bearerAuth": []
                    }
                ],
                "description": "give user info of gmail",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Gmail"
                ],
                "summary": "give user info of gmail",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/schemas.UserCredentials"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/schemas.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/schemas.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/ping": {
            "get": {
                "description": "do ping to check if the server is running",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "ping route"
                ],
                "summary": "ping example",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/schemas.Response"
                        }
                    }
                }
            }
        },
        "/reaction/info/{id}": {
            "get": {
                "security": [
                    {
                        "Bearer": []
                    },
                    {
                        "bearerAuth": []
                    }
                ],
                "description": "get reaction info of service id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Reaction"
                ],
                "summary": "get reaction info",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Service ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/schemas.Reaction"
                            }
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/schemas.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/schemas.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/service/info/": {
            "get": {
                "security": [
                    {
                        "Bearer": []
                    },
                    {
                        "bearerAuth": []
                    }
                ],
                "description": "get service info of service id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Service"
                ],
                "summary": "get service info",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/schemas.Service"
                            }
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/schemas.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/schemas.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/spotify/auth": {
            "get": {
                "description": "give url to authenticate with spotify",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Spotify"
                ],
                "summary": "give url to authenticate with spotify",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/schemas.AuthenticationUrl"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/schemas.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/spotify/auth/callback": {
            "post": {
                "description": "give url to authenticate with spotify",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Spotify"
                ],
                "summary": "give url to authenticate with spotify",
                "parameters": [
                    {
                        "description": "Callback Payload",
                        "name": "payload",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/schemas.CodeCredentials"
                        }
                    },
                    {
                        "type": "string",
                        "description": "Bearer token",
                        "name": "Authorization",
                        "in": "header"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/schemas.JWT"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/schemas.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/spotify/info/user": {
            "get": {
                "security": [
                    {
                        "Bearer": []
                    },
                    {
                        "bearerAuth": []
                    }
                ],
                "description": "give user info of spotify",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Spotify"
                ],
                "summary": "give user info of spotify",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/schemas.UserCredentials"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/schemas.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/schemas.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/user/info/all": {
            "get": {
                "security": [
                    {
                        "Bearer": []
                    },
                    {
                        "bearerAuth": []
                    }
                ],
                "description": "give user info of user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "give user info of user",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/schemas.UserAllInfo"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/schemas.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/schemas.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/user/info/user": {
            "get": {
                "security": [
                    {
                        "Bearer": []
                    },
                    {
                        "bearerAuth": []
                    }
                ],
                "description": "give user info of user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "give user info of user",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/schemas.UserCredentials"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/schemas.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/schemas.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/user/login": {
            "post": {
                "description": "Authenticates a user and provides a JWT to Authorize API calls",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "Provides a JSON Web Token",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Username",
                        "name": "username",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Password",
                        "name": "password",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/schemas.JWT"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/schemas.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/user/register": {
            "post": {
                "description": "Authenticates a user and provides a JWT to Authorize API calls",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "Provides a JSON Web Token",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Email",
                        "name": "email",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Username",
                        "name": "username",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Password",
                        "name": "password",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/schemas.JWT"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/schemas.ErrorResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "schemas.Action": {
            "type": "object",
            "required": [
                "description",
                "name",
                "option",
                "service_id"
            ],
            "properties": {
                "createdAt": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "option": {
                    "type": "string"
                },
                "service_id": {
                    "$ref": "#/definitions/schemas.Service"
                },
                "update_at": {
                    "type": "string"
                }
            }
        },
        "schemas.Area": {
            "type": "object",
            "required": [
                "action_id",
                "action_option",
                "reaction_id",
                "reaction_option",
                "user_id"
            ],
            "properties": {
                "action_id": {
                    "$ref": "#/definitions/schemas.Action"
                },
                "action_option": {
                    "type": "string"
                },
                "createdAt": {
                    "type": "string"
                },
                "enable": {
                    "type": "boolean"
                },
                "id": {
                    "type": "integer"
                },
                "reaction_id": {
                    "$ref": "#/definitions/schemas.Reaction"
                },
                "reaction_option": {
                    "type": "string"
                },
                "update_at": {
                    "type": "string"
                },
                "user_id": {
                    "$ref": "#/definitions/schemas.User"
                }
            }
        },
        "schemas.AuthenticationUrl": {
            "type": "object",
            "properties": {
                "authentication_url": {
                    "type": "string"
                }
            }
        },
        "schemas.CodeCredentials": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "string"
                }
            }
        },
        "schemas.ErrorResponse": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "string"
                }
            }
        },
        "schemas.JWT": {
            "type": "object",
            "properties": {
                "token": {
                    "type": "string"
                }
            }
        },
        "schemas.Reaction": {
            "type": "object",
            "required": [
                "description",
                "name",
                "option",
                "service_id"
            ],
            "properties": {
                "createdAt": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "option": {
                    "type": "string"
                },
                "service_id": {
                    "$ref": "#/definitions/schemas.Service"
                },
                "update_at": {
                    "type": "string"
                }
            }
        },
        "schemas.Response": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                }
            }
        },
        "schemas.Service": {
            "type": "object",
            "required": [
                "description",
                "name"
            ],
            "properties": {
                "createdAt": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "name": {
                    "$ref": "#/definitions/schemas.ServiceName"
                },
                "update_at": {
                    "type": "string"
                }
            }
        },
        "schemas.ServiceName": {
            "type": "string",
            "enum": [
                "spotify",
                "openWeatherMap",
                "timer",
                "gmail"
            ],
            "x-enum-varnames": [
                "Spotify",
                "OpenWeatherMap",
                "Timer",
                "Gmail"
            ]
        },
        "schemas.Token": {
            "type": "object",
            "properties": {
                "createdAt": {
                    "type": "string"
                },
                "expireAt": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "refresh_token": {
                    "type": "string"
                },
                "service_id": {
                    "$ref": "#/definitions/schemas.Service"
                },
                "token": {
                    "type": "string"
                },
                "updateAt": {
                    "type": "string"
                },
                "user_id": {
                    "$ref": "#/definitions/schemas.User"
                }
            }
        },
        "schemas.User": {
            "type": "object",
            "required": [
                "username"
            ],
            "properties": {
                "createdAt": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "password": {
                    "description": "can be null for Oauth2.0 users",
                    "type": "string"
                },
                "token_id": {
                    "description": "Foreign key for LinkUrl",
                    "type": "integer"
                },
                "updated_at": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "schemas.UserAllInfo": {
            "type": "object",
            "properties": {
                "tokens": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/schemas.Token"
                    }
                },
                "user": {
                    "$ref": "#/definitions/schemas.User"
                }
            }
        },
        "schemas.UserCredentials": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        }
    },
    "securityDefinitions": {
        "bearerAuth": {
            "description": "Use \"Bearer \u003ctoken\u003e\" as the format for the Authorization header",
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
	Title:            "",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
