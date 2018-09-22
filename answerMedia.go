package main

import (
	"errors"
	"github.com/gocolly/colly"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
)

type media struct {
	URL *url.URL
}

func main() {
	if len(os.Args) > 1 {
		answer := os.Args[1]
		if !strings.Contains(answer, " ") {
			media, err := getAnswerMedia(answer)
			if err != nil {
				log.Fatal(err)
			}
			media.downloadMedia()
		} else {
			log.Fatal(errors.New("answer url invalid"))
		}
	} else {
		log.Fatal(errors.New("no arguments"))
	}
}
func getAnswerMedia(URL string) (media, error) {
	c := colly.NewCollector(
		colly.UserAgent("Mozilla/5.0 (compatible; Googlebot/2.1; +http://www.google.com/bot.html)"),
	)
	var mediaArr []media
	c.OnHTML("article.item.streamItem.streamItem-single div.streamItem_visual", func(element *colly.HTMLElement) {
		element.ForEach("[src]", func(i int, element *colly.HTMLElement) {
			toBeParsed := element.Attr("src")
			if strings.Contains(element.Attr("src"), "gif") {
				toBeParsed = strings.Replace(element.Attr("src"), "/original/", "/animated/", -1)
			}
			parsedURL, _ := url.Parse(toBeParsed)
			mediaArr = append(mediaArr, media{parsedURL})
		})
	})
	c.Visit(URL)
	if len(mediaArr) > 1 {
		for _, aMedia := range mediaArr {
			if strings.Contains(aMedia.URL.String(), "mp4") {
				return aMedia, nil
			}
		}
	} else if len(mediaArr) == 0 {
		return media{}, errors.New("no media found")
	}
	return mediaArr[0], nil
}

func (media media) downloadMedia() error {
	client := http.Client{
		Timeout: 30 * time.Second,
	}
	file, _ := os.Create("./" + media.URL.String()[strings.LastIndex(media.URL.String(), "/")+1:])
	req, err := client.Get(media.URL.String())
	if err != nil {
		log.Fatal(err)
	}
	n, err := io.Copy(file, req.Body)
	if err != nil {
		log.Fatal(err)
		_ = n
	}
	file.Close()
	return err
}
func exists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return true, err
}
