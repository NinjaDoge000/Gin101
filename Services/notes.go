package services

// logic related to connecting to DB, yada, yada, goes here...
type NotesService struct {
}

type notes struct {
	Id   int
	Name string
}

func (n *NotesService) GetNotesSerivce() []notes {
	return []notes{
		{Id: 1, Name: "note1"},
		{Id: 2, Name: "note2"},
		{Id: 3, Name: "note3"},
	}
}

func (n *NotesService) CreateNotesSerivce() notes {
	return notes{
		Id:   1,
		Name: "note1",
	}
}
