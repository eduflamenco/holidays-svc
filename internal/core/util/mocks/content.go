package mocks

import (
	"encoding/json"
	"fmt"
	"payment-issuer-svc/internal/core/domain"
)

func CreatePaymentStartResponseMock() *domain.PaymentResponse {
	message := `{"status":"success","message":"Pago procesado con éxito.","transactionId":"0bb40934-f6b2-4a30-918e-6a34a5743427","bill_number":"FAC2","paymentTotal":150,"success":true,"formattedPaymentTotal":"$150.00"}`
	var response domain.PaymentResponse
	err := json.Unmarshal([]byte(message), &response)
	if err != nil {
		fmt.Println("Something went wrong" + err.Error())
	}
	return &response
}
