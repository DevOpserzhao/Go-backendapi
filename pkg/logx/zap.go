package logx

import (
	"github.com/Shopify/sarama"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"log"
	"os"
)

const _mode = "DEBUG"

type KafKaConfig struct {
	Addr        []string
	Enable      bool
	Topics      string
	RequiredAck int
	Mode        string
}

type LogConfig struct {
	Level       string
	Filepath    string
	FileMaxSize int
	FileMaxAge  int
	MaxBackups  int
	Compress    bool
	Kfc         *KafKaConfig
}

type ZapLogger struct {
	Logger *zap.Logger
}

func New(config *LogConfig) *ZapLogger {
	return &ZapLogger{Logger: makeLogger(config)}
}

func makeLogger(c *LogConfig) *zap.Logger {

	highPriority := zap.LevelEnablerFunc(func(level zapcore.Level) bool {
		return level >= zapcore.ErrorLevel
	})

	lowPriority := zap.LevelEnablerFunc(func(level zapcore.Level) bool {
		return level >= zapcore.DebugLevel
	})

	var cores []zapcore.Core

	// 日志文件
	fileWriter := zapcore.AddSync(logFileHook(c))

	// 控制台打印
	consoleDebugging := zapcore.Lock(os.Stdout)
	consoleEncoder := zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig())

	// 开发调试环境控制台输出
	if c.Level == _mode {
		cores = append(cores, zapcore.NewCore(consoleEncoder, consoleDebugging, lowPriority))
	}
	// 日志文件 生产环境 lowPriority => highPriority
	cores = append(cores, zapcore.NewCore(consoleEncoder, fileWriter, lowPriority))

	// 日志 To kafKa
	kafkaHook := NewKafKa(c.Kfc)

	if kafkaHook.Producer != nil {

		kafkaWriter := zapcore.AddSync(kafkaHook)
		kafkaEncoder := zapcore.NewJSONEncoder(zap.NewDevelopmentEncoderConfig())

		var kafkaCore zapcore.Core

		if c.Kfc.Mode == _mode {
			// 开发调试
			kafkaCore = zapcore.NewCore(kafkaEncoder, kafkaWriter, lowPriority)
		} else {
			// 生产环境
			kafkaCore = zapcore.NewCore(kafkaEncoder, kafkaWriter, highPriority)
		}

		cores = append(cores, kafkaCore)

	} else {
		log.Println("\033[1;31;31m=========== KafKa ============\033[0m")
		log.Println("\033[1;34;34m  Please Check KafKa Config\033[0m")
		log.Println("\033[1;31;31m==============================\033[0m")
	}

	core := zapcore.NewTee(cores...)

	return zap.New(core).WithOptions(zap.AddCaller())
}

func logFileHook(c *LogConfig) *lumberjack.Logger {
	hook := &lumberjack.Logger{
		Filename:   c.Filepath,
		MaxSize:    c.FileMaxSize,
		MaxAge:     c.FileMaxAge,
		MaxBackups: c.MaxBackups,
		Compress:   c.Compress,
	}
	return hook
}

type KafKa struct {
	Producer sarama.SyncProducer
	Topic    string
}

func NewKafKa(config *KafKaConfig) *KafKa {
	return &KafKa{Producer: open(config), Topic: config.Topics}
}

func open(c *KafKaConfig) sarama.SyncProducer {
	if !c.Enable || len(c.Addr) == 0 {
		return nil
	}
	kfc := sarama.NewConfig()
	kfc.Producer.RequiredAcks = sarama.WaitForLocal
	kfc.Producer.Partitioner = sarama.NewRandomPartitioner
	kfc.Producer.Return.Successes = true
	kfc.Producer.Return.Errors = true
	syncProducer, err := sarama.NewSyncProducer(c.Addr, kfc)
	if err != nil {
		log.Println("\033[1;31;31m=========== KafKa ============\033[0m")
		log.Println("\033[1;34;34mP	KafKa Connect Failed.....\033[0m")
		log.Println("\033[1;31;31m==============================\033[0m")
		os.Exit(-3)
	}
	return syncProducer
}

func (k *KafKa) Write(data []byte) (n int, err error) {
	msg := &sarama.ProducerMessage{}
	msg.Topic = k.Topic
	msg.Value = sarama.ByteEncoder(data)
	_, _, err = k.Producer.SendMessage(msg)
	if err != nil {
		return 0, err
	}
	return
}
