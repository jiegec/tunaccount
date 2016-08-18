package main

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"gopkg.in/urfave/cli.v1"
)

var (
	buildstamp = ""
	githash    = "no githash provided"
)

func main() {

	app := cli.NewApp()
	app.Name = "tunaccount"
	app.EnableBashCompletion = true
	cli.VersionPrinter = func(c *cli.Context) {
		var builddate string
		if buildstamp == "" {
			builddate = "No build data provided"
		} else {
			ts, err := strconv.Atoi(buildstamp)
			if err != nil {
				builddate = "No build data provided"
			} else {
				t := time.Unix(int64(ts), 0)
				builddate = t.String()
			}
		}
		fmt.Printf(
			"Version: %s\n"+
				"Git Hash: %s\n"+
				"Build Data: %s\n",
			c.App.Version, githash, builddate,
		)
	}

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:   "config, c",
			Value:  "/etc/tunaccountd.conf",
			Usage:  "specify configuration file",
			EnvVar: "TUNACCOUNT_CONFIG_FILE",
		},
	}

	app.Commands = []cli.Command{
		{
			Name:   "daemon",
			Usage:  "run tunaccount daemon",
			Action: startDaemon,
			Flags: []cli.Flag{
				cli.BoolFlag{
					Name:  "debug",
					Usage: "enable debug",
				},
			},
		},
		{
			Name:      "import",
			Usage:     "import json files to tunaccount",
			Action:    importFiles,
			ArgsUsage: "[files...]",
		},
		{
			Name:      "passwd",
			Usage:     "set password of a user, default is current user",
			Action:    cmdPasswd,
			ArgsUsage: "[user]",
		},
		{
			Name:      "useradd",
			Usage:     "add a user",
			Action:    cmdUseradd,
			ArgsUsage: "<username>",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "shell, s",
					Usage: "Login shell of the new account",
					Value: "/bin/bash",
				},
				cli.StringFlag{
					Name:  "name",
					Usage: "Fullname of the new account (Required)",
				},
				cli.StringFlag{
					Name:  "email, mail",
					Usage: "Email address of the new account (Required)",
				},
				cli.StringFlag{
					Name:  "phone, mobile",
					Usage: "Phone number of the new account",
				},
			},
		},
		{
			Name:  "tag",
			Usage: "tag management",
			Subcommands: []cli.Command{
				{
					Name:  "new",
					Usage: "add new tag",
				},
				{
					Name:      "user",
					Usage:     "tag a user",
					ArgsUsage: "<tag>",
				},
				{
					Name:      "group",
					Usage:     "tag a group",
					ArgsUsage: "<tag>",
				},
			},
		},
	}

	app.Run(os.Args)
}
