package structs

type ShowTable struct {
	Name         string `table:"附魔名"`
	MaterialCost string `table:"材料成本"`
	CustodyFee   string `table:"保管费"`
	AuctionFee   string `table:"拍卖费"`
	Total        string `table:"总成本价"`
}
