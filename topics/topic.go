package topics

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"

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

func checkDays(url string) bool {
	t := time.Now()
	day := fmt.Sprintf("%d", t.Day()-1)
	if len(day) < 2 {
		day = fmt.Sprintf("0%s", day)
	}
	month := fmt.Sprintf("%d", int(t.Month()))
	if len(month) < 2 {
		month = fmt.Sprintf("0%s", month)
	}

	today := fmt.Sprintf("%d/%s/%s", t.Year(), month, day)
	return strings.Contains(url, today)
}

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

	topics, err := scrapeTopics(resp.Body, baseUrl)
	if err != nil {
		return nil, errors.Wrap(err, "scrape topics err")
	}

	return topics, nil
}

func scrapeTopics(body io.Reader, baseUrl *url.URL) ([]Topic, error) {
	doc, err := goquery.NewDocumentFromReader(body)
	if err != nil {
		return nil, errors.Wrap(err, "new document from reader err")
	}

	urls := make([]string, 0, 5)
	titles := make([]string, 0, 5)
	doc.Find("h2").Each(func(_ int, srg *goquery.Selection) {
		srg.Find("a").Each(func(_ int, s *goquery.Selection) {
			href, exists := s.Attr("href")
			if exists {
				reqUrl, err := baseUrl.Parse(href)
				if !checkDays(reqUrl.String()) {
					return
				}
				if err == nil {
					urls = append(urls, reqUrl.String())
				}
			}
			titles = append(titles, s.Text())
		})
	})

	descriptions := make([]string, 0, 5)
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
