package sweeper

import (
	"os"

	"github.com/CRAB-LAB-NTNU/PPS-BS/filehandling"
)

type Sweeper struct {
	sweep                              bool
	dir                                string
	name                               string
	fr, hv, igd, archiveigd, archivehv bool
	phase                              bool
	file                               *os.File
}

func NewSweeper(sweep bool, dir, name string, phase, fr, igd, hv, archiveigd, archivehv bool) Sweeper {
	return Sweeper{
		sweep:      sweep,
		dir:        dir,
		name:       name,
		file:       filehandling.OpenFile(dir, name),
		fr:         fr,
		igd:        igd,
		hv:         hv,
		archiveigd: archiveigd,
		archivehv:  archivehv,
		phase:      phase,
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
func (s Sweeper) Phase() bool {
	return s.phase
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
func (s Sweeper) ArchiveIGD() bool {
	return s.archiveigd
}
func (s Sweeper) ArchiveHV() bool {
	return s.archivehv
}
func (s Sweeper) WriteLine(values []interface{}) {
	filehandling.WriteLine(values, *s.file)
}
