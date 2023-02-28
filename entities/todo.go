package entities

type TodosStruct struct {
	Id     string `json:"id"`
	Title  string `json:"title"`
	IsDone bool   `isDone:"isDone"`
}

var Todos = []TodosStruct{
	{
		Id:     "1",
		Title:  "Cuci tangan",
		IsDone: true,
	},
	{
		Id:     "2",
		Title:  "Jaga jarak",
		IsDone: false,
	},
}