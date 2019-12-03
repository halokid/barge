# 基于Barge的改进版管理工具，基于Portainer API

本来想选用barge项目来用的，不过发现Barge只实现比较简单的显示功能，所以决定在此基础上增加一些功能，如下：
- [ ] 容器管理

    - 改进容器信息展示, 显示容器名字， 短id， 端口等信息

      

- [ ] 自定义的服务管理

    

- [ ] 日志展示

    

## 配置文件 .barge.yaml
```markdown
# 放在 $HOME目录下， 范例如下:
PORTAINER_URL: http://pontainer_host:9000
PORTAINER_USERNAME: admin
PORTAINER_PASSWORD: 1111111 

```



感谢原有的项目Barge :-)



# ------ 下面是原来Barge的 readme -----

# Barge

Barge is a command line tool for viewing multiple Docker Swarms through Portainer.

## Usage

### Get Barge

```bash
curl -L https://github.com/BenOvermyer/barge/releases/download/0.4.0/barge-linux-amd64 -o barge
chmod +x barge
```

### Config & Run

Optionally, these values can be provided in YAML notation via `barge.yaml`.

```bash
export PORTAINER_URL=https://portainer.mysite.com
export PORTAINER_USERNAME=myuser
export PORTAINER_PASSWORD=mypass
barge -h
```

## Contributing

During local development, you can use `go run` to build from source and run the target application all in one shot.

```bash
go run .
```

If you'd like to test it with arguments, you can append them. Here's an example with the container command.

```bash
go run . container --help
```
