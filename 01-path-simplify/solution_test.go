package pathsimplify

import "testing"

func TestSimplifyPath(t *testing.T) {
	tests := []struct {
		name string
		path string
		want string
	}{
		{
			name: "exercise case",
			path: "../a/../../...././b/c/../d/../",
			want: "..../b",
		},
		{
			name: "root up has no effect",
			path: "../../../",
			want: "",
		},
		{
			name: "dots as normal name",
			path: "..../..../..",
			want: "....",
		},
		{
			name: "current dir segments ignored",
			path: "./a/././b/.",
			want: "a/b",
		},
		{
			name: "consecutive slashes collapsed",
			path: "a//b///c",
			want: "a/b/c",
		},
		{
			name: "traversal then enter",
			path: "a/b/c/../../../d",
			want: "d",
		},
		{
			name: "empty path",
			path: "",
			want: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SimplifyPath(tt.path); got != tt.want {
				t.Errorf("SimplifyPath(%q) = %q, want %q", tt.path, got, tt.want)
			}
		})
	}
}
