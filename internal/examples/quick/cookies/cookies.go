package main

import (
	"os"

	"github.com/dev-drprasad/qt/core"
	"github.com/dev-drprasad/qt/gui"
	"github.com/dev-drprasad/qt/network"
	"github.com/dev-drprasad/qt/quick"
	"github.com/dev-drprasad/qt/webengine"
)

type bridge struct {
	core.QObject

	_ func(*webengine.QQuickWebEngineProfile) `signal:"sendProfile"`
}

func main() {
	core.QCoreApplication_SetAttribute(core.Qt__AA_EnableHighDpiScaling, true)

	gui.NewQGuiApplication(len(os.Args), os.Args)
	webengine.QtWebEngine_Initialize()

	b := NewBridge(nil)
	b.ConnectSendProfile(func(v0 *webengine.QQuickWebEngineProfile) {

		store := v0.CookieStore()

		store.ConnectCookieAdded(func(cookie *network.QNetworkCookie) {
			println(cookie.ToRawForm(network.QNetworkCookie__Full).ConstData())
		})

		store.LoadAllCookies()
	})

	view := quick.NewQQuickView(nil)
	view.RootContext().SetContextProperty("bridge", b)
	view.SetSource(core.NewQUrl3("qrc:/qml/cookies.qml", 0))
	view.SetResizeMode(quick.QQuickView__SizeRootObjectToView)
	view.Show()

	gui.QGuiApplication_Exec()
}
