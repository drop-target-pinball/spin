// Judge Dredd 
// Bally 1993 
//
// https://ipdb.org/machine.cgi?id=1322
// https://www.ipdb.org/files/1322/Bally_1993_Judge_Dredd_Manual.pdf

include = [ "wpc.hcl" ]

device "judge_dredd" {
    namespace   = "jd"
}

// Switch Matrix

switch "left_fire_button" {
    address     = "11"
}

info "switch" "left_fire_button" {
    name        = "Left Fire Button"
    sort_name   = "Fire Button, Left"
}

switch "right_fire_button" {
    address     = "12"
}

info "switch" "right_fire_button" {
    name        = "Right Fire Button"
    sort_name   = "Fire Button, Right"
}

switch "start_button" {
    address     = "13"
}

info "switch" "start_button" {
    name        = "Start Button"
    manual_name = "Credit (Start)"
}

switch "tilt" {
    address     = "14"
}

info "switch" "tilt" {
    name        = "Tilt"
    manual_name = "Plumb Bob Tilt"
}

switch "left_shooter_lane" {
    address     = "15"
}

info "left_shooter_lane" {
    name        = "Left Shooter Lane"
    manual_name = "Left Shoot Lane"
    sort_name   = "Shooter Lane, Left"
}

switch "left_outlane" {
    address     = "!6"
} 

info "switch" "left_outlane" {
    name        = "Left Outlane"
    sort_name   = "Outlane, Left"
}

switch "left_inlane" {
    address     = "!7"
}

info "switch" "left_return_lane" {
    name        = "Left Return Lane"
    sort_name   = "Return Lane, Left"
}

switch "bank_targets" {
    address     = "18"
}

info "switch" "bank_targets" {
    name        = "Bank Targets"
    manual_name = "3-Bank Targets"
}

switch "slam_tilt" {
    address     = "21"
}

info "switch" "slam_tilt" {
    name        = "Slam Tilt" 
    sort_name   = "Tilt, Slam"
}

switch "door_closed" {
    address     = "22"
}

info "switch" "door_closed" {
    name        = "Door Closed"
    manual_name = "Front Door Closed"
}

switch "ticket_dispenser" {
    address     = "23"
}

info "switch" "ticket_dispenser" {
    name        = "Ticket Dispenser"
}

switch "24" {
    address     = "24"
}

info "switch" "24" {
    name        = "Always Closed"
}

switch "top_right_post" {
    address     = "25"
}

info "switch" "top_right_post" {
    name        = "Top Right Post"
    sort_name   = "Post, Top Right"
}

switch "captive_ball_1" {
    address     = "26"
}

info "switch" "captive_ball_1" {
    name        = "Captive Ball 1"
}

switch "mystery" {
    address      = "27"
}

info "switch" "mystery" {
    name        = "Mystery"
}

switch "28" {
    address     = "28"
}

switch "extra_ball_button" {
    address     = "31"
}

info "switch" "extra_ball_Button" {
    name        = "Extra Ball Button"
    manual_name = "Buy-In (Extra Ball)"
}

switch "32" {
    address     = "32"
}

switch "left_rollover" {
    address     = "33"
}

info "switch" "outer_left_loop" {
    name        = "Outer Left Loop"
    manual_name = "Left Rollover"
    sort_name   = "Loop, Outer Left" 
}

switch "inner_right_return_lane {
    address     = "34"
}

info "switch" "inner_right_return_lane" {
    name        = "Inner Right Return Lane"
    manual_name = "Inside Right Return"
    sort_name   = "Return Lane, Inner Right"
}

switch "inner_loop" {
    address     = "35"
}

info "switch" "inner_loop" {
    name        = "Inner Loop"
    manual_name = "Top Center Rollover"
    sort_name   = "Loop, Inner"
}

switch "left_post" {
    address     = "36" 
}

info "switch" "left_post" {
    name        = "Left Post"
    manual_name = "Left Score Post"
    sort_name   = "Post, Left"
}

switch "subway_enter_1" {
    address     = "37"
}

info "switch" "subway_enter_1" {
    name        = "Subway Enter 1"
}

switch "subway_enter_2" {
    address     = "38"
}

info "switch" "subway_enter_2" {
    name        = "Subway Enter 2"
}

switch "right_shooter_lane" {
    address     = "41"
}

info "switch" "right_shooter_lane" {
    name        = "Right Shooter Lane"
    manual_name = "Right Ball Shooter"
    sort_name   = "Shooter Lane, Right"
}

switch "right_outlane" {
    address     = "42"
}

info "switch" "right_outlane" {
    name        = "Right Outlane"
    sort_name   = "Outlane, Right"
}

switch "outer_right_return_lane" {
    address     = "43"
}

info "switch" "outer_right_return_lane" {
    name        = "Outer Right Return Lane"
    manual_name = "Outside Right Return"
}







