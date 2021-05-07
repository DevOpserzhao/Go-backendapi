package encrypt

import "testing"

func TestEncrypt_MD5EncryptS(t *testing.T) {
	tests := []struct {
		name string
		args Encrypt
		want bool
	}{
		{
			name: "first",
			args: Encrypt{
				Salt: "salt1",
				S:  "123456",
			},
			want: true,
		},
		{
			name: "second",
			args: Encrypt{
				Salt: "salt2",
				S:  "891011",
			},
			want: true,
		},
	}
	t.Helper()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			encryptPwd := tt.args.MD5EncryptS()
			if got := tt.args.CheckMD5EncryptS(encryptPwd); got != tt.want {
				t.Errorf("Failed: got %+v!= want %+v", got, tt.want)
			}
		})
	}
}

func BenchmarkEncrypt_MD5EncryptS(b *testing.B) {
	test := Encrypt{
		Salt:     "salt",
		S: "123456",
	}
	b.Helper()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			test.MD5EncryptS()
		}
	})
}

func BenchmarkEncrypt_CheckMD5EncryptS(b *testing.B) {
	test := Encrypt{
		Salt:     "salt",
		S: "123456",
	}
	b.Helper()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			test.CheckMD5EncryptS("")
		}
	})
}