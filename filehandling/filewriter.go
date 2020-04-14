package filehandling

import (
	"fmt"
	"log"
	"os"
)

/* Commented out because of circular dependency.
//WriteArrayToFile writes the objective values of an array of individuals to a destination file
func WriteArrayToFile(individuals []types.Individual, destination string) {

	file, err := os.Create(destination)

	checkError("cannot create file", err)
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	for _, individual := range individuals {
		objectiveValues := individual.Fitness().ObjectiveValues

		err := writer.Write([]string{
			strconv.FormatFloat(objectiveValues[0], 'E', -1, 64),
			strconv.FormatFloat(objectiveValues[1], 'E', -1, 64),
			strconv.FormatFloat(objectiveValues[2], 'E', -1, 64)})
		checkError("Cannot write to file", err)

	}
}
*/

func checkError(message string, err error) {
	if err != nil {
		log.Fatal(message, err)
	}
}

func ControlPath(path string) error {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		if innerErr := os.MkdirAll(path, 0755); innerErr != nil {
			return innerErr
		}
	}
	return nil
}

func OpenFile(dir, name string) *os.File {

	err := ControlPath(dir)
	if err != nil {
		log.Fatal(err)
	}

	if file, err := os.Create(dir + name); err == nil {
		return file
	}
	log.Fatal("Couldn't create file: ", name)

	return nil
}

func WriteLine(values []float64, file os.File) {
	var line string
	for _, value := range values {
		line += fmt.Sprint(value) + " "
	}
	if _, err := file.WriteString(line + "\n"); err != nil {
		log.Fatal("Couldn't write ", line, " to file ", file.Name())
	}
}

func RemoveFile(dir, name string) {
	err := os.Remove(dir + name)
	if err != nil {
		log.Fatal("Couldn't remove file", name)
	}
}
