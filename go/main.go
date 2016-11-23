package main

import (
	"fmt"
	"slr/model"
	"slr/data"
)

func main() {
	dataset := data.NewDataSet("../dataset.csv")

	simpleLinearRegression := model.NewSimpleLinearRegression(dataset, dataset.Columns[0], dataset.Columns[1])

	for {
		fmt.Print("Enter square meters: ")
		var squareMeters float64
		fmt.Scanf("%f", &squareMeters)
		predictedSellingPrice := simpleLinearRegression.Predict(squareMeters)
		fmt.Printf("Predicted selling price: € %.2f\n", predictedSellingPrice)
	}
}
