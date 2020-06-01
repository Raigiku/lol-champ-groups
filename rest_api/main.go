package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

type DataPoint struct {
	positions []float64
}

func newDataPoint(positions []float64) DataPoint {
	dataPoint := DataPoint{positions: positions}
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
	sum := 0.0
	for i := range dataPoint1.positions {
		sum += math.Pow(dataPoint1.positions[i]-dataPoint2.positions[i], 2)
	}
	distance := math.Sqrt(sum)
	return distance
}

func clusterContainsDataPoint(cluster Cluster, otherDataPoint DataPoint) bool {
	for _, dataPoint := range cluster.dataPoints {
		containsDatapoint := false
		for i := range dataPoint.positions {
			if dataPoint.positions[i] == otherDataPoint.positions[i] {
				containsDatapoint = true || containsDatapoint
			}
		}
		if containsDatapoint {
			return true
		}
	}
	return false
}

func randomDataPoints(totalDimensions, totalDataPoints int, minValue, maxValue float64) []DataPoint {
	dataPoints := make([]DataPoint, totalDataPoints)
	for i := range dataPoints {
		positions := make([]float64, totalDimensions)
		for j := range positions {
			positions[j] = minValue + rand.Float64()*(maxValue-minValue)
		}
		dataPoints[i] = newDataPoint(positions)
	}
	return dataPoints
}

func initialClusters(totalClusters int, dataPoints []DataPoint) []Cluster {
	clusters := make([]Cluster, totalClusters)
	totalDataPoints := len(dataPoints)
	randomIndexDataPoints := rand.Perm(totalDataPoints)
	for i := range clusters {
		randomDataPoint := dataPoints[randomIndexDataPoints[i]]
		clusters[i] = newCluster(randomDataPoint)
	}
	return clusters
}

func printResults(dataPoints []DataPoint, clusters []Cluster) {
	fmt.Println("datapoints:")
	for _, dataPoint := range dataPoints {
		fmt.Printf("%v ", dataPoint.positions)
		fmt.Println()
	}
	fmt.Println()
	fmt.Println()
	fmt.Println("clusters:")
	for _, cluster := range clusters {
		fmt.Printf("centroid: %v", cluster.centroid.positions)
		fmt.Println()
		for _, dataPoint := range cluster.dataPoints {
			fmt.Printf("%v ", dataPoint.positions)
		}
		fmt.Println()
		fmt.Println()
	}
}

func runKMeans(clusters []Cluster, dataPoints []DataPoint) {
	totalClusters := len(clusters)
	totalDimensions := len(dataPoints[0].positions)
	for {
		didClustersChange := false
		tempClusters := clusters
		for i := range tempClusters {
			tempClusters[i].dataPoints = nil
		}

		assignDataPointsToCluster := func() {
			for _, dataPoint := range dataPoints {
				nearestCluster, nearestClusterIndex := &tempClusters[0], 0
				for i := 1; i < totalClusters; i++ {
					if distanceBetweenDataPoints(tempClusters[i].centroid, dataPoint) < distanceBetweenDataPoints(nearestCluster.centroid, dataPoint) {
						nearestCluster, nearestClusterIndex = &tempClusters[i], i
					}
				}
				nearestCluster.dataPoints = append(nearestCluster.dataPoints, dataPoint)
				if !clusterContainsDataPoint(clusters[nearestClusterIndex], dataPoint) {
					didClustersChange = true
				}
			}
		}
		assignDataPointsToCluster()

		if !didClustersChange {
			break
		}

		repositionCentroid := func() {
			for _, cluster := range tempClusters {
				sumPositions := make([]float64, totalDimensions)
				for _, dataPoint := range cluster.dataPoints {
					for i, positions := range dataPoint.positions {
						sumPositions[i] += positions
					}
				}
				avgPositions := make([]float64, totalDimensions)
				for i := range avgPositions {
					avgPositions[i] = avgPositions[i] / float64(totalClusters)
				}
				cluster.centroid = newDataPoint(avgPositions)
			}
		}
		repositionCentroid()

		clusters = tempClusters
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())
	dataPoints := randomDataPoints(1, 10, -20, 20)
	clusters := initialClusters(3, dataPoints)
	runKMeans(clusters, dataPoints)
	printResults(dataPoints, clusters)
}
