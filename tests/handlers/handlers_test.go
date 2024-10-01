package handlers

import (
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/nahK994/TinyCache/pkg/config"
	"github.com/nahK994/TinyCache/pkg/handlers"
)

func TestHandleGET(t *testing.T) {
	// Key does not exist
	_, err := handlers.HandleCommand("*2\r\n$3\r\nGET\r\n$17\r\nnon_existing_key\r\n")
	if err == nil {
		t.Errorf("Expected error for non-existing key, got none")
	}

	// Key exists, type int
	handlers.HandleCommand("*3\r\n$3\r\nSET\r\n$6\r\nnumber\r\n$2\r\n10\r\n")
	resp, err := handlers.HandleCommand("*2\r\n$3\r\nGET\r\n$6\r\nnumber\r\n")
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if resp != ":10\r\n" {
		t.Errorf("Expected :10\r\n, got %s", resp)
	}

	// Key exists, type string
	handlers.HandleCommand("*3\r\n$3\r\nSET\r\n$5\r\nmykey\r\n$5\r\nhello\r\n")
	resp, err = handlers.HandleCommand("*2\r\n$3\r\nGET\r\n$5\r\nmykey\r\n")
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if resp != "$5\r\nhello\r\n" {
		t.Errorf("Expected '$5\\r\\nhello\\r\\n', got %s", resp)
	}
}

func TestHandleSET(t *testing.T) {
	resp, err := handlers.HandleCommand("*3\r\n$3\r\nSET\r\n$5\r\nmykey\r\n$5\r\nvalue\r\n")
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if resp != "+OK\r\n" {
		t.Errorf("Expected '+OK\\r\\n', got %s", resp)
	}

	resp, err = handlers.HandleCommand("*2\r\n$3\r\nGET\r\n$5\r\nmykey\r\n")
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if resp != "$5\r\nvalue\r\n" {
		t.Errorf("Expected '$5\\r\\nvalue\\r\\n', got %s", resp)
	}

	resp, err = handlers.HandleCommand("*4\r\n$3\r\nSET\r\n$4\r\nname\r\n$10\r\nShomi Khan\r\n$2\r\n10\r\n")
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if resp != "+OK\r\n" {
		t.Errorf("Expected '+OK\\r\\n', got %s", resp)
	}

	resp, err = handlers.HandleCommand("*2\r\n$3\r\nGET\r\n$4\r\nname\r\n")
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if resp != "$10\r\nShomi Khan\r\n" {
		t.Errorf("Expected '$10\\r\\nShomi Khan\\r\\n', got %s", resp)
	}

	time.Sleep(10 * time.Second)
	_, err_ttl := handlers.HandleCommand("*2\r\n$3\r\nTTL\r\n$4\r\nname\r\n")
	if err_ttl == nil {
		t.Errorf("Expected Expire key error, got %v", err)
	}
}

func TestHandleEXISTS(t *testing.T) {
	// Key does not exist
	resp, err := handlers.HandleCommand("*2\r\n$6\r\nEXISTS\r\n$6\r\nno_key\r\n")
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if resp != ":0\r\n" {
		t.Errorf("Expected ':0\\r\\n', got %s", resp)
	}

	// Key exists
	handlers.HandleCommand("*3\r\n$3\r\nSET\r\n$5\r\nmykey\r\n$5\r\nvalue\r\n")
	resp, err = handlers.HandleCommand("*2\r\n$6\r\nEXISTS\r\n$5\r\nmykey\r\n")
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if resp != ":1\r\n" {
		t.Errorf("Expected ':1\\r\\n', got %s", resp)
	}
}

func TestHandleDEL(t *testing.T) {
	// Key does not exist
	resp, err := handlers.HandleCommand("*2\r\n$3\r\nDEL\r\n$6\r\nno_key\r\n")
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if resp != ":0\r\n" {
		t.Errorf("Expected ':0\\r\\n', got %s", resp)
	}

	// Key exists, and is deleted
	handlers.HandleCommand("*3\r\n$3\r\nSET\r\n$5\r\nmykey\r\n$5\r\nvalue\r\n")
	resp, err = handlers.HandleCommand("*2\r\n$3\r\nDEL\r\n$5\r\nmykey\r\n")
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if resp != ":1\r\n" {
		t.Errorf("Expected ':1\\r\\n', got %s", resp)
	}

	// Key should no longer exist
	resp, err = handlers.HandleCommand("*2\r\n$6\r\nEXISTS\r\n$5\r\nmykey\r\n")
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if resp != ":0\r\n" {
		t.Errorf("Expected ':0\\r\\n', got %s", resp)
	}
}

