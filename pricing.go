package gooanda

type pricing connection

func NewPricingConnection(token string) *pricing {
	return &pricing{token: token}
}

func (pr *pricing) GetCandlesLatestByAccount() {}

func (pr *pricing) GetPricingInformationByAccount() {}

func (pr *pricing) GetStreamingPriceByAccount() {}

func (pr *pricing) GetCandleStickByAccount() {}
