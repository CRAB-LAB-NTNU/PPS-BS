package plotter

// Plotter is a struct used to visualise problems and a population
type Plotter interface {
	Plot()
}

type Plotter2D struct {
	title  string
	labelX string
	labelY string
}
