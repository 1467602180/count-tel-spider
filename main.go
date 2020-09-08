package main

import (
	"fmt"
	"github.com/gocolly/colly"
	"github.com/gogf/gf/container/garray"
	"github.com/gogf/gf/container/gmap"
	"github.com/gogf/gf/os/gfile"
)

func main() {
	count:=garray.NewStrArray(true)
	tel:=garray.NewStrArray(true)
	// Instantiate default collector
	c := colly.NewCollector(
		// Visit only domains: hackerspaces.org, wiki.hackerspaces.org
		colly.AllowedDomains("www.66cu.com"),
	)

	// On every a element which has href attribute call callback
	c.OnHTML("#content > table > tbody > tr > td:nth-child(1)", func(e *colly.HTMLElement) {
		count.Append(e.Text)
	})

	// On every a element which has href attribute call callback
	c.OnHTML("#content > table > tbody > tr > td:nth-child(4)", func(e *colly.HTMLElement) {
		tel.Append(e.Text)
	})

	// Before making a request print "Visiting ..."
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	c.OnScraped(func(r *colly.Response) {
		fmt.Println("Finished", r.Request.URL)
		dataTrim(count,tel)
	})

	// Start scraping on https://hackerspaces.org
	c.Visit("http://www.66cu.com/rcsh/gjqh.asp")
}

func dataTrim(count *garray.StrArray,tel *garray.StrArray)  {
	data:=gmap.NewStrStrMap(true)
	count.Remove(0)
	tel.Remove(0)
	count.Iterator(func(k int, v string) bool {
		countItem,_:=tel.Get(k)
		data.Set(v,countItem)
		return true
	})
	file,err:=gfile.Create("data.json")
	if err==nil {
		jsonData,_:=data.MarshalJSON()
		file.Write(jsonData)
	}else{
		fmt.Println(err)
	}
	fmt.Println(data.MarshalJSON())
}