package configs

type Sweeper struct {
	Sweep                              bool
	Dir                                string
	FR, IGD, HV, ArchiveIGD, ArchiveHV bool
	Phase                              bool
}
