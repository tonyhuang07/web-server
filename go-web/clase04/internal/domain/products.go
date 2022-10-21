package domain

type Product struct {
	ID        int     `json:"id" `
	Name      string  `json:"name" `
	Price     float64 `json:"price" `
	Quality   int     `json:"quality"`
	Published bool    `json:"published"`
}
