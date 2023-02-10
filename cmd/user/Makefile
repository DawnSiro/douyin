server:
	kitex --thrift-plugin validator -module douyin idl/user.thrift # execute in the project root directory
	kitex --thrift-plugin validator -module douyin -service comment -use douyin/kitex_gen ../../idl/user.thrift # execute in cmd/user