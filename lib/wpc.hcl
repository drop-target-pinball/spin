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
    address = "f3"
}

info "switch" "left_flipper_eos" {
    name        = "Left Flipper EOS"
    manual_name = "Left Flipper End of Stroke"
    sort_name   = "Flipper EOS, Left"
    wires       = [ "black-blue" ]
    jumpers     = [ "J906-3" ]
}

switch "left_flipper_button" {
    address = "f4"
    type    = "opto"
}

info "switch" "left_flipper_button" {
    name        = "Left Flipper Button"
    manual_name = "Left Flipper Opto"
    sort_name   = "Flipper Button, Left"
    wires       = [ "blue-gray" ]
    jumpers     = [ "J905-2" ]
}

switch "upper_right_flipper_eos" {
    address = "f5"
}

switch "upper_right_flipper_button" {
    address = "f6"
    type    = "opto"
}

switch "upper_left_flipper_eos" {
    address = "f7"
}

switch "upper_left_flipper_button" {
    address = "f8"
    type    = "opto"
}

// Flipper solenoids

driver "left_flipper_power" {
    address = "f1p"
    type    = "solenoid"
}

driver "left_flipper_hold" {
    address = "f1h"
    type    = "solenoid"
}

driver "right_flipper_power" {
    address = "f2p"
    type    = "solenoid"
}

driver "right_flipper_hold" {
    address = "f2h"
    type    = "solenoid"
}

driver "upper_left_flipper_power" {
    address = "f3p"
    type    = "solenoid"
}

driver "upper_left_flipper_hold" {
    address = "f3h"
    type    = "solenoid"
}

driver "upper_right_flipper_power" {
    address = "f4p"
    type    = "solenoid"
}

driver "upper_right_flipper_hold" {
    address = "f4h"
    type    = "solenoid"
}


