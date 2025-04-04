package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"agentworkshopruntime/internal/action"
)

func main() {
	executor := action.NewExecutor()

	// Basic authentication middleware
	authMiddleware := func(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			// TODO: Replace with proper authentication
			apiKey := r.Header.Get("X-API-Key")
			if apiKey == "" {
				respondWithError(w, "Unauthorized", http.StatusUnauthorized)
				return
			}
			next(w, r)
		}
	}

	// API specification endpoint (no auth required)
	http.HandleFunc("/api/spec", action.HandleSpec)

	// Universal action endpoint
	http.HandleFunc("/api/execute_action", authMiddleware(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			respondWithError(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		// First decode the action type
		var rawAction struct {
			Type    action.ActionType `json:"type"`
			Payload json.RawMessage   `json:"payload"`
		}
		if err := json.NewDecoder(r.Body).Decode(&rawAction); err != nil {
			respondWithError(w, "Invalid request body", http.StatusBadRequest)
			return
		}

		// Create the appropriate payload based on the action type
		var payload action.ActionPayload
		switch rawAction.Type {
		case action.ActionTypeReadFile:
			payload = &action.ReadFilePayload{}
		case action.ActionTypeWriteFile:
			payload = &action.WriteFilePayload{}
		case action.ActionTypeListDir:
			payload = &action.ListDirPayload{}
		case action.ActionTypeExec:
			payload = &action.ExecPayload{}
		default:
			respondWithError(w, "Unknown action type", http.StatusBadRequest)
			return
		}

		// Unmarshal the payload
		if err := json.Unmarshal(rawAction.Payload, payload); err != nil {
			respondWithError(w, "Invalid payload format", http.StatusBadRequest)
			return
		}

		// Create the action and execute it
		act := action.Action{
			Type:    rawAction.Type,
			Payload: payload,
		}

		observation := executor.Execute(act)
		if !observation.Success {
			respondWithError(w, observation.Error, http.StatusInternalServerError)
			return
		}

		respondWithSuccess(w, observation.Data)
	}))

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server starting on port %s", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}

func respondWithError(w http.ResponseWriter, message string, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(struct {
		Success bool   `json:"success"`
		Error   string `json:"error"`
	}{
		Success: false,
		Error:   message,
	})
}

func respondWithSuccess(w http.ResponseWriter, data any) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(struct {
		Success bool `json:"success"`
		Data    any  `json:"data"`
	}{
		Success: true,
		Data:    data,
	})
}
