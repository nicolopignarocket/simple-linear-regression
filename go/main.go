package main

import (
	"fmt"
	"slr/model"
	"slr/data"
)

func main() {
	dataset := data.DataSet{
		Rows: []data.Row{
			{Cells: map[string]float64{"SquareMeters": 123, "SellingPrice": 302030}},
			{Cells: map[string]float64{"SquareMeters": 86, "SellingPrice": 123000}},
			{Cells: map[string]float64{"SquareMeters": 45, "SellingPrice": 68000}},
			{Cells: map[string]float64{"SquareMeters": 200, "SellingPrice": 400876}},
			{Cells: map[string]float64{"SquareMeters": 66, "SellingPrice": 110423}},
			{Cells: map[string]float64{"SquareMeters": 90, "SellingPrice": 120432}},
		},
	}

	simpleLinearRegression := model.NewSimpleLinearRegression(&dataset, "SquareMeters", "SellingPrice")

	for {
		fmt.Print("Enter square meters: ")
		var squareMeters float64
		fmt.Scanf("%f", &squareMeters)
		predictedSellingPrice := simpleLinearRegression.Predict(squareMeters)
		fmt.Printf("Predicted selling price: %.2f\n", predictedSellingPrice)
	}
}
