package helper

import "encoding/json"

// Struct2Maps
//
//	@Description: 结构体转map
//	@param structData
//	@param returnMaps
//	@return error
func Struct2Maps(structData any, returnMaps *map[string]interface{}) error {
	byteData, err := json.Marshal(structData)
	if err != nil {
		return err
	}
	err = json.Unmarshal(byteData, returnMaps)
	if err != nil {
		return err
	}
	return nil
}
