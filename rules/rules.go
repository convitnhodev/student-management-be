package rules

import (
	"encoding/json"
	"io/ioutil"
)

const path = "rules/rules.json"

type Rules struct {
	MinAge       int
	MaxAge       int
	MaxStudent   int
	PassingGrade float64
}

// Default rules
func (c *Rules) Default() error {
	if c.MinAge == 0 {
		c.MinAge = 15
	}
	if c.MaxAge == 0 {
		c.MaxAge = 20
	}
	if c.MaxStudent == 0 {
		c.MaxStudent = 50
	}
	if c.PassingGrade == 0 {
		c.PassingGrade = 5.0
	}

	return nil
}

// Parse rules from json byte
func Parse(b []byte) (*Rules, error) {
	c := &Rules{}
	if err := json.Unmarshal(b, c); err != nil {
		return nil, err
	}
	if err := c.Default(); err != nil {
		return nil, err
	}

	return c, nil
}

// Read the rules
func Read() (*Rules, error) {
	b, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	return Parse(b)
}

// Write the rules
func Write(c *Rules) error {
	b, _ := json.MarshalIndent(c, "", "  ")
	return ioutil.WriteFile(path, b, 0644)
}

// Create the sample rules file
func Create() error {
	b, _ := json.MarshalIndent(Rules{}, "", "  ")
	return ioutil.WriteFile(path, b, 0644)
}
