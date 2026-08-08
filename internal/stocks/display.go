package stocks

type Display struct {
	Stocks []DisplayStock
	Crypto []DisplayCoin
}

type DisplayStock struct {
	Symbol        string
	MarketValue   float64
	CostBasis     float64
	UnrealizedPPC float64
	ChangeToday   float64
	Reserve       float64
}

type DisplayCoin struct {
	Currency     string
	Qty          float64
	DollarAmount float64
	Reserve      float64
}

const scaleTotal = 1000

func (s *Summary) Display() Display {
	var stockTotal, cryptoTotal float64
	for _, v := range s.Alpaca {
		stockTotal += v.MarketValue
	}
	for _, c := range s.Crypto {
		cryptoTotal += c.DollarAmount
	}
	ks, kc := factor(stockTotal), factor(cryptoTotal)

	d := Display{
		Stocks: make([]DisplayStock, 0, len(s.Alpaca)),
		Crypto: make([]DisplayCoin, 0, len(s.Crypto)),
	}

	for _, v := range s.Alpaca {
		d.Stocks = append(d.Stocks, DisplayStock{
			Symbol:        v.Symbol,
			MarketValue:   v.MarketValue * ks,
			CostBasis:     v.CostBasis * ks,
			UnrealizedPPC: v.UnrealizedPPC,
			ChangeToday:   v.ChangeToday,
			Reserve:       v.Reserve * ks,
		})
	}

	for _, c := range s.Crypto {
		d.Crypto = append(d.Crypto, DisplayCoin{
			Currency:     c.Currency,
			Qty:          c.Qty * kc,
			DollarAmount: c.DollarAmount * kc,
			Reserve:      c.Reserve * kc,
		})
	}

	return d
}

func factor(total float64) float64 {
	if total <= 0 {
		return 1
	}
	return scaleTotal / total
}
