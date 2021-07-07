package parse

import (
	"awesomeCrawl/engine"
	"regexp"
)

const regExpr = `<a href="([^"]+)" class="tag">([^"]+)</a>`
const webTag = "https://book.douban.com"

func ParseTag(content []byte) (engine.ParseResult, error){
	// <a href="/tag/哲学" class="tag">哲学</a>
	re := regexp.MustCompile(regExpr)
	matches := re.FindAllSubmatch(content, -1)

	result := engine.ParseResult{}

	for _, m := range matches {
		// fmt.Println(string(m[1]))
		// result.Items = append(result.Items, m[2])
		result.Requests = append(result.Requests, engine.Request{
			Url:       webTag + string(m[1]),
			ParseFunc: ParseBookList,
		})
	}

	return result, nil
}
