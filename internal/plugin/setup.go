package plugin

import (
	"bufio"
	"log"
	"os"
	"time"

	"github.com/caddyserver/caddy"
	"github.com/coredns/coredns/core/dnsserver"
	corednsPlugin "github.com/coredns/coredns/plugin"
	errors "golang.org/x/xerrors"
)

func Setup(c *caddy.Controller) error {
	cfg, err := parse(c)
	if err != nil {
		return errors.Errorf("parse config failed: %w", err)
	}

	coreCfg := dnsserver.GetConfig(c)

	coreCfg.AddPlugin(func(corednsPlugin.Handler) corednsPlugin.Handler {
		return NewHandler(cfg)
	})

	return nil
}

func parse(c *caddy.Controller) (cfg Config, err error) {
	cfg = NewConfig()

	var (
		chinaListPath string
		pluginCount   int
	)

	for c.Next() {
		if pluginCount > 0 {
			return Config{}, corednsPlugin.ErrOnce
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
		err = errors.Errorf("open china-list file %s failed: %w", chinaListPath, err)
		return
	}
	defer func() {
		_ = file.Close()
	}()

	log.Println("start to parse china-list")

	start := time.Now()

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

	parseUsedTime := time.Since(start)

	if scanner.Err() != nil {
		err = errors.Errorf("scan china-list failed: %w", err)
		return
	}

	log.Printf("parse done, usage time %v", parseUsedTime)

	return
}
