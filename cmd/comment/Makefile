server:
	kitex --thrift-plugin validator -module douyin idl/comment.thrift # 在项目根目录下执行
	kitex --thrift-plugin validator -module douyin -service comment -use douyin/kitex_gen ../../idl/comment.thrift # 在 cmd/comment 目录下执行