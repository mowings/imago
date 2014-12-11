package work

import (
	"github.com/twinj/uuid"
)

type Work struct {
	Id      string
	Status  string
	Infile  string
	Outfile string
	Opts    string
}

func (w *Work) Initialize() {
	u := uuid.NewV4()
	uuid.SwitchFormat(uuid.Clean)
	w.Id = u.String()
	w.Status = ""
}
