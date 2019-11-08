# 内嵌的python2解释器

- python 2.7.15

使用第三方库的时候，需要`cd py-service && pip2 install -r  requirements.txt -t ./third_party`
\
__非常诡异地__，这种情况下依然无法拉起grpc_server，需要在third_party底下，touch google/\_\_init\_\_.py才可以。不然pb文件会报`ImportError: No module named google.protobuf`


# 示例

- `go run main.go -pyfile=./py-service/main.py`, trivial sample
- `go run main.go -pyfile=./py-service/flask_server.py`， flask http server
- `go run main.go -pyfile=./py-service/grpc_server.py`， grpc server


# TODO
- [x] py文件的初始化路径，不太希望使用`os.Chdir`
- [ ] 可控的热更（如果不可控还不如把Python虚拟机重新拉起来一下）



# Memo

- 游戏服的网络操作其实是放在C++层的
- 游戏服的Python层也是有状态的
- 游戏服有可控的Python层热更
- 如果python不能调用go的基础方法，似乎和直接起一个python进程是没有什么区别的