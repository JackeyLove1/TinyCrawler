package parse

import (
	"awesomeCrawl/engine"
	"regexp"
)

const bookListRe = `<a href="([^"]+)" title="([^""]+)"`
func ParseBookList(content []byte) (engine.ParseResult, error){
	re := regexp.MustCompile(bookListRe)

	matches := re.FindAllSubmatch(content, -1)

	result := engine.ParseResult{}

	for _, m := range matches {
		// result.Items = append(result.Items, m[2])
		result.Requests = append(result.Requests, engine.Request{
			Url:       string(m[1]),
			ParseFunc: ParseBookDetail,
		})
		// fmt.Printf("m[1]: %s, m[2]: %s\n", m[1], m[2])
	}

	return result, nil
}