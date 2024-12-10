package data

type product struct {
	ID      int
	Name    string
	Type    string
	mgfdate string
}

var ProductList = []*product{
	{ID: 1, Name: "XYZ", Type: "A", mgfdate: "20/5/1996"},
	{ID: 2, Name: "XYZA", Type: "B", mgfdate: "20/5/1996"},
}

func NewProduct() []*product {
	return ProductList
}
