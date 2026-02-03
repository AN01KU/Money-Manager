package handlers

import "github.com/AN01KU/money-manager/internal/tools"

type Handlers struct {
	DB *tools.DatabaseInterface
}

func NewHandler(db *tools.DatabaseInterface) *Handlers {
	return &Handlers{DB: db}
}
