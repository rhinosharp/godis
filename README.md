# godis
## golang操作json字典的类，用最快速的json存取模块json-iterator
## 这是最小巧快速的Key-Value数据库了,是不是有点象redis啊。
## 好吧，那就叫它godis,俗称狗的屎

```
	dm := godis.NewDataMap()    //这样新建
	dm.LoadJSON("demo.json")    //载入json文件到dm这个字典中

	trans := [][]string{{"１", "２", "３", "４", "５", "６", "７", "８", "９", "０"},
		{"1", "2", "3", "4", "5", "6", "7", "8", "9", "0"}}
	dm.Add("全角数字转半角", trans)   //什么类型的变量你都可以存储进去了，只要key不相同就行了

	type User struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}
	u := User{Name: "张树新", Age: 30}
	dm.Add("user1", u)
	dm.Remove("user1")
	dm.Add("user2", u)

	dm.Add("一维", []string{"这", "是", "一维数组"})

	dm.SaveJSON("demo.json", true)   //把内存中的map变量存储到json文件中
	var ss []string
	ss = dm.List()
	sss := dm.Get("全角数字转半角")   //通过关键字取得变量
	fmt.Println("list:", ss)
	fmt.Println("全角数字转半角:", sss)
	fmt.Println("Size:", len(dm))
```
