// Package hiprus provides a Hipchat hook for the logrus loggin package.
package hiprus

import (
	"github.com/Sirupsen/logrus"
	"github.com/andybons/hipchat"
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
	c        *hipchat.Client
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
	switch e.Level {
	case logrus.DebugLevel:
		color = hipchat.ColorPurple
	case logrus.InfoLevel:
		color = hipchat.ColorGreen
	case logrus.ErrorLevel, logrus.FatalLevel, logrus.PanicLevel:
		color = hipchat.ColorRed
	default:
		color = hipchat.ColorYellow
	}

	return hh.c.PostMessage(hipchat.MessageRequest{
		RoomId:        hh.RoomName,
		From:          hh.Username,
		Message:       e.Message,
		MessageFormat: "text",
		Notify:        true,
		Color:         color,
	})
}

func (hh *HiprusHook) initClient() error {
	hh.c = &hipchat.Client{hh.AuthToken}

	if hh.Username == "" {
		hh.Username = "HipRus"
	}

	return nil
}
