package mercury

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func (m *Mercury) articleFilename(uuid string) string {
	return "/tmp/" + uuid + ".txt"
}

func (m *Mercury) convertIntoArticle(data []byte) (Article, error) {
	var v Article

	if err := json.Unmarshal(data, &v); err != nil {
		return Article{}, err
	}

	return v, nil
}

func (m *Mercury) fetchFromCache(uuid string) ([]byte, error) {
	filename := m.articleFilename(uuid)

	if _, err := os.Stat(filename); !os.IsNotExist(err) {
		return ioutil.ReadFile(filename)
	}

	return nil, errors.New("article is not cached")
}

func (m *Mercury) fetchFromWeb(link string) ([]byte, error) {
	client := &http.Client{Timeout: 5 * time.Second}
	req, err := http.NewRequest("GET", api+"?url="+link, nil)

	if err != nil {
		return nil, err
	}

	req.Header.Set("x-api-key", m.token)
	req.Header.Set("user-agent", m.agent)

	res, err := client.Do(req)

	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	return ioutil.ReadAll(res.Body)
}
