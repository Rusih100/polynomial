package polynomial

import (
	"math/big"
	"testing"
)

func TestAdd(t *testing.T) {

	arr1 := []*big.Int{
		big.NewInt(2),
		big.NewInt(-2),
		big.NewInt(4),
	}

	arr2 := []*big.Int{
		big.NewInt(0),
		big.NewInt(2),
		big.NewInt(-4),
	}

	x1 := NewPolynomial(arr1)
	x2 := NewPolynomial(arr2)

	arrWant := []*big.Int{
		big.NewInt(2),
	}

	want := NewPolynomial(arrWant)

	//

	res := new(Polynomial).Add(x1, x2)

	if res.StringCoefficients() != want.StringCoefficients() {
		t.Fatal()
	}
}

func TestSub(t *testing.T) {

	arr1 := []*big.Int{
		big.NewInt(2),
		big.NewInt(-2),
		big.NewInt(4),
	}

	arr2 := []*big.Int{
		big.NewInt(0),
		big.NewInt(2),
		big.NewInt(-4),
	}

	x1 := NewPolynomial(arr1)
	x2 := NewPolynomial(arr2)

	arrWant := []*big.Int{
		big.NewInt(2),
		big.NewInt(-4),
		big.NewInt(8),
	}

	want := NewPolynomial(arrWant)

	//

	res := new(Polynomial).Sub(x1, x2)

	if res.StringCoefficients() != want.StringCoefficients() {
		t.Fatal()
	}
}

func TestMul1(t *testing.T) {

	arr1 := []*big.Int{
		big.NewInt(5),
		big.NewInt(0),
		big.NewInt(10),
		big.NewInt(6),
	}

	arr2 := []*big.Int{
		big.NewInt(1),
		big.NewInt(2),
		big.NewInt(4),
	}

	x1 := NewPolynomial(arr1)
	x2 := NewPolynomial(arr2)

	arrWant := []*big.Int{
		big.NewInt(5),
		big.NewInt(10),
		big.NewInt(30),
		big.NewInt(26),
		big.NewInt(52),
		big.NewInt(24),
	}

	want := NewPolynomial(arrWant)

	//

	res := new(Polynomial).Mul(x1, x2)

	if res.StringCoefficients() != want.StringCoefficients() {
		t.Fatal()
	}
}

func TestMul2(t *testing.T) {

	arr1 := []*big.Int{
		big.NewInt(1),
	}

	arr2 := []*big.Int{
		big.NewInt(0),
		big.NewInt(1),
	}

	x1 := NewPolynomial(arr1)
	x2 := NewPolynomial(arr2)

	arrWant := []*big.Int{
		big.NewInt(0),
		big.NewInt(1),
	}

	want := NewPolynomial(arrWant)

	//

	res := new(Polynomial).Mul(x2, x1)

	if res.StringCoefficients() != want.StringCoefficients() {
		t.Fatal()
	}
}

func TestQuoRem1(t *testing.T) {

	arr1 := []*big.Int{
		big.NewInt(1),
		big.NewInt(2),
	}

	arr2 := []*big.Int{
		big.NewInt(0),
		big.NewInt(0),
		big.NewInt(0),
		big.NewInt(1),
	}

	x1 := NewPolynomial(arr1)
	x2 := NewPolynomial(arr2)

	arrWantQuo := []*big.Int{
		big.NewInt(0),
	}

	arrWantRem := []*big.Int{
		big.NewInt(1),
		big.NewInt(2),
	}

	wantQuo := NewPolynomial(arrWantQuo)
	wantRem := NewPolynomial(arrWantRem)

	//

	quo := new(Polynomial)
	rem := new(Polynomial)
	quo, rem = quo.QuoRem(x1, x2)

	if quo.StringCoefficients() != wantQuo.StringCoefficients() && rem.StringCoefficients() != wantRem.StringCoefficients() {
		t.Fatal()
	}
}

func TestQuoRem2(t *testing.T) {

	arr1 := []*big.Int{
		big.NewInt(1),
		big.NewInt(0),
		big.NewInt(1),
	}

	arr2 := []*big.Int{
		big.NewInt(0),
		big.NewInt(0),
		big.NewInt(1),
	}

	x1 := NewPolynomial(arr1)
	x2 := NewPolynomial(arr2)

	arrWantQuo := []*big.Int{
		big.NewInt(1),
	}

	arrWantRem := []*big.Int{
		big.NewInt(1),
	}

	wantQuo := NewPolynomial(arrWantQuo)
	wantRem := NewPolynomial(arrWantRem)

	//

	quo := new(Polynomial)
	rem := new(Polynomial)
	quo, rem = quo.QuoRem(x1, x2)

	if quo.StringCoefficients() != wantQuo.StringCoefficients() && rem.StringCoefficients() != wantRem.StringCoefficients() {
		t.Fatal()
	}
}
