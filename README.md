
### 文档地址
一个分布式 Im系统 拥有良好的扩展性和高可用性
[前端地址](https://github.com/xu756/imlogicapp)

[飞书文档](https://s1zt69w7hzv.feishu.cn/wiki/space/7240272377303629826?ccm_open_type=lark_wiki_spaceLink&open_tab_from=wiki_home?_blank)

#### 启动
使用[modd](https://github.com/cortesi/modd?_blank)部署
###### 安装
```bash
go install github.com/cortesi/modd/cmd/modd@latest
```
#### 配置 modd
创建 modd.conf
```bash
# cmd-api处理服务7080  
cmd/api/**/*.* {  
	prep: go build -o build/api/api  -v cmd/api/main.go;
	daemon +sigkill: build/api/api -f configs/dev.toml;
}
# cmd-public处理服务7081
cmd/user/**/*.* {
    prep: go build -o build/rpc/user-rpc  -v cmd/user/main.go;
    daemon +sigkill: build/rpc/user-rpc -f configs/dev.toml;
}

# cmd-im处理服务7082
cmd/im/**/*.* {
	prep: go build -o build/im/im  -v cmd/im/main.go;
	daemon +sigkill: build/im/im -f configs/dev.toml;
}
```
> `cmd/api/**/*.*` 监听 cmd/api/**/*.* 的所有文件

注意 internal 下的的文件监听不了哦
###### 启动
```bash
modd
```
#### 配置
```bash
....
├── configs
│   └── dev.toml					# 配置文件在这里
├── common	
│   ├── config
│   │   ├── config.go			#配置结构体在这里
│   │   ├── dir.go
│   │   ├── dir_test.go
│   │   ├── init.go
│   │   └── init_test.go
│   ├── db
.....
```
