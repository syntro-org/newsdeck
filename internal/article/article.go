package article

import (
	"encoding/json"
	"net/http"
	"strings"
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
var Articles map[string]Article

func GetArticleByTitle(title string) Article {
	var obj Article

	for _, v := range Articles {
		if v.Title == title {
			obj = v
			break
		}
	}

	return obj
}

func Collect() error {
	conf := config.New()
	URL = "https://newsapi.org/v2/everything?sources=bbc-news,cnn,the-washington-post,associated-press,abc-news,reuters&language=en&apiKey=" + conf.ApiKey

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

	var correctedArticles map[string]Article = make(map[string]Article)

	for _, v := range news.Articles {
		// Check if the article with the same title already exists in the map
		if _, ok := Articles[v.Title]; !ok {
			if _, ok = correctedArticles[v.Title]; !ok {
				t, err := time.Parse("2006-01-02T15:04:05.9999999Z", v.PublishedAt)

				if err != nil {
					return err
				}

				formattedTime := t.Format("2006-01-02")

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
					correctedArticles[new.Title] = new
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

				correctedArticles[new.Title] = new
			}
		}
	}

	Articles = correctedArticles
	return nil
}

func (a Article) SplitContent() []string {
	if strings.Contains(a.Content, "\n") {
		return strings.Split(a.Content, "\n")
	} else {
		return []string{a.Content}
	}
}
