package main

import (
	"fmt"
	"github.com/snarlysodboxer/k8s-spec/spec"
)

var Store EtcdStore

type EtcdStore struct {
}

func (store *EtcdStore) Create(specGroup *spec.SpecGroup) (int, error) {
	fmt.Println("create placeholder")
	return 1, nil
}

func (store *EtcdStore) Update(id int, specGroup *spec.SpecGroup) (int, error) {
	fmt.Println("update placeholder")
	return 1, nil
}

func (store *EtcdStore) Read(id int) (*spec.SpecGroup, error) {
	fmt.Println("read placeholder")
	return &spec.SpecGroup{}, nil
}

func (store *EtcdStore) List() (*map[int]string, error) {
	fmt.Println("list placeholder")
	return &map[int]string{1: "site1", 2: "site2", 3: "site3"}, nil
}

func (store *EtcdStore) Delete(id int) error {
	fmt.Println("delete placeholder")
	return nil
}
