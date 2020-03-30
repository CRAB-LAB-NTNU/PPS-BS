package filehandling

import (
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
