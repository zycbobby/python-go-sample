# 内嵌的python2解释器

- python 2.7.15

使用第三方库的时候，需要`cd py-service && pip2 install -r  requirements.txt -t ./third_party`

__非常诡异地__，这种情况下依然无法拉起grpc_server，需要在third_party底下，touch google/__init__.py才可以。不然pb文件会报`ImportError: No module named google.protobuf`


# 示例

- `go run main.go -pyfile=./py-service/main.py`, trivial sample
- `go run main.go -pyfile=./py-service/flask_server.py`， flask http server
- `go run main.go -pyfile=./py-service/grpc_server.py`， grpc server


# TODO
- [x] py文件的初始化路径，不太希望使用`os.Chdir`