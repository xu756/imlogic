
### 文档地址
一个分布式 Im系统 拥有良好的扩展性和高可用性

<a href="https://s1zt69w7hzv.feishu.cn/wiki/space/7240272377303629826?ccm_open_type=lark_wiki_spaceLink&open_tab_from=wiki_home">飞书文档</a>

#### 启动
使用<a href="https://github.com/cortesi/modd" >`modd`<a/> 部署
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
cmd/public/**/*.* {
    prep: go build -o build/rpc/public-rpc  -v cmd/public/main.go;
    daemon +sigkill: build/rpc/public-rpc -f configs/dev.toml;
}
```
> cmd/api/**/*.*_    _监听 cmd/api/**/*.* 的所有文件

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
