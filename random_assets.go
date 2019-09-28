package random_assets

import (
"fmt"
"gopkg.in/mgo.v2"
"gopkg.in/mgo.v2/bson"
"math/rand"
"mongodb/Asset"
)

func Random_Assets() {
	url := "mongodb://localhost"
	session, err := mgo.Dial(url)
	if err != nil {
		panic(err)
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)
	c := session.DB("test").C("Assets")
	var assets []ASSET.Asset
	var assets_A []ASSET.Asset
	var assets_B []ASSET.Asset
	var assets_C []ASSET.Asset
	var assets_D []ASSET.Asset
	var asset_A []ASSET.Asset
	var asset_B []ASSET.Asset
	var asset_C []ASSET.Asset
	var asset_D []ASSET.Asset
	c.Find(nil).All(&assets)
	c.Find(bson.M{"priority":"A"}).All(&assets_A)
	c.Find(bson.M{"priority":"B"}).All(&assets_B)
	c.Find(bson.M{"priority":"C"}).All(&assets_C)
	c.Find(bson.M{"priority":"D"}).All(&assets_D)
	fmt.Println("please input the the percent")
	var percentage int
	fmt.Scan(&percentage)

	number_A := len(assets_A)*percentage/100
	for i:=0;i<number_A;i++{
		temp:=rand.Intn(len(assets_A))
		asset_A=append(asset_A,assets_A[temp])
		assets_A = append(assets_A[:temp], assets_A[temp+1:]...)
	}

	number_B := len(assets_B)*percentage/100
	for i:=0;i<number_B;i++{
		temp:=rand.Intn(len(assets_B))
		asset_B=append(asset_B,assets_B[temp])
		assets_B = append(assets_B[:temp], assets_B[temp+1:]...)
	}

	number_C := len(assets_C)*percentage/100
	for i:=0;i<number_C;i++{
		temp:=rand.Intn(len(assets_C))
		asset_C=append(asset_C,assets_C[temp])
		assets_C= append(assets_C[:temp], assets_C[temp+1:]...)
	}

	number_D := len(assets_D)*percentage/100
	for i:=0;i<number_D;i++{
		temp:=rand.Intn(len(assets_D))
		asset_D=append(asset_D,assets_D[temp])
		assets_D = append(assets_D[:temp], assets_D[temp+1:]...)
	}

	fmt.Println("Priority A:",asset_A)
	fmt.Println("Priority B:",asset_B)
	fmt.Println("Priority C:",asset_C)
	fmt.Println("Priority D:",asset_D)
}
