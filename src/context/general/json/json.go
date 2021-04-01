package json

import "encoding/json"

// ToJSON transforma um objeto em json
func ToJSON(v interface{}) []byte {
	bt, err := json.Marshal(v)
	if err != nil {
		return nil
	}
	return bt
}

// FromJSON preenche um objeto baseado em um json
func FromJSON(data []byte, v interface{}) error {
	if err := json.Unmarshal(data, v); err != nil {
		return err
	}
	return nil
}
