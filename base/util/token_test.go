package util_test

import (
	"go-video/base/util"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestToken(t *testing.T) {
	// 创建一个token
	token, err := util.CreateToken("test", "jwtSecret")
	if err != nil {
		t.Errorf("Failed to create token: %v", err) // 如果创建token失败，标记测试失败
		return
	}

	// 解析token
	claims, err := util.ParseToken(token, "jwtSecret")
	if err != nil {
		t.Errorf("Failed to parse token: %v", err) // 如果解析token失败，标记测试失败
		return
	}

	// 断言：确保解析出的用户名是正确的
	assert.Equal(t, "test", claims.Username, "Token does not contain the expected username")
}
