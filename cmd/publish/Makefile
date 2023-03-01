server:
	kitex --thrift-plugin validator -module douyin idl/publish.thrift # 在项目根目录下执行
	kitex --thrift-plugin validator -module douyin -service comment -use douyin/kitex_gen ../../idl/publish.thrift # 在 cmd/publish 目录下执行