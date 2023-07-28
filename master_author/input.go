package master_author

type MasterAuthorInput struct {
	Name string `json:"name" binding:"required"`
}
