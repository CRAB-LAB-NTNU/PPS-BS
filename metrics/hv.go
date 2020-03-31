package metrics

import (
	"bytes"
	"log"
	"math"
	"os/exec"
	"strconv"
	"strings"

	"github.com/google/uuid"

	"github.com/CRAB-LAB-NTNU/PPS-BS/biooperators"
	"github.com/CRAB-LAB-NTNU/PPS-BS/filehandling"
	"github.com/CRAB-LAB-NTNU/PPS-BS/types"
	"github.com/CRAB-LAB-NTNU/PPS-BS/utils"
)

//Hypervolume requires the command line tool http://lopez-ibanez.eu/hypervolume#building
func HyperVolume(population []types.Individual, referencePoint []float64) (s float64) {

	if exists := filehandling.CommandExists("hv"); exists == false {
		log.Println("Download the HV tool to calculate Hypervolume\n", "http://lopez-ibanez.eu/hypervolume#building")
		return 0
	}

	if err := filehandling.ControlPath(".tmp"); err != nil {
		log.Fatal(err)
	}

	nonDominated := biooperators.FastNonDominatedSort(population)[0]
	// need to save data to file before sending to hv.
	id := uuid.New()

	hvFile := filehandling.OpenHVFile(id.String())

	for _, individual := range nonDominated {
		filehandling.WriteHVLine(individual.Fitness().ObjectiveValues, *hvFile)
	}
	hvFile.Sync()
	hvFile.Close()

	cmd := exec.Command("hv", "-r", utils.FloatArrayToString(referencePoint, " "), ".tmp/"+id.String())
	var stderr bytes.Buffer
	cmd.Stderr = &stderr
	output, err := cmd.Output()

	if err != nil {
		s = math.NaN()
	} else {
		s, _ = strconv.ParseFloat(strings.TrimSpace(string(output)), 64)
	}
	filehandling.RemoveHVFile(id.String())
	return s
}

func HVReferenceNadirTimes(coefficient float64, cmop types.Cmop) func() []float64 {
	return func() []float64 {
		nadir := make([]float64, cmop.ObjectiveCount)
		for _, point := range cmop.TrueParetoFront() {
			for i, cord := range point {
				nadir[i] = math.Max(cord, nadir[i])
			}
		}
		for i := range nadir {
			nadir[i] *= coefficient
		}
		return nadir
	}
}
