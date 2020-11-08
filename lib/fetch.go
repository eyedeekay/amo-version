package amo

import (
	"github.com/PuerkitoBio/goquery"

	"io"
	"log"
	"net/http"
	"os"
)

var (
	Begin    = "https://addons.mozilla.org/"
	Lang     = "en-US"
	Platform = "firefox"
	Name     = "i2p-in-private-browsing"
)

func URL() string {
	return Begin + Lang + "/" + Platform + "/addon/" + Name
}

func VersionString() string {
	doc := PageBody()
	sel := doc.Find(".AddonMoreInfo-version")
	return sel.Text()
}

func LookupVersion(name string) string {
	old := Name
	Name = name
	val := VersionString()
	Name = old
	return val
}

func DownloadURL() string {
	doc := PageBody()
	sel := doc.Find(".InstallButtonWrapper-download-link")
	val, _ := sel.Attr("href")
	//val, _ := doc.Html()
	//val, _ := sel.Html()
	return val
}

func LookupDownload(name string) string {
	old := Name
	Name = name
	val := DownloadURL()
	Name = old
	return val
}

func DownloadFile(filepath string) error {

	// Get the data
	resp, err := http.Get(DownloadURL())
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	return err
}

func PageBody() *goquery.Document {
	res, err := http.Get(URL())
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	return doc
}
