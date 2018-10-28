package test

import (
	"github.com/Penglq/QLog"
	"testing"
)

func TestFile(t *testing.T) {
	l := QLog.GetLogger()
	//l.SetConfig(time.Local.String())
	l.SetConfig(QLog.INFO, "Asia/Chongqing",
		QLog.WithFileOPT("./", "info", "log", 1,10),
		//QLog.WithFileOPT(true, "./", "info", "log", QLog.DEFAULTFILEMAXSIZE),
	)

	for i:=0;i<10;i++{
		l.Info("次数->", i,"struct--------------------------------------------")
	}
}

func TestConsole(t *testing.T) {
	st := struct {
		A string
	}{A: "it is a struct"}

	l := QLog.GetLogger()
	//l.SetConfig(time.Local.String())
	l.SetConfig(QLog.INFO, "Asia/Chongqing",
		QLog.WithConsoleOPT(),
	)

	l.Debug(`company`, "YRD", "province", "beijing")
	l.AddTextPrefix("method", "resUser")
	l.Info("name", "luqiang")
	l.Info("data", map[string]string{"info": "abc", "info1": "def"})
	l.Info("struct", st)
	// debug的不会输出
}
