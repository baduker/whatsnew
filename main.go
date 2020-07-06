package main

import (
	"flag"
	"fmt"
	"github.com/microcosm-cc/bluemonday"
	"github.com/mmcdole/gofeed"
	"unicode"
	"unicode/utf8"
)

const whatsNewRSS = "https://aws.amazon.com/about-aws/whats-new/recent/feed/"

func wordWrap(text string, lineWidth int) string {
	wrap := make([]byte, 0, len(text)+2*len(text)/lineWidth)
	eoLine := lineWidth
	inWord := false
	for i, j := 0, 0; ; {
		r, size := utf8.DecodeRuneInString(text[i:])
		if size == 0 && r == utf8.RuneError {
			r = ' '
		}
		if unicode.IsSpace(r) {
			if inWord {
				if i >= eoLine {
					wrap = append(wrap, '\n')
					eoLine = len(wrap) + lineWidth
				} else if len(wrap) > 0 {
					wrap = append(wrap, ' ')
				}
				wrap = append(wrap, text[j:i]...)
			}
			inWord = false
		} else if !inWord {
			inWord = true
			j = i
		}
		if size == 0 && r == ' ' {
			break
		}
		i += size
	}
	return string(wrap)
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
	last := len(feeds) - 1
	for index := range feeds {
		item := feeds[last-index]

		title := wordWrap(item.Title, lineLength)
		date := item.Published[:16]
		description := wordWrap(removeHTMLTags(item.Description), lineLength)

		fmt.Printf("-> %s\nPosted On: %s\n", title, date)
		fmt.Printf("%s\n\n%s\n\n", item.Link, description)
	}
}

func main() {
	count := flag.Int("c", 5, "number of feeds to show; max 100")
	wrap := flag.Int("w", 120, "word wrapping line width")
	flag.Parse()
	showFeeds(getFeeds(*count), *wrap)
}
