package resources

import (
	"context"
	"net/http"
)

type HTTP struct {
	url string
}

func NewHTTP(url string) Resource {
	return &HTTP{url}
}

func (h *HTTP) Test(ctx context.Context) error {
	req, err := http.NewRequest(http.MethodGet, h.url, nil)

	if err != nil {
		return err
	}

	client := http.Client{}
	_, err = client.Do(req.WithContext(ctx))

	return err
}
