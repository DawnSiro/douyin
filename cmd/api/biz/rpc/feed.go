package rpc

import (
	"context"

	"douyin/kitex_gen/feed"
	"douyin/kitex_gen/feed/feedservice"
	"douyin/pkg/constant"
	"douyin/pkg/mw"

	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/kitex-contrib/obs-opentelemetry/provider"
	"github.com/kitex-contrib/obs-opentelemetry/tracing"
	etcd "github.com/kitex-contrib/registry-etcd"
)

var feedClient feedservice.Client

func initFeed() {
	r, err := etcd.NewEtcdResolver([]string{constant.ETCDAddress})
	if err != nil {
		panic(err)
	}
	provider.NewOpenTelemetryProvider(
		provider.WithServiceName(constant.ApiServiceName),
		provider.WithExportEndpoint(constant.ExportEndpoint),
		provider.WithInsecure(),
	)
	c, err := feedservice.NewClient(
		constant.FeedServiceName,
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
	feedClient = c
}

func GetFeed(ctx context.Context, req *feed.DouyinFeedRequest) (*feed.DouyinFeedResponse, error) {
	return feedClient.GetFeed(ctx, req)
}
