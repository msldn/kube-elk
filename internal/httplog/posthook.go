// Package lrhook provides logrus hook for the Slack.
//
// It can post messages to slack based on the notification level of the
// logrus entry including the ability to rate limit messages.
//
// See: https://godoc.org/github.com/sirupsen/logrus#Hook
package httplog

import (
	"github.com/sirupsen/logrus"
)

var (
	// DefaultLevelColors is the default level colors used if none are present in the configuration.
	DefaultLevelColors = map[string]string{
		"debug":   "#9B30FF",
		"info":    "good",
		"warning": "danger",
		"error":   "danger",
		"fatal":   "panic",
		"panic":   "panic",
	}

	// DefaultUnknownColor is the default UnknownColor if one is not present in the configuration.
	DefaultUnknownColor = "warning"
)

// Config is the configuration of a slack logrus.Hook.
type Config struct {
	// MinLevel is the minimal level at which the hook will trigger.
	MinLevel logrus.Level

	// LevelColors is a hash of logrus level names to colors used for the attachment in the messages.
	LevelColors map[string]string

	// AttachmentText is the text message used for the attachment when fields are present in the log entry.
	AttachmentText string

	// UnknownColor is the color to use if there is no match for the log level in LevelColors.
	UnknownColor string

	// Async if true then messages are sent to slack asynchronously.
	// This means that Fire will never return an error.
	Async bool

	// Burst sets the burst limit.
	// Ignored if Limit is zero.
	Burst int

}

// Hook is a logrus hook that sends messages to Slack.
type Hook struct {
	Config
	client  Client
}

// SetConfigDefaults sets defaults on the configuration if needed to ensure the cfg is valid.
func SetConfigDefaults(cfg *Config) {
	if len(cfg.LevelColors) == 0 {
		cfg.LevelColors = DefaultLevelColors
	}
	if cfg.UnknownColor == "" {
		cfg.UnknownColor = DefaultUnknownColor
	}
}

// New returns a new Hook with the given configuration that posts messages using the webhook URL.
// It ensures that the cfg is valid by calling SetConfigDefaults on the cfg.
func NewHook(cfg Config, url string) *Hook {
	return NewClient(cfg, New(url))
}

// NewClient returns a new Hook with the given configuration using the slack.Client c.
// It ensures that the cfg is valid by calling SetConfigDefaults on the cfg.
func NewClient(cfg Config, client *Client) *Hook {
	SetConfigDefaults(&cfg)

	c := &Hook{Config: cfg, client: *client}

	return c
}

// Levels implements logrus.Hook.
// It returns the logrus.Level's that are lower or equal to that of MinLevel.
// This means setting MinLevel to logrus.ErrorLevel will send slack messages for log entries at Error, Fatal and Panic.
func (sh *Hook) Levels() []logrus.Level {
	lvls := make([]logrus.Level, 0, len(logrus.AllLevels))
	for _, l := range logrus.AllLevels {
		if sh.MinLevel >= l {
			lvls = append(lvls, l)
		}
	}

	return lvls
}

// Fire implements logrus.Hook.
// It sends a slack message for the log entry e.
func (sh *Hook) Fire(e *logrus.Entry) error {
	err := sh.client.Send(e.Message,nil)
	return err
}