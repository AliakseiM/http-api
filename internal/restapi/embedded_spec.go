// Code generated by go-swagger; DO NOT EDIT.

package restapi

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
)

var (
	// SwaggerJSON embedded version of the swagger document used at generation time
	SwaggerJSON json.RawMessage
	// FlatSwaggerJSON embedded flattened version of the swagger document used at generation time
	FlatSwaggerJSON json.RawMessage
)

func init() {
	SwaggerJSON = json.RawMessage([]byte(`{
  "swagger": "2.0",
  "info": {
    "title": "skeleton",
    "version": "0.0.1"
  },
  "basePath": "/api/v1",
  "paths": {
    "/health": {
      "get": {
        "produces": [
          "application/json"
        ],
        "tags": [
          "debug"
        ],
        "summary": "Health check",
        "operationId": "health",
        "responses": {
          "200": {
            "description": "Success",
            "schema": {
              "$ref": "#/definitions/Success"
            }
          }
        }
      }
    }
  },
  "definitions": {
    "Success": {
      "type": "object",
      "properties": {
        "message": {
          "type": "string",
          "example": "OK"
        }
      }
    }
  }
}`))
	FlatSwaggerJSON = json.RawMessage([]byte(`{
  "swagger": "2.0",
  "info": {
    "title": "skeleton",
    "version": "0.0.1"
  },
  "basePath": "/api/v1",
  "paths": {
    "/health": {
      "get": {
        "produces": [
          "application/json"
        ],
        "tags": [
          "debug"
        ],
        "summary": "Health check",
        "operationId": "health",
        "responses": {
          "200": {
            "description": "Success",
            "schema": {
              "$ref": "#/definitions/Success"
            }
          }
        }
      }
    }
  },
  "definitions": {
    "Success": {
      "type": "object",
      "properties": {
        "message": {
          "type": "string",
          "example": "OK"
        }
      }
    }
  }
}`))
}
