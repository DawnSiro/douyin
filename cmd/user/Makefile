server:
	kitex --thrift-plugin validator -module douyin idl/user.thrift # 在项目根目录下执行
	kitex --thrift-plugin validator -module douyin -service comment -use douyin/kitex_gen ../../idl/user.thrift # 在 cmd/user 目录下执行