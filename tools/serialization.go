package tools

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

// Deserialize 从文件内反序列化
func Deserialize(obj interface{}, path string) error {
	data, err := ioutil.ReadFile(path)
	if err != nil { // 如果文件读取错误，打印日志
		return fmt.Errorf("failed to open %s file: %s", path, err.Error())
	}

	nDate, err := Strings.StrictJSON(data)
	if err != nil {
		return fmt.Errorf("failed to strict %s file to json: %s", path, err.Error())
	}

	if err := json.Unmarshal(nDate, obj); err != nil {
		return fmt.Errorf("failed to deserialize %s file to json: %s", path, err.Error())
	}
	return nil
}

// Serialize 序列化数据到文件
func Serialize(obj interface{}, path string) error {
	data, err := json.MarshalIndent(obj, "", "    ")
	if err != nil {
		return fmt.Errorf("failed to serialize %s file to json: %s", path, err.Error())
	}
	if err := ioutil.WriteFile(path, data, 0664); err != nil {
		return fmt.Errorf("failed to write json to %s file: %s", path, err.Error())
	}
	return nil
}
