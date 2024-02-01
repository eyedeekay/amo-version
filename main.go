package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"

	amo "github.com/eyedeekay/amo-version/lib"
)

// input: https://addons.mozilla.org/en-US/firefox/addon/i2p-in-private-browsing/
// output: 0.82

var (
	begin = flag.String(
		"amo-url",
		"https://addons.mozilla.org/",
		"URL to check for extensions at",
	)
	lang = flag.String(
		"c",
		"en-US",
		"Language to target for the extension.",
	)
	platform = flag.String(
		"p",
		"firefox",
		"Platform to download the extension for. firefox or android.",
	)
	name = flag.String(
		"n",
		"i2p-in-private-browsing",
		"name of the extension to look up, if you don't know the ID. Must match the AMO page.",
	)
	download = flag.Bool(
		"d",
		false,
		"Download the extension.xpi file",
	)
	version = flag.Bool(
		"v",
		false,
		"display version number only",
	)
	id = flag.Bool(
		"i",
		false,
		"print the extension ID",
	)
	outpath = flag.String(
		"o",
		defaultPath(),
		"Output the file to a specific path",
	)
	outdir = flag.String(
		"od",
		defaultDir(),
		"Output the file to a specific directory",
	)
	extract = flag.Bool(
		"x",
		false,
		"extract the file after downloading it",
	)
)

func defaultDir() string {
	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	return wd
}

func defaultPath() string {
	return filepath.Join(*outdir, *name+".xpi")
}

func main() {
	flag.Parse()
	amo.Begin = *begin
	amo.Lang = *lang
	amo.Platform = *platform
	amo.Name = *name
	if !*version {
		fmt.Println(*name)
	}
	fmt.Println(amo.VersionString())
	if !*version {
		fmt.Println(amo.DownloadURL())
	}
	if *id {
		fmt.Println(amo.ExtID())
	}
	if *download {
		os.MkdirAll(*outdir, 0755)
		amo.DownloadFile(*outpath)
		if *extract {
			amo.Unzip(*outpath)
		}
	}
}
