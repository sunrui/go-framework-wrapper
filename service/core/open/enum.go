/*
 * Copyright (c) 2022 honeysense.com All rights reserved.
 * Author: sunrui
 * Date: 2022-04-25 22:23:18
 */

package open

// ApprovalStatus 审核状态
type ApprovalStatus string

const (
	ApprovalStatusWaiting ApprovalStatus = "ApprovalStatusWaiting" // 待审核
	ApprovalStatusRefuse  ApprovalStatus = "ApprovalStatusRefuse"  // 审核拒绝
	ApprovalStatusSuccess ApprovalStatus = "ApprovalStatusSuccess" // 审核成功
)
