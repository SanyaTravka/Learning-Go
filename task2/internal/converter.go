package converter

import (
 "encoding/json"
 "gopkg.in/yaml.v3"
 "io/ioutil"
 "os"
 "path/filepath"

 "models/employee"
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

 var employee1 employee.Employee
 err = json.Unmarshal(data, &employee)
 if err != nil {
  return nil, err
 }

 err = employee1.Validate()
 if err != nil {
  return nil, err
 }

 yamlData, err := yaml.Marshal(&employee1)
 if err != nil {
  return nil, err
 }

 return yamlData, nil
}
