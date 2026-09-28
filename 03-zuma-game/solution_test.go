package zuma

import "testing"

func TestEliminate(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want string
	}{
		{
			name: "simple removal leaves remainder",
			s:    "ABBBC",
			want: "AC",
		},
		{
			name: "removal merges sides into triple",
			s:    "RRBBBR",
			want: "",
		},
		{
			name: "chain reaction",
			s:    "aabbbbaa",
			want: "",
		},
		{
			name: "chain reaction with three runs",
			s:    "aaabbbaa",
			want: "aa",
		},
		{
			name: "middle removal merges sides",
			s:    "abcccb",
			want: "abb",
		},
		{
			name: "sides merge into triple",
			s:    "ccdddc",
			want: "",
		},
		{
			name: "no removal",
			s:    "aabbcc",
			want: "aabbcc",
		},
		{
			name: "single triple",
			s:    "aaa",
			want: "",
		},
		{
			name: "single char",
			s:    "a",
			want: "a",
		},
		{
			name: "empty string",
			s:    "",
			want: "",
		},
		{
			name: "four same chars removed together",
			s:    "bbbb",
			want: "",
		},
		{
			name: "two triples separated",
			s:    "aaabbb",
			want: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Eliminate(tt.s); got != tt.want {
				t.Errorf("Eliminate(%q) = %q, want %q", tt.s, got, tt.want)
			}
		})
	}
}
