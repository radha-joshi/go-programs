package insurance

type DiscountStrategy interface {
	GetDiscount() int
}

type Calculator struct {
	DiscountStrategy DiscountStrategy
}

func (c *Calculator) GetFinalAmount(amount float64) float64 {
	return amount * float64(1+c.DiscountStrategy.GetDiscount())
}
