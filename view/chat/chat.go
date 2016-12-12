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
	//nextIndex := (active + 1) % len(viewArr)
	//name := viewArr[nextIndex]
	//
	//out, err := g.View("v2")
	//if err != nil {
	//	return err
	//}
	//fmt.Fprintln(out, "Going from view "+v.Name()+" to "+name)
	//
	//if _, err := setCurrentViewOnTop(g, name); err != nil {
	//	return err
	//}
	//
	//if nextIndex == 0 || nextIndex == 3 {
	//	g.Cursor = true
	//} else {
	//	g.Cursor = false
	//}
	//
	//active = nextIndex
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

	if v, err := g.SetView("func", -1, -1, maxX/12, maxY); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}

		//v.Title="1"
		//v.Editable = true
		v.Wrap = true
		fmt.Fprint(v, "func")
	}

	// 三个类别按钮 我 : 个人信息 你: 常用联系人 他:所有联系人列表
	if v, err := g.SetView("me", maxX/50, maxY/18, maxX/50*3, maxY/24*4); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}

		v.Editable = false
		// 坑 中文字符会错行
		fmt.Fprint(v, "ME")
	}
	if v, err := g.SetView("you", maxX/50, maxY/18*3, maxX/50*3, maxY/24*8); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}

		v.Editable = false
		fmt.Fprint(v, "YOU")
	}
	if v, err := g.SetView("he", maxX/50, maxY/18*5, maxX/50*3, maxY/24*12); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}

		v.Editable = false
		fmt.Fprint(v, "HE")
	}
	if v, err := g.SetView("candy", maxX/50-1, maxY-3, maxX/50*3+1, maxY-1); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}

		v.Editable = false
		fmt.Fprint(v, "candy")
	}

	if v, err := g.SetView("search", maxX/12, -1, maxX/3-5, 2); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}

		v.Wrap = true
		v.Autoscroll = true
		fmt.Fprint(v, "search")
	}

	if v, err := g.SetView("find", maxX/3-5, -1, maxX/3, 2); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}

		v.Wrap = true
		v.Autoscroll = true
		fmt.Fprint(v, "find")
	}

	if v, err := g.SetView("list", maxX/12, 2, maxX/3, maxY); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}

		v.Wrap = true
		v.Autoscroll = true
		fmt.Fprint(v, "friend list")
	}
	if v, err := g.SetView("input", maxX/3, maxY/10*8, maxX, maxY); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}

		v.Wrap = true
		v.Autoscroll = true
		fmt.Fprint(v, "input")
	}

	if v, err := g.SetView("send", maxX/10*9+3, maxY-3, maxX/20*19+2, maxY-1); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Editable = false
		fmt.Fprintln(v, "send")
	}

	if v, err := g.SetView("notice", maxX/3,maxY/10*7, maxX, maxY/10*8); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}

		//v.Title="4"
		v.Editable = true

		//if _, err = setCurrentViewOnTop(g, "v1"); err != nil {
		//	return err
		//}
		fmt.Fprint(v, "notice")
	}
	if v, err := g.SetView("chat", maxX/3, -1, maxX, maxY/10*7); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}

		//v.Title="5"
		v.Editable = true
		fmt.Fprint(v, "chat")
	}
	return nil
}

// setCurrentViewOnTop 选中的窗口 ctl+tab 切换
func setCurrentViewOnTop(g *gocui.Gui, name string) (*gocui.View, error) {
	if _, err := g.SetCurrentView(name); err != nil {
		return nil, err
	}
	return g.SetViewOnTop(name)
}