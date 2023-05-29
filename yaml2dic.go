
package main

import (
	"io/ioutil"
	"fmt"
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

func convertToDottedNotation(data interface{}, prefix string, result map[string]interface{}) {
	switch m := data.(type) {
	case map[interface{}]interface{}:
		for k, v := range m {
			key := fmt.Sprintf("%s.%v", prefix, k)
			convertToDottedNotation(v, key, result)
		}
	default:
		result[prefix] = data
	}
}

func main() {
	// Read YAML data from standard input (stdin)
	// yamlData, err := reader := bufio.NewReader(os.Stdin)
	//reader.ReadString('\n')
	yamlData, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		log.Fatalf("Failed to read YAML data: %v", err)
	}
	// fmt.Println(yamlData)

	// Parse the YAML data into a map
	var data map[interface{}]interface{}
	err = yaml.Unmarshal([]byte(yamlData), &data)
	if err != nil {
		log.Fatalf("Failed to unmarshal YAML: %v", err)
	}

	// Convert the values to dotted notation
	result := make(map[string]interface{})
	convertToDottedNotation(data, "", result)

	// Print the result
	for k, v := range result {
		fmt.Printf("%s=%v\n", k, v)
	}
	//fmt.Printf("OUT")
}

