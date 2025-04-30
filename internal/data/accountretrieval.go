package data

import (
	"fooservice/internal/model"
	"time"
)

type AccountRetriever struct {
}

func NewAccountRetriever() *AccountRetriever {
	return &AccountRetriever{}
}

func (a *AccountRetriever) QueryData(id int64) (model.Account, error) {
	time.Sleep(2 * time.Second)
	return model.Account{
		AccountName: "Scuba Slush Fund",
		Balance:     1500.25,
	}, nil
}
