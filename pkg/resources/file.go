package resources

import (
	"context"
	"os"
)

type File struct {
	path string
}

func NewFile(path string) Resource {
	return &File{path}
}

func (f *File) Test(ctx context.Context) error {
	if ctx.Err() != nil {
		return ctx.Err()
	}

	_, err := os.Stat(f.path)

	return err
}
