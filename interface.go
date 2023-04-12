package http_client

import "context"

type IHttpClient interface {
	Request(ctx context.Context, options *Options) (*Response, error)
}
