package utils

import "math"

//https://github.com/umahmood/haversine/blob/master/haversine.go

const (
	earthRadiusMi = 3958 // radius of the earth in miles.
	earthRaidusKm = 6371 // radius of the earth in kilometers.
)


func degreesToRadians(d float64) float64 {
	return d * math.Pi / 180
}

func Haversine(la1, lo1, la2, lo2 float64) float64{

	lat1 := degreesToRadians(la1)
	lon1 := degreesToRadians(lo1)
	lat2 := degreesToRadians(la2)
	lon2 := degreesToRadians(lo2)

	diffLat := lat2 - lat1
	diffLon := lon2 - lon1

	a := math.Pow(math.Sin(diffLat/2), 2) + math.Cos(lat1)*math.Cos(lat2)*
		math.Pow(math.Sin(diffLon/2), 2)

	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))
	km := c * earthRaidusKm
	return km
}

