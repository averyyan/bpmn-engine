package bpmn_util

import (
	"embed"
	"encoding/xml"
	"io/fs"
	"os"
	"strings"

	"github.com/averyyan/bpmn-engine/bpmn/sepc/definitions"
)

func LoadFromEmbed(myfs *embed.FS) ([][]byte, error) {
	var processes [][]byte
	if err := fs.WalkDir(myfs, ".", func(path string, d fs.DirEntry, err error) error {
		if strings.Contains(path, ".bpmn") {
			data, err := myfs.ReadFile(path)
			if err != nil {
				return err
			}
			processes = append(processes, data)
		}
		return nil
	}); err != nil {
		return nil, err
	}
	return processes, nil
}

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
