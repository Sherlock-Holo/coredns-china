package plugin

import (
	"github.com/caddyserver/caddy"
)

func init() {
	caddy.RegisterPlugin("china-list", caddy.Plugin{
		ServerType: "dns",
		Action:     setup,
	})
}
