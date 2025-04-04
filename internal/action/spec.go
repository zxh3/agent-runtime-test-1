package action

import (
	"encoding/json"
	"net/http"
)

// ActionSpec represents the specification of an action
type ActionSpec struct {
	Type        string                 `json:"type"`
	Description string                 `json:"description"`
	Payload     map[string]FieldSpec   `json:"payload"`
	Example     map[string]interface{} `json:"example"`
}

// FieldSpec represents the specification of a field in an action payload
type FieldSpec struct {
	Type        string `json:"type"`
	Description string `json:"description"`
	Required    bool   `json:"required"`
}

// GenerateSpec generates the API specification for all available actions
func GenerateSpec() map[string]ActionSpec {
	return map[string]ActionSpec{
		string(ActionTypeReadFile): {
			Type:        string(ActionTypeReadFile),
			Description: "Read the contents of a file",
			Payload: map[string]FieldSpec{
				"path": {
					Type:        "string",
					Description: "Path to the file to read",
					Required:    true,
				},
			},
			Example: map[string]interface{}{
				"type": ActionTypeReadFile,
				"payload": map[string]interface{}{
					"path": "/path/to/file.txt",
				},
			},
		},
		string(ActionTypeWriteFile): {
			Type:        string(ActionTypeWriteFile),
			Description: "Write content to a file",
			Payload: map[string]FieldSpec{
				"path": {
					Type:        "string",
					Description: "Path to the file to write",
					Required:    true,
				},
				"content": {
					Type:        "string",
					Description: "Content to write to the file",
					Required:    true,
				},
			},
			Example: map[string]interface{}{
				"type": ActionTypeWriteFile,
				"payload": map[string]interface{}{
					"path":    "/path/to/file.txt",
					"content": "Hello, World!",
				},
			},
		},
		string(ActionTypeListDir): {
			Type:        string(ActionTypeListDir),
			Description: "List contents of a directory",
			Payload: map[string]FieldSpec{
				"path": {
					Type:        "string",
					Description: "Path to the directory to list (optional, defaults to current directory)",
					Required:    false,
				},
			},
			Example: map[string]interface{}{
				"type": ActionTypeListDir,
				"payload": map[string]interface{}{
					"path": "/path/to/directory",
				},
			},
		},
		string(ActionTypeExec): {
			Type:        string(ActionTypeExec),
			Description: "Execute a shell command",
			Payload: map[string]FieldSpec{
				"command": {
					Type:        "string",
					Description: "Command to execute",
					Required:    true,
				},
			},
			Example: map[string]interface{}{
				"type": ActionTypeExec,
				"payload": map[string]interface{}{
					"command": "ls -la",
				},
			},
		},
	}
}

// HandleSpec handles the API specification endpoint
func HandleSpec(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	spec := GenerateSpec()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(spec)
}
