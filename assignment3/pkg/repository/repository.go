package repository

import (
	"assignment3/pkg/entity"
	"encoding/json"
	"io/ioutil"
)

type Repository interface {
	GetData() (entity.Response, error)
	UpdateData(entity.Response) error
}

type repository struct {
	filepath string
}

func NewDataRepository(filepath string) Repository {
	return &repository{
		filepath: filepath,
	}
}

func (r *repository) GetData() (entity.Response, error) {
	res, err := ioutil.ReadFile(r.filepath)
	if err != nil {
		return entity.Response{}, err
	}

	var resonse entity.Response
	err = json.Unmarshal(res, &resonse)
	if err != nil {
		return resonse, err
	}
	return resonse, nil
}

func (r *repository) UpdateData(response entity.Response) error {
	byte, err := json.Marshal(response)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(r.filepath, byte, 0777)
	if err != nil {
		return err
	}
	return nil
}
