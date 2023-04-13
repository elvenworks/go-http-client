package http_client

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"go.elastic.co/apm/module/apmhttp/v2"
)

type Options struct {
	Method  string
	Path    string
	Body    interface{}
	Headers map[string]string
}

type Response struct {
	Body       []byte
	StatusCode int
}

type HttpClient struct {
	BaseUri string
}

func Init(baseUri string) IHttpClient {
	return &HttpClient{
		BaseUri: baseUri,
	}
}

func (h *HttpClient) Request(ctx context.Context, options *Options) (*Response, error) {
	url := h.getUrl(options)

	body, err := h.getBody(options)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, options.Method, url, body)
	if err != nil {
		return nil, err
	}

	for key, value := range options.Headers {
		req.Header.Set(key, value)
	}

	client := &http.Client{Timeout: 10 * time.Second}
	apmClient := apmhttp.WrapClient(client)
	response, err := apmClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer h.close(response.Body)

	bodyResponse, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	return &Response{
		Body:       bodyResponse,
		StatusCode: response.StatusCode,
	}, nil
}

func (h *HttpClient) getBody(options *Options) (io.Reader, error) {
	var body io.Reader

	if options.Body != nil {
		data, err := json.Marshal(options.Body)
		if err != nil {
			return nil, err
		}

		body = bytes.NewBuffer(data)
	}

	return body, nil
}

func (h *HttpClient) getUrl(options *Options) string {
	baseUri := strings.TrimSuffix(h.BaseUri, "/")
	path := strings.TrimPrefix(options.Path, "/")

	return fmt.Sprintf("%s/%s", baseUri, path)
}

func (h *HttpClient) close(body io.ReadCloser) io.ReadCloser {
	return struct {
		io.Reader
		io.Closer
	}{
		Reader: body,
		Closer: body,
	}
}
