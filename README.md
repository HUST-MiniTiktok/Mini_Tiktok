# 迷你抖音 - mini_tiktok
## 项目简介
* 本项目为第六届字节跳动青训营大项目，作者为“HUST字节青训营一简易抖音实现”团队
* 本项目为基于Golang与Hertz + Kitex + Gorm框架开发的微服务架构的极简版抖音APP后端项目
## 项目依赖
* Ubuntu 22.04 LTS
* Golang 1.20.7
* Docker 24.0.5
* FFmpeg 5.1.2
## 项目部署
### 配置
根据项目运行环境填写[**Conf包**](https://github.com/HUST-MiniTiktok/mini_tiktok/tree/master/pkg/conf)下的YAML配置文件
### 启动
```bash
chmod +x ./start-all.sh && ./start-all.sh
```
### 停止
```bash
chmod +x ./stop-all.sh && ./stop-all.sh
```
