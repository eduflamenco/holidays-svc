package client

import (
	"context"
	"net/http"
	"payment-issuer-svc/internal/core/domain"
	"payment-issuer-svc/internal/core/util/mocks"
	"reflect"
	"testing"
	"time"
)

func TestNewHttpClient(t *testing.T) {
	type args struct {
		timeout time.Duration
	}
	tests := []struct {
		name string
		args args
		want *http.Client
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewHttpClient(tt.args.timeout); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewHttpClient() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewRestClient(t *testing.T) {
	type args struct {
		http *http.Client
	}
	tests := []struct {
		name string
		args args
		want *RestClient
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewRestClient(tt.args.http); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewRestClient() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRestClient_PostJSONPayment(t *testing.T) {
	var response domain.PaymentResponse
	client := NewHttpClient(1 * time.Minute)
	svr := mocks.NewPaymentMakerServer()
	type fields struct {
		http *http.Client
	}
	type args struct {
		ctx     context.Context
		url     string
		payload any
		headers map[string]string
		out     any
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "success",
			fields: fields{
				http: client,
			},
			args: args{
				ctx:     context.Background(),
				url:     svr.URL + "/payment/start",
				payload: domain.PaymentRequest{},
				headers: map[string]string{
					"Content-Type": "application/json",
				},
				out: &response,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &RestClient{
				http: tt.fields.http,
			}
			if err := c.PostJSONPayment(tt.args.ctx, tt.args.url, tt.args.payload, tt.args.headers, tt.args.out); (err != nil) != tt.wantErr {
				t.Errorf("PostJSONPayment() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
	svr.Close()
}
