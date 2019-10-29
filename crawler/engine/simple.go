package engine

import (
	"gostudy/crawler/fetcher"
	"log"
)

/*
 *
 * <pre>
 *  Version         Date            Author          Description
 * ---------------------------------------------------------------------------------------
 *  1.0.0           2019/10/27     redli        -
 * </pre>
 * @author redli
 * @version 1.0.0 2019/10/27 10:57 AM
 * @date 2019/10/27 10:57 AM
 * @since 1.0.0
 */

type SimpleEngine struct {
}

func (SimpleEngine) Run(seeds ...Request) {
	var requests []Request

	for _, r := range seeds {
		requests = append(requests, r)
	}

	for len(requests) > 0 {
		r := requests[0]
		requests = requests[1:]

		parserResult, err := worker(r)

		if err != nil {
			continue
		}

		requests = append(requests, parserResult.Request...)

		for _, item := range parserResult.Items {
			log.Printf("Got item %v", item)
		}
	}
}

func worker(r Request) (ParserResult, error) {
	log.Printf("Fetching %s", r.Url)

	body, err := fetcher.Fetch(r.Url)

	if err != nil {
		log.Printf("Fetcher: error fetching url %s: %v", r.Url, err)
		return ParserResult{}, err
	}
	return r.ParserFunc(body), nil
}
