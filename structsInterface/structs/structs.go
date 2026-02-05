package structs

type Product struct {
	ID    uint     `json: "id"`
	Name  string   `json: "name"`
	Type  Type     `json: "type"`
	Price float64  `json: "price"`
	Count uint     `json: "count"`
	Tags  []string `json: "tags"`
}

type Type struct {
	ID          uint   `json: "id"`
	Code        string `json: "Code"`
	Description string `json: "Description"`
}

func (p Product) TotalPrice() float64 {
	return float64(p.Count) * p.Price
}

func (p *Product) SetName(name string) {
	p.Name = name
}
func (p *Product) AddTags(tags ...string) {
	p.Tags = append(p.Tags, tags...)
}

type Car struct {
	ID       uint      `json: "id"`
	UserId   uint      `json: "user_id"`
	Products []Product `json: "products"`
}

func (c *Car) AddProducts(products ...Product) {
	c.Products = append(c.Products, products...)
}

func (c *Car) Total() float64 {
	var total float64
	for _, c := range c.Products {
		total += c.TotalPrice()
	}
	return total
}

func NewCar(UserId uint) Car {
	return Car{UserId: UserId}
}