func TestHandleINCR_DECR(t *testing.T) {
	handlers.HandleCommand("*3\r\n$3\r\nSET\r\n$6\r\nnewkey\r\n$2\r\n11\r\n")
	resp, err := handlers.HandleCommand("*2\r\n$4\r\nINCR\r\n$6\r\nnewkey\r\n")
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if resp != ":12\r\n" {
		t.Errorf("Expected ':12\\r\\n', got %s", resp)
	}
	handlers.HandleCommand("*2\r\n$3\r\nDEL\r\n$6\r\nnewkey\r\n")

	resp, err = handlers.HandleCommand("*2\r\n$4\r\nINCR\r\n$6\r\nnewkey\r\n")
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if resp != ":1\r\n" {
		t.Errorf("Expected ':1\\r\\n', got %s", resp)
	}

	resp, err = handlers.HandleCommand("*2\r\n$4\r\nINCR\r\n$6\r\nnewkey\r\n")
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if resp != ":2\r\n" {
		t.Errorf("Expected ':2\\r\\n', got %s", resp)
	}

	handlers.HandleCommand("*3\r\n$3\r\nSET\r\n$5\r\nmykey\r\n$5\r\nhello\r\n")
	_, err = handlers.HandleCommand("*2\r\n$4\r\nINCR\r\n$5\r\nmykey\r\n")
	if err == nil {
		t.Errorf("Expected error for INCR on non-integer key, got none")
	}

	resp, _ = handlers.HandleCommand("*2\r\n$4\r\nDECR\r\n$6\r\nnewkey\r\n")
	if resp != ":1\r\n" {
		t.Errorf("Expected ':1\\r\\n', got %s", resp)
	}

	handlers.HandleCommand("*5\r\n$5\r\nLPUSH\r\n$4\r\nlist\r\n$1\r\na\r\n$1\r\nb\r\n$1\r\nc\r\n")
	_, err = handlers.HandleCommand("*2\r\n$4\r\nINCR\r\n$4\r\nlist\r\n")
	if err == nil {
		t.Errorf("Expected type error, got %v", err)
	}

	handlers.HandleCommand("*5\r\n$5\r\nLPUSH\r\n$4\r\nlist\r\n$1\r\na\r\n$1\r\nb\r\n$1\r\nc\r\n")
	_, err = handlers.HandleCommand("*2\r\n$4\r\nDECR\r\n$4\r\nlist\r\n")
	if err == nil {
		t.Errorf("Expected type error, got %v", err)
	}
}

func TestHandlePUSH(t *testing.T) {
	resp, err := handlers.HandleCommand("*5\r\n$5\r\nLPUSH\r\n$6\r\nmylist\r\n$1\r\na\r\n$1\r\nb\r\n$1\r\nc\r\n")
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if resp != ":3\r\n" {
		t.Errorf("Expected ':3\\r\\n', got %s", resp)
	}

	resp, err = handlers.HandleCommand("*5\r\n$5\r\nRPUSH\r\n$6\r\nmylist\r\n$1\r\na\r\n$1\r\nb\r\n$1\r\nc\r\n")
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if resp != ":6\r\n" {
		t.Errorf("Expected ':6\\r\\n', got %s", resp)
	}

	_, err = handlers.HandleCommand("*3\r\n$3\r\nSET\r\n$6\r\nmylist\r\n$3\r\nval\r\n")
	if err != nil {
		t.Errorf("Expected type error, got %v", err)
	}

	_, err = handlers.HandleCommand("*5\r\n$5\r\nRPUSH\r\n$6\r\nmylist\r\n$1\r\na\r\n$1\r\nb\r\n$1\r\nc\r\n")
	if err == nil {
		t.Errorf("Expected no error, got %v", err)
	}
}

