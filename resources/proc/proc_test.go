package proc_test

import (
	"context"
	"errors"
	"runtime"
	"testing"
	"time"

	"github.com/ziflex/waitfor/v2/resources/proc"
)

func TestProcess(t *testing.T) {
	var processName string

	if runtime.GOOS == "darwin" {
		processName = "launchd"
	} else if runtime.GOOS == "windows" {
		processName = "svchost"
	} else {
		processName = "systemd"
	}

	r := proc.New(processName)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err := r.Test(ctx)

	if err != nil {
		t.Error(err)
	}
}

func TestProcess2(t *testing.T) {
	r := proc.New("fsfsdfds")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err := r.Test(ctx)

	if err == nil {
		t.Error(errors.New("should fail"))
	}
}
