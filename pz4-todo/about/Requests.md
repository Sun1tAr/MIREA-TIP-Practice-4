```json
{
	"info": {
		"_postman_id": "41a9fb3c-2aa9-4e24-b8ce-07b1e3503fc8",
		"name": "MIREA-TIP-Practice-4",
		"schema": "https://schema.getpostman.com/json/collection/v2.1.0/collection.json",
		"_exporter_id": "42992055"
	},
	"item": [
		{
			"name": "V1",
			"item": [
				{
					"name": "V1: GET /health",
					"request": {
						"auth": {
							"type": "noauth"
						},
						"method": "GET",
						"header": [
							{
								"key": "Content-Type",
								"value": "application/json",
								"type": "default"
							},
							{
								"key": "Access-Control-Allow-Origin",
								"value": "123",
								"type": "default"
							}
						],
						"url": {
							"raw": "http://localhost:8080/api/v1/health",
							"protocol": "http",
							"host": [
								"localhost"
							],
							"port": "8080",
							"path": [
								"api",
								"v1",
								"health"
							]
						}
					},
					"response": []
				},
				{
					"name": "V1: POST /tasks + rawData",
					"request": {
						"auth": {
							"type": "noauth"
						},
						"method": "POST",
						"header": [
							{
								"key": "Content-Type",
								"value": "application/json",
								"type": "default"
							},
							{
								"key": "Access-Control-Allow-Origin",
								"value": "123",
								"type": "default"
							}
						],
						"body": {
							"mode": "raw",
							"raw": "{\"title\":\"Молоко съесть\"}"
						},
						"url": {
							"raw": "{{baseURL}}/v1/tasks",
							"host": [
								"{{baseURL}}"
							],
							"path": [
								"v1",
								"tasks"
							]
						}
					},
					"response": []
				},
				{
					"name": "V1: GET /tasks/{id}",
					"request": {
						"auth": {
							"type": "noauth"
						},
						"method": "GET",
						"header": [
							{
								"key": "Content-Type",
								"value": "application/json",
								"type": "default"
							},
							{
								"key": "Access-Control-Allow-Origin",
								"value": "123",
								"type": "default"
							}
						],
						"url": {
							"raw": "{{baseURL}}/v1/tasks/1",
							"host": [
								"{{baseURL}}"
							],
							"path": [
								"v1",
								"tasks",
								"1"
							]
						}
					},
					"response": []
				},
				{
					"name": "V1: GET /tasks",
					"request": {
						"auth": {
							"type": "noauth"
						},
						"method": "GET",
						"header": [
							{
								"key": "Content-Type",
								"value": "application/json",
								"type": "default"
							},
							{
								"key": "Access-Control-Allow-Origin",
								"value": "123",
								"type": "default"
							}
						],
						"url": {
							"raw": "{{baseURL}}/v1/tasks",
							"host": [
								"{{baseURL}}"
							],
							"path": [
								"v1",
								"tasks"
							]
						}
					},
					"response": []
				},
				{
					"name": "V1: GET /tasks doneFlag",
					"request": {
						"auth": {
							"type": "noauth"
						},
						"method": "GET",
						"header": [
							{
								"key": "Content-Type",
								"value": "application/json",
								"type": "default"
							},
							{
								"key": "Access-Control-Allow-Origin",
								"value": "123",
								"type": "default"
							}
						],
						"url": {
							"raw": "{{baseURL}}/v1/tasks?done=true",
							"host": [
								"{{baseURL}}"
							],
							"path": [
								"v1",
								"tasks"
							],
							"query": [
								{
									"key": "done",
									"value": "true"
								}
							]
						}
					},
					"response": []
				},
				{
					"name": "V1: DELETE /tasks/{id}",
					"request": {
						"auth": {
							"type": "noauth"
						},
						"method": "DELETE",
						"header": [
							{
								"key": "Content-Type",
								"value": "application/json",
								"type": "default"
							},
							{
								"key": "Access-Control-Allow-Origin",
								"value": "123",
								"type": "default"
							}
						],
						"url": {
							"raw": "{{baseURL}}/v1/tasks/1",
							"host": [
								"{{baseURL}}"
							],
							"path": [
								"v1",
								"tasks",
								"1"
							]
						}
					},
					"response": []
				},
				{
					"name": "V1: PUT /tasks/{id} + rawData",
					"request": {
						"auth": {
							"type": "noauth"
						},
						"method": "PUT",
						"header": [
							{
								"key": "Content-Type",
								"value": "application/json",
								"type": "default"
							},
							{
								"key": "Access-Control-Allow-Origin",
								"value": "123",
								"type": "default"
							}
						],
						"body": {
							"mode": "raw",
							"raw": "{\r\n    \"title\": \"Молоко съесть\",\r\n    \"done\": true\r\n}",
							"options": {
								"raw": {
									"language": "json"
								}
							}
						},
						"url": {
							"raw": "{{baseURL}}/v1/tasks/3",
							"host": [
								"{{baseURL}}"
							],
							"path": [
								"v1",
								"tasks",
								"3"
							]
						}
					},
					"response": []
				},
				{
					"name": "V1: GET /tasks pagination",
					"request": {
						"auth": {
							"type": "noauth"
						},
						"method": "GET",
						"header": [
							{
								"key": "Content-Type",
								"value": "application/json",
								"type": "default"
							},
							{
								"key": "Access-Control-Allow-Origin",
								"value": "123",
								"type": "default"
							}
						],
						"url": {
							"raw": "{{baseURL}}/v1/tasks?page=2&limit=5",
							"host": [
								"{{baseURL}}"
							],
							"path": [
								"v1",
								"tasks"
							],
							"query": [
								{
									"key": "page",
									"value": "2"
								},
								{
									"key": "limit",
									"value": "5"
								}
							]
						}
					},
					"response": []
				}
			]
		},
		{
			"name": "V2",
			"item": [
				{
					"name": "V2: GET /health",
					"request": {
						"auth": {
							"type": "noauth"
						},
						"method": "GET",
						"header": [
							{
								"key": "Content-Type",
								"value": "application/json",
								"type": "default"
							},
							{
								"key": "Access-Control-Allow-Origin",
								"value": "123",
								"type": "default"
							}
						],
						"url": {
							"raw": "{{baseURL}}/v2/health",
							"host": [
								"{{baseURL}}"
							],
							"path": [
								"v2",
								"health"
							]
						}
					},
					"response": []
				},
				{
					"name": "V2: POST /tasks + rawData",
					"request": {
						"auth": {
							"type": "noauth"
						},
						"method": "POST",
						"header": [
							{
								"key": "Content-Type",
								"value": "application/json",
								"type": "default"
							},
							{
								"key": "Access-Control-Allow-Origin",
								"value": "123",
								"type": "default"
							}
						],
						"body": {
							"mode": "raw",
							"raw": "{\"title\":\"Молоко съесть\"}"
						},
						"url": {
							"raw": "{{baseURL}}/v2/tasks",
							"host": [
								"{{baseURL}}"
							],
							"path": [
								"v2",
								"tasks"
							]
						}
					},
					"response": []
				},
				{
					"name": "V2: GET /tasks/{id}",
					"request": {
						"auth": {
							"type": "noauth"
						},
						"method": "GET",
						"header": [
							{
								"key": "Content-Type",
								"value": "application/json",
								"type": "default"
							},
							{
								"key": "Access-Control-Allow-Origin",
								"value": "123",
								"type": "default"
							}
						],
						"url": {
							"raw": "{{baseURL}}/v2/tasks/1",
							"host": [
								"{{baseURL}}"
							],
							"path": [
								"v2",
								"tasks",
								"1"
							]
						}
					},
					"response": []
				},
				{
					"name": "V2: GET /tasks",
					"request": {
						"auth": {
							"type": "noauth"
						},
						"method": "GET",
						"header": [
							{
								"key": "Content-Type",
								"value": "application/json",
								"type": "default"
							},
							{
								"key": "Access-Control-Allow-Origin",
								"value": "123",
								"type": "default"
							}
						],
						"url": {
							"raw": "{{baseURL}}/v2/tasks",
							"host": [
								"{{baseURL}}"
							],
							"path": [
								"v2",
								"tasks"
							]
						}
					},
					"response": []
				},
				{
					"name": "V2: DELETE /tasks/{id}",
					"request": {
						"auth": {
							"type": "noauth"
						},
						"method": "DELETE",
						"header": [
							{
								"key": "Content-Type",
								"value": "application/json",
								"type": "default"
							},
							{
								"key": "Access-Control-Allow-Origin",
								"value": "123",
								"type": "default"
							}
						],
						"url": {
							"raw": "{{baseURL}}/v1/tasks/1",
							"host": [
								"{{baseURL}}"
							],
							"path": [
								"v1",
								"tasks",
								"1"
							]
						}
					},
					"response": []
				},
				{
					"name": "V2: PUT /tasks/{id} + rawData",
					"request": {
						"auth": {
							"type": "noauth"
						},
						"method": "PUT",
						"header": [
							{
								"key": "Content-Type",
								"value": "application/json",
								"type": "default"
							},
							{
								"key": "Access-Control-Allow-Origin",
								"value": "123",
								"type": "default"
							}
						],
						"body": {
							"mode": "raw",
							"raw": "{\r\n    \"title\": \"Молоко съесть\",\r\n    \"done\": true\r\n}",
							"options": {
								"raw": {
									"language": "json"
								}
							}
						},
						"url": {
							"raw": "{{baseURL}}/v2/tasks/3",
							"host": [
								"{{baseURL}}"
							],
							"path": [
								"v2",
								"tasks",
								"3"
							]
						}
					},
					"response": []
				},
				{
					"name": "V2: GET /tasks doneFlag",
					"request": {
						"auth": {
							"type": "noauth"
						},
						"method": "GET",
						"header": [
							{
								"key": "Content-Type",
								"value": "application/json",
								"type": "default"
							},
							{
								"key": "Access-Control-Allow-Origin",
								"value": "123",
								"type": "default"
							}
						],
						"url": {
							"raw": "{{baseURL}}/v2/tasks?done=true",
							"host": [
								"{{baseURL}}"
							],
							"path": [
								"v2",
								"tasks"
							],
							"query": [
								{
									"key": "done",
									"value": "true"
								}
							]
						}
					},
					"response": []
				},
				{
					"name": "V2: GET /tasks pagination",
					"request": {
						"auth": {
							"type": "noauth"
						},
						"method": "GET",
						"header": [
							{
								"key": "Content-Type",
								"value": "application/json",
								"type": "default"
							},
							{
								"key": "Access-Control-Allow-Origin",
								"value": "123",
								"type": "default"
							}
						],
						"url": {
							"raw": "{{baseURL}}/v2/tasks?page=2&limit=5",
							"host": [
								"{{baseURL}}"
							],
							"path": [
								"v2",
								"tasks"
							],
							"query": [
								{
									"key": "page",
									"value": "2"
								},
								{
									"key": "limit",
									"value": "5"
								}
							]
						}
					},
					"response": []
				}
			]
		},
		{
			"name": "GET /health",
			"request": {
				"auth": {
					"type": "noauth"
				},
				"method": "GET",
				"header": [],
				"url": {
					"raw": "{{baseURL}}/health",
					"host": [
						"{{baseURL}}"
					],
					"path": [
						"health"
					]
				}
			},
			"response": []
		},
		{
			"name": "V1: PATCH /tasks/{id}",
			"request": {
				"auth": {
					"type": "noauth"
				},
				"method": "PATCH",
				"header": [
					{
						"key": "Content-Type",
						"value": "application/json",
						"type": "default"
					},
					{
						"key": "Access-Control-Allow-Origin",
						"value": "123",
						"type": "default"
					}
				],
				"url": {
					"raw": "{{baseURL}}/tasks/1",
					"host": [
						"{{baseURL}}"
					],
					"path": [
						"tasks",
						"1"
					]
				}
			},
			"response": []
		}
	],
	"variable": [
		{
			"key": "baseURL",
			"value": "",
			"type": "default"
		}
	]
}
```