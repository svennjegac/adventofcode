package main

import (
	"fmt"
	"math"
	"math/big"
	"sort"

	"github.com/svennjegac/adventofcode/2019/day10"
	"github.com/svennjegac/adventofcode/2019/day10/point"
)

const (
	expectedVaporized = 200
	stationX          = 17
	stationY          = 23
)

type bucket struct {
	points []point.Point
	degree float64
}

func main() {
	asteroids, err := day10.Asteroids2(stationX, stationY)
	if err != nil {
		panic(err)
	}

	station := point.Point{
		X: 0,
		Y: 0,
	}

	regions := make(map[int]map[string]bucket)
	for otherAsteroid := range asteroids {
		if station == otherAsteroid {
			continue
		}

		regionNum, degree := getRegionAndDegree(station, otherAsteroid)
		region, ok := regions[regionNum]
		if !ok {
			region = make(map[string]bucket)
		}

		bucket := region[rationalString(station, otherAsteroid)]
		bucket.points = append(bucket.points, otherAsteroid)
		bucket.degree = degree

		region[rationalString(station, otherAsteroid)] = bucket
		regions[regionNum] = region
	}

	sortedRegionBuckets := make(map[int][]bucket)
	for regionNum, region := range regions {
		regionBuckets := make([]bucket, 0, len(region))
		for _, bucket := range region {
			sort.Slice(bucket.points, func(i, j int) bool {
				return math.Sqrt(math.Pow(float64(bucket.points[i].X-station.X), 2)+math.Pow(float64(bucket.points[i].Y-station.Y), 2)) <
					math.Sqrt(math.Pow(float64(bucket.points[j].X-station.X), 2)+math.Pow(float64(bucket.points[j].Y-station.Y), 2))
			})
			regionBuckets = append(regionBuckets, bucket)
		}
		sort.Slice(regionBuckets, func(i, j int) bool {
			return regionBuckets[i].degree < regionBuckets[j].degree
		})
		sortedRegionBuckets[regionNum] = regionBuckets
	}

	vaporizedCounter := 0
	for i := 0; ; i = (i + 1) % 8 {
		buckets := sortedRegionBuckets[i]
		for j, bucket := range buckets {
			if len(bucket.points) == 0 {
				continue
			}

			vaporPoint := bucket.points[0]
			vaporizedCounter++
			if vaporizedCounter == expectedVaporized {
				fmt.Println("Vapor point:", stationX+vaporPoint.X, stationY-vaporPoint.Y)
				return
			}

			bucket.points = bucket.points[1:]
			buckets[j] = bucket
		}
	}
}

func getRegionAndDegree(ref, other point.Point) (int, float64) {
	dx := other.X - ref.X
	dy := other.Y - ref.Y

	if dx == 0 && dy > 0 {
		return 0, 0
	} else if dx > 0 && dy > 0 {
		return 1, float64(other.X-ref.X) / float64(other.Y-ref.Y)
	} else if dx > 0 && dy == 0 {
		return 2, 0
	} else if dx > 0 && dy < 0 {
		return 3, -float64(other.Y-ref.Y) / float64(other.X-ref.X)
	} else if dx == 0 && dy < 0 {
		return 4, 0
	} else if dx < 0 && dy < 0 {
		return 5, float64(other.X-ref.X) / float64(other.Y-ref.Y)
	} else if dx < 0 && dy == 0 {
		return 6, 0
	} else if dx < 0 && dy > 0 {
		return 7, -float64(other.Y-ref.Y) / float64(other.X-ref.X)
	}
	panic("unknown region")
}

func rationalString(ref, other point.Point) string {
	if other.X == ref.X {
		return "dx_zero"
	}
	return big.NewRat(int64(other.Y-ref.Y), int64(other.X-ref.X)).String()
}
