package logentries

import (
	"testing"
)

func TestReadFile(t *testing.T) {
	certs := GetCerts("./test.txt")
	if string(certs) != "test" {
		t.Fatalf("should have been test, but was %s", certs)
	}
	certs = GetCerts("invalid_file.txt")
	if string(certs) != string(defaultCerts) {
		t.Fatalf("should have been default certs, but was %s", certs)
	}
}
