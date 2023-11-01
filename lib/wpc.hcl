// Williams Pinball Controller 

// Dedicated Switches

switch "left_coin_slot" {
    address     = "d1"
}

info "switch" "left_coin_slot" {
    name        = "Left Coin Slot"
    manual_name = "Left Coin Chute"
    sort_name   = "Coin Slot, Left"
    wires       = [ "orange-brown" ]
    jumpers     = [ "J205-1" ]
}

switch "center_coin_slot" {
    address     = "d2"
}

info "switch" "center_coin_slot" {
    name        = "Center Coin Slot"
    manual_name = "Center Coin Chute"
    sort_name   = "Coin Slot, Center"
    wires       = [ "orange-red" ]
    jumpers     = [ "J205-2" ]
}

switch "right_coin_slot" {
    address     = "d3"
}

info "switch" "right_coin_slot" {
    name        = "Right Coin Slot"
    manual_name = "Right Coin Chute"
    sort_name   = "Coin Slot, Right"
    wires       = [ "orange-black" ]
    jumpers     = [ "J205-3" ]
}

switch "other_coin_slot" {
    address     = "d4"
}

info "switch" "other_coin_slot" {
    name        = "Other Coin Slot"
    manual_name = "4th Coin Chute"
    sort_name   = "Coin Slot, Other"
    wires       = [ "orange-yellow" ]
    jumpers     = [ "J205-4" ]
}

switch "escape_button" {
    address     = "d5"
}

info "switch" "escape_button" {
    name        = "Escape Button"
    manual_name = "Service Credits / Escape"
    sort_name   = "Service Button, Escape"
    wires       = [ "orange-green" ]
    jumpers     = [ "J205-6" ]
}

switch "down_button" {
    address     = "d6"
}

info "switch" "down_button" {
    name        = "Down Button"
    manual_name = "Volume Down / Down"
    sort_name   = "Service Button, Down"
    wires       = [ "orange-blue" ]
    jumpers     = [ "J205-7" ]
}

switch "up_button" {
    address     = "d7"
}

info "switch" "up_button" {
    name        = "Up Button"
    manual_name = "Volume Up / Up"
    sort_name   = "Service Button, Up"
    wires       = [ "orange-violet" ]
    jumpers     = [ "J205-8" ]
}

switch "enter_button" {
    address = "d8"
}

info "switch" "enter_button" {
    name        = "Enter Button"
    manual_name = "Begin Test / Enter"
    sort_name   = "Service Button, Enter"
    wires       = [ "orange-gray" ]
    jumpers     = [ "J205-9" ]
}

// Flipper grounded switches

switch "right_flipper_eos" {
    address     = "f1"
}

info "switch" "right_flipper_eos" {
    name        = "Right Flipper EOS"
    manual_name = "Right Flipper End of Stroke"
    sort_name   = "Flipper EOS, Right"
    wires       = [ "black-green" ]
    jumpers     = [ "J906-1" ]
}

switch "right_flipper_button" {
    address     = "f2"
    type        = "opto"
}

info "switch" "right_flipper_button" {
    name        = "Right Flipper Button"
    manual_name = "Right Flipper Opto"
    sort_name   = "Flipper Button, Right"
    wires       = [ "blue-violet" ]
    jumpers     = [ "J905-1" ]
}

switch "left_flipper_eos" {
    address     = "f3"
}

info "switch" "left_flipper_eos" {
    name        = "Left Flipper EOS"
    manual_name = "Left Flipper End of Stroke"
    sort_name   = "Flipper EOS, Left"
    wires       = [ "black-blue" ]
    jumpers     = [ "J906-3" ]
}

switch "left_flipper_button" {
    address     = "f4"
    type        = "opto"
}

info "switch" "left_flipper_button" {
    name        = "Left Flipper Button"
    manual_name = "Left Flipper Opto"
    sort_name   = "Flipper Button, Left"
    wires       = [ "blue-gray" ]
    jumpers     = [ "J905-2" ]
}

switch "upper_right_flipper_eos" {
    address     = "f5"
}

info "switch" "upper_right_flipper_eos" {
    name        = "Upper Right Flipper EOS"
    manual_name = "Upper Right Flipper End of Stroke"
    sort_name   = "Flipper EOS, Upper Right"
    wires       = [ "black-violet" ]
    jumpers     = [ "J906-4" ] 
}

switch "upper_right_flipper_button" {
    address     = "f6"
    type        = "opto"
}

