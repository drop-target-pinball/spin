package service

import "github.com/drop-target-pinball/spin"

const (
	ScriptFontPreview = "service.ScriptFontPreview"
	ScriptTestFrame   = "service.ScriptTestFrame"
)

func RegisterScripts(eng *spin.Engine) {
	eng.Do(spin.RegisterScript{ID: ScriptFontPreview, Script: fontPreviewScript})
	eng.Do(spin.RegisterScript{ID: ScriptTestFrame, Script: testFrameScript})
}
