package clrs

import (
	"reflect"
	"testing"
)

func TestBucketSort(t *testing.T) {
	cases := []struct {
		a    []float64
		want []float64
	}{
		{
			[]float64{0.78, 0.17, 0.39, 0.26, 0.72, 0.94, 0.21, 0.12, 0.23, 0.68},
			[]float64{0.12, 0.17, 0.21, 0.23, 0.26, 0.39, 0.68, 0.72, 0.78, 0.94},
		},
		{
			[]float64{0.79, 0.13, 0.16, 0.64, 0.39, 0.20, 0.89, 0.53, 0.71, 0.42},
			[]float64{0.13, 0.16, 0.20, 0.39, 0.42, 0.53, 0.64, 0.71, 0.79, 0.89},
		},
		{
			[]float64{0.79, 0.13, 0.16, 0.64, 0.39, 0.20, 0.89, 0.53, 0.71, 0.42, 0.99},
			[]float64{0.13, 0.16, 0.20, 0.39, 0.42, 0.53, 0.64, 0.71, 0.79, 0.89, 0.99},
		},
	}

	for _, c := range cases {
		in := make([]float64, len(c.a))
		copy(in, c.a)
		bucketSort(in)
		if !reflect.DeepEqual(in, c.want) {
			t.Errorf("sorted %v", c.a)
			t.Errorf("   got %v", in)
			t.Errorf("  want %v", c.want)
		}
	}
}
