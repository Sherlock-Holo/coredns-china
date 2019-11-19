package plugin

import (
	"reflect"
	"testing"
)

func Test_tree_add(testT *testing.T) {
	type fields struct {
		root *node
	}
	type args struct {
		words []string
	}
	tests := []struct {
		name       string
		fields     fields
		args       args
		expectTree *tree
	}{
		{
			name: "add [com, baidu, www]",

			fields: fields{root: &node{
				value:    "com",
				children: make(map[string]*node),
			}},

			args: args{words: []string{"com", "baidu", "www"}},

			expectTree: &tree{
				root: &node{
					value: "com",
					children: map[string]*node{
						"baidu": {
							value: "baidu",
							children: map[string]*node{
								"www": {
									value:    "www",
									children: make(map[string]*node),
								},
							},
						},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		testT.Run(tt.name, func(subTestT *testing.T) {
			t := &tree{
				root: tt.fields.root,
			}

			t.add(tt.args.words)

			if !reflect.DeepEqual(*t, *tt.expectTree) {
				subTestT.Errorf("got tree %v, expect tree %v", *t, *tt.expectTree)
			}
		})
	}
}

func Test_tree_get(t1 *testing.T) {
	root := &node{
		value: "com",
		children: map[string]*node{
			"baidu": {
				value: "baidu",
				children: map[string]*node{
					"www": {value: "www"},
				},
			},
		},
	}

	type fields struct {
		root *node
	}
	type args struct {
		words []string
	}
	tests := []struct {
		name      string
		fields    fields
		args      args
		wantMatch string
	}{
		{
			name:      "get [com, baidu, www]",
			fields:    fields{root: root},
			args:      args{words: []string{"com", "baidu", "www"}},
			wantMatch: "www.baidu.com",
		},
		{
			name:      "get [com, baidu, abc]",
			fields:    fields{root: root},
			args:      args{words: []string{"com", "baidu", "abc"}},
			wantMatch: "baidu.com",
		},
		{
			name:      "get [com, baidu, abc]",
			fields:    fields{root: root},
			args:      args{words: []string{"com", "qq", "www"}},
			wantMatch: "com",
		},
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := &tree{
				root: tt.fields.root,
			}
			if gotMatch := t.get(tt.args.words); gotMatch != tt.wantMatch {
				t1.Errorf("get() = %v, want %v", gotMatch, tt.wantMatch)
			}
		})
	}
}
