package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"

	_ "github.com/influxdata/influxdb1-client" // this is important because of the bug in go mod
	influxdb "github.com/influxdata/influxdb1-client/v2"
)

func main() {
	client, err := influxdb.NewHTTPClient(influxdb.HTTPConfig{
		Addr: "http://localhost:8086",
	})
	if err != nil {
		fmt.Println("Error creating InfluxDB Client: ", err.Error())
	}
	defer client.Close()

	var (
		shapes     = []string{"circle", "rectangle", "square", "triangle"}
		colors     = []string{"red", "blue", "green"}
		sampleSize = 3
		pts        = make([]*influxdb.Point, sampleSize)
	)

	rand.Seed(time.Now().UnixNano())

	t := time.Date(2019, 9, 16, 10, 5, 28, 0, time.UTC)

	for i := 0; i < sampleSize; i++ {
		pts[i], err = influxdb.NewPoint(
			"geos",
			map[string]string{
				"color": strconv.Itoa(rand.Intn(len(colors))),
				"shape": strconv.Itoa(rand.Intn(len(shapes))),
			},
			map[string]interface{}{
				"value": rand.Intn(sampleSize),
			},
			t,
		)
		if err != nil {
			panic(err)
		}
	}

	batchPointsConfig := influxdb.BatchPointsConfig{
		Precision: "s",
		Database:  "climbing",
	}

	bp, err := influxdb.NewBatchPoints(batchPointsConfig)
	if err != nil {
		panic(err)
	}

	bp.AddPoints(pts)

	err = client.Write(bp)
	if err != nil {
		panic(err)
	}
}
