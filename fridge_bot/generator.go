package main

import (
	"math/rand"
	"regexp"
	"sort"
	"time"
)

func GenerateSequence() (seqOrdered string, regSeqOrdered string) {

	seq := ""

	// Generate four random vowels
	vowelsString := "aeiouy"
	rand.Seed(int64(time.Now().Day()))
	vowelsNum := 4
	for i := 0; i < vowelsNum; i++ {
		letter := vowelsString[(rand.Intn(vowelsNum))]
		seq += string(letter)
	}

	// Generate other 16 random characters
	min := int('a')
	max := int('z')
	for i := 0; i < 20; i++ {
		letter := rune(rand.Intn(max-min+1) + min)
		seq += string(letter)
	}

	alphabet := []rune(seq)

	sort.Slice(alphabet, func(p, q int) bool {
		return (alphabet[p]) < (alphabet[q])
	})

	seqOrdered = string(alphabet)
	regSeqOrdered = "^( *)"

	for _, element := range alphabet {
		regSeqOrdered += (string(element) + "?")
	}
	regSeqOrdered += "( *)$"

	return
}

func InSequence(seqReg string, input string) bool {
	inputSplit := []rune(input)
	sort.Slice(inputSplit, func(i, j int) bool {
		return inputSplit[i] < inputSplit[j]
	})
	inputSorted := string(inputSplit)

	matched, _ := regexp.Match(seqReg, []byte(inputSorted))
	// fmt.Println("%s %s d %s d %t", seqReg, input, inputSorted, matched)
	return matched
}
