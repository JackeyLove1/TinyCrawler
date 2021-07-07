package parse

import (
	"awesomeCrawl/engine"
	"awesomeCrawl/model"
	"regexp"
	"strings"
)

var authorRe = regexp.MustCompile(`<span class="pl">作者:</span>[\d\D]*?<a.*?>([^<]+)</a>`)
var publicerRe = regexp.MustCompile(`<span class="pl">出版社:</span>[\n]*([^<]+)<br/>`)
var bookpagesRe = regexp.MustCompile(`<span class="pl">页数:</span>[\n]*([^<]+)<br/>`)
var priceRe = regexp.MustCompile(`<span class="pl">定价:</span>[\n]*([^<]+)元<br/>`)
var scoreRe = regexp.MustCompile(`<strong class="ll rating_num " property="v:average"> ([^<]+) </strong>`)
var nameRe = regexp.MustCompile(`<span property="v:itemreviewed">([^<]+)</span>`)

func ParseBookDetail(content []byte) (engine.ParseResult, error){
	bookDetail := &model.BookDetail{}
	bookDetail.Author = ExtractString(content, authorRe)
	bookDetail.Score = ExtractString(content, scoreRe)
	bookDetail.Bookpages = ExtractString(content, bookpagesRe)
	bookDetail.Price = ExtractString(content, priceRe)
	bookDetail.Publicer = ExtractString(content, publicerRe)
	bookDetail.Name = ExtractString(content, nameRe)

	// fmt.Println(bookDetail)

	result := engine.ParseResult{
		Items: []interface{}{bookDetail},
	}

	return result, nil

}

func ExtractString(content []byte, re *regexp.Regexp) string{
	parseRes := re.FindSubmatch(content)
	if len(parseRes) == 0{
		return ""
	}

	res := string(parseRes[1])
	res = strings.ReplaceAll(res, "\n", "") // 去掉字段间的换行符
	res = strings.ReplaceAll(res, " ", "") // 去掉空白格
	return res
}