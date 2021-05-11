package encrypt

import "testing"

func TestBcrypt_BcryptPassword(t *testing.T) {
	tests := []struct {
		name string
		args Bcrypt
		want bool
	}{
		{
			name: "first",
			args: Bcrypt{
				Salt: "salt1",
				Password:  "123456",
			},
			want: true,
		},
		{
			name: "second",
			args: Bcrypt{
				Salt: "salt2",
				Password:  "891011",
			},
			want: true,
		},
	}
	t.Helper()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			encryptPwd, err := tt.args.BcryptPassword()
			t.Log(encryptPwd)
			t.Log(len(encryptPwd))
			if err != nil {
				t.Errorf("Failed: got %+v!= want %+v", err.Error(), tt.want)
			}
			if got := tt.args.CheckBcryptPassword(encryptPwd); got != tt.want {
				t.Errorf("Failed: got %+v!= want %+v", got, tt.want)
			}
		})
	}
}

func BenchmarkBcrypt_BcryptPassword(b *testing.B) {
	test := Bcrypt{
		Salt:     "salt",
		Password: "123456",
	}
	b.Helper()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_, _  = test.BcryptPassword()
		}
	})
}

func BenchmarkBcrypt_CheckBcryptPassword(b *testing.B) {
	test := Bcrypt{
		Salt:     "salt",
		Password: "123456",
	}
	b.Helper()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			test.CheckBcryptPassword("")
		}
	})
}