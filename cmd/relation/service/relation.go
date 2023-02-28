package service

import (
	"context"

	"douyin/cmd/relation/pack"
	"douyin/dal/db"
	"douyin/kitex_gen/relation"
	"douyin/pkg/errno"

	"github.com/cloudwego/kitex/pkg/klog"
)

type RelationService struct {
	ctx context.Context
}

func NewRelationService(ctx context.Context) *RelationService {
	return &RelationService{
		ctx: ctx,
	}
}

func (s *RelationService) Follow(userID, toUserID uint64) (*relation.DouyinRelationActionResponse, error) {
	if userID == toUserID {
		errNo := errno.UserRequestParameterError
		errNo.ErrMsg = "不能自己关注自己哦"
		klog.Error("service.relation.Follow err:", errNo.Error())
		return nil, errNo
	}
	isFollow := db.IsFollow(s.ctx, userID, toUserID)
	if isFollow {
		klog.Error("service.relation.Follow err:", errno.RepeatOperationError)
		return nil, errno.RepeatOperationError
	}

	//关注操作
	err := db.Follow(s.ctx, userID, toUserID)
	if err != nil {
		klog.Error("service.relation.Follow err:", err.Error())
		return nil, err
	}
	return &relation.DouyinRelationActionResponse{
		StatusCode: errno.Success.ErrCode,
	}, nil
}

func (s *RelationService) CancelFollow(userID, toUserID uint64) (*relation.DouyinRelationActionResponse, error) {
	if userID == toUserID {
		errNo := errno.UserRequestParameterError
		errNo.ErrMsg = "不能自己取关自己哦"
		klog.Error("service.relation.CancelFollow err:", errNo.Error())
		return nil, errNo
	}
	//取消关注
	err := db.CancelFollow(s.ctx, userID, toUserID)
	if err != nil {
		klog.Error("service.relation.CancelFollow err:", err.Error())
		return nil, err
	}
	return &relation.DouyinRelationActionResponse{
		StatusCode: errno.Success.ErrCode,
	}, nil
}

// GetFollowList
// userID 为发送请求的用户ID，从 Token 里取到
// selectUserID 为需要查询的用户的ID，做为请求参数传递
func (s *RelationService) GetFollowList(userID, selectUserID uint64) (*relation.DouyinRelationFollowListResponse, error) {
	dbUserList, err := db.GetFollowList(s.ctx, selectUserID)
	if err != nil {
		klog.Error("service.relation.GetFollowList err:", err.Error())
		return nil, err
	}

	// 提前申请好数组大小来避免后续扩容
	userList := make([]*relation.User, 0, len(dbUserList))
	// TODO 存在循环查询DB
	for _, v := range dbUserList {
		if userID == selectUserID {
			// 自己的关注列表自己当然都关注了，无需查数据库
			userList = append(userList, pack.User(v, true))
		} else {
			// 这里要查的是，自己是否关注了查询的用户的关注列表的人
			userList = append(userList, pack.User(v, db.IsFollow(s.ctx, selectUserID, v.ID)))
		}
	}

	return &relation.DouyinRelationFollowListResponse{
		StatusCode: errno.Success.ErrCode,
		UserList:   userList,
	}, nil

}

func (s *RelationService) GetFollowerList(userID, selectUserID uint64) (*relation.DouyinRelationFollowerListResponse, error) {
	dbUserList, err := db.GetFollowerList(s.ctx, selectUserID)
	if err != nil {
		klog.Error("service.relation.GetFollowerList err:", err.Error())
		return nil, err
	}

	// 提前申请好数组大小来避免后续扩容
	userList := make([]*relation.User, 0, len(dbUserList))
	// TODO 存在循环查询DB
	for _, v := range dbUserList {
		// 这里要查的是，自己是否关注了查询的用户的粉丝列表的人
		userList = append(userList, pack.User(v, db.IsFollow(s.ctx, userID, v.ID)))
	}
	return &relation.DouyinRelationFollowerListResponse{
		StatusCode: errno.Success.ErrCode,
		UserList:   userList,
	}, nil
}

func (s *RelationService) GetFriendList(userID uint64) (*relation.DouyinRelationFriendListResponse, error) {
	userList, err := db.GetFriendList(s.ctx, userID)
	if err != nil {
		klog.Error("service.relation.GetFollowerList err:", err.Error())
		return nil, err
	}

	// TODO 存在循环查询DB
	friendUserList := make([]*relation.FriendUser, 0, len(userList))
	for _, u := range userList {
		msg, err := db.GetLatestMsg(s.ctx, userID, u.ID)
		if err != nil {
			klog.Error("service.relation.GetFollowerList err:", err.Error())
			return nil, err
		}
		friendUserList = append(friendUserList, pack.FriendUser(u, db.IsFollow(s.ctx, userID, u.ID), msg.Content, msg.MsgType))
	}

	return &relation.DouyinRelationFriendListResponse{
		StatusCode: errno.Success.ErrCode,
		UserList:   friendUserList,
	}, nil
}
