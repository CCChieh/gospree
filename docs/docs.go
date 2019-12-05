// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag at
// 2019-12-05 15:03:17.9036575 +0800 CST m=+0.065962201

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
        "license": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/HelloHandler": {
            "get": {
                "description": "这是第一个API",
                "tags": [
                    "文章"
                ],
                "summary": "第一个API",
                "responses": {
                    "200": {},
                    "400": {}
                }
            }
        },
        "/user": {
            "get": {
                "tags": [
                    "用户"
                ],
                "summary": "通过邮箱获取用户信息",
                "parameters": [
                    {
                        "type": "string",
                        "description": "用户的邮箱",
                        "name": "email",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "成功后返回值",
                        "schema": {
                            "$ref": "#/definitions/dao.User"
                        }
                    },
                    "400": {
                        "description": "fail\" \"失败后返回值",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "post": {
                "tags": [
                    "用户"
                ],
                "summary": "新建用户",
                "parameters": [
                    {
                        "description": "用户",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/user.CreateUserParams"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "成功后返回值",
                        "schema": {
                            "$ref": "#/definitions/dao.User"
                        }
                    },
                    "400": {
                        "description": "fail\" \"失败后返回值",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "dao.User": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "user.CreateUserParams": {
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
