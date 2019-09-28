package  csv_reading

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"gopkg.in/mgo.v2"
	"bufio"
	"mongodb/Asset"
)

//文件读取
func CSV_READ(){
	for true {
		fmt.Println("please enter the address of the csv file(退出为exit)")
		inputReader := bufio.NewReader(os.Stdin)
		way, err := inputReader.ReadString('\n')
		if way == "exit"{
			return
		}
		file, err := os.Open(way[:len(way)-1])				//打开文件的位置
		if err != nil {							//检查地址是否有效，无效则退出
			fmt.Println("Error:", err)
			continue
		}
		defer file.Close()
		reader := csv.NewReader(file)
		url := "mongodb://localhost"
		session, err := mgo.Dial(url)
		if err != nil {
			panic(err)
		}
		defer session.Close()
		session.SetMode(mgo.Monotonic, true)
		c := session.DB("test").C("Assets")
		flag_read := 1							//设置一个变量，检查该文件是否被读取正常
		for {
			record, err := reader.Read()
			if err == io.EOF {
				break
			} else if err != nil {
				flag_read=0
				fmt.Println("记录集错误:", err)
				break
			}
			var temp ASSET.Asset
			var assets []ASSET.Asset
			c.Find(nil).All(&assets)
			flag_repeated := 0					//设置一个变量，检查该资产ID是否在数据库中存在
			for _, ass := range assets {
				if ass.ID == record[0] {
					flag_repeated = 1
				}
			}
			if flag_repeated == 1 {
				fmt.Println("the id \"" + record[0] + "\" is repeated")		//如果该资产ID已存在，则不存储并报错
			} else {																//如果该资产ID不存在，将该资产信息存入数据库
				temp.ID = record[0]
				temp.Name = record[1]
				temp.Value = record[2]
				temp.Longitude = record[3]
				temp.Latitude = record[4]
				temp.Priority = record[5]
				c.Insert(temp)
			}
		}
		if flag_read == 1{															//如果该文件读取正常，显示文件读取完毕
			fmt.Println("this csv file has being read")
		}
	}
}
