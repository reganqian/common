package utils

import(
	"github.com/spf13/cast"
	"github.com/tealeg/xlsx"
	"github.com/astaxie/beego/utils"
	"os"
	"time"
	"reflect"
	"fmt"
)


func Struct2Map(obj interface{}) map[string]interface{} {
	t := reflect.TypeOf(obj)
	v := reflect.ValueOf(obj)

	var data = make(map[string]interface{})
	for i := 0; i < t.NumField(); i++ {
		data[t.Field(i).Name] = v.Field(i).Interface()
	}
	return data
}

func ExportExcel(dataList []map[string]interface{}, keys []string, names []string) (filename string, err error) {

	var file *xlsx.File
	var sheet *xlsx.Sheet
	
	file = xlsx.NewFile()
	sheet, _ = file.AddSheet("sheet1")
	titleRow := sheet.AddRow()
	for _, name := range names {
		cell := titleRow.AddCell()
		cell.Value = name
	}
	
	for _, data := range dataList {
		row := sheet.AddRow()
		for _, key := range keys {
			cell := row.AddCell()
			fmt.Println(key)
			fmt.Println(data)
			val := data[key]
			cell.Value =  fmt.Sprintf("%v", val)
		}
	}
	fileDir := "c://excel"
	if !utils.FileExists(fileDir) {
		os.MkdirAll(fileDir, os.ModePerm)
	}
	filename = fileDir + "/" + cast.ToString(time.Now().Unix()) + ".xlsx"
	err = file.Save(filename)
	//如果没错, 等待多久以后删除
	go RemoveExcel(filename)
	return filename, err
}

func RemoveExcel(fileName string) {
	time.Sleep(time.Second * 180)
	os.RemoveAll(fileName)
}