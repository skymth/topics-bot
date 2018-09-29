package topics

import (
	"net/http"
	"net/url"

	"github.com/PuerkitoBio/goquery"
	"github.com/pkg/errors"
)

const (
	topicsEndpoint = "https://jp.techcrunch.com/"
)

type Topic struct {
	Title       string
	Description string
	URL         string
}

// web crowler の実装
// 以前とったことがあるかの判定もかく
func crawle() ([]Topic, error) {
	baseUrl, err := url.Parse(topicsEndpoint)
	if err != nil {
		return nil, errors.Wrap(err, "url parse err")
	}

	resp, err := http.Get(baseUrl.String())
	if err != nil {
		return nil, errors.Wrap(err, "http get err")
	}
	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return nil, errors.Wrap(err, "new document from reader err")
	}

	urls := make([]string, 0, 5)
	titles := make([]string, 0, 5)
	descriptions := make([]string, 0, 5)
	doc.Find("h2").Each(func(_ int, srg *goquery.Selection) {
		srg.Find("a").Each(func(_ int, s *goquery.Selection) {
			href, exists := s.Attr("href")
			if exists {
				reqUrl, err := baseUrl.Parse(href)
				if err == nil {
					urls = append(urls, reqUrl.String())
				}
			}
			titles = append(titles, s.Text())
		})
	})

	doc.Find("p.excerpt").Each(func(_ int, s *goquery.Selection) {
		descriptions = append(descriptions, s.Text())
	})

	topics := make([]Topic, len(urls))
	for i, _ := range urls {
		topics[i].Title = titles[i]
		topics[i].URL = urls[i]
		topics[i].Description = descriptions[i]
	}

	return topics, nil
}

func GetTopics() ([]Topic, error) {
	topics, err := crawle()
	if err != nil {
		return nil, errors.Wrap(err, "crawle err")
	}

	return topics, nil
}
