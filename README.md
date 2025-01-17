# godscan
<h4 align="center">你的下一台扫描器，又何必是扫描器</h4>

<p align="center">
  <a href="https://goreportcard.com/report/github.com/godspeedcurry/godscan">
    <img src="https://goreportcard.com/badge/github.com/godspeedcurry/godscan">	
  </a>
  <a href="https://opensource.org/licenses/MIT">
    <img src="https://img.shields.io/badge/license-MIT-_red.svg">
  </a>
  <a href="https://github.com/godspeedcurry/godscan/releases">
  	<img src="https://img.shields.io/github/downloads/godspeedcurry/godscan/total">
  </a>
</p>

## Usage
### 命令自动补全
```
./godscan completion zsh > /tmp/x
source /tmp/x
```
### 目录扫描
```bash
./godscan dirbrute -u http://www.example.com
```
1. 目录扫描数量较少，只针对渗透测试中容易造成数据泄漏、命令执行的几个点进行了探测
2. 目录扫描会根据域名生成对应的备份文件路径，说不定会有意外之喜
3. 要对单一url进行大线程、多文件探测，请使用[dirsearch](https://github.com/maurosoria/dirsearch)

### 批量目录扫描+指纹识别(协程)
```bash
./godscan dirbrute --url-file url.txt
```

### 根据图标地址计算图标hash
```bash
./godscan icon -u http://www.example.com/ico.ico
```

### 弱口令生成、离线爆破
会自动识别身份证、电话号码，并根据常见的弱口令规则生成对应的弱口令
```bash
./godscan weakpass -k "张三,110101199003070759,18288888888"
# 中文会被转成英文，以一定格式生成弱口令，如干饭集团，需要自己去找一下他在网站中经常提到的一些叫法
./godscan weakpass -k "干饭,干饭集团,干饭有限公司"

# 自定义后缀
./godscan weakpass -k "张三,110101199003070759,18288888888" -suffix '123,qwe,123456'

# 查看工具默认的后缀
./godscan weakpass --show
# 更为复杂的前后缀，适合本地跑hashcat,本方法还会对字符串作变异，如o->0,i->1,a->4等等
./godscan weakpass -k '百度' --full > 1.txt  

# 自定义后缀
./godscan weakpass -k '百度,baidu.com,password,pass,root,server,qwer,admin' --prefix '@,!,",123' --suffix '!,1234,123,321' --sep '_,!,.,/,&,+' > 1.txt

# -l 获取python格式的list 如["11","222"]
# mac下拷贝至剪贴板，其余系统可自行探索哈
./godscan weakpass -k "张三,110101199003070759,18288888888" | pbcopy
```

### 爬虫递归探测URL、指纹和敏感信息
```bash
./godscan weakpass -u 'http://www.exmaple.com' 
# -d 1 可以指定爬虫的深度
```



### 爬取js中的api地址，用于未授权测试
```bash
./godscan spider -u http://example.com -d 2
# 假如知道前缀，可以得到一些路径，便于burp中爆破，这些路径可能需要GET、POST等类型去做测试 
./godscan spider -u http://example.com -d 2 --api "/api/v1"
```


## 功能详细介绍
### web应用指纹识别 
- [x] HTTP响应 Server字段
- [x] 构造404 报错 得到中间件的详情
- [x] POST请求构造报错 
- [x] 爬虫 递归访问
- [x] 正则提取注释 注释里往往有版本 github仓库等信息
- [x] 版本识别并高亮
- [x] 对注释里的内容匹配到关键字并高亮
- [x] 识别接口 主要从js里提取
- [x] url特征 人工看吧 有些组件的url是很有特征的 google: `inurl:/wh/servlet`
- [x] vue.js 前端 识别`(app|index|main|config).xxxx.js` 并使用正则提取里面的path
- [x] finger.txt来源
  * [Ehole](https://raw.githubusercontent.com/EdgeSecurityTeam/EHole/main/finger.json)
  * [nemasis](https://www.nemasisva.com/resource-library/Nemasis-Supported-Applications-Hardware-and-Platforms.pdf)
- [x] 图标哈希
  - [x] 新增可直接根据图标地址计算hash的功能

### 功能二：弱口令生成器
- [x] 在fscan的基础上新增从若干个报告中获取到的弱口令
- [x] 根据给定的keyword list 生成 逗号分隔


人工筛选时，可关注
* 域名 
  * 自身域名
  * 非自身域名 如开发商 很多默认密码和域名有千丝万缕的联系
* html注释


## 功能三：敏感信息搜集
* https://gh0st.cn/HaE/
这里面有很多现成的规则 挑了一下重点
- [x] JSON-WEB-Token
- [x] 国内手机号
- [x] 邮箱
- [x] hmtl注释
- [x] ueditor swagger 等
- [x] OSS accessKey accessId
- [x] link 识别  


---


## 跨平台编译
```bash
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-s -w " -trimpath -o godscan_linux_amd64 
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -ldflags="-s -w " -trimpath -o godscan_win_amd64
CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -ldflags="-s -w " -trimpath -o godscan_darwin_amd64
CGO_ENABLED=0 GOOS=darwin GOARCH=arm64 go build -ldflags="-s -w " -trimpath -o godscan_darwin_arm64
```

## 更新说明
* 2023-09-13 新增js中提取api路径的功能
* 2023-08-11 新增目录单线程爆破的功能，并会根据域名爆破一些备份文件
* 2023-07-03 新增直接对icon计算hash的功能
* 2023-07-02 新增批量url关键路径扫描的功能
* 2022-08-01 新增部分真实场景中得到的弱口令 新增弱口令后缀，如123,qwe等，丰富生成后的弱口令
* 2022-07-04 修复没有子路径的bug, 移除packr 改用原生的embed库进行静态资源的打包
* 2022-06-10 更新了正则 对输出的表格进行了优化
* 2022-06-09 修复了大小写导致不高亮的问题
* 2022-06-08 修复了os.Open导致找不到文件的错误，改用packr库

## 功能截图
* icon_hash计算、关键字识别
![image](https://github.com/godspeedcurry/godscan/blob/main/images/img1.jpg)

* cms高亮
![image](https://github.com/godspeedcurry/godscan/blob/main/images/img2.png)

* 敏感信息识别
![image](https://github.com/godspeedcurry/godscan/blob/main/images/img3.png)

## 开发
```
git add . && git commit -m "fix bug" && git push -u origin main
git tag -a v1.xx
git push -u origin v1.xx
# delete
git tag -d v1.xx
git push origin :refs/tags/v1.1.10
```

---
## 未来目标

### 端口扫描+协议识别 TODO
* https://github.com/4dogs-cn/TXPortMap
* https://github.com/redtoolskobe/scaninfo
