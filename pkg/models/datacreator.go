package models

import (
	"encoding/json"
)

func CreateData(subject string) (key []byte, value []byte, err error) {
	switch subject {
	case "address":
		addressInfo := createAddressData()
		value, err = json.Marshal(&addressInfo)
		return []byte(addressInfo.Zip), value, err
	default:
		return
	}
}
