package main

import (
	"os"

	deb "github.com/getcouragenow/bootstrap/tool/packager/linux"
)

func main() {
	// debian
	pwd, _ := os.Getwd()
	p := deb.Pack{
		Path:    pwd,
		Name:    "linux-deb",
		AppName: "simple",
		Ctl: deb.Control{
			Architecture: "amd64",
			Description:  "simple description",
			Maintainer:   "maintainer",
			PackageName:  "simple",
			Priority:     "optional",
			Version:      "1.0",
		},
		Desk: deb.Desktop{
			Version:    "1.0",
			Type:       "Application",
			Terminal:   "false",
			Categories: " ",
			Name:       "simple",
			Icon:       "/usr/lib/simple/assets/icon.png",
			Exec:       "/usr/lib/simple/simple",
		},
		App: map[string]string{
			pwd + "/pubspec.yaml": pwd + "/linux-deb/usr/bin",
		},
	}
	p.Build()
}
