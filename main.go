// Copyright 2014 The gocui Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	//"fmt"
	//"bufio"
	"log"
	//"strings"
	//"os"
	"github.com/jroimartin/gocui"
	"github.com/zeazen/candy-cui/candy"
	"github.com/zeazen/candy-cui/view/login"
)

// quit 退出
func quit(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}


func main() {
	g, err := gocui.NewGui()
	if err != nil {
		log.Panicln(err)
	}
	defer g.Close()
	g.Cursor = true

	if err := g.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, quit); err != nil {
		log.Panicln(err)
	}
	// 初始化 candy 客户端
	candy.CandyCUIClient = candy.NewCandyClient("172.16.23.53:9000", &candy.CuiHandler{})
	if err := candy.CandyCUIClient.Start(); err != nil {
		log.Panic(err)
	}

	// 加载程序首页 登录界面
	g.SetManagerFunc(login.LayoutLogin)
	if err := login.LoginKeybindings(g); err != nil {
		log.Panicln(err)
	}

	if err := g.MainLoop(); err != nil && err != gocui.ErrQuit {
		log.Panicln(err)
	}
}
