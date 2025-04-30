package data

import (
	"fooservice/internal/model"
	"math/rand"
	"time"
)

type AccountRetrieverImp struct {
}

func NewAccountRetrieverImp() *AccountRetrieverImp {
	return &AccountRetrieverImp{}
}

func (a *AccountRetrieverImp) QueryData(id int64) (model.Account, error) {
	randomDuration := time.Duration(rand.Intn(5001)) * time.Millisecond
	time.Sleep(randomDuration)
	return model.Account{
		AccountName: "Scuba Slush Fund",
		Balance:     1500.25,
	}, nil
}
