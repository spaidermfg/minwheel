package main

import (
	"fmt"
)

// MenuItem 表示菜单权限项
type MenuItem struct {
	Key       string
	ParentKey string
	Desc      string
}

// FindChildrenRecursive 递归查找指定父节点的所有子节点
func FindChildrenRecursive(parentKey string, allItems []*MenuItem) []*MenuItem {
	var children []*MenuItem
	for _, item := range allItems {
		if item.ParentKey == parentKey && item.ParentKey != "" && containsData(item.Desc) {
			children = append(children, item)
			children = append(children, FindChildrenRecursive(item.Key, allItems)...)
		}
	}
	return children
}

// containsData 检查描述中是否包含"数据"字样
func containsData(desc string) bool {
	return containsSubstring(desc, "数据")
}

// containsSubstring 检查字符串中是否包含子字符串
func containsSubstring(str, substr string) bool {
	return len(str) >= len(substr) && str[len(str)-len(substr):] == substr
}

// Paginate 对子菜单列表进行分页处理
func Paginate(items []*MenuItem, pageSize, page int) []*MenuItem {
	start := (page - 1) * pageSize
	end := start + pageSize
	if start >= len(items) {
		return []*MenuItem{}
	}
	if end > len(items) {
		end = len(items)
	}
	return items[start:end]
}

func main() {
	// 假设有一个包含所有节点的列表
	allItems := []*MenuItem{
		{Key: "A", ParentKey: "", Desc: ""},
		{Key: "B", ParentKey: "A", Desc: "数据"},
		{Key: "C", ParentKey: "B", Desc: "数据"},
		{Key: "D", ParentKey: "B", Desc: "其他描述"},
		{Key: "E", ParentKey: "A", Desc: "其他描述"},
		{Key: "F", ParentKey: "E", Desc: "其他描述"},
		{Key: "G", ParentKey: "A", Desc: "数据"},
		{Key: "H", ParentKey: "A", Desc: "其他描述"},
		{Key: "I", ParentKey: "C", Desc: "数据"},
		{Key: "J", ParentKey: "C", Desc: "数据"},
		{Key: "K", ParentKey: "C", Desc: "数据"},
		{Key: "L", ParentKey: "C", Desc: "其他描述"},
		{Key: "M", ParentKey: "G", Desc: "数据"},
		{Key: "N", ParentKey: "G", Desc: "其他描述"},
		{Key: "O", ParentKey: "G", Desc: "数据"},
		{Key: "P", ParentKey: "G", Desc: "其他描述"},
		{Key: "Q", ParentKey: "G", Desc: "其他描述"},
	}

	// 指定要找的父节点
	parentKey := "A"

	// 找到指定父节点的所有子节点
	children := FindChildrenRecursive(parentKey, allItems)

	fmt.Printf("节点 %s 的所有子节点中包含“数据”描述的菜单:\n", parentKey)
	for _, child := range children {
		fmt.Println(child.Key)
	}

	// 设置分页参数
	pageSize := 4
	page := 2

	// 分页处理子菜单列表
	paginatedChildren := Paginate(children, pageSize, page)
	// 输出结果
	fmt.Printf("节点 %s 的所有子节点中包含“数据”描述的菜单:\n", parentKey)
	for _, child := range paginatedChildren {
		fmt.Println(child.Key)
	}
}
