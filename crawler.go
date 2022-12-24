package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/gocolly/colly/v2"
)

func write_to_file(fileName string, content []byte) {
	file, err := os.Create(fileName)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer file.Close()
	file.Write(content)
}

func main() {
	dir, _ := os.ReadDir(".")

	for _, val := range dir {
		fileInfo, _ := val.Info()
		filename := fileInfo.Name()
		if strings.Contains(filename, ".txt") {
			os.Remove(filename)
		}
	}
	c := colly.NewCollector()

	c.OnHTML("div.sample-tests", func(h *colly.HTMLElement) {
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
				h.ForEach("pre > div", func(i int, h *colly.HTMLElement) {
					_, err := file.WriteString(fmt.Sprintf("%s\n", h.Text))
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

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting: ", r.URL)
	})

	// c.Visit("https://codeforces.com/contest/1763/problem/A")
	c.Visit("https://codeforces.com/contest/1763/problem/F")
}
