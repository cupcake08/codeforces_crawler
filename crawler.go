package codeforcescrawler

import (
	"fmt"
	"log"
	"os"

	"github.com/gocolly/colly/v2"
)

type App struct {
	collector *colly.Collector
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

func NewApp(contestId int) *App {
	return &App{
		collector: colly.NewCollector(
			colly.MaxDepth(1),
			colly.Async(true),
		),
		contestId: contestId,
	}
}

// This will crawl codeforces website and store the sample tests of [pIndex] problem
// inputs and outputs to desired files.
func (app *App) GetTestCases(pid string) {
	c := app.collector.Clone()
	uri := fmt.Sprintf("https://codeforces.com/contest/%d/problem/%s", app.contestId, pid)
	c.OnHTML("div.sample-tests", func(h *colly.HTMLElement) {
		// input handler
		h.ForEach("div.input", func(i int, h *colly.HTMLElement) {
			child := h.ChildAttrs("div", "class")
			filename := fmt.Sprintf("input_%d.txt", i)
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
		fmt.Println("Input File/s Written")
		// output handler
		h.ForEach("div.output", func(i int, h *colly.HTMLElement) {
			filename := fmt.Sprintf("output_%d.txt", i)
			content := []byte(h.ChildText("pre") + "\n")
			write_to_file(filename, content)
		})
		fmt.Println("Output File/s Written")
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting:", r.URL)
	})

	c.OnError(func(r *colly.Response, err error) {
		fmt.Println("Request URL:", r.Request.URL, "failed with response:", r, "\nError:", err)
	})

	c.Visit(uri)
	c.Wait()
}
