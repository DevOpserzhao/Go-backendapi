package logx

import "testing"

func TestNew(t *testing.T) {
	lc := &LogConfig{
		Level:       "DEBUG",
		Filepath:    "./test.log",
		FileMaxSize: 500,
		FileMaxAge:  30,
		MaxBackups:  10,
		Compress:    false,
		Kfc:         &KafKaConfig{
			Addr:        nil,
			Enable:      false,
			Topics:      "",
			RequiredAck: 1,
			Mode:        "",
		},
	}
	log := New(lc)
	log.Logger.Info("test INFO log....")
	log.Logger.Debug("test DEBUG log....")
	log.Logger.Error("test ERROR log....")
	log.Logger.Warn("test WARN log....")
}
