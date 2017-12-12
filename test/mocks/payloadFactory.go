package mocks

import (
	"github.com/uhuchain/uhuchain-core/models"
)

func createCarPayload() string {
	return `{
		"brand": "Volkswagen",
		"id": 12345,
		"model": "Sharan GTI",
		"policies": [
			{
				"claims": [
					{
						"date": "2016-11-01",
						"description": "Something bad happend",
						"id": 12345
					}
				],
				"endDate": "2017-09-01",
				"id": 12345,
				"insuranceId": 12345,
				"insuranceName": "Zurich Insurance Group",
				"startDate": "2016-09-01"
			}
		],
		"vehicleId": "THK34SDM6A2D34"
	}`
}

func createCar() models.Car {
	car := models.Car{}
	car.UnmarshalBinary([]byte(createCarPayload()))
	return car
}
