# 微助教挂机助手 Teachermate Auto

**使用场景**：

- 一人到场扫码。

- 其他人在电脑上使用客户端，实现挂机签到。

## 平台支持

- 服务端支持多平台

- 客户端仅支持 Windows (使用了 WIN32API)

## Quick start

被签到人视角：

1. 在最新的 [Releases](https://github.com/HomeArchbishop/teachermate/releases) 中下载 `client-*-*.exe` 及 `config.yaml` 两个文件。将这两个文件放在同级目录下。

2. 双击运行 `client-*-*.exe`，输入**和扫码人约定好**的课程 ID

3. 双击打开微信**文件传输助手**的**独立窗口**（因为需要获取到窗口模拟键入）

4. 挂机。保持桌面常亮，不要切换多桌面。确保客户端命令行界面没有堵塞。

扫码人视角：

1. 使用如下方式扫码即可：

    - 扫码 H5 (under deveplopment)

    - [iOS 快捷指令](https://www.icloud.com/shortcuts/6c97a56256e44ebf9faca2782694bd30)

## 原理

项目由三部分构成：

1. 客户端 Client。与服务器建立 WS 连接，接收指定课程 ID 的签到码信息。识别文件传输助手的窗口，然后模拟输入与点击完成签到。

2. 服务器 Server。接收扫码端的请求，提取签到信息，再通过 WS 通知客户端签到。

3. 扫码端 Scanner。扫码人以此扫码，并将签到码以 HTTP 请求方式通知服务器。

至于为什么要使用模拟点击这么原始的办法，以及为什么还是需要一个扫码人，请看：[微助教扫码签到分析](docs/analysis.md)

## 已知问题

- 微信可能会因为网络原因签到失败。这个不在此应用的控制范围中。建议挂机时关闭系统代理等不稳定因素。

## 开发与部署

**Client**

```sh
sh ./scripts/build-client.sh
```

**Server**

```sh
sh ./scripts/build-server.sh
```

**config 构建配置**

```sh
sh ./scripts/build-config.sh --server-port=8080 --client-api-host=api.example.com:8080
# 参数说明:
#   --server-port       服务器暴露端口号
#   --client-api-host   客户端连接服务器主机名，其端口号一般应与 --server-port 一致
```

构建文件在 `./build` 下。

注意，Client 与 Server 都会从各自同级目录下读取 config.yaml 中的相应配置。请在使用/部署时，将 config.yaml 放在同级目录下。
