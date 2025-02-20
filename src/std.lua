local pub = {}

-- Variables
pub.BALL = "ball"
pub.CREDITS = "credits"
pub.EXTRA_BALLS = "extra_balls"
pub.EXTRA_BALLS_AWARDED = "extra_balls_awarded"
pub.FREE_PLAY = "free_play"
pub.GAME_ACTIVE = "game_active"
pub.IS_EXTRA_BALL = "is_extra_ball"
pub.MAX_PLAYERS = "max_players"
pub.PLAYER = "player"
pub.PLAYER_1 = "player_1"
pub.PLAYER_2 = "player_2"
pub.PLAYER_3 = "player_3"
pub.PLAYER_4 = "player_4"
pub.PLAYER_COUNT = "player_count"
pub.REPLAY_AWARDED = "replay_awarded"
pub.SCORE = "score"

-- Scripts
pub.ADD_PLAYER = "add_player"
pub.DMD_GRADIENT = "dmd_gradient"
pub.SCORE_DRAW = "score_draw"
pub.START_GAME = "start_game"
pub.START_SERVICE = "start_service"

-- General components
pub.LEFT_FLIPPER_BUTTON = "left_flipper_button"
pub.RIGHT_FLIPPER_BUTTON = "right_flipper_button"
pub.START_BUTTON = "start_button"
pub.UPPER_LEFT_FLIPPER_BUTTON = "upper_left_flipper_button"
pub.UPPER_RIGHT_FLIPPER_BUTTON = "upper_right_flipper_button"

-- Messages
pub.BALL_ARRIVED = "ball_arrived"
pub.BALL_DEPARTED = "ball_departed"
pub.HALT = "halt"
pub.INIT = "init"
pub.NOTE = "note"
pub.MUSIC_ENDED = "music_ended"
pub.PLAY_MUSIC = "play_music"
pub.PLAY_SOUND = "play_sound"
pub.PLAY_VOCAL = "play_vocal"
pub.REJECTED = "rejected"
pub.SCRIPT_ENDED = "script_ended"
pub.SILENCE = "silence"
pub.SHUTDOWN = "shutdown"
pub.STOP_MUSIC = "stop_music"
pub.STOP_VOCAL = "stop_vocal"
pub.SWITCH_UPDATED = "switch_updated"
pub.RUN = "run"
pub.TICK = "tick"
pub.TIMER_EXPIRED = "timer_expired"
pub.VOCAL_ENDED = "vocal_ended"
pub.WAKE = "wake"

-- Fonts
pub.DMD_04B_03_7PX = "dmd_04b_03_7px"
pub.DMD_18X10 = "dmd_18x10"
pub.DMD_18X11 = "dmd_18x11"
pub.DMD_18X12 = "dmd_18x12"
pub.DMD_14X10 = "dmd_14x10"
pub.DMD_14X9 = "dmd_14x9"
pub.DMD_14X8 = "dmd_14x8"
pub.DMD_09X7 = "dmd_09x7"
pub.DMD_09X6 = "dmd_09x6"
pub.DMD_09X7 = "dmd_09x5"

-- Devices
pub.DMD = "dmd"

-- Rejections
pub.CREDITS_REQUIRED = "credits_required"
pub.GAME_FULL = "game_full"

package.loaded["std"] = pub

return pub
