package http

import (
	"context"
	"fmt"
	"net/http"
	"net/url"

	"github.com/ziflex/waitfor/v2"
)

const Scheme = "http"

type HTTP struct {
	url *url.URL
}

func Use() waitfor.ResourceConfig {
	return waitfor.ResourceConfig{
		Scheme:  Scheme,
		Factory: New,
	}
}

func New(u *url.URL) (waitfor.Resource, error) {
	if u == nil {
		return nil, fmt.Errorf("%q: %w", "url", waitfor.ErrInvalidArgument)
	}

	return &HTTP{u}, nil
}

func (h *HTTP) Test(ctx context.Context) error {
	req, err := http.NewRequest(http.MethodGet, h.url.String(), nil)

	if err != nil {
		return err
	}

	client := http.Client{}
	_, err = client.Do(req.WithContext(ctx))

	return err
}
