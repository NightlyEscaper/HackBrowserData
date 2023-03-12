package main

import (
	"os"

	"github.com/urfave/cli/v2"

	"github.com/moond4rk/HackBrowserData/browser"
	"github.com/moond4rk/HackBrowserData/log"
	"github.com/moond4rk/HackBrowserData/utils/fileutil"
)

var (
	browserName  string
	outputDir    string
	outputFormat string
	verbose      bool
	compress     bool
	profilePath  string
	isFullExport bool
)

func main() {
	Execute()
}

func Execute() {
	app := &cli.App{
		Name:      "bc",
		Usage:     "",
		UsageText: "[]\n",
		Version:   "0.4.4",
		Flags: []cli.Flag{
			&cli.BoolFlag{Name: "verbose", Aliases: []string{"vv"}, Destination: &verbose, Value: false, Usage: "verbose"},
			&cli.BoolFlag{Name: "compress", Aliases: []string{"zip"}, Destination: &compress, Value: false, Usage: "compress result to zip"},
			&cli.StringFlag{Name: "browser", Aliases: []string{"b"}, Destination: &browserName, Value: "all", Usage: "available browsers: all|" + browser.Names()},
			&cli.StringFlag{Name: "results-dir", Aliases: []string{"dir"}, Destination: &outputDir, Value: "results", Usage: "export dir"},
			&cli.StringFlag{Name: "format", Aliases: []string{"f"}, Destination: &outputFormat, Value: "csv", Usage: "file name csv|json"},
			&cli.StringFlag{Name: "profile-path", Aliases: []string{"p"}, Destination: &profilePath, Value: "", Usage: "custom profile dir path, get with chrome://version"},
			&cli.BoolFlag{Name: "full-export", Aliases: []string{"full"}, Destination: &isFullExport, Value: true, Usage: "is export full browsing data"},
		},
		HideHelpCommand: true,
		Action: func(c *cli.Context) error {
			if verbose {
				log.Init("debug")
			} else {
				log.Init("notice")
			}

			browsers, err := provider.PickBrowsers(browserName, profilePath)
			if err != nil {
				log.Error(err)
			}

			_dir := filepath.Join(".", outputDir)
			err = os.MkdirAll(_dir, os.ModePerm)
			
			if err != nil {
				log.Error(err)
			}

			for _, b := range browsers {
				data, err := b.BrowsingData()
				if err != nil {
					log.Error(err)
				}
				data.Output(outputDir, b.Name(), outputFormat)
			}
			if compress {
				if err = fileutil.CompressDir(outputDir); err != nil {
					log.Error(err)
				}
				log.Noticef("compress success")
			}
			return nil
		},
	}
	err := app.Run(os.Args)
	if err != nil {
		panic(err)
	}
}
