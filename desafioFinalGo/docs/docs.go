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
        "termsOfService": "https://publicapis.ml/terms",
        "contact": {
            "name": "Public APIs",
            "url": "https://dev.publicapis.ml"
        },
        "license": {
            "name": "Apache 2.0",
            "url": "https://www.apache.org/licenses/LICENSE-2.0.html"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/consultas": {
            "post": {
                "description": "Post a consult",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "CRUD CONSULTAS"
                ],
                "summary": "Post a consult",
                "parameters": [
                    {
                        "description": "Consulta",
                        "name": "json",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/domain.Consulta"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/domain.Consulta"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/web.errorResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/web.errorResponse"
                        }
                    }
                }
            }
        },
        "/consultas/consulta_paciente/{rg}": {
            "get": {
                "description": "Get a consult by rg",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "CRUD CONSULTAS"
                ],
                "summary": "Get a consult by rg",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Consulta RG",
                        "name": "rg",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/domain.ConsultaDTO"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/web.errorResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/web.errorResponse"
                        }
                    }
                }
            }
        },
        "/consultas/{id}": {
            "get": {
                "description": "Get a consult by id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "CRUD CONSULTAS"
                ],
                "summary": "Get a consult by id",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Consulta id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/domain.Consulta"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/web.errorResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/web.errorResponse"
                        }
                    }
                }
            },
            "put": {
                "description": "Put a consult",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "CRUD CONSULTAS"
                ],
                "summary": "Put a consult",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Consulta ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Consulta",
                        "name": "json",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/domain.Consulta"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/domain.Consulta"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/web.errorResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/web.errorResponse"
                        }
                    }
                }
            },
            "delete": {
                "description": "Delete a consult by ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "CRUD CONSULTAS"
                ],
                "summary": "Delete a consult by ID",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Consulta ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/domain.Consulta"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/web.errorResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/web.errorResponse"
                        }
                    }
                }
            },
            "patch": {
                "description": "Patch a consult",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "CRUD CONSULTAS"
                ],
                "summary": "Patch a consult",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Consulta ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Consulta",
                        "name": "json",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/domain.Consulta"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/domain.Consulta"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/web.errorResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/web.errorResponse"
                        }
                    }
                }
            }
        },
        "/dentistas": {
            "get": {
                "description": "getall dentistas",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "CRUD DENTISTAS"
                ],
                "summary": "List dentistas",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/domain.Dentista"
                        }
                    }
                }
            },
            "post": {
                "description": "Post a dentist",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "CRUD DENTISTAS"
                ],
                "summary": "Post a dentist",
                "parameters": [
                    {
                        "description": "Dentista",
                        "name": "json",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/domain.Dentista"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/domain.Dentista"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/web.errorResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/web.errorResponse"
                        }
                    }
                }
            }
        },
        "/dentistas/{id}": {
            "get": {
                "description": "Get a dentist by ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "CRUD DENTISTAS"
                ],
                "summary": "Get a dentista by ID",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Dentista ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/domain.Dentista"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/web.errorResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/web.errorResponse"
                        }
                    }
                }
            },
            "put": {
                "description": "Put a dentist",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "CRUD DENTISTAS"
                ],
                "summary": "Put a dentist",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Dentista ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Dentista",
                        "name": "json",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/domain.Dentista"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/domain.Dentista"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/web.errorResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/web.errorResponse"
                        }
                    }
                }
            },
            "delete": {
                "description": "Delete a dentist by ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "CRUD DENTISTAS"
                ],
                "summary": "Delete a dentist by ID",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Dentista ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/domain.Dentista"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/web.errorResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/web.errorResponse"
                        }
                    }
                }
            },
            "patch": {
                "description": "Patch a dentist",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "CRUD DENTISTAS"
                ],
                "summary": "Patch a dentist",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Dentista ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Dentista",
                        "name": "json",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/domain.Dentista"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/domain.Dentista"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/web.errorResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/web.errorResponse"
                        }
                    }
                }
            }
        },
        "/pacientes": {
            "get": {
                "description": "get pacientes",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "CRUD PACIENTES"
                ],
                "summary": "List pacientes",
                "parameters": [
                    {
                        "type": "string",
                        "description": "SECRET_TOKEN",
                        "name": "SECRET_TOKEN",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/domain.Paciente"
                        }
                    }
                }
            },
            "post": {
                "description": "Post a paciente",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "CRUD PACIENTES"
                ],
                "summary": "Post a paciente",
                "parameters": [
                    {
                        "type": "string",
                        "description": "SECRET_TOKEN",
                        "name": "SECRET_TOKEN",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "Paciente",
                        "name": "json",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/domain.Paciente"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/domain.Paciente"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/web.errorResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/web.errorResponse"
                        }
                    }
                }
            }
        },
        "/pacientes/{id}": {
            "get": {
                "description": "Get a paciente by ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "CRUD PACIENTES"
                ],
                "summary": "Get a paciente by ID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "SECRET_TOKEN",
                        "name": "SECRET_TOKEN",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "Paciente ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/domain.Paciente"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/web.errorResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/web.errorResponse"
                        }
                    }
                }
            },
            "put": {
                "description": "Put a paciente",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "CRUD PACIENTES"
                ],
                "summary": "Put a paciente",
                "parameters": [
                    {
                        "type": "string",
                        "description": "SECRET_TOKEN",
                        "name": "SECRET_TOKEN",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "Paciente ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Paciente",
                        "name": "json",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/domain.Paciente"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/domain.Paciente"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/web.errorResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/web.errorResponse"
                        }
                    }
                }
            },
            "delete": {
                "description": "Delete a paciente by ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "CRUD PACIENTES"
                ],
                "summary": "Delete a paciente by ID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "SECRET_TOKEN",
                        "name": "SECRET_TOKEN",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "Paciente ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/domain.Paciente"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/web.errorResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/web.errorResponse"
                        }
                    }
                }
            },
            "patch": {
                "description": "Patch a paciente",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "CRUD PACIENTES"
                ],
                "summary": "Patch a paciente",
                "parameters": [
                    {
                        "type": "string",
                        "description": "SECRET_TOKEN",
                        "name": "SECRET_TOKEN",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "Paciente ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Paciente",
                        "name": "json",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/domain.Paciente"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/domain.Paciente"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/web.errorResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/web.errorResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "domain.Consulta": {
            "type": "object",
            "properties": {
                "data": {
                    "type": "string"
                },
                "dentista": {
                    "$ref": "#/definitions/domain.Dentista"
                },
                "dentista_id": {
                    "type": "integer"
                },
                "descricao": {
                    "type": "string"
                },
                "hora": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "paciente": {
                    "$ref": "#/definitions/domain.Paciente"
                },
                "paciente_id": {
                    "type": "integer"
                }
            }
        },
        "domain.ConsultaDTO": {
            "type": "object",
            "properties": {
                "data": {
                    "type": "string"
                },
                "dentista": {
                    "$ref": "#/definitions/domain.Dentista"
                },
                "dentista_id": {
                    "type": "integer"
                },
                "descricao": {
                    "type": "string"
                },
                "hora": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "paciente": {
                    "$ref": "#/definitions/domain.Paciente"
                },
                "paciente_id": {
                    "type": "integer"
                }
            }
        },
        "domain.Dentista": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                },
                "matricula": {
                    "type": "integer"
                },
                "nome": {
                    "type": "string"
                },
                "sobrenome": {
                    "type": "string"
                }
            }
        },
        "domain.Paciente": {
            "type": "object",
            "properties": {
                "data_cadastro": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "nome": {
                    "type": "string"
                },
                "rg": {
                    "type": "integer"
                },
                "sorenome": {
                    "type": "string"
                }
            }
        },
        "web.errorResponse": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "string"
                },
                "message": {
                    "type": "string"
                },
                "status": {
                    "type": "integer"
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
	Title:            "Store API",
	Description:      "This API handle products from our store",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
