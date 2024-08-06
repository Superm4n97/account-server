package account

type User struct {
	Email    string `json:"email"`
	Name     string `json:"name"`
	Password string `json:"password"`
	//Permissions is a map of map[ClusterID]Role
	Permissions map[string]string
}
