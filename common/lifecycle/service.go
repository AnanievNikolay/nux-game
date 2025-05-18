package lifecycle

import (
	"context"

	log "github.com/sirupsen/logrus"
)

type Service interface {
	Stop(ctx context.Context) error
}

type Hub struct {
	logger      *log.Entry
	onStopStack []Service
}

func NewHub(
	logger *log.Entry,
) *Hub {
	return &Hub{
		onStopStack: []Service{},
		logger:      logger.WithField("service", "hub"),
	}
}

func (h *Hub) Register(service Service) {
	h.onStopStack = append(h.onStopStack, service)
}

func (h *Hub) Stop(ctx context.Context) {
	h.logger.Infoln("stop according to stop list")
	for _, service := range h.onStopStack {
		h.logger.Infof("stop %T service", service)
		if err := service.Stop(ctx); err != nil {
			h.logger.Errorf("error while stop entity: %s", err)
		}
	}
}
