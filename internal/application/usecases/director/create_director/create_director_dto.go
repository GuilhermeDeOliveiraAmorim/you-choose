package createdirector

import (
	"encoding/json"
	"io"
)

type InputCreateDirectorDto struct {
	Name    string `json:"name"`
	Picture string `json:"picture"`
}

func FromJSONInputCreateDirectorDto(body io.Reader) (*InputCreateDirectorDto, error) {
	inputCreateDirectorDto := InputCreateDirectorDto{}
	if err := json.NewDecoder(body).Decode(&inputCreateDirectorDto); err != nil {
		return nil, err
	}

	return &inputCreateDirectorDto, nil
}

type OutputCreateDirectorDto struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Picture string `json:"picture"`
}

func FromJSONOutputCreateDirectorDto(body io.Reader) (*OutputCreateDirectorDto, error) {
	outputCreateDirectorDto := OutputCreateDirectorDto{}
	if err := json.NewDecoder(body).Decode(&outputCreateDirectorDto); err != nil {
		return nil, err
	}

	return &outputCreateDirectorDto, nil
}
