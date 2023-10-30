component "switch" "left_coin_slot" {
    address         = "d1"
    name            = "Left Coin Slot"
    sort_name       = "Coin Slot, Left"
    groups          = [ "coin" ]
    index           = 0 
}

component "switch" "center_coin_slot" {
    address         = "d2" 
    name            = "Center Coin Slot"
    sort_name       = "Coin Slot, Center"
    groups          = [ "coin" ]
    index           = 1
}

component "switch" "right_coin_slot" {
    address         = "d3" 
    name            = "Right Coin Slot"
    sort_name       = "Coin Slot, Right"
    groups          = [ "coin" ]
    index           = 2
}

component "switch" "other_coin_slot" {
    address         = "d4"
    name            = "Other Coin Slot"
    service_name    = "4th Coin Option"
    sort_name       = "Coin Slot, 4th"
    groups          = [ "coin" ]
    index           = 3 
}

component "switch" "escape_button" {
    address         = "d5"
    name            = "Escape Button"
    service_name    = "Escape"
    sort_name       = "Service Button, Escape"
    groups          = [ "service" ]
    index           = 0 
}

component "switch" "down_button" {
    address         = "d6"
    name            = "Down Button"
    service_name    = "Down"
    sort_name       = "Service Button, Down"
    groups          = [ "service" ]
    index           = 1
}

component "switch" "up_button" {
    address         = "d7"
    name            = "Up Button"
    service_name    = "Up"
    sort_name       = "Service Button, Up" 
    groups          = [ "service" ]
    index           = 2
}

component "switch" "enter_button" {
    address         = "d8"
    name            = "Enter Button"
    service_name    = "Enter" 
    sort_name       = "Service BUtton, Enter"
    groups          = [ "service" ]
    index           = 3
}

