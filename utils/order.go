package utils

import (
	"github.com/bitly/go-simplejson"
	"fmt"
)

func DealJsonToMap(orderExt string) (data map[string]string, err error) { 
	data, err = JsonToMap(orderExt)
    if err != nil {
		fmt.Printf("Convert json to map failed with error: %+v\n", err)
		return data, err
    }
   
	return data, err
}


func DealJsonToArray(orderExt string) (extData []map[string]interface{}) {
	// listStrByte := []byte(orderExt)
	
	//jsonObj, eee := gjson.DecodeToJson(listStrByte)
	// 这里的数据类型一开始定义成了map数组，折腾半天
	// var strArr []string
	jsonstr, _ := simplejson.NewJson([]byte(orderExt))
	// eee := json.Unmarshal(listStrByte, &strArr)
	// if eee != nil {
	// 	Log.Error(eee)
	// }
	jsonArr, _ := jsonstr.Array() 
	for _, ext := range jsonArr {
		Log.Info(ext)
		extData = append(extData, ext.(map[string]interface{}))			
	}

	return extData
}