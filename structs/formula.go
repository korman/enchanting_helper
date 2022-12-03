package structs

type Formula struct {
	Name              string `json:"name"`
	Wuxianzhichen     int32  `json:"wuxianzhichen"`
	Qiangxiaoyuzhou   int32  `json:"qiangxiaoyuzhou"`
	Mengjinggsuipian  int32  `json:"mengjinggsuipian"`
	Shenyuanshuijing  int32  `json:"shenyuanshuijing"`
	Yonghengzhitu     int32  `json:"yonghengzhitu"`
	Yonghengshengming int32  `json:"yonghengshengming"`
}
