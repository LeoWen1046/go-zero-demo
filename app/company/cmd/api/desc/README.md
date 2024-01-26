1、编写api文档
2、执行命令生成一定的目录：goctl api go --api company.api --dir=../company --style=goZero
3、编写company.go文件
4、执行命令：go run company.go -f etc/company.yaml

// 生成swagger文档步骤：
1、go get -u github.com/zeromicro/goctl-swagger
2、将$GOPATH/bin中的goctl-swagger添加到环境变量
3、执行goctl api plugin -p goctl-swagger="swagger -filename company.json" --api company.api --dir .
