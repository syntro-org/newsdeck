package article

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/syntro-org/newsdeck/config"
)

type Article struct {
	Author      string `json:"author"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Content     string `json:"content"`
	Url         string `json:"url"`
	Image       string `json:"urlToImage"`
	PublishedAt string `json:"publishedAt"`
}

type Response struct {
	Articles []Article `json:"articles"`
}

var URL string
var Articles []Article

func Collect() error {
	conf := config.New()
	URL = "https://newsapi.org/v2/everything?sources=bbc-news,cnn,the-washington-post&apiKey=" + conf.ApiKey

	response, err := http.Get(URL)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	var news Response
	err = json.NewDecoder(response.Body).Decode(&news)
	if err != nil {
		return err
	}

	// Set new time format
	var correctedArticles []Article

	for _, v := range news.Articles {
		t, err := time.Parse("2006-01-02T15:04:05.9999999Z", v.PublishedAt)

		if err != nil {
			return err
		}

		formattedTime := t.Format("02.01.2006 15:04")

		new := Article{
			Author:      v.Author,
			Title:       v.Title,
			Description: v.Description,
			Content:     v.Content,
			Url:         v.Url,
			Image:       v.Image,
			PublishedAt: formattedTime,
		}

		// Get full text of article
		response, err := http.Get(v.Url)
		if err != nil {
			correctedArticles = append(correctedArticles, new)
			Articles = correctedArticles
			return err
		}
		defer response.Body.Close()

		document, err := goquery.NewDocumentFromReader(response.Body)
		if err != nil {
			return err
		}

		var text string
		document.Find("p").Each(func(index int, element *goquery.Selection) {
			text += "\n" + element.Text()
		})

		new.Content = text

		correctedArticles = append(correctedArticles, new)
	}

	Articles = correctedArticles
	return nil
}
