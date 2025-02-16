package initialize

import (
	"DeliciousTown/config"
	"DeliciousTown/global"
	"DeliciousTown/utils"
	"fmt"
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"path"
	"strings"
	"time"
)

const (
	outJson = "json"
)

func InitLogger() {
	fmt.Printf("用户参数:%v\n", global.GvaConfig.UserLog)
	fmt.Printf("默认参数:%v\n", global.GvaConfig.Log)
	global.UserLogger = GetLogger(global.GvaConfig.UserLog)
	global.DefaultLogger = GetLogger(global.GvaConfig.Log)
}

func GetLogger(logConfig config.Log) *zap.Logger {
	if exist, _ := utils.DirExist(logConfig.Path); !exist {
		_ = utils.CreateDir(logConfig.Path)
	}

	var encoder zapcore.Encoder
	if logConfig.OutFormat == outJson {
		encoder = zapcore.NewJSONEncoder(getEncoderConfig())
	} else {
		encoder = zapcore.NewConsoleEncoder(getEncoderConfig())
	}

	writeSyncer := zapcore.AddSync(getLumberjackWriteSyncer(logConfig))
	zapCore := zapcore.NewCore(encoder, writeSyncer, getLevel(logConfig))
	logger := zap.New(zapCore)
	defer logger.Sync()
	return logger
}

func getLevel(logConfig config.Log) zapcore.Level {
	levelMap := map[string]zapcore.Level{
		"debug":  zapcore.DebugLevel,
		"info":   zapcore.InfoLevel,
		"warn":   zapcore.WarnLevel,
		"error":  zapcore.ErrorLevel,
		"dpanic": zapcore.DPanicLevel,
		"panic":  zapcore.PanicLevel,
		"fatal":  zapcore.FatalLevel,
	}
	if level, ok := levelMap[logConfig.Level]; ok {
		return level
	}

	return zapcore.InfoLevel
}

func getEncoderConfig() zapcore.EncoderConfig {
	// Keys can be anything except the empty string.
	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		FunctionKey:    zapcore.OmitKey,
		MessageKey:     "msg",
		StacktraceKey:  "S",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.CapitalLevelEncoder,
		EncodeTime:     getEncodeTime,
		EncodeDuration: zapcore.StringDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}
	return encoderConfig
}

func getEncodeTime(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format("2006-01-02 15:04:05"))
}

func getLumberjackWriteSyncer(logConfig config.Log) zapcore.WriteSyncer {
	lumberjackConfig := logConfig.LumberJack
	lumberjackLogger := &lumberjack.Logger{
		Filename:   getLogFile(logConfig),       //日志文件
		MaxSize:    lumberjackConfig.MaxSize,    //单文件最大容量(单位MB)
		MaxBackups: lumberjackConfig.MaxBackups, //保留旧文件的最大数量
		MaxAge:     lumberjackConfig.MaxAge,     // 旧文件最多保存几天
		Compress:   lumberjackConfig.Compress,   // 是否压缩/归档旧文件
	}

	return zapcore.AddSync(lumberjackLogger)
}

func getLogFile(logConfig config.Log) string {
	fileFormat := time.Now().Format(logConfig.FileFormat)
	fileName := strings.Join([]string{
		logConfig.FilePrefix,
		fileFormat,
		"log"}, ".")
	return path.Join(logConfig.Path, fileName)
}
