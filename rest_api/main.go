package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

type DataPoint struct {
	position float64
}

func newDataPoint(position float64) DataPoint {
	dataPoint := DataPoint{position: position}
	return dataPoint
}

type Cluster struct {
	centroid   DataPoint
	dataPoints []DataPoint
}

func newCluster(centroid DataPoint) Cluster {
	cluster := Cluster{centroid: centroid}
	return cluster
}

func distanceBetweenDataPoints(dataPoint1, dataPoint2 DataPoint) float64 {
	distance := math.Sqrt(math.Pow(dataPoint1.position-dataPoint2.position, 2))
	return distance
}

func clusterContainsDataPoint(cluster Cluster, otherDataPoint DataPoint) bool {
	for _, dataPoint := range cluster.dataPoints {
		if dataPoint == otherDataPoint {
			return true
		}
	}
	return false
}

func main() {
	rand.Seed(time.Now().UnixNano())

	// generate input
	n := 40
	dataPoints := make([]DataPoint, n)
	minValue, maxValue := -20, 20
	randomPositions := rand.Perm(maxValue*2 + 1)
	for i, position := range randomPositions[:n] {
		dataPoints[i] = newDataPoint(float64((position % (maxValue*2 + 1)) + minValue))
	}

	// place centroids
	k := 2
	clusters := make([]Cluster, k)
	randomIndexDataPoints := rand.Perm(n)
	for i := range clusters {
		randomDataPoint := dataPoints[randomIndexDataPoints[i]]
		clusters[i] = newCluster(randomDataPoint)
	}

	// clusters[0].dataPoints = append(clusters[0].dataPoints, dataPoints[0])

	// repeat until convergence
	for {
		didClustersChange := false
		tempClusters := clusters
		for i := range tempClusters {
			tempClusters[i].dataPoints = nil
		}

		for _, dataPoint := range dataPoints {
			nearestCluster, nearestClusterIndex := &tempClusters[0], 0
			for i := 1; i < k; i++ {
				if distanceBetweenDataPoints(tempClusters[i].centroid, dataPoint) < distanceBetweenDataPoints(nearestCluster.centroid, dataPoint) {
					nearestCluster, nearestClusterIndex = &tempClusters[i], i
				}
			}
			nearestCluster.dataPoints = append(nearestCluster.dataPoints, dataPoint)
			if !clusterContainsDataPoint(clusters[nearestClusterIndex], dataPoint) {
				didClustersChange = true
			}
		}
		if !didClustersChange {
			break
		}
		for _, cluster := range tempClusters {
			sumPositions := 0.
			for _, dataPoint := range cluster.dataPoints {
				sumPositions += dataPoint.position
			}
			avgPosition := sumPositions / float64(k)
			cluster.centroid = newDataPoint(avgPosition)
		}
		clusters = tempClusters
	}

	// showing results
	fmt.Println("datapoints:")
	for _, dataPoint := range dataPoints {
		fmt.Printf("%v ", dataPoint.position)
	}
	fmt.Println()
	fmt.Println()
	fmt.Println("clusters:")
	for _, cluster := range clusters {
		fmt.Printf("centroid: %v", cluster.centroid.position)
		fmt.Println()
		for _, dataPoint := range cluster.dataPoints {
			fmt.Printf("%v ", dataPoint.position)
		}
		fmt.Println()
	}
}
