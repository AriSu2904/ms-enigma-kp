package connector

import (
	"bytes"
	"github.com/goccy/go-json"
	"io"
	"log"
	"net/http"
)

type ModelRequest struct {
	Id                string  `json:"id"`
	FullName          string  `json:"name"`
	CodingTest        float64 `json:"coding_test"`
	SkillExperience   float64 `json:"sk_exp"`
	MathSoftSkillTest float64 `json:"basic_test"`
	Status            string  `json:"status"`
}

type ModelResponse struct {
	Id              string `json:"id"`
	Name            string `json:"name"`
	PredictedStatus int8   `json:"predicted_status"`
	Status          string `json:"actual_status"`
}

type MachineLearningConnector interface {
	Predict(modelRequest []*ModelRequest) []*ModelResponse
}

type machineLearningConnector struct{}

func NewMachineLearningConnector() MachineLearningConnector {
	return &machineLearningConnector{}
}

func (m *machineLearningConnector) Predict(modelRequest []*ModelRequest) []*ModelResponse {
	serializedRequest, err := json.Marshal(modelRequest)

	if err != nil {
		log.Fatal("Failed to serialize request:", err)
	}

	response, err := http.Post("http://127.0.0.1:3000/predict", "application/json", bytes.NewBuffer(serializedRequest))

	if err != nil {
		log.Fatal("Failed to send request:", err)
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Fatal("Failed to close response body:", err)
		}
	}(response.Body)

	var responses []*ModelResponse
	err = json.NewDecoder(response.Body).Decode(&responses)
	if err != nil {
		log.Fatal("Failed to decode response:", err)
	}
	return responses
}
