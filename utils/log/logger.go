// Package log provides global context to logrus logging library
package log

import "github.com/sirupsen/logrus"

var context = logrus.Fields{}

func contextLog() *logrus.Entry {

    return logrus.WithFields(context)
}
// AddToContext add elements to a global context
// It will add the key-value pair when calling the other functions available
func AddToContext(key string, value interface{}) {
    context[key] = value
}
// RemoveFromContext remove an element from the global context
// It will remove on next call from the global context the key passed as argument
func RemoveFromContext(key string) {
    delete(context, key)
}

// ResetContext will reset the global context of the logger
func ResetContext() {
    context = logrus.Fields{}
}

// PanicOnError stops the program and logs it if the error passed as the first argument is not nil.

func PanicOnError(err error, args ...interface{}) {
    if err != nil {
        data := append([]interface{}{err}, args...)
        Fatal(data)
    }
}

// Fatal is the global Contextualisation of logrus' eponym function
func Fatal(args ...interface{}) {
    contextLog().Fatal(args...)
}

// Warning is the global Contextualisation of logrus' eponym function
func Warning(args ...interface{}) {
    contextLog().Warning(args)

}

// Error is the global Contextualisation of logrus' eponym function
func Error(args ...interface{}) {
    contextLog().Error(args...)

}

// Panic is the global Contextualisation of logrus' eponym function
func Panic(args ...interface{}) {
    contextLog().Panic(args...)

}

// Info is the global Contextualisation of logrus' eponym function
func Info(args ...interface{}) {
    contextLog().Info(args...)

}

// Debug is the global Contextualisation of logrus' eponym function
func Debug(args ...interface{}) {
    contextLog().Debug(args...)

}

// Printf is the global Contextualisation of logrus' eponym function
func Printf(format string, args ...interface{}) {
    contextLog().Printf(format, args)
}

// Trace is the global Contextualisation of logrus' eponym function
func Trace(args ...interface{}) {
    contextLog().Trace(args...)
}

// Print is the global Contextualisation of logrus' eponym function.
func Print(args ...interface{}) {
    contextLog().Print(args...)
}

// Warn is the global Contextualisation of logrus' eponym function.
func Warn(args ...interface{}) {
    contextLog().Warn(args...)
}

// Tracef is the global Contextualisation of logrus' eponym function.
func Tracef(format string, args ...interface{}) {
    contextLog().Tracef(format, args...)
}

// Debugf is the global Contextualisation of logrus' eponym function.
func Debugf(format string, args ...interface{}) {
    contextLog().Debugf(format, args...)
}

// Infof is the global Contextualisation of logrus' eponym function.
func Infof(format string, args ...interface{}) {
    contextLog().Infof(format, args...)
}

// Warnf is the global Contextualisation of logrus' eponym function.
func Warnf(format string, args ...interface{}) {
    contextLog().Warnf(format, args...)
}

// Warningf is the global Contextualisation of logrus' eponym function.
func Warningf(format string, args ...interface{}) {
    contextLog().Warningf(format, args...)
}

// Errorf is the global Contextualisation of logrus' eponym function.
func Errorf(format string, args ...interface{}) {
    contextLog().Errorf(format, args...)
}

// Panicf is the global Contextualisation of logrus' eponym function.
func Panicf(format string, args ...interface{}) {
    contextLog().Panicf(format, args...)
}

// Fatalf is the global Contextualisation of logrus' eponym function
func Fatalf(format string, args ...interface{}) {
    contextLog().Fatalf(format, args...)
}

// Traceln is the global Contextualisation of logrus' eponym function.
func Traceln(args ...interface{}) {
    contextLog().Traceln(args...)
}

// Debugln is the global Contextualisation of logrus' eponym function.
func Debugln(args ...interface{}) {
    contextLog().Debugln(args...)
}

// Println is the global Contextualisation of logrus' eponym function.
func Println(args ...interface{}) {
    contextLog().Println(args...)
}

// Infoln is the global Contextualisation of logrus' eponym function.
func Infoln(args ...interface{}) {
    contextLog().Infoln(args...)
}

// Warnln is the global Contextualisation of logrus' eponym function.
func Warnln(args ...interface{}) {
    contextLog().Warnln(args...)
}

// Warningln is the global Contextualisation of logrus' eponym function.
func Warningln(args ...interface{}) {
    contextLog().Warningln(args...)
}

// Errorln is the global Contextualisation of logrus' eponym function.
func Errorln(args ...interface{}) {
    contextLog().Errorln(args...)
}

// Panicln is the global Contextualisation of logrus' eponym function.
func Panicln(args ...interface{}) {
    contextLog().Panicln(args...)
}

// Fatalln is the global Contextualisation of logrus' eponym function
func Fatalln(args ...interface{}) {
    contextLog().Fatalln(args...)
}
