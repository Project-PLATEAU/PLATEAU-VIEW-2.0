package indexer

import (
	"gonum.org/v1/gonum/mat"
)

func eyeMat(n int) *mat.Dense {
	d := make([]float64, n*n)
	for i := 0; i < n*n; i += n + 1 {
		d[i] = 1
	}
	return mat.NewDense(n, n, d)
}

func mat4FromGltfNodeMatrix(matrix [16]float32) *mat.Dense {
	d := []float64{}
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			d = append(d, float64(matrix[i+4*j]))
		}
	}
	return mat.NewDense(4, 4, d)
}

func mat4FromRotationTranslation(rotation *mat.Dense, translation *Cartesian3) *mat.Dense {
	if translation == nil {
		translation = &Cartesian3{
			X: 0.0,
			Y: 0.0,
			Z: 0.0,
		}
	}

	res := eyeMat(4)
	res.Set(0, 0, rotation.At(0, 0))
	res.Set(0, 1, rotation.At(0, 1))
	res.Set(0, 2, rotation.At(0, 2))
	res.Set(0, 3, 0.0)
	res.Set(1, 0, rotation.At(1, 0))
	res.Set(1, 1, rotation.At(1, 1))
	res.Set(1, 2, rotation.At(1, 2))
	res.Set(1, 3, 0.0)
	res.Set(2, 0, rotation.At(2, 0))
	res.Set(2, 1, rotation.At(2, 1))
	res.Set(2, 2, rotation.At(2, 2))
	res.Set(2, 3, 0.0)
	res.Set(3, 0, translation.X)
	res.Set(3, 1, translation.Y)
	res.Set(3, 2, translation.Z)
	res.Set(3, 3, 1.0)

	return res
}

func mat4FromCartesian(cs *Cartesian3) *mat.Dense {
	return mat4FromRotationTranslation(eyeMat(3), cs)
}

func mat4MultiplyTransformation(left *mat.Dense, right *mat.Dense) *mat.Dense {
	left0 := left.At(0, 0)
	left1 := left.At(0, 1)
	left2 := left.At(0, 2)
	left4 := left.At(1, 0)
	left5 := left.At(1, 1)
	left6 := left.At(1, 2)
	left8 := left.At(2, 0)
	left9 := left.At(2, 1)
	left10 := left.At(2, 2)
	left12 := left.At(3, 0)
	left13 := left.At(3, 1)
	left14 := left.At(3, 2)

	right0 := right.At(0, 0)
	right1 := right.At(0, 1)
	right2 := right.At(0, 2)
	right4 := right.At(1, 0)
	right5 := right.At(1, 1)
	right6 := right.At(1, 2)
	right8 := right.At(2, 0)
	right9 := right.At(2, 1)
	right10 := right.At(2, 2)
	right12 := right.At(3, 0)
	right13 := right.At(3, 1)
	right14 := right.At(3, 2)

	column0Row0 := left0*right0 + left4*right1 + left8*right2
	column0Row1 := left1*right0 + left5*right1 + left9*right2
	column0Row2 := left2*right0 + left6*right1 + left10*right2

	column1Row0 := left0*right4 + left4*right5 + left8*right6
	column1Row1 := left1*right4 + left5*right5 + left9*right6
	column1Row2 := left2*right4 + left6*right5 + left10*right6

	column2Row0 := left0*right8 + left4*right9 + left8*right10
	column2Row1 := left1*right8 + left5*right9 + left9*right10
	column2Row2 := left2*right8 + left6*right9 + left10*right10

	column3Row0 :=
		left0*right12 + left4*right13 + left8*right14 + left12
	column3Row1 :=
		left1*right12 + left5*right13 + left9*right14 + left13
	column3Row2 :=
		left2*right12 + left6*right13 + left10*right14 + left14

	res := eyeMat(4)
	res.Set(0, 0, column0Row0)
	res.Set(0, 1, column0Row1)
	res.Set(0, 2, column0Row2)
	res.Set(0, 3, 0.0)
	res.Set(1, 0, column1Row0)
	res.Set(1, 1, column1Row1)
	res.Set(1, 2, column1Row2)
	res.Set(1, 3, 0.0)
	res.Set(2, 0, column2Row0)
	res.Set(2, 1, column2Row1)
	res.Set(2, 2, column2Row2)
	res.Set(2, 3, 0.0)
	res.Set(3, 0, column3Row0)
	res.Set(3, 1, column3Row1)
	res.Set(3, 2, column3Row2)
	res.Set(3, 3, 1.0)

	return res
}
