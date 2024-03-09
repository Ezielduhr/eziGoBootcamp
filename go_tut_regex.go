package main

import (
	"regexp"
	"fmt"
)

var pl = fmt.Println

func main() {
	reStr := "The ape was at the apex"
	match, _ := regexp.MatchString("(ape[^ ]?)", reStr)
	pl(match)

	reStr2 := "Cat Rat Mat Fat Pat"
	regex, _ := regexp.Compile("([^crmfp]at)")
	pl("MatchString:", regex.MatchString(reStr2))
	pl("FindString:", regex.FindString(reStr2))
	pl("FindStringIndex:", regex.FindStringIndex(reStr2))
	pl("AllFindString:", regex.FindAllString(reStr2, -1))
	pl("All Submatch index:", regex.FindAllStringSubmatchIndex(reStr2, -1))
	pl("All Submatch index:", regex.ReplaceAllString(reStr2, "Dog"))

}