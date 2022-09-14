# Firefox Addon Downloader for Go(amo-version)

Used to download Firefox extensions from the Mozilla Addons store.

 - **Status:** Maintained.

## Usage of amo-version:

```sh
  -amo-url string
    	URL to check for extensions at (default "https://addons.mozilla.org/")
  -c string
    	Language to target for the extension. (default "en-US")
  -d	Download the extension.xpi file
  -i	print the extension ID
  -n string
    	name of the extension to look up, if you don't know the ID. Must match the AMO page. (default "i2p-in-private-browsing")
  -o string
    	Output the file to a specific path (default "$PWD/i2p-in-private-browsing.xpi")
  -od string
    	Output the file to a specific directory (default "$PWD")
  -p string
    	Platform to download the extension for. firefox or android. (default "firefox")
  -v	display version number only
  -x	extract the file after downloading it
```