package utils

import (
	"testing"
)

func TestGenerateToken(t *testing.T) {
	t.Log(TransformMd5("123456Lucky.麒麟"))
}
