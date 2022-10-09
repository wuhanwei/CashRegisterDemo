package services

import (
	"CashRegisterDemo/dtos"
	"fmt"
	"strconv"
)

func GetRedeemPrice(data *dtos.StructCashRegister) (*dtos.StructResponseOrderData, error, string) {
	var (
		response_order dtos.StructResponseOrderData

		discount_points   float64
		discount_coins    float64
		redeem_points     float64 = 0
		final_order_price float64
		sale              float64 = 1
	)

	//get member data : level_bonues_member_points
	member, err := getMember(data)
	if err != nil {
		return &response_order, err, "4001"
	}

	//convert string to float64, before calculate redeem price.
	level_bonus, err := strconv.ParseFloat(member.Field_level_bonus, 64)
	if err != nil {
		return &response_order, err, "4005"
	}
	// order price convert string to float64, and initial final order price = order_price.
	order_price, err := strconv.ParseFloat(data.Order_price, 64)
	if err != nil {
		return &response_order, err, "4002"
	}
	final_order_price = order_price

	//if redeem points>0 and member_point>0, get discoint point and coins.
	if redeem_points, err = strconv.ParseFloat(data.Redeem_points, 64); err == nil && redeem_points > 0 && member.Field_points > 0 {
		if discount_points, discount_coins, err = getRedeemPoints(order_price, &member); err != nil {
			return &response_order, err, "4003"
		}
		response_order.Field_discount_points = fmt.Sprint(discount_points)
		response_order.Field_discount_coins = fmt.Sprint(discount_coins)
	}

	//calculate final price
	final_order_price = float64(order_price)*float64(level_bonus) - float64(discount_coins)

	//convert string to float64
	if data.Sale != "" {
		sale, err = strconv.ParseFloat(data.Sale, 64)
		if err != nil {
			fmt.Println("errr")
			return &response_order, err, "4006"
		}
		if sale < 1 {
			final_order_price = final_order_price * sale
		}
	}

	//set response body
	response_order.Field_level_bonus = member.Field_level_bonus
	response_order.Field_order_price = fmt.Sprint(final_order_price)
	return &response_order, nil, "200"
}

func getRedeemPoints(order_price float64, member *dtos.StructMember) (float64, float64, error) {
	var (
		coins_ratio     float64
		points_ratio    float64
		discount_points float64 = 0.0
		discount_coins  float64 = 0.0
	)

	//get coins and points ratio
	coins_ratio, points_ratio, err := GetConfigPointsRatio()
	if err != nil {
		return discount_points, discount_coins, err
	}

	//本次折抵點數
	discount_points = order_price / points_ratio
	if discount_points > float64(member.Field_points) {
		discount_points = float64(member.Field_points)
	}

	//本次折抵金額
	discount_coins = discount_points * coins_ratio
	if discount_coins > order_price {
		discount_coins = order_price
		discount_points = discount_coins / coins_ratio
	}

	return discount_points, discount_coins, nil
}
