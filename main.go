package main

import "fmt"

type RockClimber struct {
	kind        int
	rockClimbed int
	sp          SafetyPlacer
}

type SafetyPlacer struct {
	kind int
}

func (sp SafetyPlacer) placeSafeties() {
	switch sp.kind {
	case 1:

	case 2:

	case 3:
	}
	fmt.Println("Placing Safeties")
}

func (rc *RockClimber) climbRock() {
	rc.rockClimbed++
	if rc.rockClimbed%10 == 0 {
		rc.sp.placeSafeties()
	}
}

func main() {
	rc := &RockClimber{}
	for i := 0; i < 101; i++ {
		rc.climbRock()
	}
}
