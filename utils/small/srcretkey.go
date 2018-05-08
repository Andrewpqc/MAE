package small

import "os"

func GetSecretKey()string{
	//注意在生产环境需要设置SECRETKEY的值
	if secretKey:=os.Getenv("SECRETKEY");secretKey !=""{
		return secretKey
	}else{
		return "a secret key for development environment"
	}
}