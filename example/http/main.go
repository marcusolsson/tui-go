package main

import (
	"github.com/marcusolsson/tui-go"
)

func main() {
	reqParamsEdit := tui.NewTextEdit()
	reqParamsEdit.SetText("x=2")

	reqParams := tui.NewVBox(reqParamsEdit)
	reqParams.SetTitle("URL Params")
	reqParams.SetBorder(true)

	reqMethodEntry := tui.NewEntry()
	reqMethodEntry.SetText("GET")

	reqMethod := tui.NewVBox(reqMethodEntry)
	reqMethod.SetTitle("Request method")
	reqMethod.SetBorder(true)
	reqMethod.SetSizePolicy(tui.Preferred, tui.Maximum)

	reqDataEdit := tui.NewTextEdit()
	reqDataEdit.SetText(`{"id": 12}`)

	reqData := tui.NewVBox(reqDataEdit)
	reqData.SetTitle("Request body")
	reqData.SetBorder(true)

	reqHeadEdit := tui.NewTextEdit()
	reqHeadEdit.SetText("User-Agent: myBrowser")

	reqHead := tui.NewVBox(reqHeadEdit)
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

	tui.DefaultFocusChain.Set(urlEntry, reqParamsEdit, reqMethodEntry, reqDataEdit, reqHeadEdit)

	ui := tui.New(root)
	ui.SetKeybinding(tui.KeyEsc, func() { ui.Quit() })

	if err := ui.Run(); err != nil {
		panic(err)
	}
}
