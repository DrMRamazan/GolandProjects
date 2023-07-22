package main

import (
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gocolly/colly"
)

type tableData struct {
	Date, Time, Room, Lesson, Teacher, Form string
}

func main() {
	c := colly.NewCollector()

	c.WithTransport(&http.Transport{
		DialContext: (&net.Dialer{
			Timeout:   40 * time.Second,
			KeepAlive: 20 * time.Second,
			DualStack: true,
		}).DialContext,
		MaxIdleConns:          60,
		IdleConnTimeout:       50 * time.Second,
		TLSHandshakeTimeout:   10 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
	})

	var employeeData []tableData

	// c.OnRequest(func(r *colly.Request) {
	// 	fmt.Println("Scraping:", r.URL)
	// })

	c.OnResponse(func(r *colly.Response) {
		fmt.Println("Status:", r.StatusCode)
	})

	c.OnHTML("table > tbody", func(h *colly.HTMLElement) {
		h.ForEach("tr", func(_ int, el *colly.HTMLElement) {
			tableData := tableData{
				Date:    el.ChildText("th:nth-child(1)"),
				Time:    el.ChildText("td:nth-child(2) span"),
				Room:    el.ChildText("td:nth-child(3) dt span"),
				Lesson:  el.ChildText("td:nth-child(4) dd"),
				Teacher: el.ChildText("td:nth-child(4) dt"),
				Form:    el.ChildText("td:nth-child(5)"),
			}
			str := strings.Trim(tableData.Lesson, "\n")
			tableData.Lesson = str
			employeeData = append(employeeData, tableData)
		})
	})

	c.OnError(func(r *colly.Response, err error) {
		fmt.Println("Request URL:", r.Request.URL, "failed with response:", r, "\nError:", err)
	})

	c.Visit("https://itmo.ru/ru/exam/0/C42023/raspisanie_sessii_C42023.htm")

	content, err := json.Marshal(employeeData)
	if err != nil {
		fmt.Println(err.Error())
	}
	os.WriteFile("employees.json", content, 0644)
}