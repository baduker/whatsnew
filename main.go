package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/microcosm-cc/bluemonday"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
	"unicode"
	"unicode/utf8"
)

const (
	baseURL = "https://aws.amazon.com"
	path    = "/api/dirs/items/search?"
	query   = "item.directoryId=whats-new&item.locale=en_US&"
	sort    = "sort_by=item.additionalFields.postDateTime&sort_order=desc&"
)

type Data struct {
	Metadata struct {
		Count     int `json:"count"`
		TotalHits int `json:"totalHits"`
	} `json:"metadata"`
	FieldTypes struct {
		RelatedBlog  string `json:"relatedBlog"`
		PostBody     string `json:"postBody"`
		ModifiedDate string `json:"modifiedDate"`
		HeadlineURL  string `json:"headlineUrl"`
		PostDateTime string `json:"postDateTime"`
		PostSummary  string `json:"postSummary"`
		Headline     string `json:"headline"`
		ContentType  string `json:"contentType"`
	} `json:"fieldTypes"`
	// TODO: Use slice of pointers -> []*Items
	Items []struct {
		Tags []struct {
			TagNamespaceID string `json:"tagNamespaceId"`
			CreatedBy      string `json:"createdBy"`
			Name           string `json:"name"`
			DateUpdated    string `json:"dateUpdated"`
			Locale         string `json:"locale"`
			LastUpdatedBy  string `json:"lastUpdatedBy"`
			DateCreated    string `json:"dateCreated"`
			Description    string `json:"description"`
			ID             string `json:"id"`
		} `json:"tags"`
		Item struct {
			CreatedBy        string  `json:"createdBy"`
			Locale           string  `json:"locale"`
			Author           string  `json:"author"`
			DateUpdated      string  `json:"dateUpdated"`
			Score            float64 `json:"score"`
			Name             string  `json:"name"`
			NumImpressions   int     `json:"numImpressions"`
			DateCreated      string  `json:"dateCreated"`
			AdditionalFields struct {
				PostBody     string    `json:"postBody"`
				ModifiedDate time.Time `json:"modifiedDate"`
				HeadlineURL  string    `json:"headlineUrl"`
				PostDateTime time.Time `json:"postDateTime"`
				PostSummary  string    `json:"postSummary"`
				ContentType  string    `json:"contentType"`
				Headline     string    `json:"headline"`
			} `json:"additionalFields"`
			ID            string `json:"id"`
			DirectoryID   string `json:"directoryId"`
			LastUpdatedBy string `json:"lastUpdatedBy"`
		} `json:"item"`
	} `json:"items"`
}

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

func removeHTMLTags(item string) string {
	p := bluemonday.StripTagsPolicy()
	return p.Sanitize(item)
}

func fetchData(url string) []byte {
	resp, err := http.Get(url)

	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		panic(err)
	}

	return body
}

func buildURL(size string, page string) string {
	v := url.Values{}
	v.Set("size", size)
	v.Set("page", page)

	elements := []string{baseURL, path, query, sort, v.Encode()}
	fullURL, err := url.Parse(strings.Join(elements, ""))

	if err != nil {
		panic(err)
	}

	return fullURL.String()
}

func getNews(data []byte) Data {
	d := Data{}
	err := json.Unmarshal(data, &d)

	if err != nil {
		panic(err)
	}

	return d
}

func showNews(news Data, wrap int) {
	dt := time.Now()

	for index, i := range news.Items {
		date := i.Item.AdditionalFields.PostDateTime.Format("01-02-2006")
		// TODO: Add filtering
		if dt.Format("01-02-2006") == date {
			headline := i.Item.AdditionalFields.Headline
			postBody := wordWrap(removeHTMLTags(i.Item.AdditionalFields.PostBody), wrap)
			link := fmt.Sprintf("https://aws.amazon.com%s", i.Item.AdditionalFields.HeadlineURL)
			fmt.Printf("%d. %s\nPublished: %s\n%s\n\n", index+1, headline, date, link)
			fmt.Printf("%s\n\n", postBody)
		} else {
			continue
		}
	}
}

func main() {
	count := flag.String("c", "25", "number of feeds to show")
	page := flag.String("p", "0", "page number")
	wrap := flag.Int("w", 120, "line width")
	// TODO: add filtering news just for today, if available else show info: no news
	//today := flag.Bool("t", false, "shows news for today if available")
	flag.Parse()

	news := getNews(fetchData(buildURL(*count, *page)))
	showNews(news, *wrap)
}
