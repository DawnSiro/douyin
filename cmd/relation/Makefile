server:
	kitex --thrift-plugin validator -module douyin idl/relation.thrift # 在项目根目录下执行
	kitex --thrift-plugin validator -module douyin -service comment -use douyin/kitex_gen ../../idl/relation.thrift # 在 cmd/relation 目录下执行