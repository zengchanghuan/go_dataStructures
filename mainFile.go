package main

import (
	"errors"
	"fmt"
	"io/ioutil"
)

func GetFile(path string,files[] string)([] string,error)  {
	read,err := ioutil.ReadDir(path)
	if err != nil {
		return files,errors.New("文件夹不可读取")
	}
	for _,filePath := range  read {
		//判断是否是文件夹
		if filePath.IsDir() {
			//构造新的路径
			fullDir := path + "/" + filePath.Name()

			//追加路径
			files = append(files,fullDir)

			//文件夹递归处理
			files,_= GetFile(fullDir,files)

		} else {
			fullDir := path + "/" + filePath.Name()
			files = append(files,fullDir)
		}
	}
	return files,nil
}
func main()  {
	path := "/Users/zengchanghuan/Desktop/iosTips"
	var files []string

	files,_ = GetFile(path,files)

	for i := 0; i < len(files); i++ {
		fmt.Println(files[i])
	}
}