package utils

import (
	"net/http"
	"io/ioutil"
	"encoding/json"
	"net/url"
)


func DoGet(apiUrl string) (map[string]interface{}, error) {

	hc := http.Client{}
	req, err := http.NewRequest("POST", apiUrl, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	resp, err := hc.Do(req)
	if err != nil {
		return nil, err
	}
	stringData, _ := ParseResponse(resp)
	Log.Info(apiUrl, "：返回结果为：", stringData)
	return stringData, nil
}

func ParseResponse(response *http.Response) (map[string]interface{}, error){
	if response == nil {
		var result map[string]interface{}
		return result, nil
	}
	Log.Info(response.Body)
	body,err := ioutil.ReadAll(response.Body)
	stringBody := string(body)
	Log.Info("response body=", stringBody)
	// Log.Info(stringBody)
	var result map[string]interface{}
	if err == nil {
		err = json.Unmarshal(body, &result)
	}
	return result,err
}

//httpPostForm
func HttpPostForm(apiUrl string, params url.Values)  (map[string]interface{}, error) {
    resp, err := http.PostForm(apiUrl, params)
		// url.Values{"key": {"Value"}, "id": {"123"}})
    if err != nil {
		// handle error
		return nil, err
    }
 
    defer resp.Body.Close()

	stringData, _ := ParseResponse(resp)
	Log.Info(apiUrl, "：返回结果为：", stringData)
	return stringData, nil
 
}

//只管post过去, 不管对方返回信息
func PostUrl(apiUrl string, params url.Values)  (int, error) {
    resp, err := http.PostForm(apiUrl, params)
		// url.Values{"key": {"Value"}, "id": {"123"}})
	
    if err != nil {
		// handle error
		return 417, err
    }
	statusCode := resp.StatusCode 
   	return statusCode, err
 
}