package plugin

import (
	"bufio"
	"os"

	"github.com/Sherlock-Holo/coredns-china/internal/plugin"
	"github.com/Sherlock-Holo/errors"
	"github.com/caddyserver/caddy"
	"github.com/coredns/coredns/core/dnsserver"
	corednsPlugin "github.com/coredns/coredns/plugin"
)

func setup(c *caddy.Controller) error {
	cfg, err := parse(c)
	if err != nil {
		return errors.WithMessage(err, "parse config failed")
	}

	coreCfg := dnsserver.GetConfig(c)

	coreCfg.AddPlugin(func(corednsPlugin.Handler) corednsPlugin.Handler {
		return plugin.NewHandler(cfg)
	})

	return nil
}

func parse(c *caddy.Controller) (cfg plugin.Config, err error) {
	cfg = plugin.NewConfig()

	var (
		chinaListPath string
		pluginCount   int
	)

	for c.Next() {
		if pluginCount > 0 {
			return plugin.Config{}, corednsPlugin.ErrOnce
		}
		pluginCount++

		for c.NextBlock() {
			switch c.Val() {
			case "list-file":
				args := c.RemainingArgs()
				if len(args) != 1 {
					err = c.ArgErr()
					return
				}

				chinaListPath = args[0]

			case "china":
				args := c.RemainingArgs()
				if len(args) != 1 {
					err = c.ArgErr()
					return
				}

				cfg.ChinaDns = args[0]

			case "foreign":
				args := c.RemainingArgs()
				if len(args) != 1 {
					err = c.ArgErr()
					return
				}

				cfg.ForeignDns = args[0]
			}
		}
	}

	file, err := os.Open(chinaListPath)
	if err != nil {
		err = errors.Wrapf(err, "open china-list file %s failed", chinaListPath)
		return
	}
	defer func() {
		_ = file.Close()
	}()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := scanner.Text()

		switch text {
		case "":
			// ignore empty line
			continue

		default:
			cfg.Forest.Add(text)
		}
	}

	if scanner.Err() != nil {
		err = errors.Wrap(err, "scan china-list failed")
		return
	}

	return
}
