package hello

import (
	"github.com/roadrunner-server/errors"
	"go.uber.org/zap"
)

const PluginName string = "hello"

type Plugin struct {
	cfg *Config
	log *zap.Logger
}

func (p *Plugin) Name() string {
	return PluginName
}

func (p *Plugin) Init(logger Logger, cfg Configurer) error {
	const op = errors.Op("hello_plugin_init")
	if !cfg.Has(PluginName) {
		return errors.E(op, errors.Disabled)
	}

	p.cfg = &Config{}
	err := cfg.UnmarshalKey(PluginName, &p.cfg)
	if err != nil {
		return errors.E(op, err)
	}

	p.cfg.InitDefaults()
	p.log = logger.NamedLogger(PluginName)

	return nil
}

func (p *Plugin) Serve() chan error {
	errCh := make(chan error, 1)

	return errCh
}

func (p *Plugin) Stop() error {
	return nil
}

func (p *Plugin) RPC() any {
	return &rpc{p: p}
}
