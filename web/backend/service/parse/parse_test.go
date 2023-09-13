package parse

import (
	"fmt"
	"github.com/go-git/go-git/v5"
	"testing"
)

func Test_getPathNameAndVersion(t *testing.T) {
	type args struct {
		tag string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "test1",
			args: args{
				tag: "example|v1.11",
			},
			want:    "example",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := getPathName(tt.args.tag)
			if (err != nil) != tt.wantErr {
				t.Errorf("getPathNameAndVersion() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("getPathNameAndVersion() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_parse_Parse(t *testing.T) {
	p := parse{}

	repo, err := git.PlainOpen("/Users/nx/GolandProjects/OSPP2023/envoy-golang-filter-hub")
	if err != nil {
		fmt.Println("Error opening repository:", err)
		return
	}

	metadata, err := p.Parse(repo)
	if err != nil {
		panic(err)
		return
	}

	fmt.Printf("%+v\n", metadata)
}
