package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
	"strconv"
)

var rangeS ,rangeE int
var isEven bool

func main() {
	flag.IntVar(&rangeS, "rangeS", 1, "range starting")
	flag.IntVar(&rangeE, "rangeE", 999, "range end")
	flag.BoolVar(&isEven,"isEven",true,"true: num is even")
	flag.Parse()

	if rangeS > rangeE{
		os.Exit(1)
	}

	fmt.Println(Num(rangeS,rangeE,isEven))
}

func Num(rangeS int, rangeE int, isEven bool)(int) {

	var result int
	hostname := os.Getenv("HOSTNAME")

	if hostname == "" {
		os.Exit(1)
	}

	hostNum,err := refineNum(hostname)
	if err != nil {
		fmt.Errorf("refine host num error!")
		os.Exit(1)
	}

	//偶数
	if isEven {
		result = (hostNum+1)*2 + rangeS

	} else {
		result = (hostNum+1)*2 - 1 + rangeS
	}

	if (rangeS%2 > 0 && isEven) {
		result = result - 1
	} else if (rangeS%2 > 0 && !isEven) {
		result = result - 1
	} else if (rangeS%2 == 0 && isEven) {
		result = result -2
	}

	if result <= rangeE {
		return result
	} else {
		fmt.Errorf("num out of range !")
		os.Exit(1)
	}
	return -1
}

func refineNum (hostname string) (int, error) {

	splitHostName := strings.Split(hostname,"-")

	indexSHN :=len(splitHostName)

	numS := splitHostName[indexSHN-1]

	numI, err :=strconv.Atoi(numS)

	if err != nil {
		fmt.Errorf("conv error: hostnum %s",numS)
	}

	return numI,nil

}
