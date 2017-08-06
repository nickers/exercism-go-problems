package bob

import "regexp"

const testVersion = 3

var emptyWords, maybeShouting, notShouting, isQuestion *regexp.Regexp

// Hey says given text to Bob.
func Hey(words string) string {

	if emptyWords == nil {
		emptyWords, _ = regexp.Compile("^\\s*$")
		maybeShouting, _ = regexp.Compile("[A-Z]")
		notShouting, _ = regexp.Compile("[a-z]")
		isQuestion, _ = regexp.Compile("\\?\\s*$")
	}

	if emptyWords.MatchString(words) {
		return "Fine. Be that way!"
	}

	if maybeShouting.MatchString(words) && !notShouting.MatchString(words) {
		return "Whoa, chill out!"
	}

	if isQuestion.MatchString(words) {
		return "Sure."
	}

	return "Whatever."
}
