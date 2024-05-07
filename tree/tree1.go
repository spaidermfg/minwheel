package main

import "fmt"

// 角色结构体
type Role struct {
	ID   int
	Name string
}

// 资源结构体
type Resource struct {
	ID       int
	Name     string
	ParentID int
}

// 角色资源绑定关系
type RoleResource struct {
	RoleID     int
	ResourceID int
}

func main() {
	// 角色数据
	roles := []*Role{
		{ID: 1, Name: "角色A"},
		{ID: 2, Name: "角色B"},
	}

	// 资源数据
	resources := []*Resource{
		{ID: 1, Name: "资源1", ParentID: 0},
		{ID: 2, Name: "资源2", ParentID: 0},
		{ID: 3, Name: "资源3", ParentID: 1},
		{ID: 4, Name: "资源4", ParentID: 1},
		{ID: 5, Name: "资源5", ParentID: 2},
	}

	// 角色资源绑定关系数据
	roleResources := []*RoleResource{
		{RoleID: 1, ResourceID: 3},
		{RoleID: 1, ResourceID: 4},
		{RoleID: 2, ResourceID: 5},
		{RoleID: 1, ResourceID: 1},
	}

	// 打印角色数据
	fmt.Println("角色数据：")
	for _, role := range roles {
		fmt.Printf("ID: %d, Name: %s\n", role.ID, role.Name)
	}

	// 打印资源数据
	fmt.Println("\n资源数据：")
	for _, res := range resources {
		fmt.Printf("ID: %d, Name: %s, ParentID: %d\n", res.ID, res.Name, res.ParentID)
	}

	// 打印角色资源绑定关系数据
	fmt.Println("\n角色资源绑定关系数据：")
	for _, rr := range roleResources {
		fmt.Printf("RoleID: %d, ResourceID: %d\n", rr.RoleID, rr.ResourceID)
	}

	// 角色 A 绑定的资源 ID 列表
	var roleAResourceIDs []int
	for _, rr := range roleResources {
		if rr.RoleID == 1 { // 角色 A 的 ID 为 1
			roleAResourceIDs = append(roleAResourceIDs, rr.ResourceID)
		}
	}

	// 构建角色 A 绑定的资源树
	var roleAResources []*Resource
	for _, res := range resources {
		if contains(roleAResourceIDs, res.ID) {
			roleAResources = append(roleAResources, res)
		}
	}

	// 打印角色 A 绑定的资源树
	fmt.Println("角色 A 绑定的资源树：")
	printTree(roleAResources, 0)
}

// 检查切片中是否包含某个元素
func contains(slice []int, element int) bool {
	for _, e := range slice {
		if e == element {
			return true
		}
	}
	return false
}

// 打印资源树
func printTree(resources []*Resource, level int) {
	for _, res := range resources {
		fmt.Printf("%s- %s\n", getIndent(level), res.Name)
		printTree(getChildren(resources, res.ID), level+1)
	}
}

// 获取子资源
func getChildren(resources []*Resource, parentID int) []*Resource {
	var children []*Resource
	for _, res := range resources {
		if res.ParentID == parentID {
			children = append(children, res)
		}
	}
	return children
}

// 获取缩进字符串
func getIndent(level int) string {
	indent := ""
	for i := 0; i < level; i++ {
		indent += "  " // 可以根据需要调整缩进
	}
	return indent
}
