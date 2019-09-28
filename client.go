package client

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"mongodb/Asset"
	"fmt"
	)

//获得所有资产的信息
func get_all(){
	url := "mongodb://localhost"
	session, err := mgo.Dial(url)
	if err != nil {
		panic(err)
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)
	c := session.DB("test").C("Assets")
	var assets []ASSET.Asset
	c.Find(nil).All(&assets)
	fmt.Println(assets)
}

//获得指定id的资产的资产信息
func get_id(id string) {
	url := "mongodb://localhost"
	session, err := mgo.Dial(url)
	if err != nil {
		panic(err)
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)
	c := session.DB("test").C("Assets")
	var assets []ASSET.Asset
	c.Find(nil).All(&assets)
	flag := 0
	for _, item := range assets {
		if item.ID == id {
			flag = 1
			var ass []ASSET.Asset
			err = c.Find(bson.M{"id":id}).All(&ass)
			fmt.Println(ass)
		}
	}
	if flag == 0 {
		fmt.Println("This ID does not exist!")
	}
	return
}

//插入资产信息
func post(temp ASSET.Asset){
	url := "mongodb://localhost"
	session, err := mgo.Dial(url)
	if err != nil {
		panic(err)
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)
	c := session.DB("test").C("Assets")
	var assets []ASSET.Asset
	flag := 0
	c.Find(nil).All(&assets)

	for _, item := range assets {			//检查是否有重复的id
		if item.ID == temp.ID {
			flag = 1
			fmt.Println("this ID is repeated!")
		}
	}
	if flag == 1 {
		return
	}

	//如果没有重复id，则开始输入其他值
	fmt.Println("please input the name of it")
	fmt.Scan(&temp.Name)
	fmt.Println("please input the value of it")
	fmt.Scan(&temp.Value)
	fmt.Println("please enter the longitude of it")
	fmt.Scan(&temp.Longitude)
	fmt.Println("please enter the latitude of it")
	fmt.Scan( &temp.Latitude)
	fmt.Println("please enter the priority of it")
	fmt.Scan( &temp.Priority)
	c.Insert(temp)
}

//删除资产信息
func delete(delete_id string){
	url := "mongodb://localhost"
	session, err := mgo.Dial(url)
	if err != nil {
		panic(err)
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)
	c := session.DB("test").C("Assets")
	var assets []ASSET.Asset
	c.Find(nil).All(&assets)
	fmt.Println("this asset has been delete")
	_, err = c.RemoveAll(bson.M{"id": delete_id})
}


//选择所需要的客户端服务
func Client(){
	var command string
	var temp ASSET.Asset
	for true {
		fmt.Println("please input the method(get/post/delete)(退出为exit)")
		fmt.Scan(&command)
		//退出程序
		if command == "exit" {
			return
		}

		//获得资产信息
		if command == "get" {
			for true {
				fmt.Println("please input the ID of zhe asset which you want to get(全部为all,退出为exit)")
				fmt.Scan(&temp.ID) //输入id/all
				if temp.ID == "exit" {
					break
				}
				if temp.ID == "all" {
					get_all()
				} else{
					get_id(temp.ID)
				}
			}
		}

		//插入
		if command == "post" {
			for true {
				fmt.Println("please input the ID of the asset which you want to post(退出为exit)")

				fmt.Scan(&temp.ID)						//输入要插入的资产的id
				if temp.ID== "exit" {
					break
				}
				post(temp)
			}
		}

		//删除
		if command == "delete" {
			for true {
				fmt.Println("please input the ID of the asset which you want to delete(退出为exit)")
				fmt.Scan(&temp.ID)
				if temp.ID == "exit" {
					break
				}
				delete(temp.ID)
			}
		}
	}
}
