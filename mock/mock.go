package mock

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
)

type mockClient struct {
	mockDirectory string
}

type Mocker interface {
	Request(method string, data *interface{}) error
}

func NewMock(mockDirectory string) Mocker {
	return &mockClient{
		mockDirectory: mockDirectory,
	}
}

func (m *mockClient) Request(method string, data *interface{}) error {
	var err error

	path := m.mockDirectory + "/" + method + ".json"

	if _, err = os.Stat(path); os.IsNotExist(err) {
		return errors.New("mock not found path[" + path + "]")
	}

	jsonFile, err := os.Open(path)
	if err != nil {
		return err
	}

	defer func() {
		_ = jsonFile.Close()
	}()

	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		return err
	}

	err = json.Unmarshal(byteValue, &data)
	if err != nil {
		return err
	}

	return err
}
