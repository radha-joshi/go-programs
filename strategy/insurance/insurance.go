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

type SimpleDiscount struct{}

func (d *SimpleDiscount) GetDiscount() int {
	return 10
}

type ComplexDiscount struct{}

func (d *ComplexDiscount) GetDiscount() int {
	return 20
}
