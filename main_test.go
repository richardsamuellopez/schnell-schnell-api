package main

import (
	"testing"
	"fmt"
	"time"
)

func TestValidToken(t *testing.T) {
	now := time.Now()
	hour := fmt.Sprintf("%02d", now.Hour())
	minute := fmt.Sprintf("%02d", now.Minute())
	
	msg := validateToken(hour + minute)
	if !msg {
		t.Fatalf("Test Valid Token failed")
	}
}

func TestInvalidToken(t *testing.T) {
	now := time.Now()
	hour := fmt.Sprintf("%02d", now.Hour())
	minute := fmt.Sprintf("%02d", now.Add(-time.Minute * 1).Minute())
	
	msg := validateToken(hour + minute)
	if msg {
		t.Fatalf("Test Valid Token failed")
	}
}

func TestRequireAllAuthParams(t * testing.T) {
	
}
