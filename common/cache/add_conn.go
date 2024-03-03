package cache

import (
	"context"
	"time"
)

var connHeartBeatInterval = 0 * time.Second // 60s

// 获取key conn:user:userId:hostname:LinkId
func getUserCacheKey(userId, hostName, linkId string) string {
	return "conn:user:" + userId + ":" + hostName + ":" + linkId
}

// 获取key conn:link:LinkId
func getLinkCacheKey(linkId string) string {
	return "conn:link:" + linkId
}

// 获取key conn:host:hostname:linkId
func getHostCacheKey(hostName, linkId string) string {
	return "conn:host:" + hostName + ":" + linkId
}

// LinkCacheInfo 链接信息
type LinkCacheInfo struct {
	UserId    string
	HostName  string
	Device    string
	LinkId    string
	Timestamp int64
}

func (c *Client) AddConn(ctx context.Context, userId, device, linkId, hostName string, timestamp int64) error {
	linkInfo := LinkCacheInfo{
		UserId:    userId,
		HostName:  hostName,
		Device:    device,
		LinkId:    linkId,
		Timestamp: timestamp,
	}
	err := c.Set(ctx, getUserCacheKey(userId, hostName, linkId), linkInfo, connHeartBeatInterval)
	if err != nil {
		return err
	}
	err = c.Set(ctx, getLinkCacheKey(linkId), linkInfo, connHeartBeatInterval)
	if err != nil {
		return err
	}
	err = c.Set(ctx, getHostCacheKey(hostName, linkId), linkInfo, connHeartBeatInterval)
	if err != nil {
		return err
	}
	return nil
}

//// DelConn 删除链接
//func (c *Client) DelConn(ctx context.Context, linkId string) error {
//
//}
