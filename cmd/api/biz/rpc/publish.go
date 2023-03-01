package rpc

import (
	"context"
	"douyin/kitex_gen/publish"
	"douyin/kitex_gen/publish/publishservice"
	"douyin/pkg/constant"
	"douyin/pkg/mw"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/kitex-contrib/obs-opentelemetry/provider"
	"github.com/kitex-contrib/obs-opentelemetry/tracing"
	etcd "github.com/kitex-contrib/registry-etcd"
)

var publishClient publishservice.Client

func initPublish() {
	r, err := etcd.NewEtcdResolver([]string{constant.ETCDAddress})
	if err != nil {
		panic(err)
	}
	provider.NewOpenTelemetryProvider(
		provider.WithServiceName(constant.ApiServiceName),
		provider.WithExportEndpoint(constant.ExportEndpoint),
		provider.WithInsecure(),
	)
	c, err := publishservice.NewClient(
		constant.PublishServiceName,
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
	publishClient = c
}

func PublishAction(ctx context.Context, req *publish.DouyinPublishActionRequest) (*publish.DouyinPublishActionResponse, error) {
	return publishClient.PublishAction(ctx, req)
}

func GetPublishVideos(ctx context.Context, req *publish.DouyinPublishListRequest) (*publish.DouyinPublishListResponse, error) {
	return publishClient.GetPublishVideos(ctx, req)
}
