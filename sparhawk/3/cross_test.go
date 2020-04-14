package cross

import (
	"testing"
)

func TestIntersectionDistance(t *testing.T) {
	type test struct {
		wires [2]string
		want int
	}

	tests := []test {
		{
			wires: [2]string{"R75,D30,R83,U83,L12,D49,R71,U7,L72", "U62,R66,U55,R34,D71,R55,D58,R83"},
			want: 159,
		},
		{
			wires: [2]string{"R98,U47,R26,D63,R33,U87,L62,D20,R33,U53,R51", "U98,R91,D20,R16,D67,R40,U7,R15,U6,R7"},
			want: 135,
		},
	}

	for _, tc := range tests {
		t.Run("Testing Shortest Intersection Distance", func(t *testing.T) {
			got := shortestDistance(tc.wires)
			if got != tc.want {
				t.Errorf("Got %d, wanted %d", got, tc.want)
			}
		})
	}
}