info "switch" "upper_right_flipper_button" {
    name        = "Upper Right Flipper Button"
    manual_name = "Upper Right Flipper Opto"
    sort_name   = "Flipper Button, Upper Right"
    wires       = [ "black-yellow" ]
    jumpers     = [ "J905-3" ]
}

switch "upper_left_flipper_eos" {
    address = "f7"
}

info "switch" "upper_left_flipper_eos" {
    name        = "Upper Left Flipper EOS"
    manual_name = "Upper Left Flipper End of Switch"
    sort_name   = "Flipper EOS, Upper Left"
    wires       = [ "black-gray" ]
    jumpers     = [ "J906-5" ]
}

switch "upper_left_flipper_button" {
    address     = "f8"
    type        = "opto"
}

info "switch" "upper_left_flipper_button" {
    name        = "Upper Left Flipper Button"
    manual_name = "Upper Left Flipper Opto"
    sort_name   = "FLipper Button, Upper Left"
    wires       = [ "black-blue" ]
    jumpers     = [ "J905-5" ]
}

// Flipper solenoids

driver "left_flipper_power" {
    address     = "fp1"
    type        = "solenoid"
}

info "driver" "left_flipper_power" {
    name        = "Left Flipper Power"
    manual_name = "Lower Left Flipper Power"
    sort_name   = "Flipper Power, Left"
    wires       = [ "gray-yellow", "blue-gray" ]
    jumpers     = [ "J907-7", "J902-9" ]
    transistor  = "Q3"
}

driver "left_flipper_hold" {
    address     = "fh1"
    type        = "solenoid"
}

info "driver" "left_flipper_hold" { 
    name        = "Left Flipper Hold"
    manual_name = "Lower Left Flipper Hold"
    sort_name   = "Flipper Hold, Left"
    wires       = [ "gray-yellow", "orange-blue" ]
    jumpers     = [ "J907-7", "J902-7" ]
    transistor  = "Q9"
}

driver "right_flipper_power" {
    address     = "fp2"
    type        = "solenoid"
}

info "driver" "right_flipper_power" {
    name        = "Right Flipper Power"
    manual_name = "Lower Right Flipper Power"
    sort_name   = "Flipper Power, Right"
    wires       = [ "blue-yellow", "blue-violet" ]
    jumpers     = [ "J907-9", "J902-13" ]
    transistor  = "Q4"
}

driver "right_flipper_hold" {
    address     = "fh2"
    type        = "solenoid"
}

info "driver" "right_flipper_hold" {
    name        = "Right Flipper Hold"
    manual_name = "Lower Right Flipper Hold"
    sort_name   = "Flipper Hold, Right"
    wires       = [ "blue-yellow", "orange-green" ]
    jumpers     = [ "J907-9", "J902-11" ]
    transistor = "Q11"
}

driver "upper_left_flipper_power" {
    address     = "fp3"
    type        = "solenoid"
}

info "driver" "upper_left_flipper_power" {
    name        = "Upper Right Flipper Power"
    sort_name   = "Flipper Power, Upper Right"
    wires       = [ "gray-yellow", "black-blue" ] 
    jumpers     = [ "J907-1", "J902-3" ] 
    transistor  = "Q1"
}

driver "upper_left_flipper_hold" {
    address     = "fh3"
    type        = "solenoid"
}

info "driver" "upper_left_flipper_hold" {
    name        = "Upper Left Flipper Hold"
    sort_name   = "Flipper Hold, Upper Left"
    wires       = [ "gray-yellow", "orange-gray" ]
    jumpers     = [ "J907-1", "J902-1" ]
    transistor = "Q5"
}

driver "upper_right_flipper_power" {
    address     = "fp4"
    type        = "solenoid"
}

info "driver" "upper_right_flipper_power" {
    name        = "Upper Right Flipper Power"
    sort_name   = "Flipper Power, Upper Right"
    wires       = [ "blue-yellow", "black-yellow" ]
    jumpers     = [ "J907-4", "J902-6" ]
    transistor  = "Q2"
}

driver "upper_right_flipper_hold" {
    address     = "fh4"
    type        = "solenoid"
}

info "driver" "upper_right_flipper_hold" {
    name        = "Upper Right Flipper Hold"
    sort_name   = "Flipper Hold, Upper Right"
    wires       = [ "blue-yellow", "orange-violet" ]
    jumpers     = [ "J907-4", "J902-4" ]
    transistor  = "Q7"
}

// Switch Matrix 

info "switch_matrix" "row_1" {
    wires       = [ "white-brown" ]
    jumpers     = [ "J209-1" ]
}

