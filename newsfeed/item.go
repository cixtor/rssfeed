package newsfeed

import (
	"errors"
	"io/ioutil"
	"net/url"
	"os"
	"strings"

	"github.com/cixtor/rssfeed/mercury"
	"github.com/cixtor/rssfeed/parsers"
)

type Item struct {
	UUID        string `xml:"uuid"`
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	PubDate     string `xml:"pubDate"`
	Comments    string `xml:"comments"`
	Description string `xml:"description"`
}

func (v *Item) Curate() error {
	if v.isBanned() {
		return errors.New("banned")
	}

	mark := strings.Index(v.Comments, "=")
	v.UUID = v.Comments[mark+1:]

	var text string
	// text, _ := v.Download()
	data, err := mercury.Download(v.UUID, v.Link)

	if err != nil {
		return err
	}

	// if data.LeadImageURL != "" {
	// 	text += "<div><img src=\"" + data.LeadImageURL + "\"></div>"
	// }

	text += data.Content

	if v.Comments != "" {
		text += "<hr><a href=\"" + v.Comments + "\">Comments</a>"
	}

	v.Description = text

	return nil
}

func (v *Item) Download() (string, string) {
	info, err := url.Parse(v.Link)

	if err != nil {
		return err.Error(), "none"
	}

	body, err := v.CrawlContent()

	if err != nil {
		return err.Error(), "none"
	}

	return parsers.Article(info.Host, string(body))
}

func (v *Item) CrawlContent() ([]byte, error) {
	filename := "/tmp/" + v.UUID + ".txt"

	if _, err := os.Stat(filename); !os.IsNotExist(err) {
		return ioutil.ReadFile(filename)
	}

	reader, err := Curl(v.Link)

	if err != nil {
		return []byte{}, err
	}

	body, err := ioutil.ReadAll(reader)

	if err != nil {
		return []byte{}, err
	}

	if v.UUID != "" {
		_ = ioutil.WriteFile(filename, body, 0644)
	}

	return body, nil
}

func (v *Item) isBanned() bool {
	if strings.Contains(v.Link, "://jobs.lever.co") {
		return true
	}

	if strings.Contains(v.Link, "://www.themuse.com") {
		return true
	}

	if strings.Contains(v.Link, "://www.businessinsider.com") {
		return true
	}

	if strings.Contains(v.Title, "is hiring") {
		return true
	}

	if strings.Contains(v.Title, "Is Hiring") {
		return true
	}

	if strings.Contains(v.Title, "\x20(YC S") {
		return true
	}

	if strings.Contains(v.Title, "\x20(YC W") {
		return true
	}

	return false
}
