package cli

import (
	"os"

	"github.com/urfave/cli"
)

// Info defines the basic information required for the CLI.
type Info struct {
	Name        string
	Version     string
	Description string
	AuthorName  string
	AuthorEmail string
}

// Initialize and bootstrap the CLI.
func Initialize(info *Info) error {
	var secretName string
	var env string

	app := cli.NewApp()
	app.Name = info.Name
	app.Version = info.Version
	app.Usage = info.Description
	app.Authors = []cli.Author{
		cli.Author{
			Name:  info.AuthorName,
			Email: info.AuthorEmail,
		},
	}

	flags := []cli.Flag{
		cli.StringFlag{
			Name:        "secret, s",
			Usage:       "Secret's Name to fetch environment from",
			Destination: &secretName,
		},
		cli.StringFlag{
			Name:        "env, e",
			Usage:       "Environment to use the secret name from",
			Destination: &env,
		},
	}

	app.Commands = []cli.Command{
		cli.Command{
			Name:  "setup",
			Usage: "Setup envault configuration",
			Action: func(ctx *cli.Context) error {
				Setup()

				return nil
			},
		},
		cli.Command{
			Name:  "list",
			Usage: "List environment variables stored in Secrets Manager",
			Flags: flags,
			Action: func(ctx *cli.Context) error {
				List(secretName, env)

				return nil
			},
		},
		cli.Command{
			Name:      "run",
			Usage:     "Run a command with the injected env variables",
			ArgsUsage: "[command]",
			Flags:     flags,
			Action: func(ctx *cli.Context) error {
				Run(secretName, ctx.Args().Get(0), env)

				return nil
			},
		},
	}

	return app.Run(os.Args)
}
