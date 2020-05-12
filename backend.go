package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/kaigedong/bai_backend/types"
)

// zhihu: https://www.zhihu.com/api/v4/search/top_search
const zhihuTopSearchURL = "https://www.zhihu.com/api/v4/search/top_search"

func main() {

	fmt.Println("Hello world!")

	if err := fetchZhihuTop(); err != nil {
		fmt.Println(err)
	}
}

func newClient() *http.Client {
	return &http.Client{
		Timeout: 5 * time.Second,
	}
}

func fetchZhihuTop() error {
	client := newClient()
	req, err := http.NewRequest(http.MethodGet, zhihuTopSearchURL, nil)
	if err != nil {
		return err
	}

	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	var zhihuTopSearch types.ZhihuTopSearch
	if err := json.Unmarshal(body, &zhihuTopSearch); err != nil {
		return err
	}

	return nil
}