info "switch_matrix" "row_2" {
    wires       = [ "white-red" ]
    jumpers     = [ "J209-2" ]
}

info "switch_matrix" "row_3" {
    wires       = [ "white-orange" ]
    jumpers     = [ "J209-3" ]
}

info "switch_matrix" "row_4" {
    wires       = [ "white-yellow" ]
    jumpers     = [ "J209-4" ]
}

info "switch_matrix" "row_5" {
    wires       = [ "white-green" ]
    jumpers     = [ "J209-5" ]
}

info "switch_matrix" "row_6" {
    wires       = [ "white-blue" ]
    jumpers     = [ "J209-7" ]
}

info "switch_matrix" "row_7" {
    wires       = [ "white-violet" ]
    jumpers     = [ "J209-8" ]
}

info "switch_matrix" "row_8" {
    wires       = [ "white-gray" ]
    jumpers     = [ "J209-9" ]
}

info "switch_matrix" "column_1" {
    wires       = [ "green-brown" ]
    jumpers     = [ "J207-1" ]
}

info "switch_matrix" "column_2" {
    wires       = [ "green-red" ]
    jumpers     = [ "J207-2" ]
}

info "switch_matrix" "column_3" {
    wires       = [ "green-orange" ]
    jumpers     = [ "J207-3" ]
}

info "switch_matrix" "column_4" {
    wires       = [ "green-yellow" ]
    jumpers     = [ "J207-4" ]
}

info "switch_matrix" "column_5" {
    wires       = [ "green-black" ]
    jumpers     = [ "J207-5" ]
}

info "switch_matrix" "column_6" {
    wires       = [ "green-blue" ]
    jumpers     = [ "J207-6" ]
}

info "switch_matrix" "column_7" {
    wires       = [ "green-violet" ]
    jumpers     = [ "J207-7" ]
}

info "switch_matrix" "column_8" {
    wires       = [ "green-gray" ]
    jumpers     = [ "J207-9" ]
}

// Lamp Matrix 

info "lamp_matrix" "row_1" {
    wires       = [ "red-brown" ]
    jumpers     = [ "J133-1" ]
    transistor  = "Q90"
}

info "lamp_matrix" "row_2" {
    wires       = [ "red-black" ]
    jumpers     = [ "J133-2" ]
    transistor  = "Q89"
}

info "lamp_matrix" "row_3" {
    wires       = [ "red-orange" ]
    jumpers     = [ "J133-4" ]
    transistor  = "Q88"
}

info "lamp_matrix" "row_4" {
    wires       = [ "red-yellow" ]
    jumpers     = [ "J133-5" ]
    transistor  = "Q87"
}

info "lamp_matrix" "row_5" {
    wires       = [ "red-green" ]
    jumpers     = [ "J133-6" ]
    transistor  = "Q86"
}

info "lamp_matrix" "row_6" {
    wires       = [ "red-blue" ]
    jumpers     = [ "J133-7" ]
    transistor  = "Q85"
}

info "lamp_matrix" "row_7" {
    wires       = [ "red-violet" ]
    jumpers     = [ "J133-8" ]
    transistor  = "Q84"
}

info "lamp_matrix" "row_8" {
    wires       = [ "red-gray" ]
    jumpers     = [ "J133-9" ]
    transistor  = "Q83"
}

info "lamp_matrix" "column_1" {
    wires       = [ "yellow-brown"]
    jumpers     = [ "J137-1" ]
    transistor  = "Q98"
}

info "lamp_matrix" "column_2" {
    wires       = [ "yellow-red"]
    jumpers     = [ "J137-2" ]
    transistor  = "Q97"
}

info "lamp_matrix" "column_3" {
    wires       = [ "yellow-orange"]
    jumpers     = [ "J137-3" ]
    transistor  = "Q96"
}

info "lamp_matrix" "column_4" {
    wires       = [ "yellow-black"]
    jumpers     = [ "J137-4" ]
    transistor  = "Q95"
}

info "lamp_matrix" "column_5" {
    wires       = [ "yellow-green"]
    jumpers     = [ "J137-5" ]
    transistor  = "Q94"
}

info "lamp_matrix" "column_6" {
    wires       = [ "yellow-blue"]
    jumpers     = [ "J137-6" ]
    transistor  = "Q93"
}

info "lamp_matrix" "column_7" {
    wires       = [ "yellow-violet"]
    jumpers     = [ "J137-7" ]
    transistor  = "Q92"
}

info "lamp_matrix" "column_8" {
    wires       = [ "yellow-gray"]
    jumpers     = [ "J137-9" ]
    transistor  = "Q91"
}
