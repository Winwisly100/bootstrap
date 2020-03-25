package deb

import (
	"os"
	"os/exec"
	"path"
)

// Pack struct
type Pack struct {
	Path    string
	Name    string
	AppName string
	Ctl     Control
	Desk    Desktop
	App     map[string]string
}

func (p *Pack) init() {
	os.MkdirAll(p.getPath()+"/DEBIAN", 0755)
	os.MkdirAll(p.getPath()+"/usr/bin", 0755)
	os.MkdirAll(p.getPath()+"/usr/share/applications", 0755)
}

func (p *Pack) getPath() string {
	return path.Join(p.Path, p.Name)
}

// Build deb
func (p *Pack) Build() {
	p.init()
	writeFile(p.getPath()+"/DEBIAN/control", p.Ctl.encode())
	writeFile(p.getPath()+"/usr/share/applications/"+p.AppName+".desktop", p.Desk.encode())
	for k, v := range p.App {
		execute("cp", "-r", k, v)
	}
}

// Control struct
type Control struct {
	PackageName  string `json:"package"`
	Architecture string `json:"architecture"`
	Maintainer   string `json:"maintainer"`
	Priority     string `json:"priority"`
	Version      string `json:"version"`
	Description  string `json:"description"`
}

func (c *Control) encode() []byte {
	out := ""
	out += "package: " + c.PackageName + "\n"
	out += "architecture: " + c.Architecture + "\n"
	out += "maintainer: " + c.Maintainer + "\n"
	out += "priority: " + c.Priority + "\n"
	out += "version: " + c.Version + "\n"
	out += "description: " + c.Description
	return []byte(out)
}

func debControlFile(data map[string]string) *Control {
	return &Control{
		PackageName:  data["package"],
		Architecture: data["architecture"],
		Maintainer:   data["maintainer"],
		Priority:     data["priority"],
		Version:      data["version"],
		Description:  data["description"],
	}
}

// Desktop struct
type Desktop struct {
	Version    string `json:"version"`
	Type       string `json:"type"`
	Terminal   string `json:"terminal"`
	Categories string `json:"categories"`
	Name       string `json:"name"`
	Icon       string `json:"icon"`
	Exec       string `json:"exec"`
}

func (d *Desktop) encode() []byte {
	out := "[Desktop Entry]" + "\n"
	out += "Version= " + d.Version + "\n"
	out += "Type= " + d.Type + "\n"
	out += "Terminal= " + d.Terminal + "\n"
	out += "Categories= " + d.Categories + "\n"
	out += "Name= " + d.Name + "\n"
	out += "Icon= " + d.Icon + "\n"
	out += "Exec= " + d.Exec
	return []byte(out)
}

func execute(cmd string, args ...string) error {
	c := exec.Command(cmd, args...)
	c.Stdout = os.Stdout
	c.Stderr = os.Stderr
	return c.Run()
}

func writeFile(path string, data []byte) error {
	f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = f.Write(data)
	return err
}
