package handlers

import (
	"testing"

	"github.com/nahK994/TinyCache/pkg/handlers"
)

func TestHandleCommand(t *testing.T) {
	// Test GET Command
	output, err := handlers.HandleCommand("*2\r\n$3\r\nGET\r\n$3\r\nfoo\r\n")
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}
	if output != "$-1\r\n" { // foo does not exist in cache
		t.Errorf("expected $-1\r\n, got %v", output)
	}

	// Test SET Command
	output, err = handlers.HandleCommand("*3\r\n$3\r\nSET\r\n$3\r\nfoo\r\n$3\r\nbar\r\n")
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}
	if output != "+OK\r\n" {
		t.Errorf("expected +OK\r\n, got %v", output)
	}

	// Test GET after SET (string)
	output, err = handlers.HandleCommand("*2\r\n$3\r\nGET\r\n$3\r\nfoo\r\n")
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}
	if output != "$3\r\nbar\r\n" {
		t.Errorf("expected $3\r\nbar\r\n, got %v", output)
	}

	// Test EXISTS Command
	output, err = handlers.HandleCommand("*2\r\n$6\r\nEXISTS\r\n$3\r\nfoo\r\n")
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}
	if output != ":1\r\n" {
		t.Errorf("expected :1\r\n, got %v", output)
	}

	// Test EXISTS Command
	output, err = handlers.HandleCommand("*2\r\n$6\r\nEXISTS\r\n$3\r\nbar\r\n")
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}
	if output != ":0\r\n" {
		t.Errorf("expected :1\r\n, got %v", output)
	}

	// Test INCR Command on a new key
	output, err = handlers.HandleCommand("*2\r\n$4\r\nINCR\r\n$3\r\nnum\r\n")
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}
	if output != ":1\r\n" {
		t.Errorf("expected :1\r\n, got %v", output)
	}

	// Test INCR Command on existing key
	output, err = handlers.HandleCommand("*2\r\n$4\r\nINCR\r\n$3\r\nnum\r\n")
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}
	if output != ":2\r\n" {
		t.Errorf("expected :2\r\n, got %v", output)
	}

	// Test INCR Command on non-int data
	_, err = handlers.HandleCommand("*2\r\n$4\r\nINCR\r\n$3\r\nfoo\r\n")
	if err == nil {
		t.Errorf("expected error, got %v", err)
	}

	// Test DECR Command on existing key
	output, err = handlers.HandleCommand("*2\r\n$4\r\nDECR\r\n$3\r\nnum\r\n")
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}
	if output != ":1\r\n" {
		t.Errorf("expected :1\r\n, got %v", output)
	}

	// Test DECR Command on non-int data
	_, err = handlers.HandleCommand("*2\r\n$4\r\nDECR\r\n$3\r\nfoo\r\n")
	if err == nil {
		t.Errorf("expected error, got %v", err)
	}

	// Test DEL Command
	output, err = handlers.HandleCommand("*2\r\n$3\r\nDEL\r\n$3\r\nfoo\r\n")
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}
	if output != ":1\r\n" {
		t.Errorf("expected :1\r\n, got %v", output)
	}

	// Test DEL on non-existing key
	output, err = handlers.HandleCommand("*2\r\n$3\r\nDEL\r\n$3\r\nfoo\r\n")
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}
	if output != ":0\r\n" {
		t.Errorf("expected :0\r\n, got %v", output)
	}

	// Test GET after SET (int)
	output, err = handlers.HandleCommand("*2\r\n$3\r\nGET\r\n$3\r\nnum\r\n")
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}
	if output != "$1\r\n1\r\n" {
		t.Errorf("expected $1\r\n1\r\n, got %v", output)
	}

	// Test PING Command
	output, err = handlers.HandleCommand("*1\r\n$4\r\nPING\r\n")
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}
	if output != "+PONG\r\n" {
		t.Errorf("expected +PONG\r\n, got %v", output)
	}
}
