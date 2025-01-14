package theme

import (
	"encoding/hex"

	"github.com/fatih/color"
	"github.com/jesseduffield/gocui"
	"github.com/jesseduffield/lazygit/pkg/config"
)

var (
	// DefaultTextColor is the default text color
	DefaultTextColor = color.FgWhite
	// DefaultHiTextColor is the default highlighted text color
	DefaultHiTextColor = color.FgHiWhite

	// GocuiDefaultTextColor does the same as DefaultTextColor but this one only colors gocui default text colors
	GocuiDefaultTextColor gocui.Attribute

	// ActiveBorderColor is the border color of the active frame
	ActiveBorderColor gocui.Attribute

	// InactiveBorderColor is the border color of the inactive active frames
	InactiveBorderColor gocui.Attribute

	// SelectedLineBgColor is the background color for the selected line
	SelectedLineBgColor color.Attribute

	// SelectedRangeBgColor is the background color of the selected range of lines
	SelectedRangeBgColor color.Attribute

	// GocuiSelectedLineBgColor is the background color for the selected line in gocui
	GocuiSelectedLineBgColor gocui.Attribute

	OptionsFgColor color.Attribute

	OptionsColor gocui.Attribute

	DiffTerminalColor = color.FgMagenta
)

// UpdateTheme updates all theme variables
func UpdateTheme(themeConfig config.ThemeConfig) {
	ActiveBorderColor = GetGocuiColor(themeConfig.ActiveBorderColor)
	InactiveBorderColor = GetGocuiColor(themeConfig.InactiveBorderColor)
	SelectedLineBgColor = GetBgColor(themeConfig.SelectedLineBgColor)
	SelectedRangeBgColor = GetBgColor(themeConfig.SelectedRangeBgColor)
	GocuiSelectedLineBgColor = GetGocuiColor(themeConfig.SelectedLineBgColor)
	OptionsColor = GetGocuiColor(themeConfig.OptionsTextColor)
	OptionsFgColor = GetFgColor(themeConfig.OptionsTextColor)

	isLightTheme := themeConfig.LightTheme
	if isLightTheme {
		DefaultTextColor = color.FgBlack
		DefaultHiTextColor = color.FgHiBlack
		GocuiDefaultTextColor = gocui.ColorBlack
	} else {
		DefaultTextColor = color.FgWhite
		DefaultHiTextColor = color.FgHiWhite
		GocuiDefaultTextColor = gocui.ColorWhite
	}
}

// getHexColorValues returns the rgb values of a hex color
func getHexColorValues(v string) (r int32, g int32, b int32, valid bool) {
	if len(v) == 4 {
		v = string([]byte{v[0], v[1], v[1], v[2], v[2], v[3], v[3]})
	} else if len(v) != 7 {
		return
	}

	if v[0] != '#' {
		return
	}

	rgb, err := hex.DecodeString(v[1:])
	if err != nil {
		return
	}

	return int32(rgb[0]), int32(rgb[1]), int32(rgb[2]), true
}

// GetAttribute gets the gocui color attribute from the string
func GetGocuiAttribute(key string) gocui.Attribute {
	r, g, b, validHexColor := getHexColorValues(key)
	if validHexColor {
		return gocui.NewRGBColor(r, g, b)
	}

	colorMap := map[string]gocui.Attribute{
		"default":   gocui.ColorDefault,
		"black":     gocui.ColorBlack,
		"red":       gocui.ColorRed,
		"green":     gocui.ColorGreen,
		"yellow":    gocui.ColorYellow,
		"blue":      gocui.ColorBlue,
		"magenta":   gocui.ColorMagenta,
		"cyan":      gocui.ColorCyan,
		"white":     gocui.ColorWhite,
		"bold":      gocui.AttrBold,
		"reverse":   gocui.AttrReverse,
		"underline": gocui.AttrUnderline,
	}
	value, present := colorMap[key]
	if present {
		return value
	}
	return gocui.ColorWhite
}

// GetFgAttribute gets the color foreground attribute from the string
func GetFgAttribute(key string) color.Attribute {
	colorMap := map[string]color.Attribute{
		"default":   color.FgWhite,
		"black":     color.FgBlack,
		"red":       color.FgRed,
		"green":     color.FgGreen,
		"yellow":    color.FgYellow,
		"blue":      color.FgBlue,
		"magenta":   color.FgMagenta,
		"cyan":      color.FgCyan,
		"white":     color.FgWhite,
		"bold":      color.Bold,
		"reverse":   color.ReverseVideo,
		"underline": color.Underline,
	}
	value, present := colorMap[key]
	if present {
		return value
	}
	return color.FgWhite
}

// GetBgAttribute gets the color background attribute from the string
func GetBgAttribute(key string) color.Attribute {
	colorMap := map[string]color.Attribute{
		"default":   color.BgWhite,
		"black":     color.BgBlack,
		"red":       color.BgRed,
		"green":     color.BgGreen,
		"yellow":    color.BgYellow,
		"blue":      color.BgBlue,
		"magenta":   color.BgMagenta,
		"cyan":      color.BgCyan,
		"white":     color.BgWhite,
		"bold":      color.Bold,
		"reverse":   color.ReverseVideo,
		"underline": color.Underline,
	}
	value, present := colorMap[key]
	if present {
		return value
	}
	return color.FgWhite
}

// GetGocuiColor bitwise OR's a list of attributes obtained via the given keys
func GetGocuiColor(keys []string) gocui.Attribute {
	var attribute gocui.Attribute
	for _, key := range keys {
		attribute |= GetGocuiAttribute(key)
	}
	return attribute
}

// GetColor bitwise OR's a list of attributes obtained via the given keys
func GetBgColor(keys []string) color.Attribute {
	var attribute color.Attribute
	for _, key := range keys {
		attribute |= GetBgAttribute(key)
	}
	return attribute
}

// GetColor bitwise OR's a list of attributes obtained via the given keys
func GetFgColor(keys []string) color.Attribute {
	var attribute color.Attribute
	for _, key := range keys {
		attribute |= GetFgAttribute(key)
	}
	return attribute
}
