package utils

import "math"

func Haversine(lat1, lon1, lat2, lon2 float64) float64{
	// convert to radians
	// must cast radius as float to multiply later
	var la1, lo1, la2, lo2, r float64
	la1 = lat1 * math.Pi / 180
	lo1 = lon1 * math.Pi / 180
	la2 = lat2 * math.Pi / 180
	lo2 = lon2 * math.Pi / 180

	// radius of earth in km
	r = 6378.8

	// calculate
	h := math.Sin(la2-la1)*math.Cos(la2+la1) + math.Cos(la1)*math.Cos(la2)*math.Sin(lo2-lo1)

	return h * r
}

