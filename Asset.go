package ASSET

//定义固定资产的属性
type Asset struct {
	ID     string   `json:"id,omitempty"`				//ID
	Name string   `json:"name,omitempty"`				//名字
	Value  string   `json:"value,omitempty"`			//价值
	Longitude string `json:"longitude,omitempty"`		//横坐标
	Latitude string `json:"latitude,omitempty"`			//纵坐标
	Priority string `json:"priority,omitempty"`			//优先级
}
