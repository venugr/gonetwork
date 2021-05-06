package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Property map[string]string

const (
	cnt = 100
)

func main() {

	if len(os.Args) != 3 {
		fmt.Fprintf(os.Stderr, "Usage: %s <trigg-log-file1> <trigg-log-file1>\n", os.Args[0])
		os.Exit(1)
	}

	file1List, isErr := ReadFile(os.Args[1])
	if isErr != nil {
		fmt.Fprintf(os.Stderr, "File read error: %s", isErr.Error())
		os.Exit(1)
	}

	file2List, isErr := ReadFile(os.Args[2])
	if isErr != nil {
		fmt.Fprintf(os.Stderr, "File read error: %s", isErr.Error())
		os.Exit(1)
	}

	trig1_status := map[string]string{}
	trig1_owner := map[string]string{}
	trig1_type := map[string]string{}

	for _, line := range file1List {

		if !(strings.Contains(line, "DISABLED") || strings.Contains(line, "ENABLED")) {
			continue
		}
		//fmt.Println(line)

		lineSplit := strings.Split(line, ",")
		trig1_status[lineSplit[0]] = lineSplit[3]
		trig1_owner[lineSplit[0]] = lineSplit[2]
		trig1_type[lineSplit[0]] = lineSplit[1]
	}

	// for k, v := range trig1_status {

	// 	fmt.Println(k, "=", v)
	// }

	trig2_status := map[string]string{}
	trig2_owner := map[string]string{}
	trig2_type := map[string]string{}

	for _, line := range file2List {

		if !(strings.Contains(line, "DISABLED") || strings.Contains(line, "ENABLED")) {
			continue
		}
		//fmt.Println(line)

		lineSplit := strings.Split(line, ",")
		trig2_status[lineSplit[0]] = lineSplit[3]
		trig2_owner[lineSplit[0]] = lineSplit[2]
		trig2_type[lineSplit[0]] = lineSplit[1]
	}

	_, noeq, kin1, kin2 := CompareMaps(trig1_status, trig2_status)
	printMap("Keys ONLY in Map1(STATUS)", kin1)
	printMap("Keys ONLY in Map2(STATUS)", kin2)
	printMap("Keys/Values NOT matches(map1,map2)(STATUS)", noeq)

	_, noeq, kin1, kin2 = CompareMaps(trig1_owner, trig2_owner)
	printMap("Keys ONLY in Owner Map1(OWNER)", kin1)
	printMap("Keys ONLY in Owner Map2(OWNER)", kin2)
	printMap("Keys/Values NOT matches(map1,map2)(OWNER)", noeq)

	_, noeq, kin1, kin2 = CompareMaps(trig1_type, trig2_type)
	printMap("Keys ONLY in Owner Map1(TYPE)", kin1)
	printMap("Keys ONLY in Owner Map2(TYPE)", kin2)
	printMap("Keys/Values NOT matches(map1,map2)(TYPE)", noeq)

}

func CompareMaps(map1, map2 map[string]string) (map[string]string, map[string]string, map[string]string, map[string]string) {

	okKeys := map[string]string{}
	notEqKeys := map[string]string{}
	keysIn1 := map[string]string{}
	keysIn2 := map[string]string{}

	for k, v := range map1 {

		if val, ok := map2[k]; ok {
			if map1[k] == val {
				okKeys[k] = v
			} else {
				notEqKeys[k] = map1[k] + "," + val
			}
		} else {
			keysIn1[k] = v
		}

	}

	for k, v := range map2 {

		if _, ok := map1[k]; !ok {
			keysIn2[k] = v
		}

	}

	return okKeys, notEqKeys, keysIn1, keysIn2
}

func ReadFile(fileName string) ([]string, error) {

	data, err := ioutil.ReadFile(fileName)

	if err != nil {
		return nil, err
	}

	//return strings.Split(string(data), "\n"), nil

	// Remove EOF char from []string
	lines := strings.Split(string(data), "\n")
	lLen := len(lines)
	return lines[:lLen-1], nil
}

func printMap(title string, myMap map[string]string) {
	PrintProperties(title, Property(myMap))
}

// PrintProperties to print
func PrintProperties(title string, prop map[string]string) {

	mLogStr := ""
	keys := make([]string, 0, len(prop))

	for ky := range prop {
		keys = append(keys, ky)
	}

	sort.Strings(keys)
	/*
		fmt.Println()
		fmt.Println("\n" + strings.Repeat("=", cnt))
		fmt.Printf("                    %s\n", title)
		fmt.Println(strings.Repeat("=", cnt))
	*/

	mLogStr += "\n\n\n" + strings.Repeat("=", cnt)
	mLogStr += "\n                         " + title + "\n"
	mLogStr += strings.Repeat("=", cnt)

	for i, ky := range keys {
		//fmt.Printf("%v. %v=%v\n", i+1, ky, prop[ky])
		mLogStr += "\n" + strconv.Itoa(i+1) + ". " + ky + " = " + prop[ky]
		//logit(mStr)
	}

	mLogStr += "\n" + strings.Repeat("=", cnt) + "\n"
	//fmt.Println(strings.Repeat("=", cnt))
	//fmt.Println()

	logit(mLogStr)
}

func logit(logStr string) {

	//logger.Print(logStr)
	fmt.Print(logStr)

}
