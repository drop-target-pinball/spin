component "switch" "right_flipper_eos" {
    address         = "f1"
    name            = "Right Flipper End of Stroke"
    service_name    = "R. Flipper E.O.S."
    sort_name       = "Flipper EOS, Right"
    groups          = [ "flipper_eos" ]
    index           = 1 
}

component "switch" "right_flipper_button" {
    address         = "f2" 
    name            = "Right Flipper Button"
    service_name    = "R. Flipper Button"
    sort_name       = "Flipper Button, Right" 
    groups          = [ "flipper", "opto" ]
    index           = 1
}

component "switch" "left_flipper_eos" {
    address         = "f3"
    name            = "Left Flipper End of Stroke"
    service_name    = "L. Flipper E.O.S."
    sort_name       = "Flipper EOS, Left"
    groups          = [ "flipper_eos" ]
    index           = 0 
}

component "switch" "left_flipper_button" {
    address         = "f4" 
    name            = "Left Flipper Button"
    service_name    = "L. Flipper Button"
    sort_name       = "Flipper Button, Left" 
    groups          = [ "flipper", "opto" ]
    index           = 0
}

component "switch" "upper_right_flipper_eos" {
    address         = "f5"
    name            = "Upper Right Flipper End of Stroke"
    service_name    = "U.R. Flipper E.O.S."
    sort_name       = "Flipper EOS, Upper Right"
    groups          = [ "flipper_eos" ]
    index           = 3
}

component "switch" "upper_right_flipper_button" {
    address         = "f6" 
    name            = "Upper Right Flipper Button"
    service_name    = "U.R. Flipper Button"
    sort_name       = "Flipper Button, Upper Right" 
    groups          = [ "flipper", "opto" ]
    index           = 3
}

component "switch" "upper_left_flipper_eos" {
    address         = "f7"
    name            = "Upper Left Flipper End of Stroke"
    service_name    = "U.L. Flipper E.O.S."
    sort_name       = "Flipper EOS, Upper Left"
    groups          = [ "flipper_eos" ]
    index           = 2
}

component "switch" "upper_left_flipper_button" {
    address         = "f8" 
    name            = "Upper Left Flipper Button"
    service_name    = "U.L. Flipper Button"
    sort_name       = "Flipper Button, Upper Left" 
    groups          = [ "flipper", "opto" ]
    index           = 2
}

component "driver" "left_flipper_power" {
    address         = "lp"
    name            = "Left Flipper Power"
    service_name    = "L. Flip. Power"
    sort_name       = "Flipper Power, Left"
    groups          = [ "solenoid", "flipper", "power" ]
    index           = 0
}

component "driver" "left_flipper_hold" {
    address         = "lh"
    name            = "Left Flipper Hold"
    service_name    = "L. Flip. Hold"
    sort_name       = "Flipper Hold, Left"
    groups          = [ "solenoid", "flipper", "hold" ]
    index           = 0 
}

component "driver" "right_flipper_power" {
    address         = "rp"
    name            = "Right Flipper Power" 
    service_name    = "R. Flip. Power"
    sort_name       = "Flipper Power, Right"
    groups          = [ "solenoid", "flipper", "power" ]
    index           = 1
}

component "driver" "right_flipper_hold" {
    address         = "rh"
    name            = "Right Flipper Hold"
    service_name    = "R. Flip. Hold"
    sort_name       = "Flipper Hold, Right"
    groups          = [ "solenoid", "flipper", "hold" ]
    index           = 1
}

component "driver" "upper_left_flipper_power" {
    address         = "ulp"
    name            = "Upper Left Flipper Power" 
    service_name    = "U.L. Flip. Power"
    sort_name       = "Flipper Power, Left"
    groups          = [ "solenoid", "flipper", "power" ]
    index           = 2
}

component "driver" "upper_left_flipper_hold" {
    address         = "ulh"
    name            = "Upper Left Flipper Hold"
    service_name    = "U.L. Flip. Hold"
    sort_name       = "Flipper Hold, Upper Left"
    groups          = [ "solenoid", "flipper", "hold" ]
    index           = 2    
}

component "driver" "upper_right_flipper_power" {
    address         = "urp"
    name            = "Upper Right Flipper Power" 
    service_name    = "U.R. Flip. Power"
    sort_name       = "Flipper Power, Upper Right"
    groups          = [ "solenoid", "flipper", "power" ]
    index           = 3
}

component "driver" "upper_right_flipper_hold" {
    address         = "urh"
    name            = "Upper Right Flipper Hold"
    service_name    = "U.R. Flip. Hold"
    sort_name       = "Flipper Hold, Upper Right"
    groups          = [ "solenoid", "flipper", "hold" ]
    index           = 3
}

