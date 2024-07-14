package main

import (
	"testing"
)

func TestCalculateLBTTForHouseBelow145K(t *testing.T) {
	got := CalculateLBTT(144999)
	want := 0.0
	if got != want {
		t.Error("Expected no tax, got ", got)
	}
}
func TestCalculateLBTTForHouseWithin145To250KBracket(t *testing.T) {
	got := CalculateLBTT(200000)
	want := 1100.0
	if got != want {
		t.Error("Expected £1100 tax, got ", got)
	}
}
func TestCalculateLBTTForHouseWithin250To325KBracket(t *testing.T) {
	got := CalculateLBTT(300000)
	want := 4600.0
	if got != want {
		t.Error("Expected £4600 tax, got ", got)
	}
}
func TestCalculateLBTTForHouseWithin325kTo750KBracket(t *testing.T) {
	got := CalculateLBTT(550000.0)
	want := 28350.0
	if got != want {
		t.Error("Expected £28350 tax, got ", got)
	}
}
func TestCalculateLBTTForHouseInTheOver750KBracket(t *testing.T) {
	got := CalculateLBTT(1000000.0)
	want := 78350.0
	if got != want {
		t.Error("Expected £78350 tax, got ", got)
	}
}
