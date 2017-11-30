package cnpj

type CNPJ struct {
	digits []int
}

func (c *CNPJ) Generate() *CNPJ {
	return c
}
