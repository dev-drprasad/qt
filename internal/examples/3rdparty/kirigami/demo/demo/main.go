//source: https://github.com/KDE/kirigami

package main

import (
	"os"

	"github.com/dev-drprasad/qt/core"
	"github.com/dev-drprasad/qt/qml"
	"github.com/dev-drprasad/qt/quickcontrols2"
	"github.com/dev-drprasad/qt/widgets"
)

func main() {

	core.QCoreApplication_SetAttribute(core.Qt__AA_EnableHighDpiScaling, true)

	widgets.NewQApplication(len(os.Args), os.Args)

	if quickcontrols2.QQuickStyle_Name() == "" {
		quickcontrols2.QQuickStyle_SetStyle("Material")
	}

	engine := qml.NewQQmlApplicationEngine(nil)

	if os.Getenv("QT_QUICK_CONTROLS_STYLE") == "org.kde.desktop" {
		engine.Load(core.NewQUrl3("qrc:/contents/ui/DesktopExampleApp.qml", 0))
	} else {
		engine.Load(core.NewQUrl3("qrc:/contents/ui/ExampleApp.qml", 0))
	}

	widgets.QApplication_Exec()
}
