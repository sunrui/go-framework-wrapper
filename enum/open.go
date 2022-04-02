/*
 * Copyright (c) 2022 honeysense.com All rights reserved.
 * Author: sunrui
 * Date: 2022/04/03 00:18:03
 */

package enum

// OpenApprovalStatus 审核状态
type OpenApprovalStatus string

const (
	OpenApprovalWaiting = "OpenApprovalWaiting" // 待审核
	OpenApprovalRefuse  = "OpenApprovalRefuse"  // 审核拒绝
	OpenApprovalSuccess = "OpenApprovalSuccess" // 审核成功
)
