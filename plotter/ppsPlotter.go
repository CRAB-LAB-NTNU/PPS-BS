package plotter

import (
	"fmt"
	"image/color"
	"log"
	"os"
	"os/exec"
	"strconv"

	"github.com/CRAB-LAB-NTNU/PPS-BS/configs"
	"github.com/CRAB-LAB-NTNU/PPS-BS/filehandling"
	"github.com/CRAB-LAB-NTNU/PPS-BS/types"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

type PpsPlotter struct {
	Population        *[]types.Individual
	Archive           *[]types.Individual
	Generation, Stage *int
	Stages            []types.Stage
	Config            *configs.Export
	Cmop              *types.Cmop
}

func (p PpsPlotter) title() string {
	stage := p.Stages[*p.Stage].Name()
	gen := strconv.Itoa(*p.Generation)
	return p.Cmop.Name + " Stage: " + stage + " gen: " + gen
}

func (p PpsPlotter) PlotFrame() {

	plt, _ := plot.New()

	plt.Title.Text = p.title()

	plt.X.Label.Text, plt.Y.Label.Text = "f1", "f2"

	plt.X.Max, plt.Y.Max = p.Config.VideoMax, p.Config.VideoMax
	plt.X.Min, plt.Y.Min = p.Config.VideoMin, p.Config.VideoMin

	plt.Add(plotter.NewGrid())

	if paretoFront := p.Cmop.TrueParetoFront(); paretoFront != nil {
		points := p.floatScatter(paretoFront, color.RGBA{R: 0, G: 255, B: 0, A: 255})
		plt.Add(points)
		plt.Legend.Add("Pareto Front", points)
	}

	population := p.individualScatter(*p.Population, color.RGBA{R: 0, G: 0, B: 0, A: 255})
	plt.Add(population)
	plt.Legend.Add("Population", population)

	archive := p.individualScatter(*p.Archive, color.RGBA{R: 255, G: 0, B: 0, A: 255})
	plt.Add(archive)
	plt.Legend.Add("Archive", archive)

	if err := filehandling.ControlPath(p.framePath()); err != nil {
		log.Fatal(err)
	}

	if err := plt.Save(500, 500, p.frameFile()); err != nil {
		log.Fatal(err)
	}

}

func (p PpsPlotter) ExportVideo() {

	if err := filehandling.ControlPath(p.videoPath()); err != nil {
		log.Fatal(err)
	}

	if err := os.Remove(p.videoFile()); err == nil {
		fmt.Println("Removed old video file.", p.videoFile())
	}

	cmd := exec.Command("ffmpeg", "-framerate", "20", "-i", p.framePath()+"/%00d"+FrameFormat.String(), p.videoFile())
	if err := cmd.Run(); err != nil {
		fmt.Println("Feil ved laging av video, prøv \nsudo apt install ffmpeg")
		log.Fatal(err)
	}
}

func (p PpsPlotter) framePath() string {
	return FramesDir.String() + p.Cmop.Name
}

func (p PpsPlotter) frameFile() string {
	gen := strconv.Itoa(*p.Generation)
	return p.framePath() + "/" + gen + FrameFormat.String()
}

func (p PpsPlotter) videoPath() string {
	return VideosDir.String()
}

func (p PpsPlotter) videoFile() string {
	return p.videoPath() + p.Cmop.Name + VideoFormat.String()
}

func (p PpsPlotter) individualScatter(pop []types.Individual, color color.RGBA) *plotter.Scatter {
	x := convertIndividualsToPoints2D(pop, p.Config.VideoMin, p.Config.VideoMax)
	scatter, _ := plotter.NewScatter(x)
	scatter.GlyphStyle.Color = color
	scatter.GlyphStyle.Radius = vg.Points(1)
	return scatter
}

func (p PpsPlotter) floatScatter(data [][]float64, color color.RGBA) *plotter.Scatter {
	points := convertParetoToPoints(data, p.Config.VideoMin, p.Config.VideoMax)
	scatter, _ := plotter.NewScatter(points)
	scatter.GlyphStyle.Color = color
	scatter.GlyphStyle.Radius = vg.Points(1)
	return scatter
}
