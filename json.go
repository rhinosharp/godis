package godis

import (
	"errors"
	"fmt"
	"io/ioutil"
	jsoniter "luan_go/json-iterator/go"
	"os"
)

//LoadJSON  load dataMap from json file
func (dm DataMap) LoadJSON(jsonFile string) error {
	bytes, err := ioutil.ReadFile(jsonFile)
	if err != nil {
		errMsg := fmt.Sprintf("ReadFile: %v\n", err.Error())
		return errors.New(errMsg)
	}
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	if err := json.Unmarshal(bytes, &dm); err != nil {
		errMsg := fmt.Sprintf("Unmarshal: %v\n", err.Error())
		return errors.New(errMsg)
	}
	return nil
}

//SaveJSON  Save dataMap to json file,isIndent:true
func (dm DataMap) SaveJSON(jsonFile string, isIndent ...bool) error {
	var bytes []byte
	var err error

	var json = jsoniter.ConfigCompatibleWithStandardLibrary

	if len(isIndent) > 0 && isIndent[0] == true {
		bytes, err = json.MarshalIndent(dm, "", "  ")
	} else {
		bytes, err = json.Marshal(dm)
	}

	if err != nil {
		errMsg := fmt.Sprintf("Unmarshal: %v\n", err.Error())
		return errors.New(errMsg)
	}

	err = ioutil.WriteFile(jsonFile, bytes, os.ModeAppend)

	if err != nil {
		errMsg := fmt.Sprintf("WriteFile: %v\n", err.Error())
		return errors.New(errMsg)
	}
	return nil
}
