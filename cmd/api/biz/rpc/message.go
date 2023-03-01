package rpc

import (
	"context"
	"douyin/kitex_gen/message"
	"douyin/kitex_gen/message/messageservice"
	"douyin/pkg/constant"
	"douyin/pkg/mw"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/kitex-contrib/obs-opentelemetry/provider"
	"github.com/kitex-contrib/obs-opentelemetry/tracing"
	etcd "github.com/kitex-contrib/registry-etcd"
)

var messageClient messageservice.Client

func initMessage() {
	r, err := etcd.NewEtcdResolver([]string{constant.ETCDAddress})
	if err != nil {
		panic(err)
	}
	provider.NewOpenTelemetryProvider(
		provider.WithServiceName(constant.ApiServiceName),
		provider.WithExportEndpoint(constant.ExportEndpoint),
		provider.WithInsecure(),
	)
	c, err := messageservice.NewClient(
		constant.MessageServiceName,
		client.WithResolver(r),
		client.WithMuxConnection(1),
		client.WithMiddleware(mw.CommonMiddleware),
		client.WithInstanceMW(mw.ClientMiddleware),
		client.WithSuite(tracing.NewClientSuite()),
		client.WithClientBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: constant.ApiServiceName}),
	)
	if err != nil {
		panic(err)
	}
	messageClient = c
}

func SendMessage(ctx context.Context, req *message.DouyinMessageActionRequest) (*message.DouyinMessageActionResponse, error) {
	return messageClient.SendMessage(ctx, req)
}

func GetMessageChat(ctx context.Context, req *message.DouyinMessageChatRequest) (*message.DouyinMessageChatResponse, error) {
	return messageClient.GetMessageChat(ctx, req)
}
