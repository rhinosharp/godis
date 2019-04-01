package main

import (
	"fmt"

	"github.com/rhinosharp/godis"
)

func main() {

	dm := godis.New()
	dm.LoadJSON("demo.json")

	trans := [][]string{{"１", "２", "３", "４", "５", "６", "７", "８", "９", "０"},
		{"1", "2", "3", "4", "5", "6", "7", "8", "9", "0"}}
	dm.Add("全角数字转半角", trans)

	type User struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}
	u := User{Name: "张树新", Age: 30}
	dm.Add("user1", u)
	dm.Remove("user1")
	dm.Add("user2", u)

	dm.Add("一维", []string{"这", "是", "一维数组"})

	//dm.SaveJSON("demo.json", true)
	dm.SaveJSON("demo.json")
	var ss []string
	ss = dm.List()
	sss := dm.Get("全角数字转半角")
	fmt.Println("list:", ss)
	fmt.Println("全角数字转半角:", sss)
	fmt.Println("Size:", len(dm))

}
