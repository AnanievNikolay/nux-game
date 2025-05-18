package provider

import "go.uber.org/dig"

func (p *Provider) invoke(function interface{}, opts ...dig.InvokeOption) {
	if p.err != nil {
		return
	}

	p.err = p.container.Invoke(function, opts...)
}
