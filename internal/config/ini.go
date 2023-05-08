package config

import (
	"strconv"

	"gopkg.in/ini.v1"
)

const (
	iniFilePath = "config.ini"
)

type Config struct {
	Coordinates struct {
		X int `ini:"x"`
		Y int `ini:"y"`
	} `ini:"Coordinates"`
	Keystrokes struct {
		Record   []string `ini:"record"`
		Deselect []string `ini:"deselect"`
	}
}

func Setup() (*Config, error) {
	inidata, err := loadIni()
	if err != nil {
		return nil, err
	}
	return mapToConfig(inidata)
}

func (c *Config) Save() error {
	inidata, err := loadIni()
	if err != nil {
		return err
	}
	sec := inidata.Section("Coordinates")
	sec.Key("x").SetValue(strconv.Itoa(c.Coordinates.X))
	sec.Key("y").SetValue(strconv.Itoa(c.Coordinates.Y))
	return inidata.SaveTo(iniFilePath)
}

func loadIni() (*ini.File, error) {
	inidata, err := ini.Load(iniFilePath)
	if err != nil {
		inidata, err = createIniFile()
	}
	return inidata, err
}

func createIniFile() (*ini.File, error) {
	inidata := ini.Empty()
	sec, err := inidata.NewSection("Coordinates")
	if err != nil {
		return nil, err
	}
	_, err = sec.NewKey("x", "0")
	if err != nil {
		return nil, err
	}
	_, err = sec.NewKey("y", "0")
	if err != nil {
		return nil, err
	}
	sec2, err := inidata.NewSection("Keystrokes")
	if err != nil {
		return nil, err
	}
	_, err = sec2.NewKey("record", "ctrl,o")
	if err != nil {
		return nil, err
	}
	_, err = sec2.NewKey("deselect", "c")
	if err != nil {
		return nil, err
	}
	err = inidata.SaveTo("config.ini")
	if err != nil {
		return nil, err
	}
	return inidata, nil
}

func mapToConfig(inidata *ini.File) (*Config, error) {
	var cfg Config
	err := inidata.MapTo(&cfg)
	if err != nil {
		return nil, err
	}
	return &cfg, nil
}
