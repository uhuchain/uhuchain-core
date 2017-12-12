package tests

import (
	"reflect"
	"testing"

	"github.com/uhuchain/uhuchain-core/interfaces"
	"github.com/uhuchain/uhuchain-core/models"
	"github.com/uhuchain/uhuchain-core/test/mocks"
	"github.com/uhuchain/uhuchain-core/usecases"
)

func TestCarUsecase_SaveCar(t *testing.T) {
	type fields struct {
		carProvider interfaces.CarProvider
	}
	type args struct {
		car models.Car
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "Save car successfully",
			fields: fields{
				carProvider: &mocks.CarProviderMock{},
			},
			args: args{
				car: mocks.CreateCar(),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			usecase := usecases.NewCarUsecase(tt.fields.carProvider)
			if err := usecase.SaveCar(tt.args.car); (err != nil) != tt.wantErr {
				t.Errorf("CarUsecase.SaveCar() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestCarUsecase_GetCar(t *testing.T) {
	type fields struct {
		carProvider interfaces.CarProvider
	}
	type args struct {
		id int64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    models.Car
		wantErr bool
	}{
		{
			name: "Get car successfully",
			fields: fields{
				carProvider: &mocks.CarProviderMock{},
			},
			args: args{
				id: 12345,
			},
			want:    mocks.CreateCar(),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			usecase := usecases.NewCarUsecase(tt.fields.carProvider)
			got, err := usecase.GetCar(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("CarUsecase.GetCar() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CarUsecase.GetCar() = %v, want %v", got, tt.want)
			}
		})
	}
}
