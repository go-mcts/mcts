// Copyright 2021 go-mcts. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package mcts

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	defaultLogger Logger
)

func init() {
	// TODO: parse env config and add file log
	cfg := zap.NewDevelopmentConfig()
	cfg.EncoderConfig.EncodeTime = zapcore.RFC3339TimeEncoder
	cfg.EncoderConfig.EncodeCaller = nil
	zapLogger, _ := cfg.Build()
	defaultLogger = zapLogger.Sugar()
}

// LogErr prints err if it's not nil.
func LogErr(err error) {
	if err != nil {
		defaultLogger.Errorf("error occurs during runtime, %v", err)
	}
}

// Debugf logs messages at DEBUG level.
func Debugf(format string, args ...interface{}) {
	defaultLogger.Debugf(format, args...)
}

// Infof logs messages at INFO level.
func Infof(format string, args ...interface{}) {
	defaultLogger.Infof(format, args...)
}

// Warnf logs messages at WARN level.
func Warnf(format string, args ...interface{}) {
	defaultLogger.Warnf(format, args...)
}

// Errorf logs messages at ERROR level.
func Errorf(format string, args ...interface{}) {
	defaultLogger.Errorf(format, args...)
}

// Fatalf logs messages at FATAL level.
func Fatalf(format string, args ...interface{}) {
	defaultLogger.Errorf(format, args...)
}

// Logger is used for logging formatted messages.
type Logger interface {
	// Debugf logs messages at DEBUG level.
	Debugf(format string, args ...interface{})
	// Infof logs messages at INFO level.
	Infof(format string, args ...interface{})
	// Warnf logs messages at WARN level.
	Warnf(format string, args ...interface{})
	// Errorf logs messages at ERROR level.
	Errorf(format string, args ...interface{})
	// Fatalf logs messages at FATAL level.
	Fatalf(format string, args ...interface{})
}
