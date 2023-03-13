# 1. 镜像

* 镜像就是停止运行的容器
* 容器即应用
* 镜像存储在Image Registry中， 默认使用Docker Hub

## 1.1 镜像命名和标签
>
> 使用镜像的名字和标签，可以定位到一个镜像, 使用冒号分隔  

 ```docker
  docker image pull <repository>:<tag>
 ```

## 1.2 过滤docker image ls输出内容  

* 使用--filer参数:
  * dangling: 返回悬虚镜像
    * before: 返回在之前被创建的所有镜像
    * since: 返回指定镜像之后创建的所有镜像
    * label: 根据label的名称或值对镜像进行过滤

悬虚镜像：构建一个新镜像，为该镜像打了一个已经存在的标签
    移除全部悬虚镜像： docker image prune
搜索docker hub中的内容：

* docker search name

镜像和分层： 由一些松耦合的只读镜像组成

* docker image inspect

查看构建历史：

* docker image history

## 1.3 Dockerfile

```shell
MAINTAINER user_name 声明镜像的作者
ENV key value 设置环境变量 (可以写多条)
RUN command 编译镜像时运行的脚本(可以写多条)
CMD 设置容器的启动命令
ENTRYPOINT 设置容器的入口程序
ADD source_dir/file dest_dir/file 将宿主机的文件复制到容器内，如果是一个压缩文件，将会在复 制后自动解压
COPY source_dir/file dest_dir/file 和ADD相似，但是如果有压缩文件并不能解压
WORKDIR path_dir 设置工作目录
ARG 设置编译镜像时加入的参数
VOLUMN 设置容器的挂载卷
```
