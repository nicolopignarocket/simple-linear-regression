package data


type DataSet struct {
	Rows []Row
}

type Row struct {
	Cells map[string]float64
}
