// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package math

// This table might overflow 127-bit exponent representations.
// In that case, truncate it after 1.0e38.
// 此表可能从127位的指数表示中溢出，此时就会从 1.0e38 之后截断。
var pow10tab [70]float64

// Pow10 returns 10**e, the base-10 exponential of e.
//
// Special cases are:
//	Pow10(e) = +Inf for e > 309
//	Pow10(e) = 0 for e < -324

// Pow10 返回 10**e，即以 10 为底的 e 次幂。
//
// 特殊情况为：
//	对于 e >  309，有 Pow10(e) = +Inf
//	对于 e < -324，有 Pow10(e) = 0
func Pow10(e int) float64 {
	if e <= -325 {
		return 0
	} else if e > 309 {
		return Inf(1)
	}

	if e < 0 {
		return 1 / Pow10(-e)
	}
	if e < len(pow10tab) {
		return pow10tab[e]
	}
	m := e / 2
	return Pow10(m) * Pow10(e-m)
}

func init() {
	pow10tab[0] = 1.0e0
	pow10tab[1] = 1.0e1
	for i := 2; i < len(pow10tab); i++ {
		m := i / 2
		pow10tab[i] = pow10tab[m] * pow10tab[i-m]
	}
}
