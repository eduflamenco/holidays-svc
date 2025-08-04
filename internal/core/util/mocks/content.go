package mocks

import (
	"encoding/json"
	"fmt"
	"github.com/eduflamenco/holidays-svc/internal/core/domain"
)

func CreatePaymentStartResponseMock() *domain.HoliDayResponse {
	message := `{"status":"success","data":[{"date":"2025-01-01","title":"Año Nuevo","type":"Civil","inalienable":true,"extra":"Civil e Irrenunciable"},{"date":"2025-04-18","title":"Viernes Santo","type":"Religioso","inalienable":false,"extra":"Religioso"},{"date":"2025-04-19","title":"Sábado Santo","type":"Religioso","inalienable":false,"extra":"Religioso"},{"date":"2025-05-01","title":"Día Nacional del Trabajo","type":"Civil","inalienable":true,"extra":"Civil e Irrenunciable"},{"date":"2025-05-21","title":"Día de las Glorias Navales","type":"Civil","inalienable":false,"extra":"Civil"},{"date":"2025-06-20","title":"Día Nacional de los Pueblos Indígenas","type":"Civil","inalienable":false,"extra":"Civil"},{"date":"2025-06-29","title":"San Pedro y San Pablo","type":"Religioso","inalienable":false,"extra":"Religioso"},{"date":"2025-07-16","title":"Día de la Virgen del Carmen","type":"Religioso","inalienable":false,"extra":"Religioso"},{"date":"2025-08-15","title":"Asunción de la Virgen","type":"Religioso","inalienable":false,"extra":"Religioso"},{"date":"2025-09-18","title":"Independencia Nacional","type":"Civil","inalienable":true,"extra":"Civil e Irrenunciable"},{"date":"2025-09-19","title":"Día de las Glorias del Ejército","type":"Civil","inalienable":true,"extra":"Civil e Irrenunciable"},{"date":"2025-10-12","title":"Encuentro de Dos Mundos","type":"Civil","inalienable":false,"extra":"Civil"},{"date":"2025-10-31","title":"Día de las Iglesias Evangélicas y Protestantes","type":"Religioso","inalienable":false,"extra":"Religioso"},{"date":"2025-11-01","title":"Día de Todos los Santos","type":"Religioso","inalienable":false,"extra":"Religioso"},{"date":"2025-11-16","title":"Elecciones Presidenciales y Parlamentarias","type":"Civil","inalienable":true,"extra":"Civil e Irrenunciable"},{"date":"2025-12-08","title":"Inmaculada Concepción","type":"Religioso","inalienable":false,"extra":"Religioso"},{"date":"2025-12-25","title":"Navidad","type":"Religioso","inalienable":true,"extra":"Religioso e Irrenunciable"}]}`
	var response domain.HoliDayResponse
	err := json.Unmarshal([]byte(message), &response)
	if err != nil {
		fmt.Println("Something went wrong" + err.Error())
	}
	return &response
}
