//go:build windows
// +build windows

package main

import (
	"fmt"
	"github.com/fy0/systray"
	"github.com/gen2brain/beeep"
	"github.com/labstack/echo/v4"
	"github.com/lxn/win"
	"math/rand"
	"net"
	"os"
	"os/exec"
	"regexp"
	"runtime"
	"sealdice-core/dice"
	"sealdice-core/icon"
	"strings"
	"syscall"
	"time"
	"unsafe"
)

func hideWindow() {
	win.ShowWindow(win.GetConsoleWindow(), win.SW_HIDE)
}

func showWindow() {
	win.ShowWindow(win.GetConsoleWindow(), win.SW_SHOW)
}

func trayInit() {
	// 确保能收到系统消息，从而避免不能弹出菜单
	runtime.LockOSThread()
	systray.Run(onReady, onExit)
}

var (
	kernel32        = syscall.NewLazyDLL("kernel32.dll")
	procCreateMutex = kernel32.NewProc("CreateMutexW")
)

func CreateMutex(name string) (uintptr, error) {
	ret, _, err := procCreateMutex.Call(
		0,
		0,
		uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(name))),
	)
	switch int(err.(syscall.Errno)) {
	case 0:
		return ret, nil
	default:
		return ret, err
	}
}

func TestRunning() bool {
	_, err := CreateMutex("SealDice")
	if err == nil {
		return false
	}

	s1, _ := syscall.UTF16PtrFromString("SealDice 海豹已经在运作")
	s2, _ := syscall.UTF16PtrFromString("如果你想在Windows上打开多个海豹，请点“确定”，或加参数-m启动。\n如果只是打开UI界面，请在任务栏右下角的系统托盘区域找到海豹图标并右键，点“取消")
	ret := win.MessageBox(0, s2, s1, win.MB_YESNO|win.MB_ICONWARNING|win.MB_DEFBUTTON2)
	if ret == win.IDYES {
		return false
	}
	return true
}

func PortExistsWarn() {
	s1, _ := syscall.UTF16PtrFromString("SealDice 启动失败")
	s2, _ := syscall.UTF16PtrFromString("端口已被占用，建议换用其他端口")
	win.MessageBox(0, s2, s1, win.MB_OK)
	return
}

func onReady() {
	systray.SetIcon(icon.Data)
	systray.SetTitle("海豹TRPG骰点核心")
	systray.SetTooltip("海豹TRPG骰点核心")

	mOpen := systray.AddMenuItem("打开界面", "开启WebUI")
	mShowHide := systray.AddMenuItemCheckbox("显示终端窗口", "显示终端窗口", false)
	mQuit := systray.AddMenuItem("退出", "退出程序")
	mOpen.SetIcon(icon.Data)

	go beeep.Notify("SealDice", "我藏在托盘区域了，点我的小图标可以快速打开UI", "assets/information.png")

	for {
		select {
		case <-mOpen.ClickedCh:
			exec.Command(`cmd`, `/c`, `start`, `http://localhost:3211`).Start()
		case <-mQuit.ClickedCh:
			systray.Quit()
			time.Sleep(1 * time.Second)
			os.Exit(0)
		case <-mShowHide.ClickedCh:
			if mShowHide.Checked() {
				win.ShowWindow(win.GetConsoleWindow(), win.SW_HIDE)
				mShowHide.Uncheck()
			} else {
				win.ShowWindow(win.GetConsoleWindow(), win.SW_SHOW)
				mShowHide.Check()
			}
		}
	}
}

func onExit() {
	// clean up here
}

func httpServe(e *echo.Echo, dm *dice.DiceManager) {
	portStr := "3211"

	go func() {
		runtime.LockOSThread()
		for {
			time.Sleep(10 * time.Second)
			systray.SetTooltip("海豹TRPG骰点核心 #" + portStr)
		}
	}()

	var theFunc func()
	subFunc := func() {
		rePort := regexp.MustCompile(`:(\d+)$`)
		m := rePort.FindStringSubmatch(dm.ServeAddress)
		if len(m) > 0 {
			portStr = m[1]
		}

		ln, err := net.Listen("tcp", ":"+portStr)
		if err != nil {
			s1, _ := syscall.UTF16PtrFromString("海豹TRPG骰点核心")
			s2, _ := syscall.UTF16PtrFromString(fmt.Sprintf("端口 %s 已被占用，点“是”随机换一个端口，点“否”退出\n注意，此端口将被自动写入配置，后续可用启动参数改回", portStr))
			ret := win.MessageBox(0, s2, s1, win.MB_YESNO|win.MB_ICONWARNING|win.MB_DEFBUTTON2)
			if ret == win.IDYES {
				newPort := 3000 + rand.Int()%4000
				dm.ServeAddress = strings.Replace(dm.ServeAddress, portStr, fmt.Sprintf("%d", newPort), 1)
				theFunc()
				return
			} else {
				logger.Errorf("端口已被占用，即将自动退出: %s", dm.ServeAddress)
				os.Exit(1)
			}
		}
		_ = ln.Close()

		go func() {
			time.Sleep(5 * time.Second) // 先不玩花活，等5s即可
			exec.Command(`cmd`, `/c`, `start`, fmt.Sprintf(`http://localhost:%s`, portStr)).Start()
		}()

		fmt.Println("如果浏览器没有自动打开，请手动访问:")
		fmt.Println(fmt.Sprintf(`http://localhost:%s`, portStr)) // 默认:3211
		err = e.Start(dm.ServeAddress)
		if err != nil {
			logger.Errorf("端口已被占用，即将自动退出: %s", dm.ServeAddress)
			return
		}
	}

	theFunc = subFunc
	subFunc()
}
