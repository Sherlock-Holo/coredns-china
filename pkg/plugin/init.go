package plugin

import (
	"github.com/Sherlock-Holo/coredns-china/internal/plugin"
	"github.com/caddyserver/caddy"
)

func init() {
	caddy.RegisterPlugin("china-list", caddy.Plugin{
		ServerType: "dns",
		Action:     plugin.Setup,
	})
}
