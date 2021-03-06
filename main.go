package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli"
)

var build = "0" // build number set at compile time

func main() {
	app := cli.NewApp()
	app.Name = "Drone-SonarQube-Node-Plugin"
	app.Usage = "Drone plugin for SonarQube analysis of node projects."
	app.Action = run
	app.Version = fmt.Sprintf("1.0.%s", build)
	app.Flags = []cli.Flag{

		cli.StringFlag{
			Name:   "key",
			Usage:  "project key",
			EnvVar: "DRONE_REPO",
		},
		cli.StringFlag{
			Name:   "name",
			Usage:  "project name",
			EnvVar: "DRONE_REPO",
		},
		cli.StringFlag{
			Name:   "host",
			Usage:  "SonarQube host",
			EnvVar: "SONAR_HOST",
		},
		cli.StringFlag{
			Name:   "token",
			Usage:  "SonarQube token",
			EnvVar: "SONAR_TOKEN",
		},
		cli.StringFlag{
			Name:   "organization",
			Usage:  "SonarQube organization",
			EnvVar: "SONAR_ORGANIZATION",
		},

		// advanced parameters
		cli.StringFlag{
			Name:   "ver",
			Usage:  "Project version",
			EnvVar: "DRONE_BUILD_NUMBER",
		},
		cli.StringFlag{
			Name:   "timeout",
			Usage:  "Web request timeout",
			Value:  "60",
			EnvVar: "PLUGIN_TIMEOUT",
		},
		cli.StringFlag{
			Name:   "sources",
			Usage:  "analysis sources",
			Value:  ".",
			EnvVar: "PLUGIN_SOURCES",
		},
		cli.StringFlag{
			Name:   "inclusions",
			Usage:  "code inclusions",
			EnvVar: "PLUGIN_INCLUSIONS",
		},
		cli.StringFlag{
			Name:   "exclusions",
			Usage:  "code exclusions",
			EnvVar: "PLUGIN_EXCLUSIONS",
		},
		cli.StringFlag{
			Name:   "level",
			Usage:  "log level",
			Value:  "INFO",
			EnvVar: "PLUGIN_LEVEL",
		},
		cli.StringFlag{
			Name:   "showProfiling",
			Usage:  "showProfiling during analysis",
			Value:  "false",
			EnvVar: "PLUGIN_SHOWPROFILING",
		},
	}

	app.Run(os.Args)
}

func run(c *cli.Context) {
	plugin := Plugin{
		Config: Config{
			Key:          c.String("key"),
			Name:         c.String("name"),
			Host:         c.String("host"),
			Token:        c.String("token"),
			Organization: c.String("organization"),

			Version:       c.String("ver"),
			Timeout:       c.String("timeout"),
			Sources:       c.String("sources"),
			Inclusions:    c.String("inclusions"),
			Exclusions:    c.String("exclusions"),
			Level:         c.String("level"),
			showProfiling: c.String("showProfiling"),
		},
	}

	if err := plugin.Exec(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
