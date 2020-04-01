package configs

type Export struct {
	ExportVideo, PlotEval, PrintGeneration bool
	Runs                                   int
	VideoMax, VideoMin                     float64
}
