package main

import (
	"bytes"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"net/http/httputil"
	"time"

	"github.com/gogo/protobuf/proto"
	"github.com/golang/snappy"
	"github.com/prometheus/prometheus/prompb"
)

func main() {
	var s []prompb.TimeSeries
	for i := 0; i < 10; i++ {
		rand.Seed(time.Now().UnixMicro() + int64(i))
		s = append(s, prompb.TimeSeries{
			Labels: []prompb.Label{
				{Name: "__name__", Value: "testseries"},
				{Name: "foo", Value: fmt.Sprintf("bar%03d", i)},
			},
			Samples: []prompb.Sample{
				{Value: rand.Float64(), Timestamp: time.Now().Add(-10 * time.Minute).Add(time.Duration(i) * time.Second).UnixMilli()},
			},
		})
	}

	data, err := proto.Marshal(&prompb.WriteRequest{Timeseries: s})
	if err != nil {
		log.Fatal(err)
	}

	encoded := snappy.Encode(nil, data)

	body := bytes.NewReader(encoded)
	req, err := http.NewRequest("POST", "http://localhost:9009/api/v1/push", body)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/x-protobuf")
	req.Header.Set("Content-Encoding", "snappy")
	req.Header.Set("User-Agent", "testwriter/1.0.0")
	req.Header.Set("X-Prometheus-Remote-Write-Version", "0.1.0")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	dr, _ := httputil.DumpResponse(resp, true)
	fmt.Println(string(dr))
}
