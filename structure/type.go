package main

type Celsius float64
type Fabrenheit float64

const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC     Celsius = 0
	BoilingC      Celsius = 100
)

func CToF(c Celsius) Fabrenheit {
	return Fabrenheit(c*9/5 + 32)
}

func FToC(f Fabrenheit) Celsius {
	return Celsius((f - 32) * 5 / 9)
}
