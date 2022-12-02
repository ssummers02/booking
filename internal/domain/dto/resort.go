package dto

import (
	"booking/internal/delivery/api/restmodel"
	"booking/internal/domain/entity"
	"booking/internal/infrastructure/dbmodel"
)

func ResortFromDB(resort dbmodel.Resort) entity.Resort {
	return entity.Resort{
		ID:          resort.ID,
		Name:        resort.Name,
		CityID:      resort.CityID,
		OwnerID:     resort.OwnerID,
		Description: resort.Description,
		Address:     resort.Address,
		CreatedAt:   resort.CreatedAt,
		UpdatedAt:   resort.UpdatedAt,
	}
}
func ResortsFromDB(resorts []dbmodel.Resort) []entity.Resort {
	result := make([]entity.Resort, len(resorts))

	for i, r := range resorts {
		result[i] = ResortFromDB(r)
	}

	return result
}

func ResortToDB(resort entity.Resort) dbmodel.Resort {
	return dbmodel.Resort{
		ID:          resort.ID,
		Name:        resort.Name,
		CityID:      resort.CityID,
		OwnerID:     resort.OwnerID,
		Description: resort.Description,
		Address:     resort.Address,
		CreatedAt:   resort.CreatedAt,
		UpdatedAt:   resort.UpdatedAt,
	}
}

func CityFromDB(city dbmodel.City) entity.City {
	return entity.City{
		ID:   city.ID,
		Name: city.Name,
	}
}
func CitiesFromDB(cities []dbmodel.City) []entity.City {
	result := make([]entity.City, len(cities))

	for i, c := range cities {
		result[i] = CityFromDB(c)
	}

	return result
}

func CityToRest(city entity.City) restmodel.City {
	return restmodel.City{
		ID:   city.ID,
		Name: city.Name,
	}
}
func CitiesToRest(cities []entity.City) []restmodel.City {
	result := make([]restmodel.City, len(cities))

	for i, c := range cities {
		result[i] = CityToRest(c)
	}

	return result
}

func ResortToRest(resort entity.Resort) restmodel.Resort {
	return restmodel.Resort{
		ID:          resort.ID,
		Name:        resort.Name,
		CityID:      resort.CityID,
		OwnerID:     resort.OwnerID,
		Description: resort.Description,
		Address:     resort.Address,
	}
}
func ResortsToRest(resorts []entity.Resort) []restmodel.Resort {
	result := make([]restmodel.Resort, len(resorts))

	for i, r := range resorts {
		result[i] = ResortToRest(r)
	}

	return result
}

func ResortFromRest(resort restmodel.Resort) entity.Resort {
	return entity.Resort{
		ID:          resort.ID,
		Name:        resort.Name,
		CityID:      resort.CityID,
		OwnerID:     resort.OwnerID,
		Description: resort.Description,
		Address:     resort.Address,
	}
}
