package webbug

import (
	"fmt"
	"log"
	"os"

	"gopl.io/ch5/links"
)

//并发的web爬虫
func Bug(){
  worklist := make(chan []string)
  unseenLinks := make(chan string)

  go func ()  {
    worklist <- os.Args[1:] 
  }()

  for i := 0; i < 20; i++ {
    go func ()  {
     for link := range unseenLinks {
       s := crawl(link)
       go func ()  { worklist <- s}()
     } 
    }()
  }

  seen := make(map[string]bool)
  for list := range worklist {
    for _, link := range list {
      if !seen[link] {
        seen[link] = true
        //go func (link string)  {
         // worklist <- crawl(link) 
        //}(link)
        unseenLinks <- link
      }
    }
  }
}


//创建通道限制并发量
var tokens = make(chan struct{}, 20)

func crawl(url string) []string {
  fmt.Println(url)
  //tokens <- struct{}{}
  list, err := links.Extract(url)
  //<-tokens
  if err != nil {
    log.Println("[error]: ",err)
  }

  return list
}
