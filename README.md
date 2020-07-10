# whatsnew


![](https://d1.awsstatic.com/about-aws/whats-new/lpm-assets/whats-new-page-assets/Site-Merch_Amazon-New_2up.4a9e97631cc47c96d5e9fa9014fb3640b7ae2976.png)


A super simple RSS feed fetcher for 
[AWS What's New](https://aws.amazon.com/new/?nc1=f_cc&whats-new-content-all.sort-by=item.additionalFields
.postDateTime&whats-new-content-all.sort-order=desc) 
written in Go.

By default, it shows the five most recent updates (out of a hundred).

I plan to support downloading all published news.

### Usage

0. Get [Go](https://golang.org/dl/) and [install](https://golang.org/doc/install) it.
1. Clone the repo.
2. Build the file `go build -o whatsnew main.go`
3. Run `./whatsnew | less`

```
Usage of ./whatsnew:
  -c int
        number of feeds to show; max 100 (default 5)
  -w int
        word wrapping line width (default 120)

```

For example, you can dump the entire  feed to a text file with:
`./whatsnew -c 100 > rss_dump.txt`

### TODO:

- Fetch and parse old news
- Make use of API request and JSON
- Add dump command (no need for piping)
