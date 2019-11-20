package plugin

import (
	"io/ioutil"
	"reflect"
	"testing"

	"github.com/caddyserver/caddy"
)

func Test_parse(t *testing.T) {
	success, err := ioutil.ReadFile("testdata/success")
	if err != nil {
		t.Fatal(err)
	}

	successForest := new(forest)
	for _, domain := range []string{"0-100.com", "0-gold.net"} {
		successForest.Add(domain)
	}

	type args struct {
		c *caddy.Controller
	}
	tests := []struct {
		name    string
		args    args
		wantCfg Config
		wantErr bool
	}{
		{
			name: "success",
			args: args{c: caddy.NewTestController("dns", string(success))},
			wantCfg: Config{
				Forest:     successForest,
				ChinaDns:   "119.29.29.29:53",
				ForeignDns: "8.8.8.8:53",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotCfg, err := parse(tt.args.c)
			if (err != nil) != tt.wantErr {
				t.Errorf("parse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotCfg, tt.wantCfg) {
				t.Errorf("parse() gotCfg = %v, want %v", gotCfg, tt.wantCfg)
			}
		})
	}
}
