package main

// Code generated by "typedjson -interface Data *Foo *Bar"; DO NOT EDIT.

import (
	"encoding/json"
	"errors"
)

type DataTyped struct {
	Data
}

func (t DataTyped) MarshalJSON() ([]byte, error) {
	typedString := t.Data.typedjson(nil)
	wrapper := struct {
		T string
		V Data
	}{
		T: typedString,
		V: t.Data,
	}
	return json.Marshal(&wrapper)
}

func (t *DataTyped) UnmarshalJSON(src []byte) error {
	var wrapper struct {
		T string
		V json.RawMessage
	}
	err := json.Unmarshal(src, &wrapper)
	if err != nil {
		return err
	}
	switch wrapper.T {
	case "*Foo":
		t.Data = &Foo{}

	case "*Bar":
		t.Data = &Bar{}

	default:
		return errors.New("unknown type")
	}
	if err := json.Unmarshal(wrapper.V, t.Data); err != nil {
		return err
	}
	return nil
}

func (s *Foo) typedjson(*DataTyped) string {
	return "*Foo"
}

func (s *Bar) typedjson(*DataTyped) string {
	return "*Bar"
}
