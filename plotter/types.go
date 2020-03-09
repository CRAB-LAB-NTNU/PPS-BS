package plotter

// Plotter is a struct used to visualise problems and a population
type Plotter interface {
	Plot()
}

// Plotter2D is a struct incorporating the Plotter interface to plot 2D plots
type Plotter2D struct {
	Title, LabelX, LabelY, Filename string
	Min, Max                        float64
	Solution, Extremes              [][]float64
}

type GraphicsNames int

const (
	FramesDir GraphicsNames = iota
	VideosDir
	FrameFormat
	VideoFormat
)

func (g GraphicsNames) String() string {
	return [...]string{
		"graphics/frames/",
		"graphics/videos/",
		".png",
		".mp4",
	}[g]
}
