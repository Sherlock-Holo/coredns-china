package china_list

import (
	"reflect"
	"testing"
)

func TestParse(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name    string
		args    args
		want    []Domain
		wantErr bool
	}{
		{
			name: "success",
			args: args{path: "test-cn-file"},
			want: []Domain{
				"0-6.com",
				"0-gold.net",
			},
		},
		{
			name:    "file not exist",
			args:    args{path: "not-exist-file"},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Parse(tt.args.path)
			if (err != nil) != tt.wantErr {
				t.Errorf("Parse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Parse() got = %v, want %v", got, tt.want)
			}
		})
	}
}
