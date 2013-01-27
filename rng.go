package main

import (
	"fmt"
	"math/rand"
	"time"
)

var vowels = []string{"a", "e", "i", "o", "u"}
var extendedVowels = []string{"a", "e", "i", "o", "u", "a", "e", "i", "o", "y", "ea", "ou", "ue", "ia", "ie", "eo", "eu", "ui", "io", "oe"}
var consonants = []string{"b", "c", "d", "f", "g", "h", "j", "k", "l", "m", "n", "p", "q", "r", "s", "t", "v", "w", "x", "y", "z"}
var specialStartConsonants = []string{"bl", "br", "ch", "cl", "cr", "dr", "fl", "fr", "gl", "gr", "kh", "kl", "kr", "ph", "pl", "pr", "qu", "sc", "scr", "sh", "sk", "sl", "sm", "sn", "sp", "sq", "st", "str", "sw", "tr", "tw", "wh", "wr"}
var specialConsonants = []string{"ng", "gh", "ln", "rd", "gm", "ck", "wl", "ft", "rt", "ght", "ll", "rp", "mn", "ct", "rk", "wf", "tt", "rf", "ss", "dd", "bb", "ff", "gg", "mm", "nn", "pp", "rr", "zz", "ld", "lb", "ls", "vl", "vr"}

func pickString(strings ...string) (string) {
	return string(strings[rand.Int31n(int32(len(strings)))])
}

func generateNormalSyllable() (string) {
	return pickString(consonants...) + pickString(extendedVowels...)
}

func generateShortSyllable() (string) {
	return pickString(consonants...) + pickString(vowels...)
}

func generateSpecialStartSyllable() (string) {
	return pickString(specialStartConsonants...) + pickString(extendedVowels...)
}

func generateSpecialSyllable() (string) {
	return pickString(specialConsonants...) + pickString(extendedVowels...)
}

func generateName() (string) {
	word := ""
	
	num := rand.Int31n(int32(8))
	
	if num == 0 {
		word = generateSpecialStartSyllable() + generateNormalSyllable() + generateNormalSyllable()
	} else if num == 1 {
		word = generateNormalSyllable() + generateSpecialSyllable() + generateNormalSyllable()
	} else if num == 2 {
		word = generateNormalSyllable() + generateNormalSyllable() + generateSpecialSyllable()
	} else if num == 3 {
		word = generateSpecialStartSyllable() + generateNormalSyllable() + generateSpecialSyllable()
	} else if num == 4 {
		word = generateNormalSyllable() + generateNormalSyllable() + generateNormalSyllable()
	} else if num == 5 || num == 6 || num == 7 {
		word = generateShortSyllable() + generateShortSyllable() + generateShortSyllable()
	}
	
	return word
}

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	
	for i := 0; i < 50; i++ {
		fmt.Println(generateName())
	}
}

