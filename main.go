package main

import (
	"flag"
	"fmt"
	"math"
	"os"
	"strconv"

	"github.com/gocarina/gocsv"
)

const CSV_PATH = "dataset.csv"

func main() {
	lat, long := getArguments()
	stations := loadCsv(CSV_PATH)
	suitableStation, distance := findSuitable(stations, lat, long)
	if distance >= 1 {
		fmt.Println("Растояние до ближайшей станции", distance)
		return
	}

	fmt.Println(suitableStation.LineName, suitableStation.Name)
	fmt.Println("Разница:", distance)
	fmt.Println(lat, long)
	fmt.Println(suitableStation.Latitude, suitableStation.Longitude)
}

func getArguments() (float64, float64) {
	test := flag.Bool("test", true, "test")
	latitude := flag.Float64("latitude", 0, "широта")
	longitude := flag.Float64("longitude", 0, "долгота")
	if test == nil || *test == true {
		panic("Клиент отключился. Посылаем уведомление родителю. Выключение сервера.")
	}
	if latitude == nil || *latitude == 0 {
		panic("-latitude аргумент пропущен")
	}
	if longitude == nil || *longitude == 0 {
		panic("-longitude аргумент пропущен")
	}

	return *latitude, *longitude
}

func loadCsv(path string) []StationRow {
	in, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer in.Close()

	var stations []StationRow
	if err := gocsv.UnmarshalFile(in, &stations); err != nil {
		panic(err)
	}
	for idx, _ := range stations {
		stations[idx].Latitude, err = strconv.ParseFloat(stations[idx].LatitudeRaw, 64)
		stations[idx].Longitude, err = strconv.ParseFloat(stations[idx].LongitudeRaw, 64)
	}

	if len(stations) == 0 {
		panic("станции отсутствуют")
	}

	return stations
}

func findSuitable(stations []StationRow, latitude, longitude float64) (StationRow, float64) {
	getDistance := func(x1, x2, y1, y2 float64) float64 {
		return math.Sqrt(
			math.Pow(x2-x1, 2) + math.Pow(y2-y1, 2),
		)
	}
	var closestStation = stations[0]
	var leastDistance = 99999999999999999.
	for idx, _ := range stations {
		distance := getDistance(stations[idx].Latitude, latitude, stations[idx].Longitude, longitude)
		if distance <= leastDistance {
			leastDistance = distance
			closestStation = stations[idx]
		}
	}

	return closestStation, leastDistance
}
