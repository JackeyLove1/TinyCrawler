package parse

import (
	"awesomeCrawl/fetcher"
	"fmt"
	"log"
	"testing"
)

func TestExtractString(t *testing.T) {
	url := "https://book.douban.com/subject/4908885/"
	content, err := fetcher.Fetch(url)
	if err != nil{
		log.Panic(err)
	}

	bookDetail, err  := ParseBookDetail(content)
	if err != nil{
		log.Panic(err)
	}
	fmt.Println(bookDetail)
}

func TestParseBookList(t *testing.T) {
	url := "https://book.douban.com/tag/随笔"
	content, err := fetcher.Fetch(url)
	if err != nil{
		log.Panic(err)
	}

	bookList, err  := ParseBookList(content)
	if err != nil{
		log.Panic(err)
	}
	fmt.Println(bookList)
}

func TestParseTag(t *testing.T) {
	url := "https://book.douban.com/"
	content, err := fetcher.Fetch(url)
	if err != nil{
		log.Panic(err)
	}

	bookTag, err  := ParseTag(content)
	if err != nil{
		log.Panic(err)
	}
	fmt.Println(bookTag)
}