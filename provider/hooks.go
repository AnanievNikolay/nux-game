package provider

import (
	"github.com/AnanievNikolay/nux-game/common/db"
	"github.com/AnanievNikolay/nux-game/delivery/http"
)

func (p *Provider) hooks() {
	p.invoke(func(impl db.SQLiteDB) {
		p.lifecycleHub.Register(impl)
	})

	p.invoke(func(impl *http.Delivery) {
		p.lifecycleHub.Register(impl)
	})
}
