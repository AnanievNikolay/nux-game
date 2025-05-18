package http

import "context"

func (d *Delivery) Stop(ctx context.Context) error {
	if err := d.e.Shutdown(ctx); err != nil {
		d.logger.Errorf("Server shutdown failed: %v", err)
		return err
	}
	d.logger.Info("Server shutdown success")
	return nil
}
