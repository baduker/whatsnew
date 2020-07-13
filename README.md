# whatsnew


![](https://d1.awsstatic.com/about-aws/whats-new/lpm-assets/whats-new-page-assets/Site-Merch_Amazon-New_2up.4a9e97631cc47c96d5e9fa9014fb3640b7ae2976.png)


A super simple News fetcher for 
[AWS What's New](https://aws.amazon.com/new/?nc1=f_cc&whats-new-content-all.sort-by=item.additionalFields.postDateTime&whats-new-content-all.sort-order=desc) 
written in Go. It uses the aws.amazon.com API.

By default, it shows the 25 most recent updates, that basically reflects the
**What's New Feed** at the bottom of the [page](https://aws.amazon.com/new/?nc1=f_cc&whats-new-content-all.sort-by=item.additionalFields.postDateTime&whats-new-content-all.sort-order=desc)

I plan to support downloading all published news and (possibly) searching
the news items by a product or service.

### Usage

0. Get [Go](https://golang.org/dl/) and [install](https://golang.org/doc/install) it.
1. Clone the repo.
2. Build the file `go build -o whatsnew main.go`
3. Run `./whatsnew | less`

```
Usage of ./whatsnew:
  -c string
        number of feeds to show (default 25)
  -w int
        word wrapping line width (default 120)

```

For example, you can dump the output to a text file with:
`./whatsnew > aws.txt`

### TODO:

- Fetch and parse old news
- Add dump command (no need for piping)