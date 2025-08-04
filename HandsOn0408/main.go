package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	//
	//mp1 := map[string]int{
	//	"Alice":   1,
	//	"Bob":     2,
	//	"Charlie": 3,
	//}
	//mp2 := map[string]int{
	//	"Alice": 5,
	//	"Bob":   9,
	//	"Dave":  10,
	//}
	//
	//mp3 := mergeMap(mp1, mp2)
	//
	//fmt.Println(mp3)
	//
	//input := "A paragraph is a distinct unit of writing, typically composed of multiple sentences, that focuses on a single idea or topic. It serves as a building block for longer pieces of text, helping to organize and structure information in a clear and coherent manner. Each paragraph usually begins with an indentation or a new line, signaling a shift in focus or a new stage in the development of the overall argument or narrative"
	//output := Capitalize(input)
	//fmt.Println(output)

	testMap := map[string]string{
		"Hello": "World",
		"World": "New",
		"New":   "Cold",
		//"Prime": "New",
	}

	myMapInterface := make(map[interface{}]interface{})
	for k, v := range testMap {
		myMapInterface[k] = v
	}
	invertMapOther(myMapInterface)

	interfaceMap := map[interface{}]interface{}{
		"Alice":   10,
		"Bob":     20,
		"Charlei": 30,
		"Dave":    40,
		"Eve":     50,
	}

	invertMapOther(interfaceMap)
}

func invertMapOther(in map[interface{}]interface{}) {
	if len(in) == 0 {
		fmt.Println("Empty map")
	}

	res := make(map[interface{}]interface{})
	for key, val := range in {
		res[val] = key
	}

	fmt.Println(res)
}

/*Merge Two Maps
Problem:
Write a function that merges two map[string]int by adding the values for common keys. If a key only exists in one map, copy it as is.
*/

func mergeMap(m1, m2 map[string]int) map[string]int {

	for key := range m1 {
		if _, exist := m2[key]; exist {
			m1[key] += m2[key]
			delete(m2, key)
		}
	}
	for key, val := range m2 {
		m1[key] = val
	}

	return m1
}

/*Capitalize Each Word
Given a sentence, capitalize the first letter of each word.
Input: "hello world"
Output: "Hello World"
*/

func Capitalize(input string) string {

	words := strings.Split(input, " ")

	for index, word := range words {
		letters := []rune(word)
		letters[0] = unicode.ToUpper(letters[0])
		words[index] = string(letters)
	}
	return strings.Join(words, " ")
}

//
//
//Invert a Map
//Problem:
//Write a function that inverts a map[string]string so that the values become keys and the keys become values.

func InvertMap[K comparable, V comparable](inputMap map[K]V) map[V]K {
	invertedMap := make(map[V]K)

	for key, val := range inputMap {
		invertedMap[val] = key
	}

	for key, val := range invertedMap {
		fmt.Println(key, val)
	}

	return invertedMap

}
