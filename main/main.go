package main

import (
	"os"
	"path/filepath"
	"github.com/EvanStanford/cams/profiler"
	"strconv"
	"fmt"
)

func basename (filename string) string {
	ext := filepath.Ext(filename)
	return filename[0:len(filename) - len(ext)]
}

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Incorrect usage. Correct usage: main.exe file/path.csv 0.045")
		return
	}
	args := os.Args[1:]
	inputPath := args[0]
	scale, err := strconv.ParseFloat(args[1], 64)
	if err != nil {
		fmt.Println("Incorrect usage. Could not parse scale. Correct usage: main.exe file/path.csv 0.045")
		return
	}

	base := basename(filepath.Base(inputPath))

	outputStlL := filepath.Join("out", base, base + "-L.stl")
	outputStlR := filepath.Join("out", base, base + "-R.stl")

	os.MkdirAll(filepath.Join("out", base), 0755)

	cams.CreateCams(
		inputPath, //note input path must be centered around Y axis
		outputStlL,
		outputStlR,
		5,
		scale,
		scale,
		43500, //Contant: distance between cam axels, center to center
		32300, //Contant: max radius of cam. If this was any larger the cams would hit each other
		7060, //Laser pen radius
	)
}