package main

import (
	"encoding/json"
	"fmt"
)

type OrderItem struct {
	Id    string  `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

type Order struct {
	//tag（`json:"<字段名>"`）,让go语言结构体中的字段名在json格式中转换成其他语言识别的字段名，不影响程序运行
	Id   string      `json:"id"`
	Item []OrderItem `json:"items,omitempty"`
	//omitempty当没有向结构体中的某一个字段传数据时，可以省略该字段以json数据形式传输，而不是默认该字段为空
	TotalPrice float64 ` json:"total_price"`
}

func Marshal() {
	o := Order{
		Id:         "1234",
		TotalPrice: 30,
		Item: []OrderItem{
			{
				Id:    "1",
				Name:  "lreango",
				Price: 15,
			},
			{
				Id:    "2",
				Name:  "interview",
				Price: 10},
		},
	}
	b, err := json.Marshal(o)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", b)
}

func unMarshal() {
	s := `{"id":"1234","items":[{"id":"1","name":"lreango","price":15},{"id":"2","name":"interview","price":10}],"total_price":30}`
	var o Order
	err := json.Unmarshal([]byte(s), &o)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", o)
}

func parseNLP() {
	res := `{
"data":	[
	{
		"synonym":"",
		"weight":"0.100000",
		"tag":"普通词",
		"word":"请"
	},
	{
		"synonym":"",
		"weight":"0.100000",
		"tag":"普通词",
		"word":"输入"
	},
	{
		"synonym":"",
		"weight":"1.000000",
		"tag":"品类",
		"word":"文本"
	}
]
}`
	//m := make(map[string]interface{})     使用map接收值
	//使用结构体接收，相对简单
	m := struct {
		Data []struct {
			Synonym string `json:"synonym"`
			Weight  string `json:"weight"`
			Tag     string `json:"tag"`
			Word    string `json:"word"`
		}
	}{}
	err := json.Unmarshal([]byte(res), &m)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", m.Data[1].Word)
	//fmt.Printf("%+v\n", m["data"].([]interface{})[1].(map[string]interface{})["word"])    使用map接收值
}

func main() {
	//Marshal()
	//unMarshal()
	parseNLP()
}
