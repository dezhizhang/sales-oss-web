package main

import (
	"fmt"
	"os"
)

func handleError(err error) {
	fmt.Println("Error:", err)
	os.Exit(-1)
}

func main() {
	//LTAI5t6mBexkTfvNRCQ8sTAG
	endpoint := "oss-cn-hangzhou.aliyuncs.com"
	// 阿里云账号AccessKey拥有所有API的访问权限，风险很高。强烈建议您创建并使用RAM用户进行API访问或日常运维，请登录RAM控制台创建RAM用户。

	// yourBucketName填写存储空间名称。
	bucketName := "xiaozhicloud"
	// yourObjectName填写Object完整路径，完整路径不包含Bucket名称。
	objectName := "yourObjectName"
	// yourLocalFileName填写本地文件的完整路径。
	localFileName := "./image/avator.png"
	// 创建OSSClient实例。
	//client, err := oss.New(endpoint, accessKeyId, accessKeySecret)
	//if err != nil {
	//	handleError(err)
	//}
	//// 获取存储空间。
	//bucket, err := client.Bucket(bucketName)
	//if err != nil {
	//	handleError(err)
	//}
	//// 上传文件。
	//err = bucket.PutObjectFromFile(objectName, localFileName)
	//if err != nil {
	//	handleError(err)
	//}
}