server:
	kitex --thrift-plugin validator -module douyin idl/feed.thrift # 在项目根目录下执行
	kitex --thrift-plugin validator -module douyin -service comment -use douyin/kitex_gen ../../idl/feed.thrift # 在 cmd/feed 目录下执行