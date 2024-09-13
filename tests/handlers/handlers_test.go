package handlers

import (
	"testing"

	"github.com/nahK994/TinyCache/pkg/handlers"
)

func TestHandlers(t *testing.T) {
	// Test handleSET and handleGET
	t.Run("TestHandleSETAndGET", func(t *testing.T) {
		// Test SET
		resp, err := handlers.HandleCommand("*3\r\n$3\r\nSET\r\n$4\r\nname\r\n$5\r\nShomi\r\n")
		if err != nil {
			t.Errorf("handleSET returned an error: %v", err)
		}
		expectedResp := "+OK\r\n"
		if resp != expectedResp {
			t.Errorf("Expected '%s', got %s", expectedResp, resp)
		}

		// Test GET
		resp, err = handlers.HandleCommand("*2\r\n$3\r\nGET\r\n$4\r\nname\r\n")
		if err != nil {
			t.Errorf("handleGET returned an error: %v", err)
		}
		expectedResp = "$5\r\nShomi\r\n"
		if resp != expectedResp {
			t.Errorf("Expected '%s', got %s", expectedResp, resp)
		}
	})

	// Test INCR and DECR
	t.Run("TestHandleINCRAndDECR", func(t *testing.T) {
		// Test INCR
		resp, err := handlers.HandleCommand("*2\r\n$4\r\nINCR\r\n$3\r\nage\r\n")
		if err != nil {
			t.Errorf("handleINCR returned an error: %v", err)
		}
		expectedResp := ":1\r\n" // Since the age is set to 1 as it's initially absent
		if resp != expectedResp {
			t.Errorf("Expected '%s', got %s", expectedResp, resp)
		}

		// Test DECR
		resp, err = handlers.HandleCommand("*2\r\n$4\r\nDECR\r\n$3\r\nage\r\n")
		if err != nil {
			t.Errorf("handleDECR returned an error: %v", err)
		}
		expectedResp = ":0\r\n"
		if resp != expectedResp {
			t.Errorf("Expected '%s', got %s", expectedResp, resp)
		}
	})

	// Test DEL
	t.Run("TestHandleDEL", func(t *testing.T) {
		// First, set a key
		handlers.HandleCommand("*3\r\n$3\r\nSET\r\n$3\r\nkey\r\n$5\r\nvalue\r\n")

		// Test DEL
		resp, err := handlers.HandleCommand("*2\r\n$3\r\nDEL\r\n$3\r\nkey\r\n")
		if err != nil {
			t.Errorf("handleDEL returned an error: %v", err)
		}
		expectedResp := ":1\r\n"
		if resp != expectedResp {
			t.Errorf("Expected '%s', got %s", expectedResp, resp)
		}

		// Try to get the deleted key
		resp, err = handlers.HandleCommand("*2\r\n$3\r\nGET\r\n$3\r\nkey\r\n")
		if err != nil {
			t.Errorf("handleGET returned an error: %v", err)
		}
		expectedResp = "$-1\r\n"
		if resp != expectedResp {
			t.Errorf("Expected '%s', got %s", expectedResp, resp)
		}
	})
}
