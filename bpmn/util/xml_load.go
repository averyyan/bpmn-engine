package bpmn_util

import (
	"encoding/xml"
	"os"

	"github.com/averyyan/bpmn-engine/bpmn/sepc/definitions"
)

func LoadFromFile(filename string) (*definitions.TDefinitions, error) {
	xmlData, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return LoadFromByte(xmlData)
}

func LoadFromByte(xmlData []byte) (*definitions.TDefinitions, error) {
	var ds definitions.TDefinitions
	if err := xml.Unmarshal(xmlData, &ds); err != nil {
		return nil, err
	}
	return &ds, nil
}
