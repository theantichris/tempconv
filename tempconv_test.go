package tempconv_test

import (
	"testing"

	"github.com/theantichris/tempconv"
)

func TestCToF(t *testing.T) {
	t.Run("it converts absolute zero (-273.15) in Celsius to 32 Fahrenheit", func(t *testing.T) {
		got := tempconv.CToF(tempconv.AbsoluteZeroC)
		want := tempconv.Fahrenheit(-459.66999999999996)

		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	})

	t.Run("it converts the freezing point (0) in Celsius to 32 Fahrenheit", func(t *testing.T) {
		got := tempconv.CToF(tempconv.FreezingC)
		want := tempconv.Fahrenheit(32)

		if got != want {
			t.Errorf("got %.2f want %.2f", got, want)
		}
	})

	t.Run("it converts the boiling point (100) in Celsius to 212 Fahrenheit", func(t *testing.T) {
		got := tempconv.CToF(tempconv.BoilingC)
		want := tempconv.Fahrenheit(212)

		if got != want {
			t.Errorf("got %.2f want %.2f", got, want)
		}
	})
}

func TestFToC(t *testing.T) {
	t.Run("it converts -459.66999999999996 Fahrenheit to absolute zero (-273.15) in Celsius", func(t *testing.T) {
		got := tempconv.FToC(-459.66999999999996)
		want := tempconv.AbsoluteZeroC

		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	})

	t.Run("it converts 32 Fahrenheit to the (0) in Celsius", func(t *testing.T) {
		temp := tempconv.Fahrenheit(32)

		got := tempconv.FToC(temp)
		want := tempconv.FreezingC

		if got != want {
			t.Errorf("got %.2f want %.2f", got, want)
		}
	})

	t.Run("it converts the 212 Fahrenheit to the (100) in Celsius", func(t *testing.T) {
		got := tempconv.FToC(212)
		want := tempconv.BoilingC

		if got != want {
			t.Errorf("got %.2f want %.2f", got, want)
		}
	})
}
