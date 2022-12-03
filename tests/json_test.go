package tests

import (
	"encoding/json"
	"io/ioutil"
	"testing"
)

type Price struct {
	Wuxianzhichen     float32 `json:"wuxianzhichen"`
	Qiangxiaoyuzhou   float32 `json:"qiangxiaoyuzhou"`
	Mengjinggsuipian  float32 `json:"mengjinggsuipian"`
	Shenyuanshuijing  float32 `json:"shenyuanshuijing"`
	Yonghengzhitu     float32 `json:"yonghengzhitu"`
	Yonghengshengming float32 `json:"yonghengshengming"`
}

func TestLoadJson(t *testing.T) {
	f, err := ioutil.ReadFile("../data/price.json")

	if err != nil {
		t.Fatal(err)
	}

	p := &Price{}

	err = json.Unmarshal(f, p)

	if err != nil {
		t.Fatal(err)
	}
}
