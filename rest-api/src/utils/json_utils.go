package utils

import (
	"encoding/json"
)

type JsonResponse[T any] struct {
	Data      T      `json:"data"`
	ErrorText string `json:"errorText"`
}

func UnmarshalJsonResponse[T any](data []byte) (JsonResponse[T], error) {
	var t JsonResponse[T]
	err := json.Unmarshal(data, &t)

	if err != nil {
		return t, err
	}

	return t, nil
}

func MarshalJsonResponse[T any](t T) ([]byte, error) {
	jsonResponse := JsonResponse[T]{Data: t, ErrorText: ""}

	jsonData, err := json.Marshal(jsonResponse)

	if err != nil {
		return []byte(""), err
	}

	return jsonData, nil
}

func MarshalJsonErrorResponse(errorText string) ([]byte, error) {
	jsonResponse := JsonResponse[string]{Data: "", ErrorText: errorText}

	jsonData, err := json.Marshal(jsonResponse)

	if err != nil {
		return []byte(""), err
	}

	return jsonData, nil
}
