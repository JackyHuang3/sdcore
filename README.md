# sdcore

en: sdcore is the core interface api of selfdata, which supports dynamic data authorization and encryption and decryption through custom building plug-ins

zh: sdcore是selfdata核心接口api，支持通过自定义构建插件的方式进行数据动态授权和加解密

Supported API:

type T_SDDecrypt func(data []byte) []byte
type T_SDEncrypt func(data []byte) []byte
type T_SDAuth func() func(data []byte) error
type T_SDPull func(selfID string) (bool, []string, error)
type T_SDPush func(selfID, fileName string, data []byte) (bool, error)
