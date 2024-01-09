package requests

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func Post(url string, data any, queryVal map[string]interface{}, headers map[string]interface{}, timeout time.Duration) (resp *http.Response, err error) {

	//fmt.Println("query", queryVal)
	reqParam, _ := json.Marshal(data)
	reqBody := strings.NewReader(string(reqParam))
	httpReq, err := http.NewRequest("GET", url, reqBody)
	if err != nil {
		return
	}

	httpReq.Header.Add("Content-Type", "application/json")
	for key, val := range headers {
		switch v := val.(type) {
		case string:
			httpReq.Header.Add(key, v)
		case int:
			httpReq.Header.Add(key, strconv.Itoa(v))
		}
	}

	query := httpReq.URL.Query()
	id, _ := json.Marshal(queryVal["id"])
	ids := string(id)
	query.Add("id", ids)
	size, _ := json.Marshal(queryVal["size"])
	sizes := string(size)
	query.Add("size", sizes)
	//fmt.Println(httpReq.URL)
	query.Encode()
	httpReq.URL.RawQuery = query.Encode()
	replace := strings.Replace(httpReq.URL.RawQuery, "%22", "", -1)
	httpReq.URL.RawQuery = replace

	//fmt.Println(replace)
	//fmt.Println(httpReq.URL)

	client := &http.Client{
		Timeout: timeout,
	}
	httpResp, err := client.Do(httpReq)
	//fmt.Println(httpResp)
	return httpResp, err
}
