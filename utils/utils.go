/*******************************************************************************
Method: 工具函数
Author: Lemine
Langua: Golang 1.14
Modify：2020/03/07
*******************************************************************************/
package utils

import (
	"os"
)

//判断文件路径是否存在
func PathExist(path string) bool {
	_, err := os.Stat(path) //获取文件信息
	if err == nil {
		return true
	}

	if !os.IsNotExist(err) {
		return true
	}

	return false
}
