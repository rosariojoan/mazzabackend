package firebase

var AlertType = struct {
	AssignedWorkTask         string
	InvitedUserRegistration  string
	UserActiveStatusChange   string
	UserRoleChange           string
	WorkShiftApprovalRequest string
	WorkShifUpdateRequest    string
}{
	AssignedWorkTask:         "assignedWorkTask",
	InvitedUserRegistration:  "invitedUserRegistration",
	UserActiveStatusChange:   "userActiveStatusChange",
	UserRoleChange:           "userRoleChange",
	WorkShiftApprovalRequest: "workShiftApprovalRequest",
	WorkShifUpdateRequest:    "workShifUpdateRequest",
}
