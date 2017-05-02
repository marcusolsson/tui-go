package main

import (
	"github.com/marcusolsson/tui-go"
)

func main() {
	reqParamsEdit := tui.NewTextEdit()
	reqParamsEdit.SetText("x=2")
	reqParamsEdit.SetFocused(true)

	reqParams := tui.NewVBox(reqParamsEdit)
	reqParams.SetTitle("URL Params")
	reqParams.SetBorder(true)

	reqMethod := tui.NewVBox(tui.NewLabel("GET"))
	reqMethod.SetTitle("Request method")
	reqMethod.SetBorder(true)
	reqMethod.SetSizePolicy(tui.Preferred, tui.Maximum)

	reqData := tui.NewVBox(tui.NewLabel(`{"id": 12}`))
	reqData.SetTitle("Request body")
	reqData.SetBorder(true)

	reqHead := tui.NewVBox(tui.NewLabel("User-Agent: myBrowser"))
	reqHead.SetTitle("Request headers")
	reqHead.SetBorder(true)

	respHeadLbl := tui.NewLabel("HTTP/1.1 200 OK")
	respHeadLbl.SetSizePolicy(tui.Expanding, tui.Expanding)

	respHead := tui.NewVBox(respHeadLbl)
	respHead.SetTitle("Response headers")
	respHead.SetBorder(true)

	respBodyLbl := tui.NewLabel("{\n  \"args\": {\n    \"x\": 2\n  }\n}")
	respBodyLbl.SetSizePolicy(tui.Expanding, tui.Expanding)

	respBody := tui.NewVBox(respBodyLbl)
	respBody.SetTitle("Response body")
	respBody.SetBorder(true)

	req := tui.NewVBox(reqParams, reqMethod, reqData, reqHead)
	resp := tui.NewVBox(respHead, respBody)
	resp.SetSizePolicy(tui.Expanding, tui.Preferred)

	browser := tui.NewHBox(req, resp)
	browser.SetSizePolicy(tui.Preferred, tui.Expanding)

	urlEntry := tui.NewEntry()
	urlEntry.SetText("https://httpbin.org/get")

	urlBox := tui.NewHBox(urlEntry)
	urlBox.SetTitle("URL")
	urlBox.SetBorder(true)

	root := tui.NewVBox(urlBox, browser)

	ui := tui.New(root)
	ui.SetKeybinding(tui.KeyEsc, func() { ui.Quit() })

	if err := ui.Run(); err != nil {
		panic(err)
	}
}
