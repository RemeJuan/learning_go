package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	isHeistOn := true
	eludedGuards := rand.Intn(100)

	if eludedGuards >= 50 {
		fmt.Println("Looks like you've managed to make it past the guards. Good job, but remember, this is the first step.")
	} else {
		fmt.Println("Plan a better disguise next time?")
	}

	openedVault := rand.Intn(100)

	if openedVault >= 70 {
		fmt.Println("Grab and GO!")
	} else {
		isHeistOn = false
		fmt.Println("The vault cannot be opened")
	}

	leftSafely := rand.Intn(5)

	if isHeistOn {
		switch leftSafely {
		case 0:
			isHeistOn = false
			fmt.Println("Heist failed!")
		case 1:
			isHeistOn = false
			fmt.Println("Turns out vault doors don't open from the inside...")
		case 2:
			isHeistOn = false
			fmt.Println("Tripped an alarm sensor")
		case 3:
			isHeistOn = false
			fmt.Println("Forgot to fill-up the car")
		default:
			fmt.Println("Let's get out of here")
		}
	}

	if isHeistOn {
		amtStolen := 10000 + rand.Intn(1000000)
		fmt.Println("We got away with", amtStolen)
	}

	fmt.Println("Is Heist on?", isHeistOn)
}
