package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gorilla/mux"
)

// Champion league of legends
type Champion struct {
	Name  string               `json:"name"`
	Lanes []map[string]float64 `json:"lanes"`
}

// Lane league of legends
// type Lane struct {
// 	// Name           string  `json:"name"`
// 	PickPercentage float64 `json:"pick_percentage"`
// 	PhysicalDamage float64 `json:"physical_damage"`
// 	MagicDamage    float64 `json:"magic_damage"`
// 	TrueDamage     float64 `json:"true_damage"`
// 	TotalDamage    float64 `json:"total_damage"`
// 	DamageTaken    float64 `json:"damage_taken"`
// 	Healing        float64 `json:"healing"`
// 	Kills          float64 `json:"kills"`
// 	Deaths         float64 `json:"deaths"`
// 	Assists        float64 `json:"assists"`
// 	MaxKillSpree   float64 `json:"max_kill_spree"`
// 	Gold           float64 `json:"gold"`
// 	MinionsKilled  float64 `json:"minions_killed"`
// 	JungleCs       float64 `json:"jungle_cs"`
// 	EnemyJungleCs  float64 `json:"enemy_jungle_cs"`
// 	TeamJungleCs   float64 `json:"team_jungle_cs"`
// }

// DataPoint is a datapoint
type DataPoint struct {
	ChampionName string    `json:"championName"`
	Positions    []float64 `json:"positions"`
}

func newDataPoint(championName string, positions []float64) DataPoint {
	dataPoint := DataPoint{ChampionName: championName, Positions: positions}
	return dataPoint
}

// Cluster is a cluster of data points
type Cluster struct {
	Centroid   DataPoint   `json:"centroid"`
	DataPoints []DataPoint `json:"dataPoints"`
}

func newCluster(centroid DataPoint) Cluster {
	cluster := Cluster{Centroid: centroid}
	return cluster
}

func distanceBetweenDataPoints(dataPoint1, dataPoint2 DataPoint) float64 {
	sum := 0.0
	for i := range dataPoint1.Positions {
		sum += math.Pow(dataPoint1.Positions[i]-dataPoint2.Positions[i], 2)
	}
	distance := math.Sqrt(sum)
	return distance
}

func clusterContainsDataPoint(cluster Cluster, otherDataPoint DataPoint) bool {
	for _, dataPoint := range cluster.DataPoints {
		containsDatapoint := false
		for i := range dataPoint.Positions {
			if dataPoint.Positions[i] == otherDataPoint.Positions[i] {
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
		Positions := make([]float64, totalDimensions)
		for j := range Positions {
			Positions[j] = minValue + rand.Float64()*(maxValue-minValue)
		}
		dataPoints[i] = newDataPoint(string(i), Positions)
	}
	return dataPoints
}

func lolChampionsFileDataPoints(filename string, laneName string, attributes []string) []DataPoint {
	file, _ := ioutil.ReadFile(filename)
	var champions []Champion
	json.Unmarshal(file, &champions)

	laneIds := map[string]int{
		"top":     0,
		"jungle":  1,
		"middle":  2,
		"bottom":  3,
		"support": 4,
	}

	var dataPoints []DataPoint
	for _, champion := range champions {
		lane := champion.Lanes[laneIds[laneName]]
		if lane["pick_percentage"] >= 10 {
			dataPointPositions := make([]float64, 0)
			for _, attribute := range attributes {
				dataPointPositions = append(dataPointPositions, lane[attribute])
			}
			dataPoint := newDataPoint(champion.Name, dataPointPositions)
			dataPoints = append(dataPoints, dataPoint)
		}
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

func printResults(DataPoints []DataPoint, clusters []Cluster) {
	fmt.Println("datapoints:")
	for _, dataPoint := range DataPoints {
		fmt.Printf("%v ", dataPoint.Positions)
		fmt.Println()
	}
	fmt.Println()
	fmt.Println()
	fmt.Println("clusters:")
	for _, cluster := range clusters {
		fmt.Printf("Centroid: (")
		for i, position := range cluster.Centroid.Positions {
			fmt.Printf("%v", position)
			if i != len(cluster.Centroid.Positions)-1 {
				fmt.Printf(", ")
			}
		}
		fmt.Printf(")")

		fmt.Println()
		for _, dataPoint := range cluster.DataPoints {
			fmt.Printf("%v (", dataPoint.ChampionName)
			for i, position := range dataPoint.Positions {
				fmt.Printf("%v", position)
				if i != len(dataPoint.Positions)-1 {
					fmt.Printf(", ")
				}
			}
			fmt.Printf(")")
			fmt.Println()
		}
		fmt.Println()
		fmt.Println()
	}
}

func runKMeans(clusters []Cluster, dataPoints []DataPoint) {
	totalClusters := len(clusters)
	totalDimensions := len(dataPoints[0].Positions)
	for {
		didClustersChange := false

		tempClusters := make([]Cluster, len(clusters))
		copy(tempClusters, clusters)
		for i := range tempClusters {
			tempClusters[i].DataPoints = nil
		}

		assignDataPointsToCluster := func() {
			for _, dataPoint := range dataPoints {
				nearestCluster, nearestClusterIndex := &tempClusters[0], 0
				for i := 1; i < totalClusters; i++ {
					if distanceBetweenDataPoints(tempClusters[i].Centroid, dataPoint) < distanceBetweenDataPoints(nearestCluster.Centroid, dataPoint) {
						nearestCluster, nearestClusterIndex = &tempClusters[i], i
					}
				}
				nearestCluster.DataPoints = append(nearestCluster.DataPoints, dataPoint)
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
			for i := range tempClusters {
				sumPositions := make([]float64, totalDimensions)
				for _, dataPoint := range tempClusters[i].DataPoints {
					for i, Positions := range dataPoint.Positions {
						sumPositions[i] += Positions
					}
				}
				avgPositions := make([]float64, totalDimensions)
				totalDataPoints := len(tempClusters[i].DataPoints)
				for i := range avgPositions {
					avgPositions[i] = sumPositions[i] / float64(totalDataPoints)
				}
				tempClusters[i].Centroid = newDataPoint("", avgPositions)
			}
		}
		repositionCentroid()

		copy(clusters, tempClusters)
	}
}

func apiGetClusters(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	laneName := params["laneName"]

	attributesStr := r.FormValue("attributes")
	attributes := strings.Split(attributesStr, ",")

	totalClustersStr := r.FormValue("totalClusters")
	totalClusters, _ := strconv.Atoi(totalClustersStr)

	dataPoints := lolChampionsFileDataPoints("champions.json", laneName, attributes)
	clusters := initialClusters(totalClusters, dataPoints)
	runKMeans(clusters, dataPoints)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(clusters)
}

func main() {
	rand.Seed(time.Now().UnixNano())
	r := mux.NewRouter()
	r.HandleFunc("/api/lanes/{laneName}", apiGetClusters).
		Methods("GET").
		Queries("totalClusters", "{totalClusters}").
		Queries("attributes", "{attributes}")

	log.Fatal(http.ListenAndServe(":8000", r))

	// filename := "champions.json"
	// laneName := "support"
	// attributes := []string{"healing"}
	// totalClusters := 2
	// dataPoints := lolChampionsFileDataPoints(filename, laneName, attributes)
	// clusters := initialClusters(totalClusters, dataPoints)
	// runKMeans(clusters, dataPoints)
	// printResults(dataPoints, clusters)

	// dataPoints := randomDataPoints(2, 10, -20, 20)
	// clusters := initialClusters(4, dataPoints)
	// runKMeans(clusters, dataPoints)
	// printResults(dataPoints, clusters)

}
