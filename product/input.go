package product

type GetProductDetailInput struct {
	Slug string `uri:"slug" binding:"required"`
}
