package hqiniu

import (
	"github.com/qiniu/api.v7/v7/auth/qbox"
	"github.com/qiniu/api.v7/v7/storage"
)

// UploadToken 获取七牛文件上传Token
// @params accessKey secretKey七牛分配的秘钥
// @params scope  bucket:key ，当key存在时覆盖上传
func UploadToken(accessKey, secretKey, scope string) string {
	putPolicy := storage.PutPolicy{Scope: scope}
	mac := qbox.NewMac(accessKey, secretKey)
	return putPolicy.UploadToken(mac)
}
