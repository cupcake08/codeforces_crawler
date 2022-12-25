package codeforcescrawler 

import (
	"fmt"
	"log"
	"os"
	"github.com/gocolly/colly/v2"
)

type App struct {
	collector *colly.Collector
	pIndex string 
	contestId int
}

func write_to_file(fileName string, content []byte) {
	file, err := os.Create(fileName)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer file.Close()
	file.Write(content)
}

func NewApp(index string,contestId int) *App {
	return &App{
		collector: colly.NewCollector(),
		pIndex: index,
		contestId: contestId,
	}
}

func (app *App) DoWork() {
	uri := fmt.Sprintf("https://codeforces.com/contest/%d/problem/%s",app.contestId,app.pIndex)
	app.collector.OnHTML("div.sample-tests", func(h *colly.HTMLElement) {
		// input handler
		h.ForEach("div.input", func(i int, h *colly.HTMLElement) {
			child := h.ChildAttrs("div", "class")
			filename := fmt.Sprintf("input%d.txt", i)
			if len(child) > 1 {
				file, err := os.Create(filename)
				if err != nil {
					log.Fatal(err.Error())
				}
				defer file.Close()
				h.ForEach("pre > div", func(_ int, h *colly.HTMLElement) {
					_, err := file.WriteString(h.Text + "\n")
					if err != nil {
						log.Fatal(err)
					}
				})
			} else {
				content := []byte(h.ChildText("pre"))
				write_to_file(filename, content)
			}
		})

		// output handler
		h.ForEach("div.output", func(i int, h *colly.HTMLElement) {
			filename := fmt.Sprintf("output%d.txt", i)
			content := []byte(h.ChildText("pre") + "\n")
			write_to_file(filename, content)
		})
	})

	app.collector.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting:",r.URL)
	})

	app.collector.Visit(uri)
}

