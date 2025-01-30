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
        "/diagnostics": {
            "get": {
                "description": "Return person list",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "diagnostics"
                ],
                "summary": "List diagnostics",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.Diagnostic"
                            }
                        }
                    }
                }
            },
            "post": {
                "description": "Create a new diagnostic",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "diagnostics"
                ],
                "summary": "Create diagnostic",
                "parameters": [
                    {
                        "description": "Data of diagnostic",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/requests.DiagnosticRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/models.Diagnostic"
                        }
                    },
                    "400": {
                        "description": "Validation error",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    }
                }
            }
        },
        "/diagnostics/:id": {
            "get": {
                "description": "Return diagnostic by id",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "diagnostics"
                ],
                "summary": "Return diagnostic",
                "parameters": [
                    {
                        "type": "integer",
                        "example": 1,
                        "description": "Diagnostic ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Diagnostic"
                        }
                    },
                    "404": {
                        "description": "Diagnostic not found",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    }
                }
            },
            "put": {
                "description": "Update a exist diagnostic",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "diagnostics"
                ],
                "summary": "Update  diagnostic",
                "parameters": [
                    {
                        "description": "Data of diagnostic",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/requests.DiagnosticRequest"
                        }
                    },
                    {
                        "type": "integer",
                        "example": 1,
                        "description": "Diagnostic ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Diagnostic"
                        }
                    }
                }
            }
        },
        "/medications": {
            "get": {
                "description": "Return medication list",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "medications"
                ],
                "summary": "List medications",
                "responses": {
                    "200": {
                        "description": "Successful response",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.Medication"
                            }
                        }
                    }
                }
            },
            "post": {
                "description": "Create a new medication",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "medications"
                ],
                "summary": "Create medication",
                "parameters": [
                    {
                        "description": "Data of medication",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/requests.MedicationRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Successful response",
                        "schema": {
                            "$ref": "#/definitions/models.Medication"
                        }
                    },
                    "400": {
                        "description": "Validation error",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    }
                }
            }
        },
        "/medications/:id": {
            "get": {
                "description": "Return medication by id",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "medications"
                ],
                "summary": "Return medication",
                "parameters": [
                    {
                        "type": "integer",
                        "example": 1,
                        "description": "Medication ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Successful response",
                        "schema": {
                            "$ref": "#/definitions/models.Medication"
                        }
                    },
                    "404": {
                        "description": "Medication not found",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    }
                }
            },
            "put": {
                "description": "Update a exist medication",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "medications"
                ],
                "summary": "Update  medication",
                "parameters": [
                    {
                        "description": "Data of medication",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/requests.MedicationRequest"
                        }
                    },
                    {
                        "type": "integer",
                        "example": 1,
                        "description": "Medication ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Successful response",
                        "schema": {
                            "$ref": "#/definitions/models.Medication"
                        }
                    },
                    "400": {
                        "description": "Validation error",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    }
                }
            }
        },
        "/persons": {
            "get": {
                "description": "Return person list",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "persons"
                ],
                "summary": "List persons",
                "responses": {
                    "200": {
                        "description": "Successful response",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.Person"
                            }
                        }
                    }
                }
            },
            "post": {
                "description": "Create a new person",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "persons"
                ],
                "summary": "Create person",
                "parameters": [
                    {
                        "description": "Data of person",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/requests.PersonRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Successful response",
                        "schema": {
                            "$ref": "#/definitions/models.Person"
                        }
                    },
                    "400": {
                        "description": "Validation error",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    }
                }
            }
        },
        "/persons/:id": {
            "get": {
                "description": "Return person by id",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "persons"
                ],
                "summary": "Return person",
                "parameters": [
                    {
                        "type": "integer",
                        "example": 1,
                        "description": "Person ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Successful response",
                        "schema": {
                            "$ref": "#/definitions/models.Person"
                        }
                    },
                    "404": {
                        "description": "Person not found",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    }
                }
            },
            "put": {
                "description": "Update a exist person",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "persons"
                ],
                "summary": "Update  person",
                "parameters": [
                    {
                        "description": "Data of person",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/requests.PersonRequest"
                        }
                    },
                    {
                        "type": "integer",
                        "example": 1,
                        "description": "Mdication ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Successful response",
                        "schema": {
                            "$ref": "#/definitions/models.Person"
                        }
                    },
                    "400": {
                        "description": "Validation error",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    }
                }
            }
        },
        "/symptoms": {
            "get": {
                "description": "Return person list",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "symptoms"
                ],
                "summary": "List symptoms",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.Symptom"
                            }
                        }
                    }
                }
            },
            "post": {
                "description": "Create a new symptom",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "symptoms"
                ],
                "summary": "Create symptom",
                "parameters": [
                    {
                        "description": "Data of symptom",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/requests.SymptomRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/models.Symptom"
                        }
                    },
                    "400": {
                        "description": "Validation error",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    }
                }
            }
        },
        "/symptoms/:id": {
            "get": {
                "description": "Return symptom by id",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "symptoms"
                ],
                "summary": "Return symptom",
                "parameters": [
                    {
                        "type": "integer",
                        "example": 1,
                        "description": "Symptom ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Symptom"
                        }
                    },
                    "404": {
                        "description": "Symptom not found",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    }
                }
            },
            "put": {
                "description": "Update a exist symptom",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "symptoms"
                ],
                "summary": "Update  symptom",
                "parameters": [
                    {
                        "description": "Data of symptom",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/requests.SymptomRequest"
                        }
                    },
                    {
                        "type": "integer",
                        "example": 1,
                        "description": "Symptom ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Symptom"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "models.Diagnostic": {
            "type": "object",
            "required": [
                "description",
                "doctor_id",
                "patient_id",
                "symptoms"
            ],
            "properties": {
                "createdAt": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "doctor_id": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "patient_id": {
                    "type": "string"
                },
                "symptoms": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "updatedAt": {
                    "type": "string"
                }
            }
        },
        "models.Medication": {
            "type": "object",
            "required": [
                "activeIngredient",
                "companyHoldingRegistration",
                "name",
                "registrationNumber",
                "regulatoryCategory",
                "therapeuticClass"
            ],
            "properties": {
                "activeIngredient": {
                    "type": "string"
                },
                "companyHoldingRegistration": {
                    "type": "string"
                },
                "createdAt": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "registrationNumber": {
                    "type": "string"
                },
                "regulatoryCategory": {
                    "type": "string"
                },
                "therapeuticClass": {
                    "type": "string"
                },
                "updatedAt": {
                    "type": "string"
                }
            }
        },
        "models.Person": {
            "type": "object",
            "required": [
                "age",
                "name"
            ],
            "properties": {
                "age": {
                    "type": "integer"
                },
                "createdAt": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "updatedAt": {
                    "type": "string"
                }
            }
        },
        "models.Symptom": {
            "type": "object",
            "required": [
                "name"
            ],
            "properties": {
                "createdAt": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "updatedAt": {
                    "type": "string"
                }
            }
        },
        "requests.DiagnosticRequest": {
            "type": "object",
            "required": [
                "description",
                "doctor_id",
                "patient_id",
                "symptoms"
            ],
            "properties": {
                "description": {
                    "type": "string",
                    "example": "Patient has a mild fever"
                },
                "doctor_id": {
                    "type": "string",
                    "example": "1"
                },
                "patient_id": {
                    "type": "string",
                    "example": "1"
                },
                "symptoms": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    },
                    "example": [
                        "['cough'",
                        " 'headache'",
                        " 'fever']"
                    ]
                }
            }
        },
        "requests.MedicationRequest": {
            "type": "object",
            "required": [
                "CompanyHoldingRegistration",
                "activeIngredient",
                "name",
                "registrationNumber",
                "regulatoryCategory",
                "therapeuticClass"
            ],
            "properties": {
                "CompanyHoldingRegistration": {
                    "type": "string"
                },
                "activeIngredient": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "registrationNumber": {
                    "type": "string"
                },
                "regulatoryCategory": {
                    "type": "string"
                },
                "therapeuticClass": {
                    "type": "string"
                }
            }
        },
        "requests.PersonRequest": {
            "type": "object",
            "required": [
                "age",
                "name"
            ],
            "properties": {
                "age": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "requests.SymptomRequest": {
            "type": "object",
            "required": [
                "name"
            ],
            "properties": {
                "name": {
                    "type": "string",
                    "example": "fever"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8080",
	BasePath:         "/api/v1",
	Schemes:          []string{},
	Title:            "Health History API",
	Description:      "The objective this application is save history medicaments of person.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
