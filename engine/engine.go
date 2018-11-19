package engine

import (
	"log"
	"zbs/single-spider/fetcher"
)

func Run(seed ...Request) {
	var requestQueue []Request
	requestQueue = append(requestQueue, seed...)

	for len(requestQueue) > 0 {
		request := requestQueue[0]
		requestQueue = requestQueue[1:]

		log.Printf("Fetching: %v", request.Url)
		contents, err := fetcher.FetchUseHeader(request.Url)
		if err != nil {
			log.Println(err.Error())
			continue
		}

		parseResult, err := request.ParseFunc(contents)
		if err != nil {
			continue
		}

		requestQueue = append(requestQueue, parseResult.Request...)

		for _, value := range parseResult.Items {
			log.Printf("获取到Item: %s", value)
		}
	}

}
