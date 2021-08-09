package audit

const (
	SourceSelly = "selly"
)

// List actions
const (
	ActionCreate            = "create"
	ActionUpdate            = "update"
	ActionUpdatePermissions = "update-permissions"
)

// List targets
const (
	TargetSellyStaff     = "staffs"
	TargetSellyStaffRole = "staff-roles"
)

// SellyTargets ...
var SellyTargets = []string{
	TargetSellyStaff,
	TargetSellyStaffRole,
}
