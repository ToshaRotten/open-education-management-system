package models

//id: 1,
//title: 'Математика',
//start: '2023-05-18T10:00',
//end: '2023-05-18T10:45',
//backgroundColor: '#5ba8ff',
//extendedProps: {
//homework: 'Изучение дробей',
//teacher: 'Антонов Роман Александрович',
//class: '7Б',
//topic: 'Дробные числа'
//}

type ExtendedProps struct {
	Homework string `json:"homework"`
	Teacher  string `json:"teacher"`
	Class    string `json:"class"`
	Topic    string `json:"topic"`
}

type Lesson struct {
	Id              int           `json:"id"`
	Title           string        `json:"title"`
	Start           string        `json:"start"`
	End             string        `json:"end"`
	BackgroundColor string        `json:"backgroundColor"`
	ExtendedProps   ExtendedProps `json:"extendedProps"`
}
