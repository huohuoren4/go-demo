package main

import (
	"fmt"
	"github.com/spf13/viper" // 支持`json，toml，ini，yaml，hcl，env`等格式的文件内容
	"path/filepath"
)

type YamlStruct struct {
	Name    string `yaml:"name" json:"name"`
	Age     int    `yaml:"age" json:"age"`
	Address string `yaml:"address" json:"address"`
}

/**
读配置文件
@filename: 文件名
@filetype: 文件类型, json, yaml等类型
*/
func ReadConfig(filename string) interface{} {
	ext := filepath.Ext(filename)[1:]
	viper.SetConfigType(ext)
	viper.SetConfigFile(filename)
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}
	var yaml YamlStruct
	err = viper.Unmarshal(&yaml)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}
	return yaml
}

func WriteConfig(filename string, content YamlStruct) {
	ext := filepath.Ext(filename)[1:]
	viper.SetConfigType(ext)
	viper.SetConfigFile(filename)
	viper.Set("name", content.Name)
	viper.Set("age", content.Age)
	viper.Set("address", content.Address)
	err := viper.WriteConfig()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
}

func main() {
	fmt.Println(ReadConfig("test.yaml"))
	content := YamlStruct{Name: "tom", Age: 456, Address: "中关村晓坪路1203路"}
	WriteConfig("test01.json", content)
}
