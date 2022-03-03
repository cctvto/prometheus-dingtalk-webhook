package main

import (
	"flag"
	"net/http"
	"prometheus-dingtalk-webhook/model"
	"prometheus-dingtalk-webhook/notifier"
	ding "prometheus-dingtalk-webhook/signurl"

	"github.com/gin-gonic/gin"
)

var (
	h            bool
	defaultRobot string
	enableSign   bool
	secretRobot  string
	accessToken  string
)

func init() {
	flag.BoolVar(&h, "h", false, "help")
	flag.BoolVar(&enableSign, "enableSign", false, "sign to dingtalk url")
	flag.StringVar(&defaultRobot, "defaultRobot", "", "global dingtalk robot webhook")
	flag.StringVar(&secretRobot, "secretRobot", "", " dingtalk robot webhook sgin secret")
	flag.StringVar(&accessToken, "accessToken", "", " dingtalk robot webhook accessToken")
}

func main() {
	flag.Parse()

	if h {
		flag.Usage()
		return
	}
	if enableSign {
		//对url加签名
		d := ding.Webhook{
			AccessToken: accessToken,
			Secret:      secretRobot,
		}
		defaultRobot = d.GetURL()
	}

	//创建路由
	router := gin.Default()
	//绑定路由gin.Context，封装了request和response
	router.POST("/webhook", func(c *gin.Context) {

		//声明接收的变量
		var notification model.Notification
		//将request的body中的数据按照json格式解析到结构体
		err := c.BindJSON(&notification)
		//如果发送的不是json格式输出错误信息
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		//转发到webhook
		err = notifier.Send(notification, defaultRobot)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		}
		//输出结果
		c.JSON(http.StatusOK, gin.H{"message": " successful receive alert notification message!"})
	})
	router.Run()
}
