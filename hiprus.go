// Package hiprus provides a Hipchat hook for the logrus loggin package.
package hiprus

import (
	"net/url"

	"github.com/Sirupsen/logrus"
	"github.com/tbruyelle/hipchat-go/hipchat"
)

const (
	VERSION     = "2.0.0"
	ColorYellow = "yellow"
	ColorRed    = "red"
	ColorGreen  = "green"
	ColorPurple = "purple"
	ColorGray   = "gray"
	ColorRandom = "random"
)

// HiprusHook is a logrus Hook for dispatching messages to the specified
// channel on Hipchat.
type HiprusHook struct {
	// Messages with a log level not contained in this array
	// will not be dispatched. If nil, all messages will be dispatched.
	AcceptedLevels []logrus.Level
	AuthToken      string
	RoomName       string
	// If empty, "Hiprus" will be used.
	Username string
	// If empty, will point to hipchat cloud
	BaseURL string
	c       *hipchat.Client
}

func (hh *HiprusHook) Levels() []logrus.Level {
	if hh.AcceptedLevels == nil {
		return AllLevels
	}
	return hh.AcceptedLevels
}

func (hh *HiprusHook) Fire(e *logrus.Entry) error {
	if hh.c == nil {
		if err := hh.initClient(); err != nil {
			return err
		}
	}

	color := ""
	notify := false
	switch e.Level {
	case logrus.DebugLevel:
		color = ColorPurple
	case logrus.InfoLevel:
		color = ColorGreen
	case logrus.ErrorLevel, logrus.FatalLevel, logrus.PanicLevel:
		color = ColorRed
		notify = true
	default:
		color = ColorYellow
		notify = true
	}

	_, err := hh.c.Room.Notification(hh.RoomName, &hipchat.NotificationRequest{
		From:          hh.Username,
		Message:       e.Message,
		MessageFormat: "text",
		Notify:        notify,
		Color:         color,
	})

	return err
}

func (hh *HiprusHook) initClient() error {
	c := hipchat.NewClient(hh.AuthToken)

	if hh.BaseURL != "" {
		hipchatUrl, _ := url.Parse(hh.BaseURL)
		c.BaseURL = hipchatUrl
	}

	hh.c = c

	if hh.Username == "" {
		hh.Username = "HipRus"
	}

	return nil
}
