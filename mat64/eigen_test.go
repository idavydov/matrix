// Copyright ©2013 The gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package mat64

import (
	"math"
	"reflect"
	"testing"
)

func TestEigen(t *testing.T) {
	for i, test := range []struct {
		a *Dense

		epsilon float64

		e, d []float64
		v    *Dense
	}{
		{
			a: NewDense(3, 3, []float64{
				1, 2, 1,
				6, -1, 0,
				-1, -2, -1,
			}),

			epsilon: math.Pow(2, -52.0),

			d: []float64{3.0000000000000044, -4.000000000000003, -1.0980273383714707e-16},
			e: []float64{0, 0, 0},
			v: NewDense(3, 3, []float64{
				-0.48507125007266627, 0.41649656391752204, 0.11785113019775795,
				-0.7276068751089995, -0.8329931278350428, 0.7071067811865481,
				0.48507125007266627, -0.4164965639175216, -1.5320646925708532,
			}),
		},
		{
			a: NewDense(3, 3, []float64{
				1, 6, -1,
				6, -1, -2,
				-1, -2, -1,
			}),

			epsilon: math.Pow(2, -52.0),

			d: []float64{-6.240753470718579, -1.3995889142010132, 6.640342384919599},
			e: []float64{0, 0, 0},
			v: NewDense(3, 3, []float64{
				-0.6134279348516111, -0.31411097261113, -0.7245967607083111,
				0.7697297716508223, -0.03251534945303795, -0.6375412384185983,
				0.17669818159240022, -0.9488293044247931, 0.2617263908869383,
			}),
		},
		{ // Jama pvals
			a: NewDense(3, 3, []float64{
				4, 1, 1,
				1, 2, 3,
				1, 3, 6,
			}),

			epsilon: math.Pow(2, -52.0),
		},
		{ // Jama evals
			a: NewDense(4, 4, []float64{
				0, 1, 0, 0,
				1, 0, 2e-7, 0,
				0, -2e-7, 0, 1,
				0, 0, 1, 0,
			}),

			epsilon: math.Pow(2, -52.0),
		},
		{ // Jama badeigs
			a: NewDense(5, 5, []float64{
				0, 0, 0, 0, 0,
				0, 0, 0, 0, 1,
				0, 0, 0, 1, 0,
				1, 1, 0, 0, 1,
				1, 0, 1, 0, 1,
			}),

			epsilon: math.Pow(2, -52.0),
		},
	} {
		ef := eigen(DenseCopyOf(test.a), test.epsilon)
		if test.d != nil {
			if !reflect.DeepEqual(ef.d, test.d) {
				t.Errorf("unexpected d for test %d", i)
			}
		}
		if test.e != nil {
			if !reflect.DeepEqual(ef.e, test.e) {
				t.Errorf("unexpected e for test %d", i)
			}
		}

		if test.v != nil {
			if !Equal(ef.V, test.v) {
				t.Errorf("unexpected v for test %d", i)
			}
		}

		test.a.Mul(test.a, ef.V)
		ef.V.Mul(ef.V, ef.D())
		if !EqualApprox(test.a, ef.V, 1e-12) {
			t.Errorf("unexpected factor product for test %d", i)
		}
	}
}
