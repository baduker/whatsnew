package main

import (
	"flag"
	"fmt"
	"github.com/microcosm-cc/bluemonday"
	"github.com/mmcdole/gofeed"
	"strings"
)

const whatsNewRSS = "https://aws.amazon.com/about-aws/whats-new/recent/feed/"

func wordWrap(text string, lineWidth int) (wrapped string) {
	words := strings.Fields(strings.TrimSpace(text))
	if len(words) == 0 {
		return text
	}
	wrapped = words[0]
	spaceLeft := lineWidth - len(wrapped)
	for _, word := range words[1:] {
		if len(word)+1 > spaceLeft {
			wrapped += "\n" + word
			spaceLeft = lineWidth - len(word)
		} else {
			wrapped += " " + word
			spaceLeft -= 1 + len(word)
		}
	}
	return
}

func getFeeds(count int) []*gofeed.Item {
	fp := gofeed.NewParser()
	feed, _ := fp.ParseURL(whatsNewRSS)
	return feed.Items[:count]
}

func removeHTMLTags(item string) string {
	p := bluemonday.StripTagsPolicy()
	return p.Sanitize(item)
}

func showFeeds(feeds []*gofeed.Item, lineLength int) {
	for index, item := range feeds {
		fmt.Printf("%d. %s | %s\n", index + 1, item.Title, item.Published)
		description := wordWrap(removeHTMLTags(item.Description), lineLength)
		fmt.Printf("%s\n\n%s\n\n", item.Link, description)
	}
}

func main() {
	count := flag.Int("c", 5, "number of feeds to show; max 100")
	wrap := flag.Int("w", 120, "word wrapping line width")
	flag.Parse()
	showFeeds(getFeeds(*count), *wrap)
}
