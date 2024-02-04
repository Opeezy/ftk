package ftk

import (
	"errors"
	"math"
)

// Trees per acre
func TPA(denominator int, plotnum int, treestotal int, round int) (n float32, err error) {
	//0 = Round down
	//1 = Round up
	//2 = No rounding
	var (
		eFactor float64 = float64(denominator) / float64(plotnum)
		tpa     float64 = eFactor * float64(treestotal)
		result  float32
	)

	if round == 0 {
		result = float32(math.Floor(tpa))
	}
	if round == 1 {
		result = float32(math.Ceil(tpa))
	}
	if round == 2 {
		result = float32(tpa)
	}
	if round != 1 && round != 0 && round != 2 {
		return 0, errors.New("Invalid input for param 'round'\nEnter 1 for round up\nEnter 2 for round down")
	}
	return result, nil
}

// Basal area
func BA(dbh float64, denominator int, plotnum int, round int) (n float32, err error) {
	//0 = Round down
	//1 = Round up
	//2 = No rounding
	const x float64 = .005454
	var (
		eFactor float64 = float64(denominator) / float64(plotnum)
		dbhSqr  float64 = dbh * dbh
		tba     float64 = dbhSqr * x
		bpa     float64 = eFactor * tba
		result  float32
	)

	if round == 0 {
		result = float32(math.Floor(bpa))
	}
	if round == 1 {
		result = float32(math.Ceil(bpa))
	}
	if round == 2 {
		result = float32(bpa)
	}
	if round != 1 && round != 0 && round != 2 {
		return 0, errors.New("Invalid input for param 'round'\nEnter 1 for round up\nEnter 2 for round down")
	}
	return result, nil
}
