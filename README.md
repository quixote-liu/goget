# gget
一个类似于`wget`的工具，用来对指定的`URL`下载文件

## 为什么开发这个工具？
在`linux`下面安装`golang`环境时使用到了`wget`工具，个人感觉这个工具还是挺有意思的，所以这个工具就是重复造轮子。

## 如何做？
语言：目前准备使用`golang`去做，后面如果使用其他语言会再次开辟另一个库去做

1. 先做技术调研，学习前辈们的思路想法

2. 做技术取舍，第一个版本没有必要实现完整的、很复杂的功能，实现核心功能即可

3. 开发（包含测试代码）

4. 测试，在`linux`、`macOS`、`windows`上做正反测试

## 支持的功能（v1.0）
1. 语法基于`wget`不变：`wget [options] [url]`

2. 支持协议：`HTTP`,`HTTPS`,`FTP`

3. 支持使用不同的文件名：`gget -o email-cli https://github.com/quixote-liu/email`

4. 支持下载到指定的文件夹中：`gget -p /usr/local/ https://github.com/quixote-liu/email`

5. 下载多个文件：`gget -i linux-distros.txt`。其中的linux-distros.txt文件中的每行可以写入一个文件url

## 软件设计
总目标为扩展性、可用性强，为后面的扩展铺路，通过总结设计，整体软件模型可大致分为如下几层：

1. 命令解析层：通用命令，http命令，https命令，ftp命令

3. 命令执行层：解析完成后执行该命令，执行的实例有http、https、ftp

## 补充
目前暂时没有时间去做开发，先将想法写出来，等到后面有时间了再进行开发
