package main

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
)

/*
*******************************************************
Author: min
Date: 2023-12-29 21:58
Name: 文件目录正确性校验程序

校验文件是否为空
校验文件是否存在
校验二进制文件是否有可执行权限

功能项：

	  基础功能：
		扫描记录目录内容
		展示已记录目录条目
		展示已记录文件树
	  比对功能：
		查找缺失目录
		扫描结果展示


json

********************************************************
*/

const (
	WorkSpace = ".filetwins"
	FilePath  = ""
)

type Target struct {
	FileNum  int               // 文件数
	DirNum   int               // 目录总数
	NullFile map[string]string // 空文件
	NewFile  map[string]string // 新文件
}

func main() {
	roll()
}

func roll() {
	filepath.Walk(FilePath, func(path string, info fs.FileInfo, err error) error {

		return nil
	})
	filepath.WalkDir(FilePath, func(path string, d fs.DirEntry, err error) error {

		if err != nil {
			return err
		}
		return nil
	})

}

func (t *Target) save() {
	if err := os.MkdirAll(WorkSpace, 0777); err != nil {
		fmt.Println(err)
	}
}
