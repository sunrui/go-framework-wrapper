/*
 * Copyright (c) 2022 honeysense All rights reserved.
 * Author: sunrui
 * Date: 2022/01/31 20:31:31
 */

package enum

// OpenApprovalStatus 审核状态
type OpenApprovalStatus int

const (
	OpenApprovalWaiting = iota // 待审核
	OpenApprovalRefuse         // 审核拒绝
	OpenApprovalSuccess        // 审核成功
)
