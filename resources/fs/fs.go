package fs

import (
	"context"
	"fmt"
	"net/url"
	"os"
	"strings"

	"github.com/ziflex/waitfor/v2"
)

const Scheme = "file"

type File struct {
	path string
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

	path := strings.TrimPrefix(u.String(), "file://")

	return &File{path}, nil
}

func (f *File) Test(ctx context.Context) error {
	if ctx.Err() != nil {
		return ctx.Err()
	}

	_, err := os.Stat(f.path)

	return err
}
