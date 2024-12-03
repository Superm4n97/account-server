package permission

type PermissionRef struct {
	UserId    string `json:"userid"`
	Role      string `json:"role"`
	ClusterID string `json:"clusterid"`
}
