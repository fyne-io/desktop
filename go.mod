module fyne.io/desktop

go 1.12

require (
	fyne.io/fyne v1.2.4-0.20200328093604-76394fca3cc4
	github.com/BurntSushi/freetype-go v0.0.0-20160129220410-b763ddbfe298 // indirect
	github.com/BurntSushi/graphics-go v0.0.0-20160129215708-b43f31a4a966 // indirect
	github.com/BurntSushi/xgb v0.0.0-20160522181843-27f122750802
	github.com/BurntSushi/xgbutil v0.0.0-20160919175755-f7c97cef3b4e
	github.com/jackmordaunt/icns v1.0.0
	github.com/stretchr/testify v1.5.1
	howett.net/plist v0.0.0-20181124034731-591f970eefbb
)

replace github.com/BurntSushi/xgbutil => github.com/fyne-io/xgbutil v0.0.0-20191220152344-7d838166824d
