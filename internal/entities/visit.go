package entities

import(
	"gorm.io/gorm"
	"github.com/oklog/ulid/v2"
	"time"
)

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

type VisitStatus string

const (
	VisitStatusCompleted VisitStatus = "completed"
	VisitStatusAdsPending VisitStatus = "ads_pending"
)

type Visit struct {
	gorm.Model
	ID ulid.ULID
	LinkId  uint
	UserId  uint
	OS      OS
	Browser Browser
	IP      string
	RedirectedAt time.Time
	Status VisitStatus
}
