package data

type State int

const (
	Pending State = iota
	InProgress
	Completed
)

var ValidStatus = []string{"all", "pending", "inprogress", "completed"}

func (st State) String() string {
	return [...]string{"Pendiente", "En progreso", "Completado"}[st]
}

const Route = "./tasks/data.json"
const ConfigRoute = "./tasks/config.json"

type Task struct {
	ID          int    `json:"ID"`
	Name        string `json:"Nombre"`
	Created     string `json:"Creado el"`
	State       State  `json:"Estado"`
	Update      string `json:"Actualizado el,omitempty"`
	Description string `json:"Descripción"`
}

type Config struct {
	ActualID int `json:"ActualID"`
}
