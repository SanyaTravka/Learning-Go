package models

import (
 "errors"
)

type Employee struct {
 Id       int
 Name     string
 Phone    string
 Birthday *string
}

func (e *Employee) Validate() error {
 if e.Id <= 0 {
  return errors.New("Id must be greater than 0")
 }
 if e.Name == "" {
  return errors.New("Name is required")
 }
 if e.Phone == "" {
  return errors.New("Phone is required")
 }
 return nil
}
