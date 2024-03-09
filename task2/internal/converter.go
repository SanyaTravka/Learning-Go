package internal

import (
 "encoding/json"
 "gopkg.in/yaml.v3"
 "io/ioutil"

 "models"
)

func ConvertJsonToYaml(filename string) ([]byte, error) {
 file, err := os.Open(filename)
 if err != nil {
  return nil, err
 }
 defer file.Close()

 data, err := ioutil.ReadAll(file)
 if err != nil {
  return nil, err
 }

 var employee models.Employee
 err = json.Unmarshal(data, &employee)
 if err != nil {
  return nil, err
 }

 err = employee.Validate()
 if err != nil {
  return nil, err
 }

 yamlData, err := yaml.Marshal(&employee)
 if err != nil {
  return nil, err
 }

 return yamlData, nil
}
