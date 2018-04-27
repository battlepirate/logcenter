package common

/**
 * 文件名: config.go
 * 创建时间: 2016年6月12日-下午7:13:13
 * 简介:
 * 详情: 序列化/反序列化配置文件
 * Copyright (C) 2013 duhaibo0404@gmail.com. All Rights Reserved.
 */
import (
	"encoding/xml"
	"io/ioutil"
	"os"
)

//These includes are needed by each config.
func LoadConfigWithDefault(fileName, defaultName string) string {
	if _, err := os.Stat(fileName); err == nil {
		return fileName
	}
	return defaultName
}

func LoadConfig(filename string, v interface{}) error {
	if contents, err := ioutil.ReadFile(filename); err != nil {
		return err
	} else {
		if err = xml.Unmarshal(contents, v); err != nil {
			return err
		}
		return nil
	}
}

func SaveConfig(filename string, v interface{}) error {
	if contents, err := xml.Marshal(v); err != nil {
		return err
	} else {
		if err = ioutil.WriteFile(filename, contents, 0644); err != nil {
			return err
		}
		return nil
	}
}
