package fetcher

import (
	"bufio"
	"fmt"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

var rateLimiter = time.Tick(100 * time.Millisecond)

func Fetch(url string) ([]byte, error) {
	<-rateLimiter
	//response, err := http.Get(url)
	request, _ := http.NewRequest(http.MethodGet, url, nil)
	request.Header.Add("user-agent",
		"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_0) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/77.0.3865.90 Safari/537.36")
	request.Header.Add("cookie", "sid=aa957f82-7f16-4997-9cad-5249de3a2e5c; Hm_lvt_2c8ad67df9e787ad29dbd54ee608f5d2=1572143153; FSSBBIl1UgzbN7N443S=3oZJt8Jfh.Zv4UtN1FYDjORNozv93cq8HYOtKkAFSD1rN63zxdekjMptummyHnCb; token=1819687631.1572315819740.2a049c100b07bdfe176ed7a799abbb3b; _pc_myzhenai_showdialog_=1; _pc_myzhenai_memberid_=%22%2C1819687631%22; Hm_lpvt_2c8ad67df9e787ad29dbd54ee608f5d2=1572315845; FSSBBIl1UgzbN7N443T=461OAAFdsm11cyFVPRxEWctsMEq2zdILcm8TeT4ArsNtqoGdd18dPtUlLYazCbnDRVL0PHVKL87J__4iTFkvPKala5xb.ttaTyTVF3v1wr3.mRTtxFwpz.tXOpFoPYFHIljIth0cWwYY39Y7.4WpVLNQEFYY6CgJrl6shrsRZUwX8vp4bzY.QIc9tXO1rT551LltPCvJVc6CjBuKFz7O4SR8ae_TuERrWyVqc_VOfRir.2y7Ta5qWVRonT6zfiHX3RxYvpKWC9EMQbbcvRvtq_lpvZ_Eq287ihob4pKKvcm51ebmv.7N4FsqLjhQCJJOW0lL")
	client := http.Client{
		CheckRedirect: func(
			req *http.Request,
			via []*http.Request) error {
			return nil
		},
	}
	response, err := client.Do(request)

	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("wrong status code: %d", response.StatusCode)
	}

	bodyReader := bufio.NewReader(response.Body)

	e := determineEncoding(bodyReader)
	utf8Reader := transform.NewReader(bodyReader, e.NewDecoder())

	return ioutil.ReadAll(utf8Reader)
}

func determineEncoding(reader *bufio.Reader) encoding.Encoding {
	bytes, err := bufio.NewReader(reader).Peek(1024)
	if err != nil {
		log.Printf("Fetcher error")
		return unicode.UTF8
	}

	e, _, _ := charset.DetermineEncoding(bytes, "")

	return e
}
