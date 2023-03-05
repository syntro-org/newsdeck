package main

import (
	"log"
	"time"

	"github.com/joho/godotenv"
	"github.com/syntro-org/newsdeck/internal/article"
	"github.com/syntro-org/newsdeck/internal/server"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env files found")
	}
}

func main() {
	err := article.Collect()

	if err != nil {
		log.Println(err.Error())
	}

	s := new(server.Server)
	s.Run(server.MakeAddr())

	ticker := time.NewTicker(24 * time.Hour)
	go func() {
		for {
			select {
			case <-ticker.C:
				err := article.Collect()

				if err != nil {
					log.Println(err.Error())
				}

				log.Println(article.Articles)
			}
		}
	}()

	done := make(chan bool)
	<-done

}
