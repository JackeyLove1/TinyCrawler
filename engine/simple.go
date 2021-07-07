package engine

import "log"

type SimpleEngine struct {

}

// 单线程队列处理
func (s *SimpleEngine) Run(seeds ... Request){
	var request []Request

	for _, r := range seeds{
		request = append(request, r)
	}

	for len(request) > 0{
		r := request[0]
		request = request[1:]

		log.Printf("Fetching ... %s\n", r.Url)
		parseResult , err := worker(r)
		if err != nil{
			continue
		}

		request = append(request, parseResult.Requests ... )

		for _, item := range parseResult.Items{
			log.Printf("Got Item %v\n", item)
		}
	}
}