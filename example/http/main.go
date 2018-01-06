package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"sort"
	"strings"

	"github.com/marcusolsson/tui-go"
)

var (
	method  = "GET"
	params  = "x=2&y=3"
	payload = `{"id": 12}`
	headers = "User-Agent: myBrowser"
)

func main() {
	reqParamsEdit := tui.NewTextEdit()
	reqParamsEdit.SetText(params)
	reqParamsEdit.OnTextChanged(func(e *tui.TextEdit) {
		params = e.Text()
	})

	reqParams := tui.NewVBox(reqParamsEdit)
	reqParams.SetTitle("URL Params")
	reqParams.SetBorder(true)

	reqMethodEntry := tui.NewEntry()
	reqMethodEntry.SetText(method)
	reqMethodEntry.OnChanged(func(e *tui.Entry) {
		method = e.Text()
	})

	reqMethod := tui.NewVBox(reqMethodEntry)
	reqMethod.SetTitle("Request method")
	reqMethod.SetBorder(true)
	reqMethod.SetSizePolicy(tui.Preferred, tui.Maximum)

	reqDataEdit := tui.NewTextEdit()
	reqDataEdit.SetText(payload)
	reqDataEdit.OnTextChanged(func(e *tui.TextEdit) {
		payload = e.Text()
	})

	reqData := tui.NewVBox(reqDataEdit)
	reqData.SetTitle("Request body")
	reqData.SetBorder(true)

	reqHeadEdit := tui.NewTextEdit()
	reqHeadEdit.SetText(headers)
	reqHeadEdit.OnTextChanged(func(e *tui.TextEdit) {
		headers = e.Text()
	})

	reqHead := tui.NewVBox(reqHeadEdit)
	reqHead.SetTitle("Request headers")
	reqHead.SetBorder(true)

	respHeadLbl := tui.NewLabel("")
	respHeadLbl.SetSizePolicy(tui.Expanding, tui.Expanding)

	respHead := tui.NewVBox(respHeadLbl)
	respHead.SetTitle("Response headers")
	respHead.SetBorder(true)

	respBodyLbl := tui.NewLabel("")
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
	urlEntry.OnSubmit(func(e *tui.Entry) {
		req, err := http.NewRequest(method, e.Text(), strings.NewReader(payload))
		if err != nil {
			return
		}
		req.URL.RawQuery = params

		for _, h := range strings.Split(headers, "\n") {
			kv := strings.Split(h, ":")
			if len(kv) == 2 {
				req.Header.Set(kv[0], kv[1])
			}
		}

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			return
		}
		defer resp.Body.Close()

		var headers []string
		for k, v := range resp.Header {
			headers = append(headers, k+": "+strings.Join(v, ";"))

		}
		sort.Strings(headers)

		respHeadLbl.SetText(strings.Join(headers, "\n"))

		b, _ := ioutil.ReadAll(resp.Body)
		respBodyLbl.SetText(string(b))
	})

	urlBox := tui.NewHBox(urlEntry)
	urlBox.SetTitle("URL")
	urlBox.SetBorder(true)

	root := tui.NewVBox(urlBox, browser)

	tui.DefaultFocusChain.Set(urlEntry, reqParamsEdit, reqMethodEntry, reqDataEdit, reqHeadEdit)

	theme := tui.NewTheme()
	theme.SetStyle("box.focused.border", tui.Style{Fg: tui.ColorYellow, Bg: tui.ColorDefault})

	ui, err := tui.New(root)
	if err != nil {
		log.Fatal(err)
	}

	ui.SetTheme(theme)
	ui.SetKeybinding("Esc", func() { ui.Quit() })

	if err := ui.Run(); err != nil {
		log.Fatal(err)
	}
}
