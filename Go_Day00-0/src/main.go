package main

import (
	"bufio"
	"flag"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
)

func mostFrequent(arr []int) int {
	m := map[int]int{}
	var mfreq, maxf int

	if len(arr) == 1 {
		return arr[0]
	}
	for i, a := range arr {
		m[a]++
		if i == 0 {
			maxf = a
			mfreq = 1
		} else {
			if m[a] > mfreq {
				maxf = a
				mfreq = m[a]
			} else if m[a] == mfreq && a < maxf {
				maxf = a
			}
		}
	}
	return maxf
}

func main() {
	in := bufio.NewScanner(os.Stdin)
	var arr []int
	var sum, i, mean, median, sd float64
	var Mean, Median, Mode, SD bool

	/*recovery*/

	defer func() {
		if err := recover(); err != nil {
			fmt.Println("Unknown panic happend, but it's recovered already ^__^ :", err)
		}
	}()

	/*flags*/

	flag.BoolVar(&Mean, "Mean", false, "display in Mean")
	flag.BoolVar(&Median, "Median", false, "display in Median")
	flag.BoolVar(&Mode, "Mode", false, "display in Mode")
	flag.BoolVar(&SD, "SD", false, "display in SD")
	flag.Parse()

	/*scanning*/

	for in.Scan() {
		txt := in.Text()
		if txt == "" {
			fmt.Println("Empty input")
			os.Exit(3)
		}
		num, err := strconv.Atoi(txt)
		if err != nil {
			fmt.Println("Non-int input")
			os.Exit(4)
		}
		if num > 100000 || num < -100000 {
			fmt.Println("Out-of-bounds value")
			os.Exit(5)
		}
		arr = append(arr, num)
		sum += float64(num)
		i++
	}

	/*sorting*/

	sort.Ints(arr)

	/*mean calculation*/

	mean = sum / i

	/*median calculation*/

	ii := int(i)
	if ii%2 != 0 {
		median = float64(arr[ii/2])
	} else {
		median = float64(arr[ii/2]+arr[ii/2-1]) / 2.0
	}

	/*mode calculation*/

	mode := mostFrequent(arr)

	/*sd calculation*/

	for _, j := range arr {
		sd += math.Pow(float64(j)-mean, 2)
	}
	sd = math.Sqrt(sd / i)

	/*boolean output*/

	if Mean {
		fmt.Printf("Mean: %.2f\n", mean)
	}
	if Median {
		fmt.Printf("Median: %.2f\n", median)
	}
	if Mode {
		fmt.Printf("Mode: %d\n", mode)
	}
	if SD {
		fmt.Printf("SD: %.2f\n", sd)
	}
	if !Mean && !Median && !Mode && !SD {
		fmt.Printf("Mean: %.2f\nMedian: %.2f\nMode: %d\nSD: %.2f\n", mean, median, mode, sd)
	}
}
