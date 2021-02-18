// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag

package docs

import (
	"bytes"
	"encoding/json"
	"strings"

	"github.com/alecthomas/template"
	"github.com/swaggo/swag"
)

var doc = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{.Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/event": {
            "put": {
                "description": "Create an event in a specific organization",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Event"
                ],
                "summary": "Create event in organization",
                "parameters": [
                    {
                        "description": "Event Information",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.EventApi"
                        }
                    }
                ]
            }
        },
        "/event/interview/all/{eid}": {
            "get": {
                "description": "Get brief information of all interviews in a specific event",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Interview"
                ],
                "summary": "Get all interviews in event",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Event ID",
                        "name": "eid",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/model.Brief"
                            }
                        }
                    }
                }
            }
        },
        "/event/interview/{eid}{iid}": {
            "get": {
                "description": "Get information of an interview in a specific event",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Interview"
                ],
                "summary": "Get interview in event",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Event ID",
                        "name": "eid",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "Interview ID",
                        "name": "iid",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.InterviewApi"
                        }
                    }
                }
            }
        },
        "/event/{eid}": {
            "get": {
                "description": "Get information of an event",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Event"
                ],
                "summary": "Get event",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Event ID",
                        "name": "eid",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.EventApi"
                        }
                    }
                }
            },
            "post": {
                "description": "Update an event",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Event"
                ],
                "summary": "Update event",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Event ID",
                        "name": "eid",
                        "in": "query",
                        "required": true
                    },
                    {
                        "description": "Event Information",
                        "name": "data",
                        "in": "body",
                        "schema": {
                            "$ref": "#/definitions/model.EventApi"
                        }
                    }
                ]
            }
        },
        "/interview": {
            "put": {
                "description": "Create an events in a specific organization",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Interview"
                ],
                "summary": "Create interview in event",
                "parameters": [
                    {
                        "description": "Interview Information",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.InterviewApi"
                        }
                    }
                ]
            }
        },
        "/interview/{iid}": {
            "get": {
                "description": "Get information of an interview",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Interview"
                ],
                "summary": "Get interview",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Interview ID",
                        "name": "iid",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.InterviewApi"
                        }
                    }
                }
            },
            "post": {
                "description": "Update an interview",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Interview"
                ],
                "summary": "Update interview",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Interview ID",
                        "name": "iid",
                        "in": "query",
                        "required": true
                    },
                    {
                        "description": "Interview Information",
                        "name": "data",
                        "in": "body",
                        "schema": {
                            "$ref": "#/definitions/model.InterviewApi"
                        }
                    }
                ]
            }
        },
        "/message": {
            "put": {
                "description": "send a message",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Message"
                ],
                "summary": "Send a message",
                "parameters": [
                    {
                        "description": "Message Information",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.MessageRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.MessageAPI"
                        }
                    }
                }
            }
        },
        "/message/template": {
            "put": {
                "description": "Add a message template",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "MessageTemplate"
                ],
                "summary": "Add a message template",
                "parameters": [
                    {
                        "description": "Message Template Infomation",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.MessageTemplateRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.MessageTemplateAPI"
                        }
                    }
                }
            }
        },
        "/message/template/all/{oid}": {
            "get": {
                "description": "Get information of all message templates of a specific organization",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "MessageTemplate"
                ],
                "summary": "Get all message templates",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Organization ID",
                        "name": "oid",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/model.AllMessageTemplateAPI"
                            }
                        }
                    }
                }
            }
        },
        "/message/template/{tid}": {
            "get": {
                "description": "Get information of a specific message template",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "MessageTemplate"
                ],
                "summary": "Get a message template",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Message Template ID",
                        "name": "tid",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.MessageTemplateAPI"
                        }
                    }
                }
            },
            "post": {
                "description": "Update a message template",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "MessageTemplate"
                ],
                "summary": "Update a message template",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Message Template ID",
                        "name": "tid",
                        "in": "query",
                        "required": true
                    },
                    {
                        "description": "Message Template Infomation",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.MessageTemplateRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.MessageTemplateAPI"
                        }
                    }
                }
            }
        },
        "/message/{mid}": {
            "get": {
                "description": "Get information of a specific message",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Message"
                ],
                "summary": "Get a message",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Message ID",
                        "name": "mid",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.MessageAPI"
                        }
                    }
                }
            }
        },
        "/organization/all": {
            "get": {
                "description": "Get brief information of all organizations",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Organization"
                ],
                "summary": "Get all organizations",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/model.Brief"
                            }
                        }
                    }
                }
            }
        },
        "/organization/department/all/{oid}": {
            "get": {
                "description": "Get brief information of all departments in a specific organization",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Department"
                ],
                "summary": "Get all departments in organization",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Organization ID",
                        "name": "oid",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/model.Brief"
                            }
                        }
                    }
                }
            }
        },
        "/organization/department/{oid}{did}": {
            "get": {
                "description": "Get information of a department in a specific organization",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Department"
                ],
                "summary": "Get department in organization",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Organization ID",
                        "name": "oid",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "Department ID",
                        "name": "did",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.DepartmentApi"
                        }
                    }
                }
            }
        },
        "/organization/event/all/{oid}": {
            "get": {
                "description": "Get brief information of all events in a specific organization",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Event"
                ],
                "summary": "Get all events in organization",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Organization ID",
                        "name": "oid",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/model.Brief"
                            }
                        }
                    }
                }
            }
        },
        "/organization/event/{oid}{eid}": {
            "get": {
                "description": "Get information of an event in a specific organization",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Event"
                ],
                "summary": "Get event in organization",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Organization ID",
                        "name": "oid",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "Event ID",
                        "name": "eid",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.EventApi"
                        }
                    }
                }
            }
        },
        "/organization/{oid}": {
            "get": {
                "description": "Get information of a specific organization",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Organization"
                ],
                "summary": "Get information of organization",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Organization ID",
                        "name": "oid",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.OrganizationApi"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "model.AllMessageTemplateAPI": {
            "type": "object",
            "properties": {
                "description": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "status": {
                    "type": "integer"
                }
            }
        },
        "model.Brief": {
            "type": "object",
            "properties": {
                "description": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "model.DepartmentApi": {
            "type": "object",
            "properties": {
                "description": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "organizationID": {
                    "type": "integer"
                }
            }
        },
        "model.EventApi": {
            "type": "object",
            "required": [
                "EndTime",
                "Name",
                "OrganizationID",
                "StartTime"
            ],
            "properties": {
                "Description": {
                    "type": "string"
                },
                "EndTime": {
                    "description": "request string must be in RFC 3339 format",
                    "type": "string"
                },
                "Name": {
                    "type": "string"
                },
                "OrganizationID": {
                    "type": "integer"
                },
                "OtherInfo": {
                    "type": "string"
                },
                "StartTime": {
                    "description": "request string must be in RFC 3339 format",
                    "type": "string"
                },
                "Status": {
                    "description": "0 disabled (default), 1 testing, 2 running",
                    "type": "integer"
                },
                "id": {
                    "type": "integer"
                }
            }
        },
        "model.InterviewApi": {
            "type": "object",
            "required": [
                "DepartmentID",
                "EndTime",
                "EventID",
                "Name",
                "StartTime"
            ],
            "properties": {
                "DepartmentID": {
                    "type": "integer"
                },
                "Description": {
                    "type": "string"
                },
                "EndTime": {
                    "description": "request string must be in RFC 3339 format",
                    "type": "string"
                },
                "EventID": {
                    "type": "integer"
                },
                "Location": {
                    "type": "string"
                },
                "MaxInterviewee": {
                    "description": "default 6",
                    "type": "integer"
                },
                "Name": {
                    "type": "string"
                },
                "OtherInfo": {
                    "type": "string"
                },
                "StartTime": {
                    "description": "request string must be in RFC 3339 format",
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                }
            }
        },
        "model.MessageAPI": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                },
                "receiverID": {
                    "type": "integer"
                },
                "reply": {
                    "type": "string"
                },
                "senderID": {
                    "type": "integer"
                },
                "text": {
                    "type": "string"
                }
            }
        },
        "model.MessageRequest": {
            "type": "object",
            "required": [
                "MessageTemplateID",
                "ReceiverID",
                "SenderID"
            ],
            "properties": {
                "CrossInterviewID": {
                    "description": "TODO(TO/GA): wait for logic",
                    "type": "integer"
                },
                "FormID": {
                    "type": "integer"
                },
                "InterviewID": {
                    "type": "integer"
                },
                "MessageTemplateID": {
                    "type": "integer"
                },
                "ReceiverID": {
                    "type": "integer"
                },
                "SenderID": {
                    "type": "integer"
                }
            }
        },
        "model.MessageTemplateAPI": {
            "type": "object",
            "properties": {
                "description": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "organizationID": {
                    "type": "integer"
                },
                "status": {
                    "type": "integer"
                },
                "text": {
                    "type": "string"
                }
            }
        },
        "model.MessageTemplateRequest": {
            "type": "object",
            "required": [
                "Description",
                "Text"
            ],
            "properties": {
                "Description": {
                    "type": "string"
                },
                "OrganizationID": {
                    "description": "not required because it might be 0",
                    "type": "integer"
                },
                "Text": {
                    "type": "string"
                }
            }
        },
        "model.OrganizationApi": {
            "type": "object",
            "properties": {
                "description": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                }
            }
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
	Version:     "0.1",
	Host:        "rop-neo-staging.zjuqsc.com",
	BasePath:    "/api",
	Schemes:     []string{},
	Title:       "Recruit Open Platform API",
	Description: "This API will be used under staging environment.",
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
	swag.Register(swag.Name, &s{})
}