func TestHandlePOP(t *testing.T) {
	handlers.HandleCommand("*5\r\n$5\r\nLPUSH\r\n$7\r\nmylist2\r\n$3\r\none\r\n$3\r\ntwo\r\n$5\r\nthree\r\n")
	handlers.HandleCommand("*5\r\n$5\r\nRPUSH\r\n$7\r\nmylist2\r\n$1\r\n1\r\n$1\r\n2\r\n$1\r\n3\r\n")
	response, err := handlers.HandleCommand("*2\r\n$4\r\nLPOP\r\n$7\r\nmylist2\r\n")
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	expected := "$5\r\nthree\r\n"
	if response != expected {
		t.Errorf("Expected %s, got %s", expected, response)
	}

	response, err = handlers.HandleCommand("*2\r\n$4\r\nRPOP\r\n$7\r\nmylist2\r\n")
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	expected = "$1\r\n3\r\n"
	if response != expected {
		t.Errorf("Expected %s, got %s", expected, response)
	}

	response, err = handlers.HandleCommand("*2\r\n$4\r\nLPOP\r\n$7\r\nmylist2\r\n")
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	expected = "$3\r\ntwo\r\n"
	if response != expected {
		t.Errorf("Expected %s, got %s", expected, response)
	}

	response, err = handlers.HandleCommand("*2\r\n$4\r\nRPOP\r\n$7\r\nmylist2\r\n")
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	expected = "$1\r\n2\r\n"
	if response != expected {
		t.Errorf("Expected %s, got %s", expected, response)
	}

	response, err = handlers.HandleCommand("*2\r\n$4\r\nLPOP\r\n$7\r\nmylist2\r\n")
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	expected = "$3\r\none\r\n"
	if response != expected {
		t.Errorf("Expected %s, got %s", expected, response)
	}

	response, err = handlers.HandleCommand("*2\r\n$4\r\nRPOP\r\n$7\r\nmylist2\r\n")
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	expected = "$1\r\n1\r\n"
	if response != expected {
		t.Errorf("Expected %s, got %s", expected, response)
	}

	_, err = handlers.HandleCommand("*2\r\n$4\r\nLPOP\r\n$7\r\nmylist2\r\n")
	if err == nil {
		t.Errorf("Expected empty list error, got %v", err)
	}
}

func TestHandleLRANGE(t *testing.T) {
	_, err := handlers.HandleCommand("*4\r\n$6\r\nLRANGE\r\n$7\r\nmylist1\r\n$1\r\n0\r\n$2\r\n-1\r\n")
	if err == nil {
		t.Errorf("Expected empty list error, got %v", err)
	}

	handlers.HandleCommand("*5\r\n$5\r\nLPUSH\r\n$7\r\nmylist1\r\n$1\r\na\r\n$1\r\nb\r\n$1\r\nc\r\n")
	resp, err1 := handlers.HandleCommand("*4\r\n$6\r\nLRANGE\r\n$7\r\nmylist1\r\n$1\r\n0\r\n$2\r\n-1\r\n")
	if err1 != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if !contains(resp, "$1\r\na\r\n") {
		t.Errorf("Expected response to contain '$1\\r\\na\\r\\n', got %s", resp)
	}

}

func TestHandleEXPIRE(t *testing.T) {
	_, err := handlers.HandleCommand("*3\r\n$6\r\nEXPIRE\r\n$7\r\nexp_key\r\n$1\r\n5\r\n")
	if err == nil {
		t.Errorf("Expected type err, got %v", err)
	}

	_, err = handlers.HandleCommand("*3\r\n$3\r\nSET\r\n$7\r\nexp_key\r\n$5\r\nvalue\r\n")
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	resp, err := handlers.HandleCommand("*3\r\n$6\r\nEXPIRE\r\n$7\r\nexp_key\r\n$1\r\n5\r\n")
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if resp != "+OK\r\n" { // 1 if the key was set to expire
		t.Errorf("Expected +OK\r\n, got %s", resp)
	}

	// Check the TTL
	resp, err = handlers.HandleCommand("*2\r\n$3\r\nTTL\r\n$7\r\nexp_key\r\n")
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if resp != ":4\r\n" { // Should reflect the remaining time
		t.Errorf("Expected TTL response to be ':5\\r\\n', got %s", resp)
	}

	// Wait for expiration
	time.Sleep(6 * time.Second)
	_, err = handlers.HandleCommand("*2\r\n$3\r\nTTL\r\n$7\r\nexp_key\r\n")
	if err == nil {
		t.Errorf("Expected type error error, got %v", err)
	}
}

