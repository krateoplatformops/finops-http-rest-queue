package main

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/nats-io/nats.go"

	"github.com/krateoplatformops/finops-http-rest-queue/pkg/utils"
)

var (
	c *nats.EncodedConn
)

func upload(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		data, err := io.ReadAll(r.Body)
		utils.Fatal(err)

		topic := r.URL.Query().Get("topic")

		fmt.Println("Upload received")
		w.Write([]byte("Upload received"))

		parsedData := utils.ParseOptimization(data)
		err = publish(parsedData, topic)
		utils.Fatal(err)
	} else {
		http.Error(w, "only POST allowed", http.StatusMethodNotAllowed)
	}
}

func publish(parsedData utils.OptimizationRequest, topic string) error {
	fmt.Println("Publishing on topic:", topic)
	err := c.Publish(topic, &parsedData)
	if err != nil {
		return err
	}
	return c.Flush()
}

func main() {
	fmt.Println("Connecting to NATS server in:", os.Getenv("NATS_SERVICE_HOST")+":"+os.Getenv("NATS_SERVICE_PORT"))
	nc, err := nats.Connect(os.Getenv("NATS_SERVICE_HOST") + ":" + os.Getenv("NATS_SERVICE_PORT"))
	if err != nil {
		utils.Fatal(err)
	}
	defer nc.Close()

	c, err = nats.NewEncodedConn(nc, nats.JSON_ENCODER)
	if err != nil {
		utils.Fatal(err)
	}
	defer c.Close()

	http.HandleFunc("/upload", upload)
	utils.Fatal(http.ListenAndServe(":8080", nil))
}
