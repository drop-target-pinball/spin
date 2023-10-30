package spin

type Config struct {
	Component []Component `hcl:"component,block"`
}

type Component struct {
	Type        string   `hcl:"type,label"`
	ID          string   `hcl:"id,label"`
	Address     string   `hcl:"address"`
	Name        string   `hcl:"name,optional"`
	ServiceName string   `hcl:"service_name,optional"`
	SortName    string   `hcl:"sort_name,optional"`
	Groups      []string `hcl:"groups,optional"`
	Index       int      `hcl:"index,optional"`
}
