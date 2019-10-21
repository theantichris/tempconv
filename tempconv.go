package tempconv

// Celsius holds a Celsius temperature.
type Celsius float64

// Fahrenheit holds a Fahrenheit temperature.
type Fahrenheit float64

const (
	// AbsoluteZeroC is absolute zero in Celsius.
	AbsoluteZeroC Celsius = -273.15
	// FreezingC is the freezing point of water in Celsius.
	FreezingC Celsius = 0
	// BoilingC is the boiling point of water in Celsius.
	BoilingC Celsius = 100
)

// CToF converts a Celsius temperate to Fahrenheit.
func CToF(c Celsius) Fahrenheit {
	return Fahrenheit(c*9/5 + 32)
}

// FToC converts a Fahrenheit temperature to Celsius.
func FToC(f Fahrenheit) Celsius {
	return Celsius((f - 32) * 5 / 9)
}
