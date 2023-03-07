module github.com/eyedeekay/amo-version

go 1.19

require github.com/PuerkitoBio/goquery v1.8.1

require (
	github.com/andybalholm/cascadia v1.3.1 // indirect
	github.com/artdarek/go-unzip v1.0.1-0.20210323073738-f9883ad8bd15
	golang.org/x/net v0.8.0 // indirect
)

replace github.com/artdarek/go-unzip => github.com/eyedeekay/go-unzip v0.0.0-20220914222511-f2936bba53c2
