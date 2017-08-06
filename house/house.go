package house

const testVersion = 1

var parts = []string{
	"the horse and the hound and the horn\nthat belonged to ",
	"the farmer sowing his corn\nthat kept ",
	"the rooster that crowed in the morn\nthat woke ",
	"the priest all shaven and shorn\nthat married ",
	"the man all tattered and torn\nthat kissed ",
	"the maiden all forlorn\nthat milked ",
	"the cow with the crumpled horn\nthat tossed ",
	"the dog\nthat worried ",
	"the cat\nthat killed ",
	"the rat\nthat ate ",
	"the malt\nthat lay in ",
	"the house that Jack built.",
}

func Song() string {
	_, s := getPart(len(parts))
	return s
}

func Verse(nr int) string {
	p, _ := getPart(nr)
	return "This is " + p
}

func getPart(nr int) (string, string) {
	p, s := "", ""
	if nr > 1 {
		p, s = getPart(nr - 1)
		s += "\n\n"
	}

	return parts[len(parts)-nr] + p, s + "This is " + parts[len(parts)-nr] + p
}
