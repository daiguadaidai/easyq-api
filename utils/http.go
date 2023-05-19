package utils

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

const (
	DefaultTimeout = 30
)

// Get请求
func GetURLRaw(url, query string, hearders map[string]string) ([]byte, error) {
	if len(query) != 0 {
		url += "?" + query
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("新建get请求失败. %s", err.Error())
	}
	req.Close = true
	for key, value := range hearders {
		req.Header.Add(key, value)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("获取返回数据失败. %v", err)
	}

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("code: %d. %s", resp.StatusCode, string(body))
	}

	return body, nil
}

// Post请求
func PostURLRaw(url string, data interface{}, hearders map[string]string, timeout int64) ([]byte, error) {
	if timeout <= 0 {
		timeout = DefaultTimeout
	}

	jsonRaw, err := json.Marshal(data)
	if err != nil {
		return nil, fmt.Errorf("post, 序列化数据失败: %s", err.Error())
	}

	ctx, _ := context.WithTimeout(context.Background(), time.Duration(timeout)*time.Second)
	req, err := http.NewRequestWithContext(ctx, "POST", url, bytes.NewBuffer(jsonRaw))
	if err != nil {
		return nil, fmt.Errorf("新建post请求失败. %s", err.Error())
	}
	req.Close = true
	for key, value := range hearders {
		req.Header.Add(key, value)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("获取返回数据失败. %v", err)
	}

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("code: %d. %s", resp.StatusCode, string(body))
	}

	return body, nil
}

func PostURLRaw2xx(url string, data interface{}, hearders map[string]string) ([]byte, error) {
	jsonRaw, err := json.Marshal(data)
	if err != nil {
		return nil, fmt.Errorf("post, 序列化数据失败: %s", err.Error())
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonRaw))
	if err != nil {
		return nil, fmt.Errorf("新建post请求失败. %s", err.Error())
	}
	req.Close = true
	for key, value := range hearders {
		req.Header.Add(key, value)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("获取返回数据失败. %v", err)
	}

	if resp.StatusCode < 200 && resp.StatusCode >= 300 {
		return nil, fmt.Errorf("code: %d. %s", resp.StatusCode, string(body))
	}

	return body, nil
}

func InterfaceToUrlQuery(data interface{}, ignoreEmpty bool) (string, error) {
	raw, err := json.Marshal(data)
	if err != nil {
		return "", fmt.Errorf("Url Query struct数据转化为 json字符串出错. %s", err.Error())
	}

	dataMap := make(map[string]interface{})
	if err = json.Unmarshal(raw, &dataMap); err != nil {
		return "", fmt.Errorf("Url Query json数据转化为Map出错. data: %s,  %s", string(raw), err.Error())
	}

	querys := make([]string, 0, len(dataMap))
	for key, value := range dataMap {
		switch v := value.(type) {
		case string:
			if strings.TrimSpace(v) == "" {
				continue
			}
			querys = append(querys, fmt.Sprintf("%s=%s", key, v))
		default:
			querys = append(querys, fmt.Sprintf("%s=%v", key, v))
		}
	}

	return strings.Join(querys, "&"), nil
}

func PostUploadRaw(url string, filename string, hearders map[string]string) ([]byte, error) {
	data, err := ReadBinaryFile(filename)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("新建post upload请求失败. %s", err.Error())
	}
	req.Close = true
	for key, value := range hearders {
		req.Header.Add(key, value)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("获取返回数据失败. %v", err)
	}

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("code: %d. %s", resp.StatusCode, string(body))
	}

	return body, nil
}
