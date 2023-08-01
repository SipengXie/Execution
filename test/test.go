package main

import "encoding/json"

type struct1 struct {
	ES2 *struct2 `json:"s2,omitempty"`
}

type struct2 struct {
	GData string `json:"data,omitempty"`
}

func main() {
	s2 := &struct2{GData: "data"}
	s1 := &struct1{ES2: s2}

	ret, err := json.Marshal(s1)
	if err != nil {
		panic(err)
	}
	println(string(ret))

	s1_ := new(struct1)
	err = json.Unmarshal(ret, s1_)
	if err != nil {
		panic(err)
	}
}
