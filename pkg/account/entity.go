package account

const (
	roleContributor = "Contributor"
	roleAdmin       = "Admin"
	roleOwner       = "Owner"
)

func init() {
	role = []string{
		roleContributor,
		roleAdmin,
		roleOwner,
	}
}

var role []string

// getRole returns the roles that the requester can have,
// if the requester role is empty, it will return all roles and if the requester
// role is invalid it will return nil
func getRole(requesterRole string) []string {
	switch requesterRole {
	case roleContributor:
		return []string{roleContributor}
	case roleAdmin:
		return []string{roleContributor, roleAdmin}
	case roleOwner:
		return []string{roleContributor, roleAdmin, roleOwner}
	case "":
		return []string{roleContributor, roleAdmin, roleOwner}
	default:
		return nil
	}
}

type permission struct {
	ClusterID string `json:"clusterID"`
	Role      string `json:"role"`
}
