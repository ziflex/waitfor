package http_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/ziflex/waitfor/v2/resources/http"
)

func TestHTTP(t *testing.T) {
	r := http.New("https://www.github.com")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err := r.Test(ctx)

	if err != nil {
		t.Error(err)
	}
}

func TestHTTP2(t *testing.T) {
	r := http.New("https://localhost:1000")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := r.Test(ctx)

	if err == nil {
		t.Error(errors.New("should fail"))
	}
}
