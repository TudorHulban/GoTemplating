package main

func generatorProducts(no uint) []*Product {
	result := make([]*Product, int(no))
	// TODO: logic for products generation

	p1 := &Product{
		ID:              1,
		Name:            "Jelly Bean",
		SalesPriceCents: 99,
	}

	p2 := &Product{
		ID:              1,
		Name:            "Honey",
		SalesPriceCents: 299,
	}

	result[0] = p1
	result[1] = p2

	return result
}
