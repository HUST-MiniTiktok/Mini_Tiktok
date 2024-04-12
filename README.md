# 迷你抖音 - Mini_Tiktok
[查看项目文档](https://p01un12ypkm.feishu.cn/docx/Wb1ldLO6xo3LZzx8xR0chpiYntg)
## 主要贡献者
* 队长：余天泽 [@yutianzeabc](https://github.com/yutianzeabc)，负责技术选型、整体架构设计与服务部署，负责API网关、基础接口、社交微服务开发
* 队员：杨涛瑞 [@alteracd](https://github.com/alteracd)，负责Redis缓存设计、功能、性能测试，负责互动微服务开发
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
