package model

import (
	"slr/data"
	"math"
	"fmt"
)

func NewSimpleLinearRegression(ds *data.DataSet, inputColumn, outputColumn string) *SimpleLinearRegression {
	model := &SimpleLinearRegression{
		dataSet: ds,
		inputColumn: inputColumn,
		outputColumn: outputColumn,
	}

	model.findCoefficients()

	return model
}

type SimpleLinearRegression struct {
	dataSet *data.DataSet
	inputColumn string
	outputColumn string
	w0 float64
	w1 float64
}

func (m *SimpleLinearRegression) Coefficients() (float64, float64) {
	return m.w0, m.w1
}

func (m *SimpleLinearRegression) Predict(squareMeters float64) float64 {
	return m.model(squareMeters)
}

func (m *SimpleLinearRegression) findCoefficients() {
	fmt.Println("Finding coefficients...")

	var estimatedW0, estimatedW1, previousW0, previousW1 float64
	previousMagnitude := -1.

	acceptableError := 1.
	iterations := 0
	stepSize := 1.

	for {
		dRssW0, dRssW1, dRssMagnitude := m.dRss(estimatedW0, estimatedW1)

		if dRssMagnitude < acceptableError {
			break
		}

		if previousMagnitude >= 0 {
			if dRssMagnitude > previousMagnitude {
				estimatedW0 = previousW0
				estimatedW1 = previousW1
				stepSize *= .5
			} else {
				previousMagnitude = dRssMagnitude
				stepSize *= 1.05
			}
		} else {
			previousMagnitude = dRssMagnitude
		}


		fmt.Printf(
			"Iteration %d: est. w0 %g, est. w1 %g, dRssW0 %g, dRssW1 %g, magnitude %g\n",
			iterations, estimatedW0, estimatedW1, dRssW0, dRssW1, dRssMagnitude)

		iterations++

		fmt.Printf("Next step size: %e\n", stepSize)

		previousW0 = estimatedW0
		previousW1 = estimatedW1

		estimatedW0 = estimatedW0 - (stepSize * dRssW0)
		estimatedW1 = estimatedW1 - (stepSize * dRssW1)
	}

	m.w0 = estimatedW0
	m.w1 = estimatedW1

	fmt.Printf("Optimal coefficients found: w0 = %f, w1 = %f\n", m.w0, m.w1)
}

func (m *SimpleLinearRegression) dRss(w0, w1 float64) (float64, float64, float64) {
	var sumW0, sumW1 float64

	for _, row := range m.dataSet.Rows {
		xi := row.Cells[m.inputColumn]
		yi := row.Cells[m.outputColumn]

		partialTerm := yi - w0 - (w1 * xi)
		sumW0 += partialTerm
		sumW1 += partialTerm * xi
	}

	dRssW0 := -2 * sumW0
	dRssW1 := -2 * sumW1
	dRssMagnitude := math.Sqrt(
		math.Pow(dRssW0, 2) + math.Pow(dRssW1, 2),
	)

	return dRssW0, dRssW1, dRssMagnitude
}

func (m *SimpleLinearRegression) model(input float64) float64 {
	return m.w0 + (m.w1 * input)
}
