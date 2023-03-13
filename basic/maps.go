package main

import (
	"fmt"
)

var pl = fmt.Println

func main() {

	pl("Hello World")

	// var mapName map[Key Type]Value Type
	var heroes map[string]string
	heroes = make(map[string]string)

	heroes["Batman"] = "Burce Wayne"
	heroes["Super Man"] = "Clark Kent"
	heroes["Flash"] = "Bary Allen"
	heroes["Captain America"] = "Shaktiman"

	//other way to make a map
	villans := make(map[string]string)

	villans["Lex Luther"] = "Lex Luther"
	villans["penguin"] = "penguuuuu"

	// other way to make a map

	superPets := map[int]string{1: "cat", 2: "dog", 3: "penguin???"}

	pl(superPets[1], "- shows its existence")

	// when a key does not exist, it returns a nil character

	for key, value := range heroes {
		pl("%s is %s", key, value)
	}

	//  for deleting

	// delete(mapName, keyName)
	delete(heroes, "Flash")

}
