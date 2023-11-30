package stan

import (
	"encoding/json"
	"time"

	"github.com/ell1jah/show_order/internal/config"
	"github.com/ell1jah/show_order/internal/logic"
	"github.com/ell1jah/show_order/internal/model"
	"github.com/nats-io/stan.go"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

type StanManager interface {
	Run() error
	Stop() error
}

type stanManager struct {
	logic  logic.Logic
	conn   stan.Conn
	logger zap.Logger
	cfg    *config.Config
}

func NewStanManager(logic logic.Logic, logger *zap.Logger, cfg *config.Config) StanManager {
	return &stanManager{
		logic:  logic,
		logger: logger,
		cfg:    cfg,
	}
}

func (sm *stanManager) Run() error {
	conn, err := stan.Connect(
		sm.cfg.Stan.ClusterID, sm.cfg.Stan.ClientID,
		stan.NatsURL(sm.cfg.Stan.URL),
	)
	if err != nil {
		return errors.Wrap(err, "stan connect error")
	}

	sm.conn = conn

	_, err = conn.Subscribe(
		sm.cfg.Stan.Subject, func(m *stan.Msg) {
			var order model.Order
			err := json.Unmarshal(m.Data, &order)
			if err != nil {
				sm.logger.Error(err.Error())
				return
			}
			err = sm.logic.Create(&order)
			if err != nil {
				sm.logger.Error(err.Error())
				return
			}
		}, stan.StartAtTimeDelta(time.Minute*time.Duration(sm.cfg.Stan.StartDeltaMin)),
		stan.DurableName(sm.cfg.Stan.Durable),
	)
	if err != nil {
		return errors.Wrap(err, "stan subscribe error")
	}

	return nil
}

func (sm *stanManager) Stop() error {
	return sm.conn.Close()
}
