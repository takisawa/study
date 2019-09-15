package main

import (
	"fmt"

	_ "github.com/influxdata/influxdb1-client" // this is important because of the bug in go mod
	client "github.com/influxdata/influxdb1-client/v2"
)

func main() {
	c, err := client.NewHTTPClient(client.HTTPConfig{
		Addr: "http://localhost:8086",
	})
	if err != nil {
		panic(err)
	}
	defer c.Close()

	q := client.NewQuery("SELECT * FROM geos", "climbing", "")
	response, err := c.Query(q)
	if err != nil {
		panic(err)
	}

	if response.Error() != nil {
		panic(response.Error())
	}

	fmt.Printf("%#v\n", response.Results)
}
