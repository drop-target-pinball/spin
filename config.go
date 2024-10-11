package spin

type ConfigFile struct {
	ItemC []ItemC `hcl:"item_c,block"`
	ItemD []ItemD `hcl:"item_d,block"`
}

type Base struct {
	A string `hcl:"a,optional"`
	B string `hcl:"b,optional"`
}

type ItemC struct {
	Base `hcl:",remain"`
	C    string `hcl:"c,optional"`
}

type ItemD struct {
	Base `hcl:",remain"`
	D    string `hcl:"d,optional"`
}
