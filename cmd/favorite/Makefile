server:
	kitex --thrift-plugin validator -module douyin idl/favorite.thrift # 在项目根目录下执行
	kitex --thrift-plugin validator -module douyin -service comment -use douyin/kitex_gen ../../idl/favorite.thrift # 在 cmd/favorite 目录下执行