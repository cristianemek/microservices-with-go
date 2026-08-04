package main

import (
	"encoding/json"
	"log"
	"net/http"
	"ride-sharing/services/api-gateway/grpc_clients"
	"ride-sharing/shared/contracts"
)

func handleTripStart(w http.ResponseWriter, r *http.Request) {
	var reqBody startTripRequest

	decoder := json.NewDecoder(r.Body)

	decoder.DisallowUnknownFields()

	if err := decoder.Decode(&reqBody); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	defer r.Body.Close()
	//validation

	if reqBody.UserID == "" {
		http.Error(w, "user ID is required", http.StatusBadRequest)
		return
	}

	tripService, err := grpc_clients.NewTripServiceClient()
	if err != nil {
		log.Fatal(err)
	}

	defer tripService.Close()

	trip, err := tripService.Client.CreateTrip(r.Context(), reqBody.toProto())

	if err != nil {
		log.Printf("Failed to start a trip: %v", err)
		http.Error(w, "Failed to start trip", http.StatusInternalServerError)
		return
	}

	response := contracts.APIResponse{Data: trip}

	writeJson(w, http.StatusCreated, response)
}

func handleTripPreview(w http.ResponseWriter, r *http.Request) {
	var reqBody previewTripRequest

	decoder := json.NewDecoder(r.Body)

	decoder.DisallowUnknownFields()

	if err := decoder.Decode(&reqBody); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	defer r.Body.Close()
	//validation

	if reqBody.UserID == "" {
		http.Error(w, "user ID is required", http.StatusBadRequest)
		return
	}

	tripService, err := grpc_clients.NewTripServiceClient()
	if err != nil {
		log.Fatal(err)
	}

	defer tripService.Close()

	tripPreview, err := tripService.Client.PreviewTrip(r.Context(), reqBody.toProto())

	if err != nil {
		log.Printf("Failed to preview a trip: %v", err)
		http.Error(w, "Failed to preview trip", http.StatusInternalServerError)
		return
	}

	response := contracts.APIResponse{Data: tripPreview}

	writeJson(w, http.StatusCreated, response)
}
