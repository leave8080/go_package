package util

import (
	"github/leave8080/go_package/pkg/log"
	"testing"
)

func TestVerifyPhoneNumber(t *testing.T) {
	log.Debug(VerifyPhoneNumber("11111111111"))
}
