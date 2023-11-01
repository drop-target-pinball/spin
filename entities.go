package spin

type Driver struct {
	ID      string `hcl:"id,label" json:"id"`
	Address string `hcl:"address" json:"address"`
	Type    string `hcl:"type,optional" json:"type,omitempty"`
}

type Device struct {
	ID        string `hcl:"id,label" json:"id"`
	Namespace string `hcl:"namespace" json:"namespace"`
}

type Info struct {
	Type       string   `hcl:"type,label" json:"type"`
	ID         string   `hcl:"id,label" json:"id"`
	Name       string   `hcl:"name,optional" json:"name,omitempty"`
	MenuName   string   `hcl:"menu_name,optional" json:"menu_name,omitempty"`
	ManualName string   `hcl:"manual_name,optional" json:"manual_name,omitempty"`
	SortName   string   `hcl:"sort_name,optional" json:"sort_name,omitempty"`
	Wires      []string `hcl:"wires,optional" json:"wires,omitempty"`
	Jumpers    []string `hcl:"jumpers,optional" json:"jumpers,omitempty"`
	Transistor string   `hcl:"transistor,optional" json:"transistor,omitempty"`
}

type Switch struct {
	ID      string `hcl:"id,label" json:"id"`
	Address string `hcl:"address" json:"address"`
	Type    string `hcl:"type,optional" json:"type,omitempty"`
}
