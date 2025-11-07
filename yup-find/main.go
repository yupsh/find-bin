package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli/v2"

	yup "github.com/gloo-foo/framework"
	. "github.com/yupsh/find"
)

const (
	flagName           = "name"
	flagType           = "type"
	flagSize           = "size"
	flagMaxDepth       = "maxdepth"
	flagFollowSymlinks = "follow"
)

func main() {
	app := &cli.App{
		Name:  "find",
		Usage: "search for files in a directory hierarchy",
		UsageText: `find [PATH...] [EXPRESSION]

   Search for files in a directory hierarchy.`,
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  flagName,
				Usage: "base of file name matches shell pattern PATTERN",
			},
			&cli.StringFlag{
				Name:  flagType,
				Usage: "file is of type TYPE (f=file, d=directory, l=link)",
			},
			&cli.StringFlag{
				Name:  flagSize,
				Usage: "file uses SIZE units of space",
			},
			&cli.IntFlag{
				Name:  flagMaxDepth,
				Usage: "descend at most LEVELS (a non-negative integer) levels",
			},
			&cli.BoolFlag{
				Name:    flagFollowSymlinks,
				Aliases: []string{"L"},
				Usage:   "follow symbolic links",
			},
		},
		Action: action,
	}

	if err := app.Run(os.Args); err != nil {
		fmt.Fprintf(os.Stderr, "find: %v\n", err)
		os.Exit(1)
	}
}

func action(c *cli.Context) error {
	var params []any

	// Add all arguments as paths
	for i := 0; i < c.NArg(); i++ {
		params = append(params, c.Args().Get(i))
	}

	// Add flags based on CLI options
	if c.IsSet(flagName) {
		params = append(params, Name(c.String(flagName)))
	}
	if c.IsSet(flagType) {
		// Map the type string to the fileType constant
		params = append(params, c.String(flagType))
	}
	if c.IsSet(flagSize) {
		params = append(params, Size(c.String(flagSize)))
	}
	if c.IsSet(flagMaxDepth) {
		params = append(params, MaxDepth(c.Int(flagMaxDepth)))
	}
	if c.Bool(flagFollowSymlinks) {
		params = append(params, FollowSymlinks)
	}

	// Create and execute the find command
	cmd := Find(params...)
	return yup.Run(cmd)
}
