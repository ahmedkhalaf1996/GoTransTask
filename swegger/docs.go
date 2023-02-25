package docs

import "github.com/swaggo/swag"

const docTemplate = `
{
    "swagger": "2.0",
    "info": {
      "title": "Scoring App",
      "description": "Authentication token should start with Bearer [space] token || example > Bearer eyJhbGciOiJIUzI1NiIsInR5.....",
      "version": "1.0"
    },
    "host": "localhost:3000",
    "basePath": "/",
    "schemes": [
      "http"
    ],
    "securityDefinitions": {
        "Bearer": {
          "type": "apiKey",
          "name": "Authorization",
          "in": "header"
        }
      },







      
    "paths": {






				
		"/credit": {
			"post": {
			  "tags": [
				"add new credit"
			  ],
			  "summary": "signup",
			  "description": "",
			  "operationId": "addPet",
			  "consumes": [
				"application/json",
				"application/xml"
			  ],
			  "produces": [
				"application/xml",
				"application/json"
			  ],
			  "parameters": [
				{
				  "in": "body",
				  "name": "body",
				  "description": "Pet object that needs to be added to the database",
				  "required": true,
				  "schema": {
					"type": "object",
					"properties": {
					  "credit": {
						"type": "number"
					  },
					  "description": {
						"type": "string"
					  }
					},
					"example": {
					  "credit": 5000,
					  "description": "Misc Expenses"
					}
				  }
				}
			  ],
			  "responses": {
				"208": {
				  "description": "username Already Exist"
				},
				"400": {
				  "description": "Invalid input"
				},
				"500": {
				  "description": "Server Error Can't crate Now User"
				}
			  }
			}
		  },


		  "/Getrecoreds": {
			"get": {
			  "tags": [
				"add new Getrecoreds"
			  ],
			  "summary": "signup",
			  "description": "",
			  "operationId": "addPet",
			  "consumes": [
				"application/json",
				"application/xml"
			  ],
			  "produces": [
				"application/xml",
				"application/json"
			  ],
			  "responses": {
				"208": {
				  "description": "username Already Exist"
				},
				"400": {
				  "description": "Invalid input"
				},
				"500": {
				  "description": "Server Error Can't crate Now User"
				}
			  }
			}
		  },



				
		  "/debit": {
			"post": {
			  "tags": [
				"Add debit"
			  ],
			  "summary": "debit",
			  "description": "",
			  "operationId": "addPet",
			  "consumes": [
				"application/json",
				"application/xml"
			  ],
			  "produces": [
				"application/xml",
				"application/json"
			  ],
			  "parameters": [
				{
				  "in": "body",
				  "name": "body",
				  "description": "Pet object that needs to be added to the database",
				  "required": true,
				  "schema": {
					"type": "object",
					"properties": {
					  "debit": {
						"type": "number"
					  },
					  "description": {
						"type": "string"
					  }
					},
					"example": {
					  "debit": 5000,
					  "description": "Misc Expenses"
					}
				  }
				}
			  ],
			  "responses": {
				"208": {
				  "description": "username Already Exist"
				},
				"400": {
				  "description": "Invalid input"
				},
				"500": {
				  "description": "Server Error Can't crate Now User"
				}
			  }
			}
		  },











    }








  }
`

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
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
