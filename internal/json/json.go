package json

import "Traffic-data-processing/model"

/**
* @Author: super
* @Date: 2020-09-17 07:10
* @Description:
**/

func MarshJson(i interface{}) (string, error) {
	var result string
	var err error
	switch i {
	case i.(model.ShortRoad):
		bytes, err := i.(model.ShortRoad).MarshalJSON()
		if err == nil{
			result = string(bytes)
		}
	case i.(model.Road):
		bytes, err := i.(model.Road).MarshalJSON()
		if err == nil{
			result = string(bytes)
		}
	}
	return result, err
}
