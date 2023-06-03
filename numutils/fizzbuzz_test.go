package numutils

import "testing"

func TestFizzBuzz(t *testing.T) {
	type args struct {
		num int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "1",
			args: args{
				num: 1,
			},
			want: "1",
		},
		{
			name: "2",
			args: args{
				num: 2,
			},
			want: "2",
		},
		{
			name: "3",
			args: args{
				num: 3,
			},
			want: "Fizz",
		},
		{
			name: "5",
			args: args{
				num: 5,
			},
			want: "Buzz",
		},
		{
			name: "15",
			args: args{
				num: 15,
			},
			want: "Fizz Buzz",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FizzBuzz(tt.args.num); got != tt.want {
				t.Errorf("FizzBuzz() = %v, want %v", got, tt.want)
			}
		})
	}
}
