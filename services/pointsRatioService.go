package services

import "errors"

func GetConfigPointsRatio() (float64, float64, error) {
	var ok bool
	var coins float64 = 0
	var points float64 = 0

	points_ratio_map, err := GetViper("points_coins_ratio")
	if err != nil {
		return coins, points, err
	}

	if x, found := points_ratio_map["coins_ratio"]; found {
		if coins, ok = x.(float64); !ok {
			return coins, points, errors.New("coins_ratio doesn't exist")
		}
	} else {
		return coins, points, errors.New("coins_ratio doesn't exist")
	}

	if x, found := points_ratio_map["points_ratio"]; found {
		if points, ok = x.(float64); !ok {
			return coins, points, errors.New("points_ratio doesn't exist")
		}
	} else {
		return coins, points, errors.New("points_ratio doesn't exist")
	}

	return coins, points, nil
}
