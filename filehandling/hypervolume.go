package filehandling

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

const dir = ".tmp/"

func CommandExists(cmd string) bool {
	_, err := exec.LookPath(cmd)
	return err == nil
}

func OpenHVFile(name string) *os.File {
	if file, err := os.Create(dir + name); err == nil {
		return file
	}
	log.Fatal("Couldn't create file named ", name)
	return nil
}

func WriteHVLine(values []float64, file os.File) {
	var s string
	for _, value := range values {
		s += (fmt.Sprint(value) + " ")
	}
	if _, err := file.WriteString(s + "\n"); err != nil {
		log.Fatal("Couldnt write ", s, " to file ", file.Name())
	}
}

func RemoveHVFile(name string) {
	err := os.Remove(dir + name)
	if err != nil {
		log.Fatal("Couldnt remove file", name)
	}

}
