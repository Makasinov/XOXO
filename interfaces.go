package main

type StationRow struct {
	LineName     string  `csv:"line_name"`
	LineNumber   string  `csv:"line_number"`
	LineHexColor string  `csv:"line_hex_color"`
	Name         string  `csv:"station_name"`
	Number       string  `csv:"station_number"`
	LatitudeRaw  string  `csv:"latitude"`
	LongitudeRaw string  `csv:"longitude"`
	Latitude     float64 `json:"latitude"`
	Longitude    float64 `json:"longitude"`
}
