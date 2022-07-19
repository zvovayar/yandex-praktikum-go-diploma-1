package businesslogic

import (
	"fmt"
	"os"
	"testing"

	"github.com/zvovayar/yandex-praktikum-go-diploma-1/internal/accrualclient"
	config "github.com/zvovayar/yandex-praktikum-go-diploma-1/internal/config/cls"
	"github.com/zvovayar/yandex-praktikum-go-diploma-1/internal/storage"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func TestMain(m *testing.M) {
	fmt.Print("TestMain run\n")

	// config test database URI
	config.ConfigCLS.DataBaseURI = "postgres://postgres:qweasd@localhost:5432/yandexdiploma1_test?sslmode=disable"

	os.Exit(m.Run())

}
func TestBusinessSession_RegisterNewUser(t *testing.T) {

	// clear table users
	db, _ := gorm.Open(postgres.Open("postgres://postgres:qweasd@localhost:5432/yandexdiploma1_test?sslmode=disable"),
		&gorm.Config{})
	db.Exec("truncate gorm_users")

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
		{
			name: "add user again",
			fields: fields{
				AccrualClient: accrualclient.Accrual{
					Address: "",
				},
			},
			args: args{
				u: storage.User{
					Login:      "user",
					PasswdHash: "password2",
				},
			},
			wantStatus: 409,
			wantErr:    true,
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
		{
			name: "login user \"user\" with bad password",
			fields: fields{
				AccrualClient: accrualclient.Accrual{
					Address: "",
				},
			},
			args: args{
				u: storage.User{
					Login:      "user",
					PasswdHash: "password2",
				},
			},
			wantStatus: 401,
			wantErr:    true,
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
