package provider

import "go.uber.org/dig"

func (p *Provider) provide(constructor interface{}, opts ...dig.ProvideOption) {
	if p.err != nil {
		return
	}

	p.err = p.container.Provide(constructor, opts...)
}
