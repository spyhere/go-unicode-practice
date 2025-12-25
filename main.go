package main

import (
	"bufio"
	"fmt"
	"log"
	"math/bits"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	getBytesStatistics("Привет", true)
	for {
		fmt.Print("> ")
		ok := scanner.Scan()
		if !ok {
			log.Fatal("failed to read")
		}
		if scanner.Text() == "q" {
			getBytesStatistics("Пока", true)
			return
		}
		arguments := strings.Split(
			strings.TrimSpace(scanner.Text()),
			" ",
		)
		if len(arguments) < 2 {
			fmt.Println("usage: {text} {bool (display as rune?)}")
		} else {
			argsLen := len(arguments)
			parsedBool, err := strconv.ParseBool(arguments[argsLen-1])
			if err != nil {
				fmt.Println("usage: {text} {bool (display as rune?)}")
				return
			}
			text := strings.Join(arguments[0:argsLen-1], " ")
			getBytesStatistics(text, parsedBool)
		}
	}
}

func strlen(input string) int {
	return strings.Count(input, "") - 1
}

func getRoundedBitsNum(n int) int {
	if n == 0 {
		return 1
	}
	if n&(n-1) == 0 {
		return n
	}
	return 1 << bits.Len(uint(n))
}

func getBytesStatistics(text string, showWithRune bool) {
	var listR []string
	var listB []string
	if showWithRune {
		listR, listB = []string{}, []string{}
		for _, r := range text {
			listR = append(listR, string(r))
			curRune := fmt.Sprintf("%b", r)
			runeLen := getRoundedBitsNum(len(curRune))
			listB = append(listB, fmt.Sprintf("%0*s", runeLen, curRune))
		}
	} else {
		textLen := len(text)
		listR, listB = make([]string, textLen), make([]string, textLen)
		for idx := range textLen {
			listR[idx] = string(text[idx])
			listB[idx] = fmt.Sprintf("%08b", text[idx])
		}
	}
	fRow := []string{text, strconv.Itoa(strlen(text))}
	bytesJoined := strings.Join(listB, "")
	sRow := []string{bytesJoined, strconv.Itoa(strlen(bytesJoined))}
	table, err := createStringifiedTable([]string{"Данные", "Длина"}, [][]string{fRow, sRow})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("> %s (as rune: %t)\n", strings.Join(listR, ""), showWithRune)
	fmt.Println(table)
	table, err = createStringifiedTable(listR, [][]string{listB})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(table)
	fmt.Println()
}

func createStringifiedTable(title []string, content [][]string) (string, error) {
	if len(content) == 0 {
		return "", nil
	}
	if len(title) != len(content[0]) {
		return "", fmt.Errorf("Title cannot has les entries that content! Title %v\nContent %v", title, content)
	}
	measures := make([]int, len(content[0]))
	for _, entry := range content {
		entryLen := len(entry)
		lastEntryIdx := entryLen - 1
		for idx, it := range entry {
			bordersChars := 2
			if idx != lastEntryIdx {
				bordersChars = 1
			}
			curLen := max(strlen(it)+bordersChars, strlen(title[idx])+bordersChars)
			if curLen > measures[idx] {
				measures[idx] = curLen
			}
		}
	}
	titleBody := assembleFormattedRow(title, measures, "|", "-")
	var contentBody strings.Builder
	for _, entry := range content {
		entryStr := assembleFormattedRow(entry, measures, " ", " ")
		contentBody.WriteString(entryStr + "\n")
	}
	return titleBody + "\n" + contentBody.String(), nil
}

func assembleFormattedRow(row []string, measures []int, border string, padding string) string {
	var res strings.Builder
	rowLen := len(row)
	lastIndex := rowLen - 1
	for idx, it := range row {
		leftBorder := border
		rightBorder := border
		if idx != lastIndex {
			rightBorder = ""
		}
		rem := measures[idx] - strlen(it) - (len(leftBorder) + len(rightBorder))
		if rem < 0 {
			res.WriteString(it)
			continue
		}
		half := int(float64(rem) / 2)
		paddingLen := len(padding)
		fmt.Fprintf(&res, "%v%s%s%s%v",
			leftBorder,
			strings.Repeat(padding, half/paddingLen),
			it,
			strings.Repeat(padding, (half+rem-(2*half))/paddingLen),
			rightBorder,
		)
	}
	return res.String()
}
