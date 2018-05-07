package jwt
import (

	"testing"
)


func TestJWTTool(t *testing.T){
	// 签名字符串
	sign := "fDEtrkpbQbocVxYRLZrnkrXDWJzRZMfO"

	token := NewJWToken(sign)

	// -----------  生成jwt token -----------
	tokenString, _ := token.GenJWToken(map[string]interface{}{
		"name": "root",
	})


	// -----------  解析 jwt token -----------
	got,_:=token.ParseJWToken(tokenString)
	want:="root"
	if got["name"] != want{
		t.Errorf("jwt tool failed!")
	}
}
