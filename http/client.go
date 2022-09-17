package http

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/taobig/go-helper/io"
	stdio "io"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
)

type RequestOption struct {
	timeout time.Duration
	params  map[string]string
	headers map[string]string
}

func NewRequestOption(timeout time.Duration) *RequestOption {
	return &RequestOption{timeout: timeout}
}

func (r *RequestOption) Params() map[string]string {
	return r.params
}

func (r *RequestOption) SetParams(params map[string]string) {
	r.params = params
}

func (r *RequestOption) Headers() map[string]string {
	return r.headers
}

func (r *RequestOption) SetHeaders(headers map[string]string) {
	r.headers = headers
}

func PostForm(urlString string, requestOption RequestOption) (*http.Response, error) {
	client := &http.Client{
		Timeout: requestOption.timeout * time.Second,
	}
	params := url.Values{}
	for key, val := range requestOption.params {
		params.Add(key, val)
	}
	//response, err := client.PostForm(urlString, params)
	//return response, err

	request, err := http.NewRequest("POST", urlString, strings.NewReader(params.Encode()))
	if err != nil {
		return nil, err
	}
	for key, value := range requestOption.headers {
		request.Header.Add(key, value)
	}
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	return client.Do(request)
}

func PostJson(urlString string, requestOption RequestOption) (*http.Response, error) {
	client := &http.Client{
		Timeout: requestOption.timeout * time.Second,
	}

	content, _ := json.Marshal(requestOption.params)
	request, err := http.NewRequest("POST", urlString, bytes.NewBuffer(content))
	if err != nil {
		return nil, err
	}
	for key, value := range requestOption.headers {
		request.Header.Add(key, value)
	}
	request.Header.Set("Content-Type", "application/json")
	return client.Do(request)
}

func Get(urlString string, requestOption RequestOption) (*http.Response, error) {
	client := &http.Client{
		Timeout: requestOption.timeout * time.Second,
	}

	request, err := http.NewRequest("GET", urlString, nil)
	if err != nil {
		return nil, err
	}
	// if you appending to existing query this works fine
	q := request.URL.Query()
	for key, val := range requestOption.params {
		q.Add(key, val)
	}
	//// or you can create new url.Values struct and encode that like so
	//q := url.Values{}
	//q.Add("api_key", "key_from_environment_or_flag")
	//q.Add("another_thing", "foo & bar")
	request.URL.RawQuery = q.Encode()

	//request.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	//request.Header.Add("Accept-Language", "zh-CN,zh;q=0.8,en-US;q=0.5,en;q=0.3")
	//request.Header.Add("Cookie", cookies)
	//request.Header.Add("Host", "")
	//request.Header.Add("Referer", "")
	//request.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; WOW64; rv:47.0) Gecko/20100101 Firefox/47.0")
	for key, value := range requestOption.headers {
		request.Header.Add(key, value)
	}
	//fmt.Println((request.Header))
	response, err := client.Do(request)
	return response, err
}

func ReadResponseBody(resp *http.Response) (*bytes.Buffer, error) {
	defer resp.Body.Close()
	//系统对bytes.Buffer提供了比[]byte更多的操作方法
	//body, err := ioutil.ReadAll(resp.Body)
	//return body, err

	buf := new(bytes.Buffer)
	_, err := buf.ReadFrom(resp.Body)
	//fmt.Println(buf.String())
	//return buf.String()
	return buf, err
}

func DownloadFile(url, dir string) (string, error) {
	var localFilePath string
	if dir != "" {
		err := os.MkdirAll(dir, os.ModePerm)
		if err != nil {
			return localFilePath, err
		}
	}

	//根据url地址分析文件保存到本地的文件名
	startIndex := strings.LastIndex(url, "/")
	filename := url[startIndex+1:]
	if len(filename) == 0 {
		filename = time.Now().Format("20060102_150405")
	} else {
		if strings.ContainsAny(filename, "?/|\\\"*:<>") {
			filename = time.Now().Format("20060102_150405")
		}
	}

	//如果同名文件已经存在，则变化保存文件名
	filenameAppendIndex := 1
	filenameNew := filename
	for {
		//fmt.Println(filenameNew)
		var tempFilePath string
		if dir != "" {
			tempFilePath = fmt.Sprintf("%s/%s", dir, filenameNew)
		} else {
			tempFilePath = filenameNew
		}
		b, err := io.PathExist(tempFilePath)
		if err != nil {
			return localFilePath, err
		}
		if b {
			suffixStartIndex := strings.LastIndex(filename, ".")
			if suffixStartIndex == -1 { //without ext name
				filenameNew = fmt.Sprintf("%s(%d)", filename, filenameAppendIndex)
			} else {
				filenameNew = fmt.Sprintf("%s(%d)%s", filename[:suffixStartIndex], filenameAppendIndex, filename[suffixStartIndex:])
			}
			filenameAppendIndex++
			continue
		}
		break
	}

	if dir != "" {
		localFilePath = fmt.Sprintf("%s/%s", dir, filenameNew)
	} else {
		localFilePath = filenameNew
	}
	//log.Println("save to ==> " + localFilePath)

	requestOption := RequestOption{}
	response, err := Get(url, requestOption)
	if err != nil {
		return localFilePath, err
	}
	if response.StatusCode != 200 {
		return localFilePath, errors.New("download file error")
	}
	totalBytes := response.ContentLength

	localFile, err := os.Create(localFilePath)
	if err != nil {
		return localFilePath, err
	}
	defer localFile.Close()

	var downloadBytes int64 = 0
	var downloadRate int64
	bufferBytes := 1024 * 1024 * 10
	buf := make([]byte, bufferBytes)
	var lastRecordDt = time.Now().Unix()
	for {
		size, e := response.Body.Read(buf)
		//读到文件结尾
		if e != nil {
			if e == stdio.EOF {
				if size == 0 {
					//log.Println("download complete")
					break
				}
				// size > 0
				//log.Println("read file end, but size > 0", size)
			} else {
				return localFilePath, err
			}
		}
		//log.Println("read size:", size)
		downloadBytes = downloadBytes + int64(size)
		_, err = localFile.Write(buf[:size])
		if err != nil {
			return localFilePath, err
		}
		downloadRate = downloadBytes * 100 / totalBytes
		now := time.Now()
		if now.Unix() >= lastRecordDt+5 {
			lastRecordDt = now.Unix()
			log.Printf("%d / %d     ==>%d%%[%s]\n", downloadBytes, totalBytes, downloadRate, now.Format(time.RFC3339))
		}
	}

	return localFilePath, nil
}
