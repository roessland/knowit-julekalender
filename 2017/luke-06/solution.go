package main

import "fmt"
import "encoding/csv"
import "sort"
import "os"
import "strconv"
import "log"
import "github.com/golang/geo/s2"

/*
0 Landskode
1 Stadnamn nynorsk
2 Stadnamn bokmÃ¥l
3 Stadnamn engelsk
4 Geonames-ID int
5 Stadtype nynorsk
6 Stadtype bokmÃ¥l
7 Stadtype engelsk
8 Landsnamn nynorsk
9 Landsnamn bokmÃ¥l
10 Landsnamn engelsk
11 Folketal	int
12 Lat	double
13 Lon	double
14 HÃ¸gd over havet	double
15 Lenke til nynorsk-XML
16 Lenke til bokmÃ¥ls-XML
17 Lenke til engelsk-XML
*/

var EarthRadius float64 = 6371

type Capital struct {
	Name   string
	LatLng s2.LatLng
}

var Oslo Capital = Capital{"Oslo", s2.LatLngFromDegrees(59.911491, 10.757933)}

func NewCapital(record []string) Capital {
	c := Capital{}
	c.Name = record[3]
	latDegs, err := strconv.ParseFloat(record[12], 64)
	if err != nil {
		log.Fatal(err)
	}
	lngDegs, err := strconv.ParseFloat(record[13], 64)
	if err != nil {
		log.Fatal(err)
	}
	c.LatLng = s2.LatLngFromDegrees(latDegs, lngDegs)
	return c
}

func (c1 Capital) Distance(c2 Capital) float64 {
	angle := c1.LatLng.Distance(c2.LatLng)
	km := angle.Radians() * EarthRadius
	return km
}

func GetCapitals() []Capital {
	f, err := os.Open("verda.txt")
	if err != nil {
		log.Fatal(err)
	}

	reader := csv.NewReader(f)
	reader.Comma = '\t'
	records, err := reader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	alreadyAdded := make(map[string]bool)

	capitals := []Capital{}
	for _, record := range records {
		if record[7] == "capital" {
			capital := NewCapital(record)
			if !alreadyAdded[capital.Name] {
				capitals = append(capitals, capital)
				alreadyAdded[capital.Name] = true
			}
		}
	}
	return capitals
}

func main() {
	capitals := GetCapitals()
	sort.Slice(capitals, func(i, j int) bool {
		return Oslo.Distance(capitals[i]) < Oslo.Distance(capitals[j])
	})

	counter := 1
	time := 24.0
	speed := 7274.0
	for time > 0 {
		distance := 2 * Oslo.Distance(capitals[counter])
		flytime := distance / speed
		time = time - flytime

		if time > 0 {
			counter++
		}
	}
	fmt.Println(counter)
}
