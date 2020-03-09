package filehandling

import (
	"bufio"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func ParseDatFile(path string) ([][]float64, error) {
	var bucket [][]float64
	if file, err := os.Open(path); err != nil {
		return nil, err
	} else {
		defer file.Close()
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			line := scanner.Text()
			cords := strings.Fields(line)
			var point []float64
			for _, f := range cords {
				if v, err := parseFloat(f); err != nil {
					log.Fatal(err)
				} else {
					point = append(point, v)
				}
			}
			bucket = append(bucket, point)
		}
		return bucket, nil
	}
}

func parseFloat(str string) (float64, error) {
	val, err := strconv.ParseFloat(str, 64)
	if err == nil {
		return val, nil
	}

	//Some number may be seperated by comma, for example, 23,120,123, so remove the comma firstly
	str = strings.Replace(str, ",", "", -1)

	//Some number is specifed in scientific notation
	pos := strings.IndexAny(str, "eE")
	if pos < 0 {
		return strconv.ParseFloat(str, 64)
	}

	var baseVal float64
	var expVal int64

	baseStr := str[0:pos]
	baseVal, err = strconv.ParseFloat(baseStr, 64)
	if err != nil {
		return 0, err
	}

	expStr := str[(pos + 1):]
	expVal, err = strconv.ParseInt(expStr, 10, 64)
	if err != nil {
		return 0, err
	}

	return baseVal * math.Pow10(int(expVal)), nil
}
