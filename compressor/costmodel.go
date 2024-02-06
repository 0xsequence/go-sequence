package compressor

import "strings"

type CostModel struct {
	ZeroByteCost int
	OneByteCost  int

	FixedCost          int
	OverHeadPersentage int // 0 - 1000
}

func NewCostModel(zeroByteCost, oneByteCost, fixedCost, overHeadPersentage int) *CostModel {
	return &CostModel{
		ZeroByteCost:       zeroByteCost,
		OneByteCost:        oneByteCost,
		FixedCost:          fixedCost,
		OverHeadPersentage: overHeadPersentage,
	}
}

func GetCostModel(network string) *CostModel {
	// Convert to lowerscase, trim spaces
	n := strings.ToLower(strings.TrimSpace(network))

	switch n {
	case "mainnet":
		return NewCostModel(4, 16, 2300, 5)
	default:
		return GetCostModel("mainnet")
	}
}

func (cm *CostModel) Cost(data []byte) int {
	var calldataCost int

	for _, b := range data {
		if b == 0 {
			calldataCost += cm.ZeroByteCost
		} else {
			calldataCost += cm.OneByteCost
		}
	}

	return calldataCost
}
