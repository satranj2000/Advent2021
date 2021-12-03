package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func readInput3() [][12]int {
	file, err := os.Open("input3.txt")

	if err != nil {
		log.Fatalf("failed to open")

	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var bitCnt [12]int
	var bitCntArr [][12]int
	for scanner.Scan() {
		//println(scanner.Text())
		val := scanner.Text()
		for i := 0; i < 12; i++ {
			if byte(val[i]) == 49 {
				bitCnt[i] = 1
			} else {
				bitCnt[i] = 0
			}
		}
		bitCntArr = append(bitCntArr, bitCnt)
	}
	return bitCntArr
}

func Puzzle5() int {

	bitCntArr := readInput3()

	var oneArrCnt [12]int
	for i := 0; i < len(bitCntArr); i++ {
		for j := 0; j < 12; j++ {
			if bitCntArr[i][j] == 1 {
				oneArrCnt[j]++
			}
		}
	}

	var zeroArrCnt [12]int
	for i := 0; i < 12; i++ {
		zeroArrCnt[i] = len(bitCntArr) - oneArrCnt[i]
	}

	var gammaratestr [12]int
	var espilonratestr [12]int
	for i := 0; i < 12; i++ {
		if oneArrCnt[i] > zeroArrCnt[i] {
			gammaratestr[i] = 1
			espilonratestr[i] = 0
		} else {
			gammaratestr[i] = 0
			espilonratestr[i] = 1
		}
	}

	gammarate, _ := strconv.ParseInt(convertIntArrtoString(gammaratestr), 2, 64)

	espilonrate, _ := strconv.ParseInt(convertIntArrtoString(espilonratestr), 2, 64)

	return int(gammarate) * int(espilonrate)

}

func convertIntArrtoString(arr [12]int) string {

	//var str string
	str := fmt.Sprintf("%d%d%d%d%d%d%d%d%d%d%d%d", arr[0], arr[1], arr[2],
		arr[3], arr[4], arr[5], arr[6], arr[7], arr[8], arr[9], arr[10], arr[11])
	return str
}

func Puzzle6() int {

	bitCntArr := readInput3()
	includeItems := make([]bool, len(bitCntArr))

	for i := 0; i < len(bitCntArr); i++ {
		includeItems[i] = true
	}
	var oxygenrate int64
	var co2scrubrate int64
	var posid int

	for i := 0; i < 12; i++ {
		ones := 0
		loop := 0
		for j := 0; j < len(bitCntArr); j++ {
			if !includeItems[j] {
				continue
			}
			if bitCntArr[j][i] == 1 {
				ones++
			}
			loop++
			if loop == 1 {
				posid = j
			}
		}
		// only one value left
		if loop == 1 {
			oxygenratestr := convertIntArrtoString(bitCntArr[posid])
			oxygenrate, _ = strconv.ParseInt(oxygenratestr, 2, 64)
			break
		}
		zeros := loop - ones
		var bIgnoreval int
		if ones < zeros {
			bIgnoreval = 1
		} else {
			bIgnoreval = 0
		}

		for j := 0; j < len(bitCntArr); j++ {
			if bitCntArr[j][i] == bIgnoreval {
				includeItems[j] = false
			}
		}
	}

	// it has passed all values but no single value was found, loop and find the only
	// row which was not set to false.
	if oxygenrate == 0 {
		for i := 0; i < len(includeItems); i++ {
			if includeItems[i] {
				oxygenratestr := convertIntArrtoString(bitCntArr[i])
				oxygenrate, _ = strconv.ParseInt(oxygenratestr, 2, 64)
			}
		}
	}

	// find the c02 scrubbing rate.
	// first consider all items.
	for i := 0; i < len(bitCntArr); i++ {
		includeItems[i] = true
	}
	posid = 0
	for i := 0; i < 12; i++ {
		ones := 0
		loop := 0
		//ignore those that were marked as false
		for j := 0; j < len(bitCntArr); j++ {
			if !includeItems[j] {
				continue
			}
			if bitCntArr[j][i] == 1 {
				ones++
			}
			loop++
			if loop == 1 {
				posid = j
			}
		}
		// only one value left
		if loop == 1 {
			co2scrubratestr := convertIntArrtoString(bitCntArr[posid])
			co2scrubrate, _ = strconv.ParseInt(co2scrubratestr, 2, 64)
			break
		}
		// get the count of number of zeros
		zeros := loop - ones
		var bIgnoreval int
		if ones < zeros {
			bIgnoreval = 0
		} else {
			bIgnoreval = 1
		}
		for j := 0; j < len(bitCntArr); j++ {
			if bitCntArr[j][i] == bIgnoreval {
				includeItems[j] = false
			}
		}
	}

	return int(oxygenrate) * int(co2scrubrate)
}
