package main

import (
	"math"
	"testing"
)

func TestDivide(t *testing.T) {
	t.Run("ValidDivision", func(t *testing.T) {
		got, err := Divide(20, 4)
		if err != nil {
			t.Fatalf("expected no error, got: %v", err)
		}
		if got != 5 {
			t.Errorf("expected 5, got %v", got)
		}
	})

	t.Run("DivisionByZero", func(t *testing.T) {
		_, err := Divide(20, 0)
		if err == nil {
			t.Fatal("expected error for division by zero, got nil")
		}
		expectedErr := "division by zero: cannot divide 20 by zero"
		if err.Error() != expectedErr {
			t.Errorf("expected %q, got %q", expectedErr, err.Error())
		}
	})
}

func TestSquareRoot(t *testing.T) {
	t.Run("ValidSquareRoot", func(t *testing.T) {
		got, err := SquareRoot(64)
		if err != nil {
			t.Fatalf("expected no error, got: %v", err)
		}
		if math.Abs(got-8) > 1e-9 {
			t.Errorf("expected 8, got %v", got)
		}
	})

	t.Run("NegativeSquareRoot", func(t *testing.T) {
		_, err := SquareRoot(-16)
		if err == nil {
			t.Fatal("expected error for negative square root, got nil")
		}
		expectedErr := "negative square root: value -16 must be non-negative"
		if err.Error() != expectedErr {
			t.Errorf("expected %q, got %q", expectedErr, err.Error())
		}
	})
}

func TestProcessInput(t *testing.T) {
	t.Run("ValidInput", func(t *testing.T) {
		got, err := ProcessInput("welcome")
		if err != nil {
			t.Fatalf("expected no error, got: %v", err)
		}
		if got != "WELCOME" {
			t.Errorf("expected 'WELCOME', got %v", got)
		}
	})

	t.Run("EmptyInput", func(t *testing.T) {
		_, err := ProcessInput("   ")
		if err == nil {
			t.Fatal("expected error for empty input, got nil")
		}
		expectedErr := "empty input: input string cannot be empty or whitespace only"
		if err.Error() != expectedErr {
			t.Errorf("expected %q, got %q", expectedErr, err.Error())
		}
	})
}
