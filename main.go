package main

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"image/color"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"
)

type myTheme struct{}

var _ fyne.Theme = (*myTheme)(nil)

func (m myTheme) Color(name fyne.ThemeColorName, variant fyne.ThemeVariant) color.Color {
	return theme.DefaultTheme().Color(name, variant)
}

func (m myTheme) Font(style fyne.TextStyle) fyne.Resource {
	if style.Bold {
		return resourceMsyhbdTtc
	}
	if style.Italic {
		return theme.DefaultTheme().Font(style)
	}
	if style.Monospace {
		return theme.DefaultTheme().Font(style)
	}
	return resourceMsyhTtc
}

func (m myTheme) Icon(name fyne.ThemeIconName) fyne.Resource {
	return theme.DefaultTheme().Icon(name)
}

func (m myTheme) Size(name fyne.ThemeSizeName) float32 {
	return theme.DefaultTheme().Size(name)
}

var tk = widget.NewLabel("")
var cs = widget.NewLabel("")
var bd = widget.NewLabel("")
var myApp = app.New()

func yuming(tk string) string {
	url := "https://cat-match.easygame2021.com/sheep/v1/game/game_over?rank_score=1&rank_state=1&rank_time=44&rank_role=1&skin=1"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal("Error reading request. ", err)
	}
	//req.Header.Set("Host", "site.ip138.com")
	req.Header.Set("User-Agent", " Mozilla/5.0 (iPhone; CPU iPhone OS 15_4_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Mobile/15E148 MicroMessenger/8.0.28(0x18001c22) NetType/WIFI Language/zh_CN")
	req.Header.Set("T", tk)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Referer", "https://servicewechat.com/wx141bfb9b73c970a9/19/page-frame.html")

	client := &http.Client{Timeout: time.Second * 1000}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Error reading response. ", err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("Error reading body. ", err)
	}
	fmt.Println(string(body))
	return string(body)

}

func info(res string, res2 int) fyne.CanvasObject {

	screen := widget.NewForm(
		&widget.FormItem{Text: "闯关状况", Widget: bd},
	)
	if res == "" {
		bd.SetText("请输入微信token")
	} else {
		bd.SetText("请稍后")
		var yy int
		yy = 0
		for i := 0; i < res2; i++ {

			if yuming(res) == "{\"err_code\":0,\"err_msg\":\"\",\"data\":0}" {
				fmt.Println(yy)
				yy = yy + 1
				bd.SetText("目前已成功闯过第二关-->" + strconv.Itoa(yy) + "次")
			} else {
				fmt.Println(yy)
				bd.SetText("目前网络异常，请切换网络进行重试已完成次数-->" + strconv.Itoa(yy) + "次")
			}

		}
	}

	return screen
}

func query() fyne.CanvasObject {
	tk := widget.NewEntry()
	tk.SetPlaceHolder("请填写数据")
	cs := widget.NewEntry()
	cs.SetPlaceHolder("请填写数据")
	form := widget.NewForm(
		&widget.FormItem{Text: "请输入微信token:", Widget: tk, HintText: "必填"},
		&widget.FormItem{Text: "成功次数:", Widget: cs, HintText: "必填"},
	)
	form.OnCancel = func() {
		myApp.Quit()
	}

	form.OnSubmit = func() {
		test1, err := strconv.Atoi(cs.Text)
		if err != nil {
			fmt.Println("can't convert to int")
		}
		info(tk.Text, test1)
	}
	form.Resize(fyne.NewSize(800, 200))
	return form
}

func main() {
	myApp.Settings().SetTheme(&myTheme{})
	myWindow := myApp.NewWindow("羊了个羊闯光秘籍")

	myWindow.Resize(fyne.Size{Width: 800, Height: 600})

	myWindow.SetContent(fyne.NewContainerWithLayout(layout.NewGridLayoutWithRows(2), query(), info("", 0)))

	myWindow.ShowAndRun()
}
