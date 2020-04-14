package sweeper

import (
	"os"

	"github.com/CRAB-LAB-NTNU/PPS-BS/filehandling"
)

type Sweeper struct {
	sweep       bool
	dir         string
	name        string
	fr, hv, igd bool
	file        *os.File
}

func NewSweeper(sweep bool, dir, name string, fr, igd, hv bool) Sweeper {
	return Sweeper{
		sweep: sweep,
		dir:   dir,
		name:  name,
		file:  filehandling.OpenFile(dir, name),
		fr:    fr,
		igd:   igd,
		hv:    hv,
	}
}

func (s Sweeper) Sweep() bool {
	return s.sweep
}

func (s Sweeper) Dir() string {
	return s.dir
}
func (s Sweeper) Name() string {
	return s.name
}

func (s Sweeper) FR() bool {
	return s.fr
}
func (s Sweeper) IGD() bool {
	return s.igd
}
func (s Sweeper) HV() bool {
	return s.hv
}
func (s Sweeper) WriteLine(values []float64) {
	filehandling.WriteLine(values, *s.file)
}
