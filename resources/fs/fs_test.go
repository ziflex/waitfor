package fs_test

import (
	"context"
	"errors"
	"fmt"
	"net/url"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/ziflex/waitfor/v2/resources/fs"
)

func TestFile(t *testing.T) {
	name := fmt.Sprintf("waitfor_%s.txt", time.Now())
	fileName := filepath.Join(os.TempDir(), name)

	file, err := os.Create(fileName)

	if err != nil {
		t.Error(err)
	}

	defer file.Close()
	defer os.Remove(fileName)

	u, err := url.Parse("file://" + fileName)

	if err != nil {
		t.Error(err)
	}

	r, err := fs.New(u)

	if err != nil {
		t.Error(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = r.Test(ctx)

	if err != nil {
		t.Error(err)
	}
}

func TestFile_FileNotExists(t *testing.T) {
	u, err := url.Parse("file://" + filepath.Join(os.TempDir(), "fdsfsdfds"))

	if err != nil {
		t.Error(err)
	}

	r, err := fs.New(u)

	if err != nil {
		t.Error(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = r.Test(ctx)

	if err == nil {
		t.Error(errors.New("should fail"))
	}
}

func TestFile_MissedURL(t *testing.T) {
	_, err := fs.New(nil)

	if err == nil {
		t.Error(errors.New("should fail"))
	}
}
