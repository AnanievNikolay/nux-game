package http

import (
	"errors"
	"fmt"
	"net/http"
)

func (d *Delivery) Start() error {
	d.logger.Infof("Listening on http://%s", d.addr)

	if err := d.e.Start(d.addr); !errors.Is(err, http.ErrServerClosed) {
		err = fmt.Errorf("error while StartServer: %w", err)
		d.logger.Error(err)

		return err
	}

	return nil
}