// Helper function to check if a string contains another string
func contains(resp, substr string) bool {
	return strings.Contains(resp, substr)
}

func TestHandleTTL(t *testing.T) {
	_, err := handlers.HandleCommand("*2\r\n$3\r\nTTL\r\n$7\r\nttl_key\r\n")
	if err == nil {
		t.Errorf("Expected type error, got %v", err)
	}

	handlers.HandleCommand("*3\r\n$3\r\nSET\r\n$7\r\nttl_key\r\n$5\r\nvalue\r\n")
	response, _ := handlers.HandleCommand("*2\r\n$3\r\nTTL\r\n$7\r\nttl_key\r\n")
	expected := ":0\r\n"
	if response != expected {
		t.Errorf("Expected %s, got %s", expected, response)
	}
	handlers.HandleCommand("*3\r\n$6\r\nEXPIRE\r\n$7\r\nttl_key\r\n$2\r\n10\r\n")

	// Check TTL value
	resp, err := handlers.HandleCommand("*2\r\n$3\r\nTTL\r\n$7\r\nttl_key\r\n")
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	// Expected response should be a positive integer representing remaining seconds
	if !isPositiveInteger(resp) {
		t.Errorf("Expected positive integer, got %s", resp)
	}

	// Wait for the key to expire
	time.Sleep(10 * time.Second)
	_, err = handlers.HandleCommand("*2\r\n$3\r\nTTL\r\n$7\r\nttl_key\r\n")
	if err == nil {
		t.Errorf("Expected type error, got %v", err)
	}
}

// Helper function to check if the response is a positive integer
func isPositiveInteger(resp string) bool {
	if len(resp) < 3 {
		return false
	}
	if resp[0] != ':' {
		return false
	}
	num, err := strconv.Atoi(resp[1 : len(resp)-2]) // Remove ':', and CRLF
	return err == nil && num > 0
}

func TestHandlePERSIST(t *testing.T) {
	_, err := handlers.HandleCommand("*2\r\n$7\r\nPERSIST\r\n$11\r\npersist_key\r\n")
	if err == nil {
		t.Errorf("Expected type error, got %v", err)
	}

	handlers.HandleCommand("*3\r\n$3\r\nSET\r\n$11\r\npersist_key\r\n$3\r\nval\r\n")
	resp, err := handlers.HandleCommand("*2\r\n$7\r\nPERSIST\r\n$11\r\npersist_key\r\n")
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if resp != "+OK\r\n" {
		t.Errorf("Expected +OK\r\n, got %s", resp)
	}

	resp, err = handlers.HandleCommand("*2\r\n$3\r\nTTL\r\n$11\r\npersist_key\r\n")
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if resp != ":0\r\n" {
		t.Errorf("Expected :0\r\n, got %s", resp)
	}
}

func TestHandleCommand_FLUSHALL(t *testing.T) {
	app := &config.App
	app.IsAsyncFlush = false

	handlers.HandleCommand("*3\r\n$3\r\nSET\r\n$6\r\nnumber\r\n$2\r\n10\r\n")
	response, err := handlers.HandleCommand("*1\r\n$8\r\nFLUSHALL\r\n")
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	expectedResponse := "+OK\r\n"
	if response != expectedResponse {
		t.Errorf("Expected response %v, got %v", expectedResponse, response)
	}

	resp, err := handlers.HandleCommand("*2\r\n$6\r\nEXISTS\r\n$6\r\nnumber\r\n")
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if resp != ":0\r\n" {
		t.Errorf("Expected ':0\\r\\n', got %s", resp)
	}

	go func() {
		for {
			select {
			case <-app.FlushCh:
				app.Cache.FLUSHALL()
			}
		}
	}()
	app.IsAsyncFlush = true

	handlers.HandleCommand("*3\r\n$3\r\nSET\r\n$6\r\nnumber\r\n$2\r\n10\r\n")
	handlers.HandleCommand("*1\r\n$8\r\nFLUSHALL\r\n")

	time.Sleep(15 * time.Second)
	resp, err = handlers.HandleCommand("*2\r\n$6\r\nEXISTS\r\n$6\r\nnumber\r\n")
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if resp != ":0\r\n" {
		t.Errorf("Expected ':0\\r\\n', got %s", resp)
	}
}
