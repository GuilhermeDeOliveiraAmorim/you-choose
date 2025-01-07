// Package api Code generated by swaggo/swag. DO NOT EDIT
package api

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "termsOfService": "http://swagger.io/terms/",
        "contact": {
            "name": "You Choose",
            "url": "http://www.youchoose.com.br",
            "email": "contato@youchoose.com.br"
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
        "/lists": {
            "get": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Get a list of movies and votes",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Lists"
                ],
                "summary": "Get List",
                "parameters": [
                    {
                        "type": "string",
                        "description": "List id",
                        "name": "list_id",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/usecases.GetListByUserIDOutputDTO"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/util.ProblemDetails"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/util.ProblemDetails"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/util.ProblemDetails"
                        }
                    }
                }
            },
            "post": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Registers a new list in the system",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Lists"
                ],
                "summary": "Create a new list",
                "parameters": [
                    {
                        "description": "List data",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/usecases.List"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/usecases.CreateListOutputDTO"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/util.ProblemDetails"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/util.ProblemDetails"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/util.ProblemDetails"
                        }
                    }
                }
            }
        },
        "/lists/movies": {
            "post": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Add new movies to list",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Lists"
                ],
                "summary": "Add movies to list",
                "parameters": [
                    {
                        "description": "AddMoviesList data",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/usecases.Movies"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/usecases.AddMoviesListOutputDTO"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/util.ProblemDetails"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/util.ProblemDetails"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/util.ProblemDetails"
                        }
                    }
                }
            }
        },
        "/login": {
            "post": {
                "description": "Authenticates a user and returns a JWT token",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Authentication"
                ],
                "summary": "Login a user",
                "parameters": [
                    {
                        "description": "User credentials",
                        "name": "LoginRequest",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/usecases.LoginInputDto"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/usecases.LoginOutputDto"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/util.ProblemDetails"
                        }
                    }
                }
            }
        },
        "/movies": {
            "post": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Registers a new movie in the system",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Movies"
                ],
                "summary": "Create a new movie",
                "parameters": [
                    {
                        "description": "Movie data",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/usecases.Movie"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/usecases.CreateMovieOutputDTO"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/util.ProblemDetails"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/util.ProblemDetails"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/util.ProblemDetails"
                        }
                    }
                }
            }
        },
        "/signup": {
            "post": {
                "description": "Registers a new user in the system",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Authentication"
                ],
                "summary": "Create a new user",
                "parameters": [
                    {
                        "description": "User data",
                        "name": "CreateUserRequest",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/usecases.CreateUserInputDto"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/usecases.CreateUserOutputDto"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/util.ProblemDetails"
                        }
                    }
                }
            }
        },
        "/votes": {
            "post": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Registers a new vote in the system",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Votes"
                ],
                "summary": "Create a new vote",
                "parameters": [
                    {
                        "description": "Vote data",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/usecases.Vote"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/usecases.VoteOutputDTO"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/util.ProblemDetails"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/util.ProblemDetails"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/util.ProblemDetails"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "entities.Combination": {
            "type": "object",
            "properties": {
                "first_movie": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "list_id": {
                    "type": "string"
                },
                "second_movie": {
                    "type": "string"
                }
            }
        },
        "entities.List": {
            "type": "object",
            "properties": {
                "active": {
                    "type": "boolean"
                },
                "combinations": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/entities.Combination"
                    }
                },
                "created_at": {
                    "type": "string"
                },
                "deactivated_at": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "movies": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/entities.Movie"
                    }
                },
                "name": {
                    "type": "string"
                },
                "updated_at": {
                    "type": "string"
                }
            }
        },
        "entities.Movie": {
            "type": "object",
            "properties": {
                "active": {
                    "type": "boolean"
                },
                "created_at": {
                    "type": "string"
                },
                "deactivated_at": {
                    "type": "string"
                },
                "external_id": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "poster": {
                    "type": "string"
                },
                "updated_at": {
                    "type": "string"
                },
                "year": {
                    "type": "integer"
                }
            }
        },
        "entities.Vote": {
            "type": "object",
            "properties": {
                "active": {
                    "type": "boolean"
                },
                "combination_id": {
                    "type": "string"
                },
                "created_at": {
                    "type": "string"
                },
                "deactivated_at": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "user_id": {
                    "type": "string"
                },
                "winner_id": {
                    "type": "string"
                }
            }
        },
        "usecases.AddMoviesListOutputDTO": {
            "type": "object",
            "properties": {
                "content_message": {
                    "type": "string"
                },
                "success_message": {
                    "type": "string"
                }
            }
        },
        "usecases.CreateListOutputDTO": {
            "type": "object",
            "properties": {
                "content_message": {
                    "type": "string"
                },
                "success_message": {
                    "type": "string"
                }
            }
        },
        "usecases.CreateMovieOutputDTO": {
            "type": "object",
            "properties": {
                "content_message": {
                    "type": "string"
                },
                "success_message": {
                    "type": "string"
                }
            }
        },
        "usecases.CreateUserInputDto": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                }
            }
        },
        "usecases.CreateUserOutputDto": {
            "type": "object",
            "properties": {
                "content_message": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "success_message": {
                    "type": "string"
                }
            }
        },
        "usecases.GetListByUserIDOutputDTO": {
            "type": "object",
            "properties": {
                "list": {
                    "$ref": "#/definitions/entities.List"
                },
                "votes": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/entities.Vote"
                    }
                }
            }
        },
        "usecases.List": {
            "type": "object",
            "properties": {
                "movies": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "usecases.LoginInputDto": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                }
            }
        },
        "usecases.LoginOutputDto": {
            "type": "object",
            "properties": {
                "access_token": {
                    "type": "string"
                },
                "content_message": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "success_message": {
                    "type": "string"
                }
            }
        },
        "usecases.Movie": {
            "type": "object",
            "properties": {
                "external_id": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "poster": {
                    "type": "string"
                },
                "year": {
                    "type": "integer"
                }
            }
        },
        "usecases.Movies": {
            "type": "object",
            "properties": {
                "list_id": {
                    "type": "string"
                },
                "movies": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                }
            }
        },
        "usecases.Vote": {
            "type": "object",
            "properties": {
                "combination_id": {
                    "type": "string"
                },
                "list_id": {
                    "type": "string"
                },
                "winner_id": {
                    "type": "string"
                }
            }
        },
        "usecases.VoteOutputDTO": {
            "type": "object",
            "properties": {
                "content_message": {
                    "type": "string"
                },
                "success_message": {
                    "type": "string"
                }
            }
        },
        "util.ProblemDetails": {
            "type": "object",
            "properties": {
                "detail": {
                    "type": "string"
                },
                "instance": {
                    "type": "string"
                },
                "status": {
                    "type": "integer"
                },
                "title": {
                    "type": "string"
                },
                "type": {
                    "type": "string"
                }
            }
        }
    },
    "securityDefinitions": {
        "BearerAuth": {
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8080",
	BasePath:         "/",
	Schemes:          []string{"http"},
	Title:            "You Choose API",
	Description:      "This is an API for managing expenses.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
