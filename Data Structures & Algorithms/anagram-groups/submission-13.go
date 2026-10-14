func groupAnagrams(strs []string) [][]string { 
	//create hashmap
	hashmap := make(map[string][]string)

	//iterate through strs,
	for _, str := range strs {
	//->sort each string
	stringChars := sortStrings(str)
	hashmap[stringChars] = append(hashmap[stringChars], str)
	}

	//now how do we get a 2d array out of this? 
	//create one..
	var array2d [][]string
	for _, group := range hashmap {
		array2d = append(array2d, group)
	}

	return array2d
} 

func sortStrings(str string) string {
	chars := []rune(str)
	sort.Slice(chars, func(i, j int) bool {
		return chars[i] < chars[j]
	})
	//->append string to hashmap at hashmap[sortedString]
	return string(chars)
}