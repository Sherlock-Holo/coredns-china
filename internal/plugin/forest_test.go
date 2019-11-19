package plugin

import (
	"reflect"
	"testing"
)

func Test_forest_Add(t *testing.T) {
	type args struct {
		domains []string
	}
	tests := []struct {
		name       string
		args       args
		wantForest forest
	}{
		{
			name: "add www.baidu.com",
			args: args{
				domains: []string{
					"www.baidu.com",
				},
			},
			wantForest: forest{
				trees: []*tree{
					{
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
			},
		},
		{
			name: "add www.baidu.com, baidu.com",
			args: args{
				domains: []string{
					"www.baidu.com",
					"baidu.com",
				},
			},
			wantForest: forest{
				trees: []*tree{
					{
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
			},
		},
		{
			name: "add www.baidu.com, abc.baidu.com",
			args: args{
				domains: []string{
					"www.baidu.com",
					"abc.baidu.com",
				},
			},
			wantForest: forest{
				trees: []*tree{
					{
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
										"abc": {
											value:    "abc",
											children: make(map[string]*node),
										},
									},
								},
							},
						},
					},
				},
			},
		},
		{
			name: "add www.baidu.com, www.abc.baidu.com",
			args: args{
				domains: []string{
					"www.baidu.com",
					"www.abc.baidu.com",
				},
			},
			wantForest: forest{
				trees: []*tree{
					{
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
										"abc": {
											value: "abc",
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
					},
				},
			},
		},
		{
			name: "add www.baidu.com, www.baidu.cn",
			args: args{
				domains: []string{
					"www.baidu.com",
					"www.baidu.cn",
				},
			},
			wantForest: forest{
				trees: []*tree{
					{
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
					{
						root: &node{
							value: "cn",
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
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var f forest

			for _, domain := range tt.args.domains {
				f.Add(domain)
			}

			if !reflect.DeepEqual(f, tt.wantForest) {
				t.Errorf("got forest %v, want forest %v", f, tt.wantForest)
			}
		})
	}
}

func Test_forest_Get(t *testing.T) {
	type fields struct {
		trees []*tree
	}
	type args struct {
		domain string
	}
	tests := []struct {
		name         string
		fields       fields
		args         args
		wantMatch    string
		wantRootWord string
	}{
		{
			name: "one tree get www.baidu.com",
			fields: fields{
				trees: []*tree{
					{
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
			},
			args:         args{domain: "www.baidu.com"},
			wantMatch:    "www.baidu.com",
			wantRootWord: "com",
		},
		{
			name: "one tree get baidu.com",
			fields: fields{
				trees: []*tree{
					{
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
			},
			args:         args{domain: "baidu.com"},
			wantMatch:    "baidu.com",
			wantRootWord: "com",
		},
		{
			name:         "no tree get www.baidu.com",
			args:         args{domain: "www.baidu.com"},
			wantMatch:    "",
			wantRootWord: "",
		},
		{
			name: "one tree get www.baidu.cn",
			fields: fields{
				trees: []*tree{
					{
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
			},
			args:         args{domain: "www.baidu.cn"},
			wantMatch:    "",
			wantRootWord: "",
		},
		{
			name: "multi trees get www.baidu.com",
			fields: fields{
				trees: []*tree{
					{
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
					{
						root: &node{
							value: "cn",
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
			},
			args:         args{domain: "www.baidu.com"},
			wantMatch:    "www.baidu.com",
			wantRootWord: "com",
		},
		{
			name: "multi trees get www.baidu.org",
			fields: fields{
				trees: []*tree{
					{
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
					{
						root: &node{
							value: "cn",
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
			},
			args:         args{domain: "www.baidu.org"},
			wantMatch:    "",
			wantRootWord: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := &forest{
				trees: tt.fields.trees,
			}
			gotMatch, gotRootWord := f.Get(tt.args.domain)
			if gotMatch != tt.wantMatch {
				t.Errorf("Get() gotMatch = %v, want %v", gotMatch, tt.wantMatch)
			}
			if gotRootWord != tt.wantRootWord {
				t.Errorf("Get() gotRootWord = %v, want %v", gotRootWord, tt.wantRootWord)
			}
		})
	}
}
