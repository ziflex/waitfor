package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli/v2"

	"github.com/ziflex/waitfor/v2"
	"github.com/ziflex/waitfor/v2/resources/fs"
	"github.com/ziflex/waitfor/v2/resources/http"
	"github.com/ziflex/waitfor/v2/resources/mongodb"
	"github.com/ziflex/waitfor/v2/resources/mysql"
	"github.com/ziflex/waitfor/v2/resources/postgres"
	"github.com/ziflex/waitfor/v2/resources/proc"
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
				Aliases:  []string{"r"},
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

			program := waitfor.Program{
				Executable: "",
				Args:       nil,
				Resources:  ctx.StringSlice("resource"),
			}

			args := ctx.Args().Slice()

			program.Executable = args[0]

			if len(args) > 1 {
				program.Args = args[1:]
			}

			runner := waitfor.New(
				fs.Use(),
				http.Use(),
				mongodb.Use(),
				proc.Use(),
				mysql.Use(),
				postgres.Use(),
			)

			out, err := runner.Run(
				ctx.Context,
				program,
				waitfor.WithAttempts(ctx.Uint64("attempts")),
				waitfor.WithInterval(ctx.Uint64("interval")),
				waitfor.WithMaxInterval(ctx.Uint64("max-interval")),
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
