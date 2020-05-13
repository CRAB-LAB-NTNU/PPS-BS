package configs

type Sweeper struct {
	Sweep      bool
	Dir        string
	FR         bool
	CD         bool
	IGD        bool
	HV         bool
	ArchiveIGD bool
	ArchiveHV  bool
	Phase      bool
}
