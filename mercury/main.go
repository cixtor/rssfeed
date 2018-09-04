package mercury

import (
	"errors"
	"io/ioutil"
	"time"
)

// api defines the base URL for the Mercurity API service.
//
// See: https://mercury.postlight.com/web-parser/
const api string = "https://mercury.postlight.com/parser"

// useragent defines the human readable description for the HTTP client.
const useragent string = "Mozilla/5.0 (KHTML, like Gecko) Safari/537.36"

type Mercury struct {
	token string
	agent string
}

type Article struct {
	Title         string    `json:"title"`
	Content       string    `json:"content"`
	Author        string    `json:"author"`
	Dek           string    `json:"dek"`
	Url           string    `json:"url"`
	Domain        string    `json:"domain"`
	Excerpt       string    `json:"excerpt"`
	Direction     string    `json:"direction"`
	LeadImageURL  string    `json:"lead_image_url"`
	NextPageURL   string    `json:"next_page_url"`
	DatePublished time.Time `json:"date_published"`
	RenderedPages int       `json:"rendered_pages"`
	TotalPages    int       `json:"total_pages"`
	WordCount     int       `json:"word_count"`
}

func New() *Mercury {
	m := new(Mercury)
	m.agent = useragent
	return m
}

func (m *Mercury) SetToken(token string) {
	m.token = token
}

func (m *Mercury) SetAgent(agent string) {
	m.agent = agent
}

func (m *Mercury) HasToken() bool {
	return m.token != ""
}

func (m *Mercury) Download(uuid string, link string) (Article, error) {
	if data, err := m.fetchFromCache(uuid); err == nil {
		return m.convertIntoArticle(data)
	}

	if data, err := m.fetchFromWeb(link); err == nil {
		filename := m.articleFilename(uuid)
		_ = ioutil.WriteFile(filename, data, 0644)
		return m.convertIntoArticle(data)
	}

	return Article{}, errors.New("article is not available")
}
