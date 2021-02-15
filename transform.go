package gtrimstrings

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"reflect"
	"strings"

	"github.com/gin-gonic/gin"
)

var Excepts = map[string]bool{
	"password": true,
}

func Transform(c *gin.Context) {
	var body map[string]interface{}
	_ = c.Bind(&body)

	cleanBody(body)
	log.Printf("Request transformed : %v", body)

	newBody, _ := json.Marshal(body)
	c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(newBody))
}

func cleanBody(objects map[string]interface{}) {
	for key, value := range objects {
		reflectValue := reflect.ValueOf(value)
		if !isValueInMap(key) {
			objects[key] = transform(reflectValue, value)
		}
	}
}

func cleanMapString(datas map[string]interface{}) interface{} {
	for key, value := range datas {
		reflectValue := reflect.ValueOf(value)
		if !isValueInMap(key) {
			datas[key] = transform(reflectValue, value)
		}
	}
	return datas
}

func cleanSlice(datas []interface{}) interface{} {
	for key, value := range datas {
		reflectValue := reflect.ValueOf(value)
		datas[key] = transform(reflectValue, value)
	}
	return datas
}

func transform(reflectValue reflect.Value, value interface{}) interface{} {

	if reflectValue.Kind() == reflect.Slice {
		return cleanSlice(value.([]interface{}))
	}
	if reflectValue.Kind() == reflect.Map {
		return cleanMapString(value.(map[string]interface{}))
	}

	if reflectValue.Kind() == reflect.String {
		ret := strings.TrimSpace(value.(string))
		//ret = strings.ToLower(ret)
		return ret
	}
	return value
}

func isValueInMap(value string) bool {
	return Excepts[value]
}
