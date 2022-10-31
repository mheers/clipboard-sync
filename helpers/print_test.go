package helpers

import "testing"

func TestMaskString(t *testing.T) {
	type args struct {
		s     string
		mask  rune
		start int
		end   int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "test mask string",
			args: args{
				s:     "1234567890",
				mask:  '*',
				start: 3,
				end:   6,
			},
			want: "123***7890",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MaskString(tt.args.s, tt.args.mask, tt.args.start, tt.args.end); got != tt.want {
				t.Errorf("MaskString() = %v, want %v", got, tt.want)
			}
		})
	}
}
