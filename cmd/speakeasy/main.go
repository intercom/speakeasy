package main

import (
	"fmt"
	"github.com/urfave/cli"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"time"
)

const speakeasyPackage = "github.com/intercom/speakeasy"

func main() {
	app := cli.NewApp()
  app.Name = "Speakeasy"
  app.Usage = "Build embeddable servers for mobile apps"
  app.Version = "0.0.1"
  app.UsageText = "speakeasy build {package-name}"

  flags := []cli.Flag {
    cli.StringFlag{
      Name:        "name, n",
      Value:       "speakeasy",
      Usage:       "Name for the output artifacts",
    },
  }
	app.Commands = []cli.Command{
		{
			Name:  "build",
			Usage: "Builds the Android and iOS output for a given Go package",
			Flags: flags,
			Action: func(c *cli.Context) error {
				go tickDots()
				packageName := c.Args().First()
				buildAndroidFramework(packageName, c.String("name"))
				buildAppleFramework(packageName, c.String("name"))
				fmt.Println("\nDone!")
				return nil
			},
		},
		{
			Name:  "android",
			Usage: "Builds the Android output for a given Go package",
			Flags: flags,
			Action: func(c *cli.Context) error {
				go tickDots()
				packageName := c.Args().First()
				buildAndroidFramework(packageName, c.String("name"))
				fmt.Println("\nDone!")
				return nil
			},
		},
		{
			Name:  "ios",
			Usage: "Builds the iOS output for a given Go package",
			Flags: flags,
			Action: func(c *cli.Context) error {
				go tickDots()
				packageName := c.Args().First()
				buildAppleFramework(packageName, c.String("name"))
				fmt.Println("\nDone!")
				return nil
			},
		},
		{
			Name:  "clean",
			Usage: "Deletes the output directory",
			Action: func(c *cli.Context) error {
				panicIfError(os.RemoveAll("output"))
				return nil
			},
		},
	}

	app.Run(os.Args)
}

func panicIfError(e interface{}) {
	if e != nil {
		panic(e)
	}
}

func buildAndroidFramework(packageName string, outputName string) {
	panicIfError(os.MkdirAll("output/android", os.ModePerm))

	fmt.Println("\nBuilding Android .aar 🤖")
	androidCmd := exec.Command("gomobile",
		"bind",
		"-target=android/arm,android/386",
		`-ldflags=-s`,
		"-o", fmt.Sprintf("output/android/%v.aar", outputName),
		packageName,
		speakeasyPackage)
	androidCmd.Stdout = os.Stdout
	androidCmd.Stderr = os.Stderr
	androidCmd.Run()
}

func buildAppleFramework(packageName string, outputName string) {
	panicIfError(os.MkdirAll("output/ios", os.ModePerm))

  capitalisedName := strings.Title(outputName)
	if runtime.GOOS == "darwin" {
		fmt.Println("\nBuilding iOS framework ")
		iosCmd := exec.Command("gomobile",
			"bind",
			"-target=ios",
			`-ldflags=-s`,
			"-o", fmt.Sprintf("output/ios/%v.framework", capitalisedName),
			packageName,
			speakeasyPackage)
		iosCmd.Stdout = os.Stdout
		iosCmd.Stderr = os.Stderr
		iosCmd.Run()
	} else {
		fmt.Println("\nNot building iOS framework as we're not on macOS 🐧")
	}
}

func tickDots() {
	tick := time.Tick(1000 * time.Millisecond)
	for {
		select {
		case <-tick:
			fmt.Print(".")
		}
	}
}
