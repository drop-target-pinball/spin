package spin

type Audio struct {
	// Unique identifier for this audio clip.
	ID string `hcl:"id,label" json:"id"`

	// ID of the AudioDevice responsible for this asset
	Device string `hcl:"device,optional" json:"device,omitempty"`

	// ID of the Module responsible for this asset
	Module string `hcl:"module,optional" json:"module,omitempty"`

	// Type of audio, either "music", "sound", or "voice"
	Type string `hcl:"type" json:"type"`

	// Path to the asset.
	File string `hcl:"file" json:"file"`
}

type AudioDevice struct {
	// Unique identifier for this audio device.
	ID string `hcl:"id,label" json:"id"`

	// Handler for this audio device, such as "sdl"
	Handler string `hcl:"handler" json:"handler"`
}

type Driver struct {
	// Unique identifier for this driver.
	ID string `hcl:"id,label" json:"id"`

	// Address used to reference this switch. This should match notation in
	// official manuals when possible.
	Address string `hcl:"address" json:"address"`

	// Type of driver such as "solenoid", "lamp", "flasher", "motor", or
	// "magnet".
	Type string `hcl:"type,optional" json:"type,omitempty"`
}

type Info struct {
	// Type is the what is being described which is an entity defined
	// by a top-level HCL block. Examples: "switch", "driver".
	Type string `hcl:"type,label" json:"type"`

	// Identifier of the entity being described.
	ID string `hcl:"id,label" json:"id"`

	// Name suitable for general, user-friendly display.
	Name string `hcl:"name,optional" json:"name,omitempty"`

	// Name to be used in limited space contexts such as in service
	// mode menus. If empty, the general name should be used instead.
	MenuName string `hcl:"menu_name,optional" json:"menu_name,omitempty"`

	// Name as described in the official manual if different from the general
	// name.
	ManualName string `hcl:"manual_name,optional" json:"manual_name,omitempty"`

	// Name suitable for use in a sorted context. For example, name might be
	// "Left Flipper" but the sort name is "Flipper, Left" to keep the
	// flipper entries together. If blank, the general name should be used.
	SortName string `hcl:"sort_name,optional" json:"sort_name,omitempty"`

	// Color of the wires used to connect this component.
	Wires []string `hcl:"wires,optional" json:"wires,omitempty"`

	// Jumpers where the wires are connected.
	Jumpers []string `hcl:"jumpers,optional" json:"jumpers,omitempty"`

	// Transistor used to power this component.
	Transistor string `hcl:"transistor,optional" json:"transistor,omitempty"`
}

type Switch struct {
	// Unique identifier for this switch.
	ID string `hcl:"id,label" json:"id"`

	// Address used to reference this switch. This should match notation in
	// official manuals when possible.
	Address string `hcl:"address" json:"address"`

	// Set to "opto" if this switch is normally closed when no ball is present.
	Type string `hcl:"type,optional" json:"type,omitempty"`
}
