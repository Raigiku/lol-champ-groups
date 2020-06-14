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

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

// Champion league of legends
type Champion struct {
	Name  string               `json:"name"`
	Lanes []map[string]float64 `json:"lanes"`
}

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

func euclideanDistanceBetweenDataPoints(dataPoint1, dataPoint2 DataPoint) float64 {
	sum := 0.0
	for i := range dataPoint1.Positions {
		sum += math.Pow(dataPoint1.Positions[i]-dataPoint2.Positions[i], 2)
	}
	distance := math.Sqrt(sum)
	return distance
}

func manhattanDistanceBetweenDataPoints(dataPoint1, dataPoint2 DataPoint) float64 {
	distance := 0.0
	for i := range dataPoint1.Positions {
		distance += math.Abs(dataPoint1.Positions[i] - dataPoint2.Positions[i])
	}
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

func parallelLocalSum(sumsChannel chan []float64, endChannel chan bool, dataPoints []DataPoint, id, totalDimensions, totalThreads int) {
	localSum := make([]float64, totalDimensions)
	for i := id; i < len(dataPoints); i += totalThreads {
		for j, position := range dataPoints[i].Positions {
			localSum[j] += position
		}
	}
	sumsChannel <- localSum
	endChannel <- true
}

func parallelDataPointSum(totalThreads, totalDimensions int, sumsChannel chan []float64, endChannel chan bool, dataPoints []DataPoint) {
	for th := 0; th < totalThreads; th++ {
		go parallelLocalSum(sumsChannel, endChannel, dataPoints, th, totalDimensions, totalThreads)
	}
	for i := 0; i < totalThreads; i++ {
		<-endChannel
	}
	close(sumsChannel)
}

func repositionCentroid(finishChannel chan bool, tempClusters []Cluster, totalDimensions int) {
	numThreads := 12
	for k := range tempClusters {
		// Canal que almacenara 12 numeros que representan la suma de las posiciones
		// de los datapoints del cluster actual, cada suma es calculado por cada hilo
		sumsChannel := make(chan []float64)
		endChannel := make(chan bool)
		go parallelDataPointSum(numThreads, totalDimensions, sumsChannel, endChannel, tempClusters[k].DataPoints)

		// Obteniendo la suma de todas las posiciones de los datapoints del cluster actual
		totalSum := make([]float64, totalDimensions)
		for sum := range sumsChannel {
			for i, position := range sum {
				totalSum[i] += position
			}
		}

		// Obteniendo la posicion promedio calculada con la suma de todas las posiciones y
		// dividiendo cada dimension por el total de datapoints en el cluster
		avgPositions := make([]float64, totalDimensions)
		totalDataPoints := len(tempClusters[k].DataPoints)
		for i := range avgPositions {
			avgPositions[i] = totalSum[i] / float64(totalDataPoints)
		}
		// Reposicionando el centroide
		tempClusters[k].Centroid = newDataPoint("", avgPositions)
	}
	finishChannel <- true
}

func runKMeans(distanceMethod func(dataPoint1, dataPoint2 DataPoint) float64, clusters []Cluster, dataPoints []DataPoint) {
	totalClusters := len(clusters)
	totalDimensions := len(dataPoints[0].Positions)
	for {
		didClustersChange := false

		// Creando copia para luego verificar si los clusters han cambiado
		tempClusters := make([]Cluster, len(clusters))
		copy(tempClusters, clusters)
		for i := range tempClusters {
			tempClusters[i].DataPoints = nil
		}

		// Asignando datapoints hacia clusters
		assignDataPointsToCluster := func() {
			for _, dataPoint := range dataPoints {
				nearestCluster, nearestClusterIndex := &tempClusters[0], 0
				for i := 1; i < totalClusters; i++ {
					if distanceMethod(tempClusters[i].Centroid, dataPoint) < distanceMethod(nearestCluster.Centroid, dataPoint) {
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

		finishChannel := make(chan bool)
		go repositionCentroid(finishChannel, tempClusters, totalDimensions)
		<-finishChannel

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

	var distanceMethod func(dataPoint1, dataPoint2 DataPoint) float64
	distanceMethodStr := r.FormValue("distanceMethod")
	if distanceMethodStr == "manhattan" {
		distanceMethod = manhattanDistanceBetweenDataPoints
	} else if distanceMethodStr == "euclidean" {
		distanceMethod = euclideanDistanceBetweenDataPoints
	}
	runKMeans(distanceMethod, clusters, dataPoints)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(clusters)
}

func main() {
	rand.Seed(time.Now().UnixNano())
	r := mux.NewRouter()
	r.HandleFunc("/api/lanes/{laneName}", apiGetClusters).
		Methods("GET").
		Queries("totalClusters", "{totalClusters}").
		Queries("attributes", "{attributes}").
		Queries("distanceMethod", "{distanceMethod}")

	headersOk := handlers.AllowedHeaders([]string{"*"})
	originsOk := handlers.AllowedOrigins([]string{"*"})
	methodsOk := handlers.AllowedMethods([]string{"GET"})

	log.Fatal(http.ListenAndServe(":8000", handlers.CORS(headersOk, originsOk, methodsOk)(r)))

	// filename := "champions.json"
	// laneName := "jungle"
	// attributes := []string{"enemy_jungle_cs"}
	// totalClusters := 4
	// dataPoints := lolChampionsFileDataPoints(filename, laneName, attributes)
	// clusters := initialClusters(totalClusters, dataPoints)
	// runKMeans(clusters, dataPoints)
	// printResults(dataPoints, clusters)

	// dataPoints := randomDataPoints(2, 10, -20, 20)
	// clusters := initialClusters(3, dataPoints)
	// runKMeans(euclideanDistanceBetweenDataPoints, clusters, dataPoints)
	// printResults(dataPoints, clusters)

}
