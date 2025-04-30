package data

import (
	"fooservice/internal/model"
	"fooservice/internal/server"
	"github.com/timpellison/flaggy/client"
	"log/slog"
)

type AccountService struct {
	flag       string
	flagClient client.FlaggyClient
	logger     *slog.Logger
}

func NewAccountService(flag string, flagClient client.FlaggyClient) *AccountService {
	return &AccountService{
		flag:       flag,
		flagClient: flagClient,
		logger:     slog.Default(),
	}
}

func (s *AccountService) QueryData(id int64) (model.Account, error) {
	var d server.DataService
	fl, e := s.flagClient.GetFlag(s.flag)
	if e != nil || !fl.Value {
		s.logger.Info("Use account retriever.")
		d = NewAccountRetriever()
	} else {
		s.logger.Info("Use NEW! IMPROVED! account retriever.")
		d = NewAccountRetrieverImp()
	}
	return d.QueryData(id)
}
