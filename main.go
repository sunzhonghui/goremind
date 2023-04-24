package main

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/gen2brain/beeep"
	"goremind/res"
	"goremind/theme"
	"strconv"
	"time"
)

func main() {
	a := app.NewWithID("com.idmiss.timer")
	a.Settings().SetTheme(theme.MyTheme{})
	w := a.NewWindow("我的提醒")
	w.Resize(fyne.NewSize(400, 200))

	txentry := widget.NewEntry()
	txentry.SetPlaceHolder("输入提醒内容")

	entry := widget.NewEntry()
	entry.SetPlaceHolder("输入时间（分钟）")

	startBtn := widget.NewButton("开始", func() {

		minutes, err := strconv.Atoi(entry.Text)
		if err != nil {
			showDialog(w, "请输入正确的分钟数")
			return
		}

		duration := time.Duration(minutes) * time.Minute
		w.Hide()
		time.AfterFunc(duration, func() {
			sendNotification("时间到了！", fmt.Sprintf("你设定的 %d 分钟已经到了！该去%s了", minutes, txentry.Text), w)
		})

		fmt.Printf("设置了 %d 分钟\n", minutes)
	})

	content := container.NewVBox(txentry, entry, startBtn)
	w.SetIcon(res.MyIcon)
	w.SetContent(content)
	w.CenterOnScreen()
	w.ShowAndRun()
}
func init() {

}
func sendNotification(title, message string, w fyne.Window) {
	w.Show()
	showDialog(w, message)
	err := beeep.Notify(title, message, "")
	if err != nil {
		fmt.Println("通知失败")
	}
}

func showDialog(win fyne.Window, msg string) {
	// 创建弹窗内容
	label := widget.NewLabel(msg)

	closeButton := widget.NewButton("关闭", func() {

	})
	// 将内容添加到容器中
	content := container.NewVBox(
		label,
		closeButton,
	)

	// 创建并显示弹窗
	dialog := widget.NewModalPopUp(content, win.Canvas())
	closeButton.OnTapped = func() {
		dialog.Hide()
	}
	dialog.Show()
}
