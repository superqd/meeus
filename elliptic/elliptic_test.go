package elliptic_test

import (
	"fmt"
	"math"

	"github.com/soniakeys/meeus/base"
	"github.com/soniakeys/meeus/elliptic"
	"github.com/soniakeys/meeus/julian"
	pp "github.com/soniakeys/meeus/planetposition"
)

func ExamplePosition() {
	// Example 33.a, p. 225.  VSOP87 result p. 227.
	earth, err := pp.LoadPlanet(pp.Earth, "")
	if err != nil {
		fmt.Println(err)
		return
	}
	venus, err := pp.LoadPlanet(pp.Venus, "")
	if err != nil {
		fmt.Println(err)
		return
	}
	α, δ := elliptic.Position(venus, earth, 2448976.5)
	fmt.Printf("α = %.3d\n", base.NewFmtRA(α))
	fmt.Printf("δ = %.2d\n", base.NewFmtAngle(δ))
	// Output:
	// α = 21ʰ4ᵐ41ˢ.454
	// δ = -18°53′16″.84
}

func ExampleElements_Position() {
	// Example 33.b, p. 232.
	earth, err := pp.LoadPlanet(pp.Earth, "")
	if err != nil {
		fmt.Println(err)
		return
	}
	k := &elliptic.Elements{
		TimeP: julian.CalendarGregorianToJD(1990, 10, 28.54502),
		Axis:  2.2091404,
		Ecc:   .8502196,
		Inc:   11.94524 * math.Pi / 180,
		Node:  334.75006 * math.Pi / 180,
		ArgP:  186.23352 * math.Pi / 180,
	}
	j := julian.CalendarGregorianToJD(1990, 10, 6)
	α, δ, ψ := k.Position(j, earth)
	fmt.Printf("α = %.1d\n", base.NewFmtRA(α))
	fmt.Printf("δ = %.0d\n", base.NewFmtAngle(δ))
	fmt.Printf("ψ = %.2f\n", ψ*180/math.Pi)
	// Output:
	// α = 10ʰ34ᵐ14ˢ.2
	// δ = 19°9′31″
	// ψ = 40.51
}

func ExampleVelocity() {
	// Example 33.c, p. 238
	fmt.Printf("%.2f\n", elliptic.Velocity(17.9400782, 1))
	// Output:
	// 41.53
}

func ExampleVPerihelion() {
	// Example 33.c, p. 238
	fmt.Printf("%.2f\n", elliptic.VPerihelion(17.9400782, .96727426))
	// Output:
	// 54.52
}

func ExampleVAphelion() {
	// Example 33.c, p. 238
	fmt.Printf("%.2f\n", elliptic.VAphelion(17.9400782, 0.96727426))
	// Output:
	// 0.91
}

func ExampleLength1() {
	// Example 33.d, p. 239
	fmt.Printf("%.2f\n", elliptic.Length1(17.9400782, 0.96727426))
	// Output:
	// 77.06
}

func ExampleLength2() {
	// Example 33.d, p. 239
	fmt.Printf("%.2f\n", elliptic.Length2(17.9400782, 0.96727426))
	// Output:
	// 77.09
}

/* func ExampleLength3() {
	// Example 33.d, p. 239
	fmt.Printf("%.2f\n", elliptic.Length3(17.9400782, 0.96727426))
	// Output:
	// 77.07
} */

func ExampleLength4() {
	// Example 33.d, p. 239
	fmt.Printf("%.2f\n", elliptic.Length4(17.9400782, 0.96727426))
	// Output:
	// 77.07
}
