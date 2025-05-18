package provider

import "context"

func (p *Provider) Stop(ctx context.Context) {
	p.lifecycleHub.Stop(ctx)
}
