package data

// declare struct for data
type Course struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Category string `json:"category"`
}

type Teacher struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Gender     string `json:"gender"`
	Speciality string `json:"speciality"`
}

type Student struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Age    int    `json:"age"`
	Gender string `json:"gender"`
}

// initialize struct variable for base data
var Students = []Student{
	{ID: 1, Name: "Shikadai Nara", Age: 14, Gender: "M"},
	{ID: 2, Name: "Mirai Sarutobi", Age: 18, Gender: "F"},
	{ID: 3, Name: "Sarada Uchiha", Age: 14, Gender: "F"},
}

var Teachers = []Teacher{
	{ID: 1, Name: "Hanabi Hyuga", Gender: "F", Speciality: "Taijutsu"},
	{ID: 2, Name: "Konohamaru Sarutobi", Gender: "M", Speciality: "Ninjutsu"},
	{ID: 3, Name: "Udon Ise", Gender: "M", Speciality: "Genjutsu"},
}

var Courses = []Course{
	{ID: 1, Name: "Kagebunshin No Jutsu", Category: "Ninjutsu"},
	{ID: 2, Name: "Flower Petal Escape", Category: "Genjutsu"},
	{ID: 3, Name: "Konoha Senpoo", Category: "Taijutsu"},
}