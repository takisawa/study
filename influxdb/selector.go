package main

import (
	"fmt"

	_ "github.com/influxdata/influxdb1-client" // this is important because of the bug in go mod
	influxdb "github.com/influxdata/influxdb1-client/v2"
)

func main() {
	client, err := influxdb.NewHTTPClient(influxdb.HTTPConfig{
		Addr: "http://localhost:8086",
	})
	if err != nil {
		panic(err)
	}
	defer client.Close()

	q := influxdb.NewQuery("SELECT * FROM geos", "climbing", "")
	response, err := client.Query(q)
	if err != nil {
		panic(err)
	}

	if response.Error() != nil {
		panic(response.Error())
	}

	fmt.Printf("%#v\n", response.Results)
}
