package main

import (
 "flag"
 "fmt"
 "io/ioutil"
 "internal/converter"
)

func main() {
 filenamePtr := flag.String("filename", "", "Filename to convert from JSON to YAML")
 flag.Parse()

 if *filenamePtr == "" {
  fmt.Println("Please provide a filename using the -filename flag")
  return
 }

 yamlData, err := converter.ConvertJsonToYaml(*filenamePtr)
 if err != nil {
  fmt.Println("Error converting JSON to YAML:", err)
  return
 }

 outputFilename := "output.yaml"
 err = ioutil.WriteFile(outputFilename, yamlData, 0644)
 if err != nil {
  fmt.Println("Error writing YAML data to file:", err)
  return
 }

 fmt.Println("Conversion successful. Output file:", outputFilename)
}
