package fetcher

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
)

// 为全局变量但是限定于包内, Tick会定时向通道发送信号
// 所有worker回去抢占rateLimiter，
// Tick的时间设置过长会导致goroutine退化成单线程的
var rateLimiter = time.Tick(time.Millisecond)

func Fetch(url string) ([]byte, error) {

	// resp, err := http.Get(url)
	<-rateLimiter
	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	request.Header.Add("User-Agent",
		"Mozilla/5.0 (iPhone; CPU iPhone OS 11_0 like Mac OS X) AppleWebKit/604.1.38 (KHTML, like Gecko) Version/11.0 Mobile/15A372 Safari/604.1",
	)

	resp, err := http.DefaultClient.Do(request)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Error Status code %d", resp.StatusCode)

	}
	bodyReader := bufio.NewReader(resp.Body)
	e := determineEncoding(bodyReader)
	// e := determineEncoding(resp.Body)
	// utf8Reader := transform.NewReader(resp.Body, simplifiedchinese.GBK.NewDecoder())
	utf8Reader := transform.NewReader(bodyReader, e.NewDecoder())
	return ioutil.ReadAll(utf8Reader)

}

func determineEncoding(r *bufio.Reader) encoding.Encoding {

	// 直接传resp.body会在读的时候导致句柄后移1024
	// 原因在于Peek会调用Reader.fill方法去填充buf，导致resp.Body的文件指针移动
	// Peek本身是不会改变文件或者缓冲区指针的
	bytes, err := r.Peek(1024)
	if err != nil {
		log.Printf("Fetcher determineEncoding error %s", err)
		return unicode.UTF8
	}

	e, _, _ := charset.DetermineEncoding(bytes, "")
	return e

}
