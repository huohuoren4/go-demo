package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"path/filepath"
	"strconv"
	"strings"
)

type User struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

// restful-API路由
func main() {
	router := gin.Default()

	// 查询用户列表: GET /hello?name=123&password=12121
	router.Handle("GET", "/hello", func(context *gin.Context) {
		var user User
		err := context.ShouldBindQuery(&user)
		if err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{"msg": "请求信息错误!!!"})
			return
		}
		fmt.Println("请求参数: name=", context.Query("name"), ", password=", context.Query("password"))
		context.JSON(http.StatusOK, gin.H{"msg": "查询用户列表成功 !"})
	})

	// 查询用户信息: GET /hello/{id}
	router.Handle("GET", "/hello/:id", func(context *gin.Context) {
		fmt.Println("路径参数: id=", context.Param("id"))
		context.JSON(http.StatusOK, gin.H{"msg": "查询用户信息成功 !"})
	})

	// 创建用户: POST /hello  json: { "name": "123", "password": "123456" }
	router.Handle("POST", "/hello", func(context *gin.Context) {
		var user User
		err := context.ShouldBindJSON(&user)
		if err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{"msg": "请求信息错误!!!"})
			return
		}
		fmt.Println("请求体: name=", user.Name, ", password=", user.Password)
		context.JSON(http.StatusOK, gin.H{"msg": "创建用户成功 !"})
	})

	// 上传用户头像: POST /hello/{id}/upload  file=xxxx
	router.Handle("POST", "/hello/:id/upload", func(context *gin.Context) {
		fmt.Println("路径参数: id=", context.Param("id"))
		file, err := context.FormFile("image_file")
		fmt.Println("上传的文件名:", file.Filename)
		if err != nil {
			fmt.Println(err.Error())
			context.JSON(http.StatusInternalServerError, gin.H{"msg": "文件上传失败!!!"})
			return
		}
		// 文件大小: 2K
		if fileSize := int64(2 << 20); file.Size > fileSize {
			context.JSON(http.StatusInternalServerError, gin.H{"msg": "文件大小不能超过 " + strconv.Itoa(int(fileSize)) + "K !!!"})
			return
		}
		// 文件类型
		if ext := "jpg;jpeg;png"; !strings.Contains(ext, strings.ToLower(filepath.Ext(file.Filename))[1:]) {
			fmt.Println(strings.ToLower(filepath.Ext(file.Filename)))
			context.JSON(http.StatusInternalServerError, gin.H{"msg": "文件类型只支持" + ext + " !!!"})
			return
		}
		// 保存文件
		if err = context.SaveUploadedFile(file, "./test.png"); err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{"msg": "文件保存失败!!!"})
			return
		}
		context.JSON(http.StatusOK, gin.H{"msg": "用户头像上传成功 !"})
	})

	// 修改用户信息: PUT /hello  json: { "name": "123" }
	router.Handle("PUT", "/hello/:id", func(context *gin.Context) {
		fmt.Println("路径参数: id=", context.Param("id"))
		var user User
		err := context.ShouldBindJSON(&user)
		if err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{"msg": "请求信息错误!!!"})
			return
		}
		fmt.Println("请求体: name=", user.Name, ", password=", user.Password)
		context.JSON(http.StatusOK, gin.H{"msg": "修改用户信息成功 !"})
	})

	// 删除用户: DELETE /hello/{id}
	router.Handle("DELETE", "/hello/:id", func(context *gin.Context) {
		fmt.Println("路径参数: id=", context.Param("id"))
		context.JSON(http.StatusOK, gin.H{"msg": "删除用户成功 !"})
	})

	err := router.Run("0.0.0.0:8080")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
}
