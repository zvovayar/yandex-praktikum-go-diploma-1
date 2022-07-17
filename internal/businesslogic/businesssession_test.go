package businesslogic

import (
	"testing"

	"github.com/zvovayar/yandex-praktikum-go-diploma-1/internal/accrualclient"
	"github.com/zvovayar/yandex-praktikum-go-diploma-1/internal/storage"
)

func TestBusinessSession_RegisterNewUser(t *testing.T) {
	type fields struct {
		AccrualClient accrualclient.Accrual
	}
	type args struct {
		u storage.User
	}
	tests := []struct {
		name       string
		fields     fields
		args       args
		wantStatus int
		wantErr    bool
	}{
		{
			name: "add user",
			fields: fields{
				AccrualClient: accrualclient.Accrual{
					Address: "",
				},
			},
			args: args{
				u: storage.User{
					Login:      "user",
					PasswdHash: "password",
				},
			},
			wantStatus: 200,
			wantErr:    false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bs := &BusinessSession{
				AccrualClient: tt.fields.AccrualClient,
			}
			gotStatus, err := bs.RegisterNewUser(tt.args.u)
			if (err != nil) != tt.wantErr {
				t.Errorf("BusinessSession.RegisterNewUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotStatus != tt.wantStatus {
				t.Errorf("BusinessSession.RegisterNewUser() = %v, want %v", gotStatus, tt.wantStatus)
			}
		})
	}
}

func TestBusinessSession_UserLogin(t *testing.T) {
	type fields struct {
		AccrualClient accrualclient.Accrual
	}
	type args struct {
		u storage.User
	}
	tests := []struct {
		name       string
		fields     fields
		args       args
		wantStatus int
		wantErr    bool
	}{
		// TODO: Add test cases.
		{
			name: "login user \"user\"",
			fields: fields{
				AccrualClient: accrualclient.Accrual{
					Address: "",
				},
			},
			args: args{
				u: storage.User{
					Login:      "user",
					PasswdHash: "password",
				},
			},
			wantStatus: 200,
			wantErr:    false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bs := &BusinessSession{
				AccrualClient: tt.fields.AccrualClient,
			}
			gotStatus, err := bs.UserLogin(tt.args.u)
			if (err != nil) != tt.wantErr {
				t.Errorf("BusinessSession.UserLogin() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotStatus != tt.wantStatus {
				t.Errorf("BusinessSession.UserLogin() = %v, want %v", gotStatus, tt.wantStatus)
			}
		})
	}
}
