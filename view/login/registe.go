package login

import (
	"fmt"
	//"io"
	//"log"
	//"strings"
	//	"errors"
	"github.com/jroimartin/gocui"
	"github.com/zeazen/candy-cui/candy"
	"log"
)

// showRegisteLayout 切换到 registe 界面
func showRegisteLayout(g *gocui.Gui, v *gocui.View) error {
	g.SetManagerFunc(LayoutRegiste)
	if err := registeKeybindings(g); err != nil {
		log.Panicln(err)
		return err
	}
	return nil
}

// registeNextView registe 内部控件 tab 切换
func registeNextView(g *gocui.Gui, v *gocui.View) error {
	if v == nil || v.Name() == "registeEmailTextField" {
		_, err := g.SetCurrentView("registePasswdTextField")
		return err
	}

	if v == nil || v.Name() == "registePasswdTextField" {
		_, err := g.SetCurrentView("passwdRepeatTextField")
		return err
	}

	if v == nil || v.Name() == "passwdRepeatTextField" {
		_, err := g.SetCurrentView("registeCallButton")
		return err
	}

	if v == nil || v.Name() == "registeCallButton" {
		_, err := g.SetCurrentView("registeCancleButton")
		return err
	}

	if v == nil || v.Name() == "registeCancleButton" {
		_, err := g.SetCurrentView("registeEmailTextField")
		return err
	}

	_, err := g.SetCurrentView("registeEmailTextField")
	return err
}

// callRegiste 临时 调用 registe
func callRegiste(g *gocui.Gui, v *gocui.View) error {
	//获取用户输入
	v, err := g.View("registeEmailTextField")
	if err != nil {
		return err
	}
	emailStr := v.Buffer()
	v, err = g.View("registePasswdTextField")
	if err != nil {
		return err
	}
	passwd1 := v.Buffer()
	v.Clear()
	v, err = g.View("passwdRepeatTextField")
	if err != nil {
		return err
	}
	passwd2 := v.Buffer()
	v.Clear()
	if passwd1 != passwd2 {
		fmt.Fprint(v, emailStr)
		//return errors.New("密码不匹配")
	}
	_, err = candy.CandyCUIClient.Register(emailStr, passwd1)
	if err != nil {
		//e := candy.ErrorParse(err.Error())
		//		log.Errorf("Register code:%v error:%v", e.Code, e.Msg)
		return err
	}
	//	log.Debugf("Register success, userID:%v userName:%v userPassword:%v", id, emailStr, passwd1)
	return nil
}

// registeKeybindings registe 界面按键绑定
func registeKeybindings(g *gocui.Gui) error {
	// Registe 界面的 Tab 切换 binding
	if err := g.SetKeybinding("registeEmailTextField", gocui.KeyTab, gocui.ModNone, registeNextView); err != nil {
		return err
	}
	if err := g.SetKeybinding("registePasswdTextField", gocui.KeyTab, gocui.ModNone, registeNextView); err != nil {
		return err
	}
	if err := g.SetKeybinding("passwdRepeatTextField", gocui.KeyTab, gocui.ModNone, registeNextView); err != nil {
		return err
	}
	if err := g.SetKeybinding("registeCallButton", gocui.KeyTab, gocui.ModNone, registeNextView); err != nil {
		return err
	}
	if err := g.SetKeybinding("registeCancleButton", gocui.KeyTab, gocui.ModNone, registeNextView); err != nil {
		return err
	}

	// 各按钮功能部分
	if err := g.SetKeybinding("registeCallButton", gocui.KeyEnter, gocui.ModNone, callRegiste); err != nil {
		return err
	}
	if err := g.SetKeybinding("registeCancleButton", gocui.KeyEnter, gocui.ModNone, backToLoginLayout); err != nil {
		return err
	}

	return nil
}

// LayoutRegiste registe 布局
func LayoutRegiste(g *gocui.Gui) error {
	maxX, maxY := g.Size()
	if v, err := g.SetView("Logo", -1, -1, maxX, maxY); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Highlight = true
		v.SelBgColor = gocui.ColorBlack
		v.SelFgColor = gocui.ColorWhite
		showLogo(v)

	}
	if v, err := g.SetView("registeEmailTextField", maxX/2-20, maxY/2, maxX/2+18, maxY/2+2); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Editable = true
		g.SetCurrentView("registeEmailTextField")
	}
	if v, err := g.SetView("registePasswdTextField", maxX/2-20, maxY/2+3, maxX/2+18, maxY/2+5); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Editable = true
		v.Mask = '*'
	}
	if v, err := g.SetView("passwdRepeatTextField", maxX/2-20, maxY/2+6, maxX/2+18, maxY/2+8); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Editable = true
		v.Mask = '*'
	}
	if v, err := g.SetView("registeCallButton", maxX/2-14, maxY/2+9, maxX/2-6, maxY/2+11); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Editable = false
		fmt.Fprintln(v, "Registe")
	}

	if v, err := g.SetView("registeCancleButton", maxX/2+2, maxY/2+9, maxX/2+9, maxY/2+11); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Editable = false
		fmt.Fprintln(v, "Cancle")
	}
	return nil
}
