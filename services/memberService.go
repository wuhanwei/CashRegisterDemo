package services

import (
	"CashRegisterDemo/dtos"
	"encoding/json"
	"errors"
	"fmt"
)

func getMember(data *dtos.StructCashRegister) (dtos.StructMember, error) {
	var (
		member dtos.StructMember
		ok     bool
		field  interface{}
	)

	members_map, err := GetViper("members_data")
	if err != nil {
		return member, err
	}

	if field, ok = members_map[data.Name]; !ok {
		return member, errors.New("member not found")
	}
	//get member level and points(會員目前可用剩餘點數)
	jsonString, err := json.Marshal(field)
	if err != nil {
		return member, errors.New("json error")
	}
	json.Unmarshal(jsonString, &member)

	//get level_bonus map
	level_bonus_map, err := GetViper("level_bonus_map")
	if err != nil {
		return member, err
	}

	//check level bonus
	if level, found := level_bonus_map[member.Field_level]; found {
		if member.Field_level_bonus, ok = level.(string); !ok {
			return member, errors.New("member_level doesn't exist")
		}
		fmt.Println("get member level bonus : ", member.Field_level_bonus)
	} else {
		return member, errors.New("member_level doesn't exist")
	}

	return member, nil
}
