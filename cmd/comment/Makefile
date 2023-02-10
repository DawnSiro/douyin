server:
	kitex --thrift-plugin validator -module douyin idl/comment.thrift # execute in the project root directory
	kitex --thrift-plugin validator -module douyin -service comment -use douyin/kitex_gen ../../idl/comment.thrift # execute in cmd/comment