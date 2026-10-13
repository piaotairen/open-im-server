// Copyright © 2023 OpenIM. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package msgtransfer

import (
	"context"

	cbapi "github.com/openimsdk/open-im-server/v3/pkg/callbackstruct"
	"github.com/openimsdk/protocol/constant"
	"github.com/openimsdk/protocol/sdkws"
	"github.com/openimsdk/tools/mcontext"
	"google.golang.org/protobuf/proto"
)

// zhilian v2.2.0 A1: the rpc-level afterSend webhook fires before seq is
// allocated (msg is only queued to Kafka at that point), so the callback
// always carried openImSeq=0. Here in msgtransfer the messages have just been
// assigned their authoritative seq by BatchInsertChat2Cache, so we re-emit the
// same afterSend callback commands with the real seq filled in.
//
// Operators must disable the rpc-level afterSend webhooks (env
// IMENV_WEBHOOKS_AFTERSENDSINGLEMSG_ENABLE / _AFTERSENDGROUPMSG_ENABLE) and
// enable this one (IMENV_WEBHOOKS_AFTERMSGPERSIST_ENABLE) to avoid duplicates.
func (och *OnlineHistoryRedisConsumerHandler) webhookAfterMsgPersist(ctx context.Context, msg *sdkws.MsgData) {
	if msg.ContentType == constant.Typing {
		return
	}
	after := &och.webhooksConfig.AfterMsgPersist
	if !after.Enable {
		return
	}
	common := cbapi.CommonCallbackReq{
		SendID:           msg.SendID,
		ServerMsgID:      msg.ServerMsgID,
		ClientMsgID:      msg.ClientMsgID,
		OperationID:      mcontext.GetOperationID(ctx),
		SenderPlatformID: msg.SenderPlatformID,
		SenderNickname:   msg.SenderNickname,
		SessionType:      msg.SessionType,
		MsgFrom:          msg.MsgFrom,
		ContentType:      msg.ContentType,
		Status:           msg.Status,
		SendTime:         msg.SendTime,
		CreateTime:       msg.CreateTime,
		AtUserIDList:     msg.AtUserIDList,
		SenderFaceURL:    msg.SenderFaceURL,
		Content:          persistCallbackContent(msg),
		Seq:              uint32(msg.Seq),
		Ex:               msg.Ex,
	}
	switch msg.SessionType {
	case constant.ReadGroupChatType:
		cbReq := &cbapi.CallbackAfterSendGroupMsgReq{
			CommonCallbackReq: common,
			GroupID:           msg.GroupID,
		}
		och.webhookClient.AsyncPost(ctx, cbReq.GetCallbackCommand(), cbReq, &cbapi.CallbackAfterSendGroupMsgResp{}, after)
	default:
		cbReq := &cbapi.CallbackAfterSendSingleMsgReq{
			CommonCallbackReq: common,
			RecvID:            msg.RecvID,
		}
		och.webhookClient.AsyncPost(ctx, cbReq.GetCallbackCommand(), cbReq, &cbapi.CallbackAfterSendSingleMsgResp{}, after)
	}
}

func persistCallbackContent(msg *sdkws.MsgData) string {
	if msg.ContentType >= constant.NotificationBegin && msg.ContentType <= constant.NotificationEnd {
		var tips sdkws.TipsComm
		_ = proto.Unmarshal(msg.Content, &tips)
		return tips.JsonDetail
	}
	return string(msg.Content)
}
