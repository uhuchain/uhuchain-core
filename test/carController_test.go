package tests

import (
	"testing"

	"github.com/uhuchain/uhuchain-core/controller"
	"github.com/uhuchain/uhuchain-core/interfaces"
	"github.com/uhuchain/uhuchain-core/test/mocks"
)

func TestCarController_Run(t *testing.T) {
	type fields struct {
		provider interfaces.CarProvider
	}
	type args struct {
		function string
		args     []string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []byte
		wantErr bool
	}{
		{
			name: "Run controller with getCar successfully",
			fields: fields{
				provider: &mocks.CarProviderMock{},
			},
			args: args{
				function: "getCar",
				args:     []string{"12345"},
			},
			want:    []byte(mocks.CreateCarPayload()),
			wantErr: false,
		},
		{
			name: "Run controller with saveCar successfully",
			fields: fields{
				provider: &mocks.CarProviderMock{},
			},
			args: args{
				function: "saveCar",
				args:     []string{mocks.CreateCarPayload()},
			},
			want:    nil,
			wantErr: false,
		},
		{
			name: "Throw error on unknown fuction",
			fields: fields{
				provider: &mocks.CarProviderMock{},
			},
			args: args{
				function: "noneExistingFunction",
				args:     []string{""},
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			controller := controller.NewCarController(tt.fields.provider)
			_, err := controller.Run(tt.args.function, tt.args.args)
			if (err != nil) != tt.wantErr {
				t.Errorf("CarController.Run() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
