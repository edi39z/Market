package app

import "golang/app/models"

type Model struct {
	Model interface{}
}

func RegisterModels() []Model {
	return []Model{
		{Model: models.Pedagang{}},
		{Model: models.JenisDagang{}},
		{Model: models.Lapak{}},
		{Model: models.Transaksi{}},
		{Model: models.Admin{}},
	}
}
