package main

import (
	"fmt"
	"mongodb/client_select"
	"random_assets"
)

func main() {
	for true {
		// 输入需要选择的功能（生成随机资产组/查看更改数据库内容）

		fmt.Println("please choose the features which do you want to use")
		fmt.Println("select_client please enter 1, get random assets please enter 2, exit please enter 0")
		var choice int
		fmt.Scan(&choice)
		if choice == 0{
			break
		}else  if choice == 1{
			select_client.Select_client()			//进行数据库查看修改功能的客户端选择
		}else if choice == 2{
			random_assets.Random_Assets()			//生成随机资产组
		}
	}
}
