package main

import (
	"fmt"
	"os"

	"github.com/codegangsta/cli"
)

func main() {
	app := cli.NewApp()

	app.Action = func(c *cli.Context) error {
		fmt.Printf("Hello %q\n", c.Args().Get(0))
		return nil
	}

	err := app.Run(os.Args)

	if err != nil {
		fmt.Println(err)
		os.Exit(1001)
	}
}
