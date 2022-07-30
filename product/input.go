package product

type GetProductDetailInput struct {
	Slug string `uri:"slug" binding:"required"`
}

type DummyBenchmark struct {
	Benchmark string `form:"benchmark"`
}
