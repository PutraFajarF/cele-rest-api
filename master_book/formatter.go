package master_book

import "project-rest-api/entities"

type MasterBookFormatter struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Amount int    `json:"amount"`
	Price  int    `json:"price"`
}

func FormatMasterBook(masterBook entities.MasterBook) MasterBookFormatter {
	formatter := MasterBookFormatter{
		ID:     int(masterBook.ID),
		Name:   masterBook.Name,
		Amount: masterBook.Amount,
		Price:  masterBook.Price,
	}

	return formatter
}
