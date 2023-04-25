package hello

type rpc struct {
	p *Plugin
}

func (r *rpc) Hello(input string, output *string) error {
	r.p.log.Debug("Inside rpc Hello method")
	*output = r.p.cfg.prefix + input

	return nil
}
