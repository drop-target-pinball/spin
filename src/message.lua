local pub = {}

pub.HALT = "halt"
pub.INIT = "init"
pub.NOTE = "note"
pub.MUSIC_ENDED = "music_ended"
pub.PLAY_MUSIC = "play_music"
pub.PLAY_SOUND = "play_sound"
pub.PLAY_VOCAL = "play_vocal"
pub.SCRIPT_ENDED = "script_ended"
pub.SILENCE = "silence"
pub.SHUTDOWN = "shutdown"
pub.STOP_MUSIC = "stop_music"
pub.STOP_VOCAL = "stop_vocal"
pub.RUN = "run"
pub.TICK = "tick"
pub.VOCAL_ENDED = "vocal_ended"

package.loaded['message'] = pub

return pub

