package main

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"github.com/ziflex/waitfor/pkg/runner"
	"os"
)

var version string

func main() {
	app := &cli.App{
		Name:        "waitfor",
		Usage:       "Tests and waits on the availability of a remote resource",
		Description: "Tests and waits on the availability of a remote resource before executing a command with exponential backoff",
		Version:     version,
		Flags: []cli.Flag{
			&cli.StringSliceFlag{
				Name:     "resource",
				Aliases:  []string{"d"},
				Usage:    "-r http://localhost:8080",
				EnvVars:  []string{"WAITFOR_RESOURCE"},
				Required: true,
			},
			&cli.Uint64Flag{
				Name:    "attempts",
				Aliases: []string{"a"},
				Usage:   "amount of attempts",
				EnvVars: []string{"WAITFOR_ATTEMPTS"},
				Value:   5,
			},
			&cli.Uint64Flag{
				Name:    "interval",
				Usage:   "interval between attempts (sec)",
				EnvVars: []string{"WAITFOR_INTERVAL"},
				Value:   5,
			},
			&cli.Uint64Flag{
				Name:    "max-interval",
				Usage:   "maximum interval between attempts (sec)",
				EnvVars: []string{"WAITFOR_MAX_INTERVAL"},
				Value:   60,
			},
		},
		Action: func(ctx *cli.Context) error {
			if ctx.NArg() == 0 {
				return cli.NewExitError("executable is required", 1)
			}

			program := runner.Program{
				Executable: "",
				Args:       nil,
				Resources:  ctx.StringSlice("resource"),
			}

			args := ctx.Args().Slice()

			program.Executable = args[0]

			if len(args) > 1 {
				program.Args = args[1:]
			}

			out, err := runner.Run(
				ctx.Context,
				program,
				runner.WithAttempts(ctx.Uint64("attempts")),
				runner.WithInterval(ctx.Uint64("interval")),
				runner.WithMaxInterval(ctx.Uint64("max-interval")),
			)

			if out != nil {
				fmt.Println(string(out))
			}

			return err
		},
	}

	err := app.Run(os.Args)

	if err != nil {
		fmt.Println(err)

		os.Exit(1)
	}
}
