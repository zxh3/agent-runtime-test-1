package action

import (
	"fmt"
)

// ActionType represents the type of action to be executed
type ActionType string

const (
	ActionTypeReadFile  ActionType = "read_file"
	ActionTypeWriteFile ActionType = "write_file"
	ActionTypeListDir   ActionType = "list_dir"
	ActionTypeExec      ActionType = "exec"
)

// ActionPayload is an interface that all action payloads must implement
type ActionPayload interface {
	Validate() error
}

// Action represents a generic action request
type Action struct {
	Type    ActionType    `json:"type"`
	Payload ActionPayload `json:"payload"`
}

// ReadFilePayload represents the payload for reading a file
type ReadFilePayload struct {
	Path string `json:"path"`
}

func (p ReadFilePayload) Validate() error {
	if p.Path == "" {
		return fmt.Errorf("path is required for read_file action")
	}
	return nil
}

// WriteFilePayload represents the payload for writing to a file
type WriteFilePayload struct {
	Path    string `json:"path"`
	Content string `json:"content"`
}

func (p WriteFilePayload) Validate() error {
	if p.Path == "" {
		return fmt.Errorf("path is required for write_file action")
	}
	return nil
}

// ListDirPayload represents the payload for listing directory contents
type ListDirPayload struct {
	Path string `json:"path"`
}

func (p ListDirPayload) Validate() error {
	return nil // Path is optional, defaults to current directory
}

// ExecPayload represents the payload for executing a command
type ExecPayload struct {
	Command string `json:"command"`
}

func (p ExecPayload) Validate() error {
	if p.Command == "" {
		return fmt.Errorf("command is required for exec action")
	}
	return nil
}

// Observation represents the result of an action execution
type Observation struct {
	Success bool   `json:"success"`
	Error   string `json:"error,omitempty"`
	Data    any    `json:"data,omitempty"`
}
