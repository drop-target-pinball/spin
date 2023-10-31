package spin

type Driver struct {
	ID      string `hcl:"id,label"`
	Address string `hcl:"address"`
	Type    string `hcl:"type,optional"`
}

type Device struct {
	ID        string `hcl:"id,label"`
	Namespace string `hcl:"namespace"`
}

type Info struct {
	Type       string   `hcl:"type,label"`
	ID         string   `hcl:"id,label"`
	Name       string   `hcl:"name,optional"`
	MenuName   string   `hcl:"menu_name,optional"`
	ManualName string   `hcl:"manual_name,optional"`
	SortName   string   `hcl:"sort_name,optional"`
	Wires      []string `hcl:"wires,optional"`
	Jumpers    []string `hcl:"jumpers,optional"`
}

type Switch struct {
	ID      string `hcl:"id,label"`
	Address string `hcl:"address"`
	Type    string `hcl:"type,optional"`
}
