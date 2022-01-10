# sdcore

en: **sdcore is the core interface api for private data selfdata, which supports dynamic authorization and data encryption and decryption through custom building plug-ins**

zh: **sdcore是私有数据selfdata核心接口api，支持通过自定义构建插件的方式进行动态授权和数据加解密**

## [refitself][1]

## [selfdata][2]

## Supported API:

- func SDDecrypt(data []byte) []byte
- func SDEncrypt(data []byte) []byte
- func SDAuth() func(data []byte) error
- func SDPull(selfID string) (bool, []string, error)
- func SDPush(selfID, fileName string, data []byte) (bool, error)

## Development

setup for sdcore
1. install go

build sdcore as plugin:
1. git clone https://github.com/refitor/sdcore.git
2. cd sdcore && go mod init sdcore
4. cd ./plugin
4. go build -buildmode=plugin -o sdcore.so work.go

usage for sdcore:
selfdata write/read --plugin=sdcore.so

[1]: https://www.refitself.cn
[2]: https://www.refitself.cn/selfdata
