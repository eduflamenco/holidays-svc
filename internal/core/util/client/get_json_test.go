package client

import (
	"context"
	"github.com/eduflamenco/holidays-svc/internal/core/domain"
	"github.com/eduflamenco/holidays-svc/internal/core/util/mocks"
	"net/http"
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
		{
			name: "new http client",
			args: args{
				timeout: 1 * time.Second,
			},
			want: &http.Client{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewHttpClient(tt.args.timeout)
			if got == nil {
				t.Errorf("NewHttpClient() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewRestClient(t *testing.T) {
	cli := &http.Client{}
	type args struct {
		http *http.Client
	}
	tests := []struct {
		name string
		args args
		want *RestClient
	}{
		{
			name: "new rest client",
			args: args{
				http: cli,
			},
			want: &RestClient{http: cli},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewRestClient(tt.args.http); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewRestClient() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRestClient_GetRequest(t *testing.T) {
	var response domain.HoliDayResponse
	client := NewHttpClient(1 * time.Minute)
	svr := mocks.NewPaymentMakerServer()
	type fields struct {
		http *http.Client
	}
	type args struct {
		ctx     context.Context
		url     string
		headers map[string]string
		result  any
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name:   "Get Request",
			fields: fields{http: client},
			args: args{
				ctx:     context.Background(),
				url:     svr.URL + "/api/holidays",
				headers: map[string]string{},
				result:  &response,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &RestClient{
				http: tt.fields.http,
			}
			if err := c.GetRequest(tt.args.ctx, tt.args.url, tt.args.headers, tt.args.result); (err != nil) != tt.wantErr {
				t.Errorf("GetRequest() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
	svr.Close()
}
