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

func (v *Item) Curate(client *mercury.Mercury) error {
	if v.isBanned() {
		return errors.New("banned")
	}

	mark := strings.Index(v.Comments, "=")
	v.UUID = v.Comments[mark+1:]

	text, _ := v.Download(client)
	forum := "<hr><a href=\"" + v.Comments + "\">Comments</a>"
	v.Description = text + forum

	return nil
}

func (v *Item) Download(client *mercury.Mercury) (string, string) {
	info, err := url.Parse(v.Link)

	if err != nil {
		return err.Error(), "none"
	}

	if !v.HasBlockedMercury(info.Host) {
		data, err := client.Download(v.UUID, v.Link)

		if err == nil && !data.Unauthorized() {
			return data.Content, "mercury"
		}
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

func (v *Item) HasBlockedMercury(host string) bool {
	if host == "developer.apple.com" {
		return true
	}

	if host == "www.bloomberg.com" {
		return true
	}

	if host == "github.com" {
		return true
	}

	return false
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
