
Package go-http-client implements the possibility to provides to make HTTP requests in a better way, making it possible to use mocks and interfaces.

## Installation
Use go get.
```
go get github.com/elvenworks/go-http-client
```
Then import the validator package into your own code.
```
import "github.com/elvenworks/go-http-client"
```

## Usage
Sample code:
```go
package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	httpClient "github.com/elvenworks/go-http-client"
)

func  main() {
	client := httpClient.Init("https://myapp.com")
	options := &httpClient.Options{
		Path: "/api/v1/users",
		Method: http.MethodPost,
		Body: []byte(`{"user": "Loren Ipsun"}`),
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
	}
	
	response, err := client.Request(context.Background(), options)
	if err != nil {
		log.Fatal(err)
	}

	// {"id": 1, "user": "Loren Ipsun"}
	// 201 created
	fmt.Println(string(response.Body), response.StatusCode)
}
```

## Options:
| Parameter | Description | Value |
| :-------------: |:--------:| :-------------: |
| Method | The HTTP method to be used in the request | A string that can be GET, POST, PUT, PATCH, DELETE, etc. |
| Path | The API endpoint path for the request | A string that contains the endpoint path, for example, /users/123. |
| Body | The data in the request body | It should be a JSON object that contains the data to be sent in the request body. The JSON object can contain any type of data, such as strings, numbers, arrays, or other JSON objects. |
| Headers | The additional HTTP headers for the request | A string-key and value map, where each key is the header name and each value is the header value. For example, {"Authorization": "Bearer xyz123"}. |

## Response:

| Parameter | Description | Value |
| :-------------: |:--------:| :-------------: |
| Body | The response body data | A byte array that contains the response body data. |
| StatusCode | The HTTP status code of the response | An integer that represents the HTTP status code returned by the server. For example, 200 for a successful request or 404 for a resource not found. |
