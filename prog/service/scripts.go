package service

import "github.com/drop-target-pinball/spin"

const (
	ScriptFontPreview = "service.ScriptFontPreview"
)

func RegisterScripts(eng *spin.Engine) {
	eng.Do(spin.RegisterScript{ID: ScriptFontPreview, Script: fontPreviewScript})
}