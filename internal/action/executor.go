package action

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

// Executor handles the execution of different types of actions
type Executor struct{}

// NewExecutor creates a new action executor
func NewExecutor() *Executor {
	return &Executor{}
}

// Execute processes an action and returns an observation
func (e *Executor) Execute(action Action) Observation {
	// Validate the payload first
	if err := action.Payload.Validate(); err != nil {
		return Observation{
			Success: false,
			Error:   err.Error(),
		}
	}

	switch action.Type {
	case ActionTypeReadFile:
		if payload, ok := action.Payload.(*ReadFilePayload); ok {
			return e.executeReadFile(payload)
		}
	case ActionTypeWriteFile:
		if payload, ok := action.Payload.(*WriteFilePayload); ok {
			return e.executeWriteFile(payload)
		}
	case ActionTypeListDir:
		if payload, ok := action.Payload.(*ListDirPayload); ok {
			return e.executeListDir(payload)
		}
	case ActionTypeExec:
		if payload, ok := action.Payload.(*ExecPayload); ok {
			return e.executeCommand(payload)
		}
	}

	return Observation{
		Success: false,
		Error:   fmt.Sprintf("invalid payload type for action: %s", action.Type),
	}
}

func (e *Executor) executeReadFile(payload *ReadFilePayload) Observation {
	content, err := os.ReadFile(payload.Path)
	if err != nil {
		return Observation{
			Success: false,
			Error:   fmt.Sprintf("error reading file: %v", err),
		}
	}

	return Observation{
		Success: true,
		Data:    string(content),
	}
}

func (e *Executor) executeWriteFile(payload *WriteFilePayload) Observation {
	if err := os.WriteFile(payload.Path, []byte(payload.Content), 0644); err != nil {
		return Observation{
			Success: false,
			Error:   fmt.Sprintf("error writing file: %v", err),
		}
	}

	return Observation{
		Success: true,
		Data:    "file written successfully",
	}
}

func (e *Executor) executeListDir(payload *ListDirPayload) Observation {
	path := payload.Path
	if path == "" {
		path = "."
	}

	entries, err := os.ReadDir(path)
	if err != nil {
		return Observation{
			Success: false,
			Error:   fmt.Sprintf("error listing directory: %v", err),
		}
	}

	var files []map[string]interface{}
	for _, entry := range entries {
		info, err := entry.Info()
		if err != nil {
			continue
		}
		files = append(files, map[string]interface{}{
			"name":     entry.Name(),
			"is_dir":   entry.IsDir(),
			"size":     info.Size(),
			"mod_time": info.ModTime(),
		})
	}

	return Observation{
		Success: true,
		Data:    files,
	}
}

func (e *Executor) executeCommand(payload *ExecPayload) Observation {
	parts := strings.Fields(payload.Command)
	if len(parts) == 0 {
		return Observation{
			Success: false,
			Error:   "invalid command",
		}
	}

	cmd := exec.Command(parts[0], parts[1:]...)
	var stdout, stderr strings.Builder
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	err := cmd.Run()
	if err != nil {
		return Observation{
			Success: false,
			Error:   fmt.Sprintf("command execution failed: %v\nStderr: %s", err, stderr.String()),
		}
	}

	return Observation{
		Success: true,
		Data: map[string]string{
			"stdout": stdout.String(),
			"stderr": stderr.String(),
		},
	}
}
