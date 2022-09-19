package entity

type Users struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	City string `json:"city"`
	Age  int    `json:"age"`
}

var Data = []Users{
	{"E001", "John", "New York", 12},
	{"E002", "Doe", "California", 23},
	{"E003", "Harry", "London", 22},
	{"E004", "Maguire", "Stockholms", 25},
	{"E005", "Gerry", "Tokyo", 16},
	{"E006", "Picky", "Singapore", 18},
	{"E007", "Davies", "New Zealand", 19},
	{"E008", "Alex", "Sydney", 17},
}
