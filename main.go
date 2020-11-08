package main

import (
	"flag"
	"fmt"

	"github.com/eyedeekay/amo-version/lib"
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
	end = "/addon/"
)

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
		amo.DownloadFile(*name + ".xpi")
	}
}
