package resources_test

import (
	"context"
	"errors"
	"fmt"
	"github.com/ziflex/waitfor/pkg/resources"
	"os"
	"path/filepath"
	"testing"
	"time"
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

	r := resources.NewFile(fileName)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = r.Test(ctx)

	if err != nil {
		t.Error(err)
	}
}

func TestFile2(t *testing.T) {
	r := resources.NewFile(filepath.Join(os.TempDir(), "fdsfsdfds"))
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err := r.Test(ctx)

	if err == nil {
		t.Error(errors.New("should fail"))
	}
}
