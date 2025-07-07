package iherbApi

import (
	"context"
	"github.com/gocolly/colly"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/zakharova-e/iherb-scraper-server/internal/catalog/models"
	"reflect"
	"testing"
)

func TestNewIherbApiRepository(t *testing.T) {
	type args struct {
		//no args at the moment
	}
	tests := []struct {
		name  string
		args  args
		isNil bool
	}{
		{"default behavior", args{}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewIherbApiRepository(); !assert.NotNil(t, got) && tt.isNil {
				t.Errorf("NewIherbApiRepository() returns nil, want not nil link")
			}
		})
	}
}

func TestNewIherbCollector(t *testing.T) {
	type args struct {
		options []option
	}
	option1 := option(func(collector *iherbCollector) {
		if collector == nil {
			return
		}
		collector.OnRequest(func(r *colly.Request) {
			r.Headers.Set("platform", "Linux")
			r.Headers.Set("regiontype", "GLOBAL")
		})
	})
	option2 := option(func(collector *iherbCollector) {
		if collector == nil {
			return
		}
		collector.OnRequest(func(r *colly.Request) {
			r.Headers.Set("platform", "Windows")
		})
	})
	tests := []struct {
		name string
		args args
	}{
		{"without options", args{}},
		{"with one option",
			args{
				[]option{
					option1,
				},
			},
		},
		{"with two options",
			args{
				[]option{
					option1,
					option2,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewIherbCollector(tt.args.options...)
			//colly don't provide functionaluty for checking on request headers
			require.NotNil(t, got, "NewIherbCollector() returned nil value, wanted not nil ")
		})
	}
}

func TestIherbApiRepository_GetProductData(t *testing.T) {
	type args struct {
		ctx       context.Context
		productId uint32
	}
	tests := []struct {
		name    string
		repo    *IherbApiRepository
		args    args
		want    *models.ProductData
		wantErr bool
	}{
		{"with invalid productId", NewIherbApiRepository(), args{context.Background(), 0}, nil, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.repo.GetProductData(tt.args.ctx, tt.args.productId)
			if (err != nil) != tt.wantErr {
				t.Errorf("IherbApiRepository.GetProductData() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("IherbApiRepository.GetProductData() = %v, want %v", got, tt.want)
			}
		})
	}
}
