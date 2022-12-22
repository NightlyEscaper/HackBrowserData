package main
import "C"
import (
	"os"
	"strings"
	"path/filepath"

	"hack-browser-data/internal/log"
	"hack-browser-data/internal/provider"
	"hack-browser-data/internal/utils/fileutil"

	"github.com/urfave/cli/v2"
)

var (
	browserName  string
	outputDir    string
	outputFormat string
	verbose      bool
	compress     bool
	profilePath  string
)

func main(){
	run()
}

//export run
func run() {
	app := &cli.App{
		Name:      "bc",
		Usage:     "",
		UsageText: "[]\n",
		Version:   "0.4.4",
		Flags: []cli.Flag{
			&cli.BoolFlag{Name: "verbose", Aliases: []string{"vv"}, Destination: &verbose, Value: false, Usage: "verbose"},
			&cli.BoolFlag{Name: "compress", Aliases: []string{"zip"}, Destination: &compress, Value: true, Usage: "compress result to zip"},
			&cli.StringFlag{Name: "browser", Aliases: []string{"b"}, Destination: &browserName, Value: "all", Usage: "available browsers: all|" + strings.Join(provider.ListBrowsers(), "|")},
			&cli.StringFlag{Name: "results-dir", Aliases: []string{"dir"}, Destination: &outputDir, Value: "results", Usage: "export dir"},
			&cli.StringFlag{Name: "format", Aliases: []string{"f"}, Destination: &outputFormat, Value: "csv", Usage: "file name csv|json"},
			&cli.StringFlag{Name: "profile-path", Aliases: []string{"p"}, Destination: &profilePath, Value: "", Usage: "custom profile dir path, get with chrome://version"},
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
