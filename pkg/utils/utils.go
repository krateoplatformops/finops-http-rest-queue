package utils

import (
	"encoding/json"
	"fmt"
	"log"
)

type OptimizationRequest struct {
	ResourceId   string       `json:"resourceId"`
	Optimization Optimization `json:"optimization"`
}

type Optimization struct {
	ResourceName  string     `json:"resourceName"`
	ResourceDelta int        `json:"resourceDelta"`
	TypeChange    TypeChange `json:"typeChange"`
}

type TypeChange struct {
	Cyclic string `json:"cyclic"`
	From   string `json:"from"`
	To     string `json:"to"`
}

func ParseOptimization(data []byte) OptimizationRequest {
	parsedData := OptimizationRequest{}
	err := json.Unmarshal(data, &parsedData)
	Fatal(err)

	return parsedData
}

func Fatal(err error) {
	if err != nil {
		log.Fatalln(err)
		fmt.Println(err)
	}
}
