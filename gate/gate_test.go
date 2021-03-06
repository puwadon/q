package gate

import (
	"fmt"
	"testing"

	"github.com/itsubaki/q/matrix"
)

func TestSwap(t *testing.T) {
	swap := Swap(2)

	if swap[0][0] != 1 {
		t.Error(swap[0][0])
	}

	if swap[1][2] != 1 {
		t.Error(swap[1][2])
	}

	if swap[2][1] != 1 {
		t.Error(swap[2][1])
	}

	if swap[3][3] != 1 {
		t.Error(swap[3][3])
	}
}

func TestCNOT(t *testing.T) {
	g0 := matrix.TensorProduct(I().Add(Z()), I())
	g1 := matrix.TensorProduct(I().Sub(Z()), X())
	CN := g0.Add(g1).Mul(0.5)

	if !CNOT().Equals(CN) {
		t.Error(CN)
	}

}

func TestToffoli(t *testing.T) {
	g := make([]matrix.Matrix, 13)

	g[0] = matrix.TensorProduct(I(2), H())
	g[1] = matrix.TensorProduct(I(), CNOT())
	g[2] = matrix.TensorProduct(I(2), T().Dagger())
	g[3] = CNOTc1t3()
	g[4] = matrix.TensorProduct(I(2), T())
	g[5] = matrix.TensorProduct(I(), CNOT())
	g[6] = matrix.TensorProduct(I(2), T().Dagger())
	g[7] = CNOTc1t3()
	g[8] = matrix.TensorProduct(I(), T().Dagger(), T())
	g[9] = matrix.TensorProduct(CNOT(), H())
	g[10] = matrix.TensorProduct(I(), T().Dagger(), I())
	g[11] = matrix.TensorProduct(CNOT(), I())
	g[12] = matrix.TensorProduct(T(), S(), I())

	toffoli := I(3)
	for _, gate := range g {
		toffoli = toffoli.Apply(gate)
	}

	if !CNOT(3).Equals(toffoli, 1e-13) {
		t.Error(toffoli)
	}
}

func TestIsHermite(t *testing.T) {
	if !H().IsHermite() {
		t.Error(H())
	}

	if !X().IsHermite() {
		t.Error(X())
	}

	if !Y().IsHermite() {
		t.Error(Y())
	}

	if !Z().IsHermite() {
		t.Error(Z())
	}

}

func TestIsUnitary(t *testing.T) {
	if !H().IsUnitary(1e-13) {
		t.Error(H())
	}

	if !X().IsUnitary() {
		t.Error(X())
	}

	if !Y().IsUnitary() {
		t.Error(Y())
	}

	if !Z().IsUnitary() {
		t.Error(Z())
	}

	u := U(1, 2, 3, 4)
	if !u.IsUnitary(1e-13) {
		t.Error(u)
	}

}

func TestTrace(t *testing.T) {
	trA := I().Trace()
	if trA != complex(2, 0) {
		t.Error(trA)
	}

	trH := H().Trace()
	if trH != complex(0, 0) {
		t.Error(trH)
	}
}

func TensorProductProductXY(t *testing.T) {
	x := X()
	y := Y()

	m, n := x.Dimension()
	tmp := []matrix.Matrix{}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			tmp = append(tmp, y.Mul(x[i][j]))
		}
	}

	fmt.Printf("%v %v %v %v\n", tmp[0][0][0], tmp[0][0][1], tmp[1][0][0], tmp[1][0][1])
	fmt.Printf("%v %v %v %v\n", tmp[0][1][0], tmp[0][1][1], tmp[1][1][0], tmp[1][1][1])
	fmt.Printf("%v %v %v %v\n", tmp[2][0][0], tmp[2][0][1], tmp[3][0][0], tmp[3][0][1])
	fmt.Printf("%v %v %v %v\n", tmp[2][1][0], tmp[2][1][1], tmp[3][1][0], tmp[3][1][1])
	fmt.Println()
}

func TensorProductProductXXY(t *testing.T) {
	xx := X().TensorProduct(X())
	y := Y()

	m, n := xx.Dimension()
	tmp := []matrix.Matrix{}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			tmp = append(tmp, y.Mul(xx[i][j]))
		}
	}

	fmt.Printf("%v %v %v %v %v %v %v %v\n", tmp[0][0][0], tmp[0][0][1], tmp[1][0][0], tmp[1][0][1], tmp[2][0][0], tmp[2][0][1], tmp[3][0][0], tmp[3][0][1])
	fmt.Printf("%v %v %v %v %v %v %v %v\n", tmp[0][1][0], tmp[0][1][1], tmp[1][1][0], tmp[1][1][1], tmp[2][1][0], tmp[2][1][1], tmp[3][1][0], tmp[3][1][1])
	fmt.Printf("%v %v %v %v %v %v %v %v\n", tmp[4][0][0], tmp[4][0][1], tmp[5][0][0], tmp[5][0][1], tmp[6][0][0], tmp[6][0][1], tmp[7][0][0], tmp[7][0][1])
	fmt.Printf("%v %v %v %v %v %v %v %v\n", tmp[4][1][0], tmp[4][1][1], tmp[5][1][0], tmp[5][1][1], tmp[6][1][0], tmp[6][1][1], tmp[7][1][0], tmp[7][1][1])
	fmt.Printf("%v %v %v %v %v %v %v %v\n", tmp[8][0][0], tmp[8][0][1], tmp[9][0][0], tmp[9][0][1], tmp[10][0][0], tmp[10][0][1], tmp[11][0][0], tmp[11][0][1])
	fmt.Printf("%v %v %v %v %v %v %v %v\n", tmp[8][1][0], tmp[8][1][1], tmp[9][1][0], tmp[9][1][1], tmp[10][1][0], tmp[10][1][1], tmp[11][1][0], tmp[11][1][1])
	fmt.Printf("%v %v %v %v %v %v %v %v\n", tmp[12][0][0], tmp[12][0][1], tmp[13][0][0], tmp[13][0][1], tmp[14][0][0], tmp[14][0][1], tmp[15][0][0], tmp[15][0][1])
	fmt.Printf("%v %v %v %v %v %v %v %v\n", tmp[12][1][0], tmp[12][1][1], tmp[13][1][0], tmp[13][1][1], tmp[14][1][0], tmp[14][1][1], tmp[15][1][0], tmp[15][1][1])
	fmt.Println()
}
