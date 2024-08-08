package logic

import (
	"context"
	"imlogic/kitex_gen/im"
	"log"
)

// HandlerGroupMessage implements im.ImHandler.
func (i *ImRpcImpl) HandlerGroupMessage(ctx context.Context, req *im.Message) (res *im.MessageRes, err error) {
	userIds, err := i.Model.GetGroupUserIdsByGroupId(ctx, req.GroupId)
	if err != nil {
		return nil, err
	}
	users, err := i.Model.GetUsersByIds(ctx, userIds)
	if err != nil {
		return nil, err
	}
	for _, user := range users {
		conns, err := i.Model.GetUserConnsByUserId(ctx, user.ID)
		if err != nil {
			return nil, err
		}
		for _, conn := range conns {
			err = i.PrivateMq.PublishPrivateMessage(conn.LinkID, conn.HostName, req)
			if err != nil {
				log.Print("publish group message error: ", err)
			}
		}
	}
	return &im.MessageRes{
		Success: true,
	}, nil
}
