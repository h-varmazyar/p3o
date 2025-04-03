package entities

import "gorm.io/gorm"

type OS string

const (
	OSLinux        OS = "linux"
	OSWindows      OS = "windows"
	OSMac          OS = "mac"
	OSAndroid      OS = "android"
	IOS            OS = "ios"
	OSWindowsPhone OS = "windows-phone"
)

type Browser string

const (
	BrowserGoogleChrome     Browser = "google-chrome"
	BrowserMozilla          Browser = "mozilla"
	BrowserEdge             Browser = "edge"
	BrowserInternetExplorer Browser = "internet-explorer"
)

type Visit struct {
	gorm.Model
	LinkId  uint
	UserId  uint
	OS      OS
	Browser Browser
	IP      string
}
