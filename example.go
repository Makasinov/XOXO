package main

import (
	"fmt"
	"strings"
	"time"
)

type Req struct {
	lat  float64
	long float64
}

type Res struct {
	msg       string
	next_ping time.Duration
}

type Step struct {
	req Req
	res Res
}

func findEntry(s1, s2 string) bool {
	return strings.Index(s1, s2) >= 0
}

// Find takes a slice and looks for an element in it. If found it will
// return it's key, otherwise it will return -1 and a bool of false.
func isTrustedLocation(slice []string, val string) bool {
	for _, item := range slice {
		if findEntry(val, item) {
			return true
		}
	}
	return false
}

func init() {
	dialog := []Step{
		{
			req: Req{55.649799, 37.595258},
			res: Res{
				"",
				time.Second * 5,
			},
		},
		{
			req: Req{55.649789, 37.595208},
			res: Res{
				"",
				time.Second * 5,
			},
		},
		{
			req: Req{55.649759, 37.595238},
			res: Res{
				"",
				time.Second * 5,
			},
		},
		{
			req: Req{55.647660, 37.596395},
			res: Res{
				"",
				time.Second * 5,
			},
		},
		{
			req: Req{55.647663, 37.596355},
			res: Res{
				"",
				time.Second * 5,
			},
		},
		{
			req: Req{55.646827, 37.596722},
			res: Res{
				"",
				time.Second * 5,
			},
		},
		{
			req: Req{55.645827, 37.596722},
			res: Res{
				"",
				time.Second * 5,
			},
		},
		{
			req: Req{55.645827, 37.596722},
			res: Res{
				"",
				time.Second * 5,
			},
		},
		{
			req: Req{55.645827, 37.596722},
			res: Res{
				"",
				time.Second * 5,
			},
		},
		{
			req: Req{55.645822, 37.596723},
			res: Res{
				"",
				time.Second * 5,
			},
		},
	}
	trustedLocations := []string{"Дом", "дом друга"}
	trustedLocationsPingDuration := time.Second * 10
	locations := loadCsv(CSV_PATH)
	fmt.Println("Сервер запущен")
	time.Sleep(time.Second)
	fmt.Println("Ребёнок установил приложение")
	time.Sleep(time.Second * 2)
	for idx, _ := range dialog {
		suitableLocation, distance := findSuitable(locations, dialog[idx].req.lat, dialog[idx].req.long)
		fmt.Print("--> ")
		time.Sleep(time.Millisecond * 700)
		fmt.Println(suitableLocation.Name, "( Широта:", dialog[idx].req.lat, "Долгота:", dialog[idx].req.long, "Расстояние", distance, ")")
		time.Sleep(time.Millisecond * 300)
		if isTrustedLocation(trustedLocations, suitableLocation.Name) {
			dialog[idx].res.next_ping = trustedLocationsPingDuration
			fmt.Println("<-- Безопасное место, следующий пинг через", dialog[idx].res.next_ping)
		} else {
			fmt.Println("<-- На улице, следующий пинг через", dialog[idx].res.next_ping)
		}
		time.Sleep(dialog[idx].res.next_ping)
	}
}
