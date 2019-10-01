package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var ang, rise_time, kp, targ float64
	print("Overshoot: ")
	angStr, _ := reader.ReadString('\n')
	ang, _ = strconv.ParseFloat(strings.TrimSpace(angStr), 64)
	print("Rise Time: ")
	rise_timeStr, _ := reader.ReadString('\n')
	rise_time, _ = strconv.ParseFloat(strings.TrimSpace(rise_timeStr), 64)
	print("Kp: ")
	kpStr, _ := reader.ReadString('\n')
	kp, _ = strconv.ParseFloat(strings.TrimSpace(kpStr), 64)
	print("Target: ")
	targStr, _ := reader.ReadString('\n')
	targ, _ = strconv.ParseFloat(strings.TrimSpace(targStr), 64)

	overshoot := abs((ang - targ) / targ)
	zeta := math.Sqrt((math.Log(overshoot) * math.Log(overshoot)) / (3.14*3.14 + math.Log(overshoot)*math.Log(overshoot)))
	nat_freq := 3.14 / (rise_time * math.Sqrt(1-zeta*zeta))
	tau := 1 / (2 * nat_freq * zeta)
	k1 := nat_freq * nat_freq * tau / kp
	println()
	fmt.Println("Overshoot Percent:", overshoot)
	fmt.Println("Zeta:", zeta)
	fmt.Println("Natural Frequency:", nat_freq)
	fmt.Println("Tau:", tau)
	fmt.Println("K1:", k1)
	reader.ReadString('\n')
}

func abs(val float64) float64 {
	if val < 0.0 {
		return -val
	}

	return val
}
