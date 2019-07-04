package main

import (
	"errors"
	"fmt"
	"io"
	"strings"

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
	if v.isIrrelevant() {
		return errors.New("irrelevant content")
	}

	mark := strings.Index(v.Comments, "=")
	v.UUID = v.Comments[mark+1:]

	var err error
	var rdr io.Reader
	var doc readability.Article

	if rdr, err = Curl(v.Link); err != nil {
		return err
	}

	if doc, err = readability.New().Parse(rdr, v.Link); err != nil {
		return err
	}

	v.Description = fmt.Sprintf(
		"%s<hr><a href=\"%s\">Comments</a>",
		doc.Content,
		v.Comments,
	)

	return nil
}

func (v *Item) isIrrelevant() bool {
	return strings.Contains(v.Link, "://jobs.lever.co") ||
		strings.Contains(v.Link, "://www.themuse.com") ||
		strings.Contains(v.Link, "://www.businessinsider.com") ||
		strings.Contains(v.Title, "is hiring") ||
		strings.Contains(v.Title, "Is Hiring") ||
		strings.Contains(v.Title, "\x20(YC S") ||
		strings.Contains(v.Title, "\x20(YC W")
}
