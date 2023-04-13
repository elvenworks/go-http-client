package http_client

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestShouldSuccessResponse(t *testing.T) {
	options := &Options{
		Method: http.MethodPost,
		Body:   []byte(`{"name": "loren ipsun"}`),
	}

	want := &Response{
		Body:       []byte(`{"id": 1, "name": "loren ipsun"}`),
		StatusCode: http.StatusCreated,
	}

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(want.StatusCode)

		w.Write(want.Body)
	}))

	defer server.Close()

	client := Init(server.URL)

	response, err := client.Request(context.Background(), options)
	if (err != nil) != false {
		t.Errorf("error = %v, wantErr %v", err, false)

		return
	}

	if diff := cmp.Diff(response, want); diff != "" {
		t.Error(diff)
	}
}

func TestShouldMarshalError(t *testing.T) {
	options := &Options{
		Body: make(chan int, 0),
	}

	client := Init("")

	_, err := client.Request(context.Background(), options)
	if (err != nil) != true {
		t.Errorf("error = %v, wantErr %v", err, true)

		return
	}
}

func TestShouldNewRequestWithContextError(t *testing.T) {
	options := &Options{}

	client := Init("")

	_, err := client.Request(nil, options)
	if (err != nil) != true {
		t.Errorf("error = %v, wantErr %v", err, true)

		return
	}
}

func TestShouldDoError(t *testing.T) {
	options := &Options{}

	client := Init("http://invalid")

	_, err := client.Request(context.Background(), options)
	if (err != nil) != true {
		t.Errorf("error = %v, wantErr %v", err, true)

		return
	}
}

func TestShouldReadAllError(t *testing.T) {
	options := &Options{
		Method: http.MethodGet,
		Headers: map[string]string{
			"Content-Length": "1",
		},
	}

	want := &Response{
		Body: []byte(`{"id": 1, "name": "loren ipsun"}`),
	}

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		for key, value := range options.Headers {
			w.Header().Set(key, value)
		}

		w.Write(want.Body)
	}))

	defer server.Close()

	client := Init(server.URL)

	_, err := client.Request(context.Background(), options)
	if (err != nil) != true {
		t.Errorf("error = %v, wantErr %v", err, true)

		return
	}
}
