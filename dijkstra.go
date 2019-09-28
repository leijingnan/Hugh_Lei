package dijkstra

import (
	"fmt"
	"mongodb/Asset"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"math/rand"
	"strconv"
)

func make(x float64) float64 {
	if x>=0{
		return x
	}else {
		return float64(-x)
	}
}

func Dijkstra(assets []ASSET.Asset,asset ASSET.Asset) {
	rand.Seed(100000)
	url := "mongodb://localhost"
	session, err := mgo.Dial(url)
	if err != nil {
		panic(err)
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)
	c := session.DB("test").C("Assets")
	var assets []ASSET.Asset
	var asset []ASSET.Asset
	c.Find(nil).All(&assets)
	a:=rand.Intn(len(assets))
	ID := strconv.Itoa(a)
	err = c.Find(bson.M{"id":ID}).All(&asset)
	dijkstra.Dijkstra(assets,asset[0])
	var location [2][1000]float64
	var graph [1000][1000]float64
	var v [1000][1000]float64
	temp_1, err := strconv.ParseFloat(asset.Longitude, 32)
	temp_2, err := strconv.ParseFloat(asset.Latitude, 32)
	if err!=nil{	}
	location[0][0]=temp_1
	location[1][0]=temp_2
	fmt.Println(temp_1,temp_2)
	fmt.Println(assets)
	num:=1
	for _,item := range assets {
			temp_1, err := strconv.ParseFloat(item.Longitude, 32)
			temp_2, err := strconv.ParseFloat(item.Latitude, 32)
			if err!=nil{	}
			if temp_1 != location[0][0]&&temp_2 != location[1][0]{
			location[0][num]=temp_1
			location[1][num]=temp_2
			num++
		}
	}
	for i:=0;i<len(assets);i++{
			fmt.Println(location[0][i],location[1][i])
	}

	for i:=0;i<len(assets);i++{
		for j:=0;j<len(assets);j++  {
			v[i][j]=make((location[0][i]-location[0][j]))+make((location[1][i]-location[1][j]))
		}
	}


	for i:=0;i<len(assets);i++{
		for j:=0;j<len(assets);j++  {
			fmt.Print(v[i][j]," ")
		}
		fmt.Println()
	}


	for i:=0;i<len(assets);i++{
		for j:=0;j<len(assets);j++  {
			graph[i][j]=v[i][j]
		}
	}


	var TablePathMin float64       	//存放shortTablePath中,未遍历的最小结点的值
	var Vx int                 		//存放shortTablePath中,未遍历的最小结点的下标
	var isgetPath [1000]bool 		//记录结点是否已经找到v0到vx的最小路径
	var shortTablePath []float64
	// 获取v0这一行的权值数组
	for v := 0; v <len(assets); v++ {
		shortTablePath = append(shortTablePath,graph[0][v])
	}
	shortTablePath[0] = 0
	isgetPath[0] = true

	//遍历v1 ~ v8
	for v := 1; v <len(assets); v++ {
		TablePathMin = 10000

		//找出shortTablePath中,未遍历的最小结点的值
		for w := 0; w <len(assets); w++ {
			if !isgetPath[w] && shortTablePath[w] < TablePathMin {
				Vx = w
				TablePathMin = shortTablePath[w]
			}
		}
		isgetPath[Vx] = true
		for j := 0; j < len(assets); j++ {
			if !isgetPath[j] && TablePathMin+graph[Vx][j] < shortTablePath[j] {
				shortTablePath[j] = TablePathMin + graph[Vx][j]
			}
		}

		fmt.Println("遍历完V", v, "后:", shortTablePath)

	}
	//输出
	for i := 1; i < len(assets); i++ {
		fmt.Println("V0到V", i, "最小路径:", shortTablePath[i])
	}

}


