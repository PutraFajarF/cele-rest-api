package master_book

type MasterBookInput struct {
	Name     string `json:"name" binding:"required"`
	Amount   int    `json:"amount" binding:"required"`
	Price    int    `json:"price" binding:"required"`
	AuthorID uint   `json:"author_id" binding:"required"`
}
