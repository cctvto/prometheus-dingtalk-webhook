# prometheus-dingtalk-webhook
# 以https://github.com/yunlzheng/alertmanaer-dingtalk-webhook为基础修改的
---
# 使用方法
```shell
cd cmd/webhook
#go build
webhook -defaultRobot=https://oapi.dingtalk.com/robot/send?access_token=xxxx -accessToken=xxx -secretRobot=xxx
#go run
go run webhook.go -defaultRobot=https://oapi.dingtalk.com/robot/send?access_token=xxxx -accessToken=xxx -secretRobot=xxx
```
- defaultRobot: 钉钉机器人自定义webhook url
- accessToken:  钉钉机自定义webhook url中的token字符串
- secretRobot:  钉钉自定义webhook 签名密钥
### 说明: 如果机器人没有启用加签功能就不用accessToken、secretRobot这两个参数
## 另外可以在Prometheus alert配置中覆盖webhookurl
```shell
groups:
- name: hostStatsAlert
  rules:
  - alert: hostCpuUsageAlert
    expr: sum(avg without (cpu)(irate(node_cpu{mode!='idle'}[5m]))) by (instance) > 0.85
    for: 1m
    labels:
      severity: page
    annotations:
      summary: "Instance {{ $labels.instance }} CPU usgae high"
      description: "{{ $labels.instance }} CPU usage above 85% (current value: {{ $value }})"
      dingtalkRobot: "https://oapi.dingtalk.com/robot/send?access_token=xxxx"
```
