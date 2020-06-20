# whatsnew
a dead simple AWS What's New RSS feed grabber

A super simple RSS feed fetcher for 
[AWS What's new](https://aws.amazon.com/new/?nc1=f_cc&whats-new-content-all.sort-by=item.additionalFields.postDateTime&whats-new-content-all.sort-order=desc) written in Go.

### Usage

1. Clone the repo
2. Build the file `go build -o whatsnew main.go`
3. Run `./whatsnew -c 5 -w 80`
