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
	stationX = 17
	stationY = 23
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

		regionNum := determineRegion(station, otherAsteroid)
		region, ok := regions[regionNum]
		if !ok {
			region = make(map[string]bucket)
		}

		bucket := region[rationalString(regionNum, otherAsteroid, station)]
		bucket.points = append(bucket.points, otherAsteroid)
		bucket.degree = determineDegree(regionNum, station, otherAsteroid)

		region[rationalString(regionNum, otherAsteroid, station)] = bucket
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
				fmt.Println("Vapor point:", stationX + vaporPoint.X, stationY - vaporPoint.Y)
				return
			}

			bucket.points = bucket.points[1:]
			buckets[j] = bucket
		}
	}
}

func determineRegion(ref point.Point, other point.Point) int {
	dx := other.X - ref.X
	dy := other.Y - ref.Y

	if dx == 0 && dy > 0 {
		return 0
	} else if dx > 0 && dy > 0 {
		return 1
	} else if dx > 0 && dy == 0 {
		return 2
	} else if dx > 0 && dy < 0 {
		return 3
	} else if dx == 0 && dy < 0 {
		return 4
	} else if dx < 0 && dy < 0 {
		return 5
	} else if dx < 0 && dy == 0 {
		return 6
	} else if dx < 0 && dy > 0 {
		return 7
	}
	panic("unknown region")
}

func determineDegree(region int, ref, other point.Point) float64 {
	switch region {
	case 0:
		return 0
	case 1:
		return float64(other.X-ref.X) / float64(other.Y-ref.Y)
	case 2:
		return 0
	case 3:
		return -float64(other.Y-ref.Y) / float64(other.X-ref.X)
	case 4:
		return 0
	case 5:
		return float64(other.X-ref.X) / float64(other.Y-ref.Y)
	case 6:
		return 0
	case 7:
		return -float64(other.Y-ref.Y) / float64(other.X-ref.X)
	}
	panic("unknown region")
}

func rationalString(region int, other, aster point.Point) string {
	switch region {
	case 0, 4:
		return "dx_zero"
	default:
		return big.NewRat(int64(other.Y-aster.Y), int64(other.X-aster.X)).String()
	}
}
