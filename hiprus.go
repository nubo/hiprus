package hiprus

import (
	"github.com/Sirupsen/logrus"
	"github.com/andybons/hipchat"
)

type HiprusHook struct {
	AuthToken string
	RoomName  string
	Username  string
	c         *hipchat.Client
}

func (hh *HiprusHook) Levels() []logrus.Level {
	return []logrus.Level{
		logrus.DebugLevel,
		logrus.InfoLevel,
		logrus.WarnLevel,
		logrus.ErrorLevel,
		logrus.FatalLevel,
		logrus.PanicLevel,
	}
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
	case logrus.FatalLevel, logrus.PanicLevel:
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
