package store

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/jay-dee7/sia-box/config"
	"github.com/jay-dee7/sia-box/types"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
)

type Store struct {
	state     types.StoreState
	storeFile string
}

func CreateStore() (*Store, error) {
	cfg, err := config.Read()
	if err != nil {
		return nil, err
	}

	fileName := fmt.Sprintf("%s/store.yaml", cfg.Path)
	storeFile, err := os.OpenFile(fileName, os.O_CREATE|os.O_RDWR, os.ModePerm)
	if err != nil {
		return nil, fmt.Errorf("error opening data store: %w", err)
	}
	data, err := ioutil.ReadAll(storeFile)
	if err != nil {
		return nil, err
	}

	var state types.StoreState
	if err = yaml.Unmarshal(data, &state); err != nil {
		return nil, err
	}

	return &Store{state, storeFile.Name()}, nil

}

func (s *Store) Read() (types.StoreState, error) {
	return s.state, nil
}

func (s *Store) Update(key, value string) {
	if value == s.state.Files[key] {
		color.Yellow("state has not changed for the file: %s", key)
		return
	}

	s.state.Files[key] = value
	color.Yellow("state updated for file: %s", key)
}

func (s *Store) Flush() error {
	data, err := yaml.Marshal(s)
	if err != nil {
		return err
	}

	return ioutil.WriteFile(s.storeFile, data, os.ModePerm)
}
