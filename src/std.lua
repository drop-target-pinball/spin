local pub = {}

pub.ADD_PLAYER = "add_player"
pub.CREDITS = "credits"
pub.CREDITS_REQUIRED = "credits_required"
pub.FREE_PLAY = "free_play"
pub.GAME_ACTIVE = "game_active"
pub.GAME_FULL = "game_full"
pub.MAX_PLAYERS = "max_players"
pub.PLAYER = "player"
pub.PLAYER_COUNT = "player_count"
pub.START_BUTTON = "start_button"
pub.START_SERVICE = "start_service"
pub.START_GAME = "start_game"

package.loaded["std"] = pub

return pub
