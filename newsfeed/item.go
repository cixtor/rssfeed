package newsfeed

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"time"

	"github.com/cixtor/readability"
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

	doc, err := readability.FromURL(v.Link, 10*time.Second)

	if err != nil {
		return err
	}

	v.Description = fmt.Sprintf(
		"%s<hr><a href=\"%s\">Comments</a>",
		doc.Content,
		v.Comments,
	)

	return nil
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
