package chat

import (
	"fmt"

	"github.com/jroimartin/gocui"
	//"github.com/zeazen/candy-cui/candy"
	"log"
)

// showRegisteLayout 切换到 chat 界面
func ShowChatLayout(g *gocui.Gui, v *gocui.View) error {
	g.SetManagerFunc(LayoutChat)
	if err := chatKeybindings(g); err != nil {
		log.Panicln(err)
		return err
	}
	return nil
}

// chatNextView 聊天窗口使用 ctl+tab 切换
func chatNextView(g *gocui.Gui, v *gocui.View) error {

	return nil
}

// registeKeybindings registe 界面按键绑定
func chatKeybindings(g *gocui.Gui) error {
	// Registe 界面的 Tab 切换 binding
	//if err := g.SetKeybinding("registeEmailTextField", gocui.KeyTab, gocui.ModNone, registeNextView); err != nil {
	//	return err
	//}
	//if err := g.SetKeybinding("registePasswdTextField", gocui.KeyTab, gocui.ModNone, registeNextView); err != nil {
	//	return err
	//}
	//if err := g.SetKeybinding("passwdRepeatTextField", gocui.KeyTab, gocui.ModNone, registeNextView); err != nil {
	//	return err
	//}
	//if err := g.SetKeybinding("registeCallButton", gocui.KeyTab, gocui.ModNone, registeNextView); err != nil {
	//	return err
	//}
	//if err := g.SetKeybinding("registeCancleButton", gocui.KeyTab, gocui.ModNone, registeNextView); err != nil {
	//	return err
	//}
	//
	//// 各按钮功能部分
	//if err := g.SetKeybinding("registeCallButton", gocui.KeyEnter, gocui.ModNone, callRegiste); err != nil {
	//	return err
	//}
	//if err := g.SetKeybinding("registeCancleButton", gocui.KeyEnter, gocui.ModNone, backToLoginLayout); err != nil {
	//	return err
	//}

	return nil
}

// LayoutRegiste registe 布局
func LayoutChat(g *gocui.Gui) error {
	maxX, maxY := g.Size()

	// geek 极简版本
	if v, err := g.SetView("chat", -1, -1, maxX/6*5, maxY-3); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}

		v.Editable = false
		v.Frame = false
		v.Autoscroll = true
		fmt.Fprint(v, "chat")
	}
	if v, err := g.SetView("input", -1, maxY-3, maxX, maxY); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}

		v.Autoscroll = false
		v.Editable = true
		fmt.Fprint(v, "input")
	}
	if v, err := g.SetView("list", maxX/6*5, -1, maxX, maxY-3); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}

		v.Editable = false
		v.Frame = false
		v.Autoscroll = true
		fmt.Fprint(v, "list")
	}

	g.SetCurrentView("input")
	return nil
}

// setCurrentViewOnTop 选中的窗口 ctl+tab 切换
func setCurrentViewOnTop(g *gocui.Gui, name string) (*gocui.View, error) {
	if _, err := g.SetCurrentView(name); err != nil {
		return nil, err
	}
	return g.SetViewOnTop(name)
}
