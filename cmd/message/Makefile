server:
	kitex --thrift-plugin validator -module douyin idl/message.thrift # 在项目根目录下执行
	kitex --thrift-plugin validator -module douyin -service comment -use douyin/kitex_gen ../../idl/message.thrift # 在 cmd/message 目录下执行