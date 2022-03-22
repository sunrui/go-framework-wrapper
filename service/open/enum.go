/*
 * Copyright (c) 2022 honeysense All rights reserved.
 * Author: sunrui
 * Date: 2022/01/31 20:36:31
 */

package open

// 审核状态
type ApprovalStatus int

const (
	ApprovalWaiting = iota // 待审核
	ApprovalRefuse         // 审核拒绝
	ApprovalSuccess        // 审核成功
)
