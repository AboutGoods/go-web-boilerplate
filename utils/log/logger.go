package log

import "github.com/sirupsen/logrus"

var context = logrus.Fields{}

func contextLog() *logrus.Entry {

    return logrus.WithFields(context)
}
func AddToContext(key string, value interface{}) {
    context[key] = value
}
func RemoveFromContext(key string) {
    delete(context, key)
}

func ResetContext() {
    context = logrus.Fields{}
}

func PanicOnError(err error, args ...interface{}) {
    if err != nil {
        data := append([]interface{}{err}, args...)
        contextLog().Fatal(data)
    }
}

func Fatal(args ...interface{}) {
    contextLog().Fatal(args...)
}

func Warning(args ...interface{}) {
    contextLog().Warning(args)

}

func Error(args ...interface{}) {
    contextLog().Error(args...)

}

func Panic(args ...interface{}) {
    contextLog().Panic(args...)

}

func Info(args ...interface{}) {
    contextLog().Info(args...)

}

func Debug(args ...interface{}) {
    contextLog().Debug(args...)

}

func Printf(format string, args ...interface{}) {
    contextLog().Printf(format, args)
}

func Trace(args ...interface{}) {
    contextLog().Trace(args...)
}

// Print logs a message at level Info on the standard logger.
func Print(args ...interface{}) {
    contextLog().Print(args...)
}

// Warn logs a message at level Warn on the standard logger.
func Warn(args ...interface{}) {
    contextLog().Warn(args...)
}

// Tracef logs a message at level Trace on the standard logger.
func Tracef(format string, args ...interface{}) {
    contextLog().Tracef(format, args...)
}

// Debugf logs a message at level Debug on the standard logger.
func Debugf(format string, args ...interface{}) {
    contextLog().Debugf(format, args...)
}

// Infof logs a message at level Info on the standard logger.
func Infof(format string, args ...interface{}) {
    contextLog().Infof(format, args...)
}

// Warnf logs a message at level Warn on the standard logger.
func Warnf(format string, args ...interface{}) {
    contextLog().Warnf(format, args...)
}

// Warningf logs a message at level Warn on the standard logger.
func Warningf(format string, args ...interface{}) {
    contextLog().Warningf(format, args...)
}

// Errorf logs a message at level Error on the standard logger.
func Errorf(format string, args ...interface{}) {
    contextLog().Errorf(format, args...)
}

// Panicf logs a message at level Panic on the standard logger.
func Panicf(format string, args ...interface{}) {
    contextLog().Panicf(format, args...)
}

// Fatalf logs a message at level Fatal on the standard logger then the process will exit with status set to 1.
func Fatalf(format string, args ...interface{}) {
    contextLog().Fatalf(format, args...)
}

// Traceln logs a message at level Trace on the standard logger.
func Traceln(args ...interface{}) {
    contextLog().Traceln(args...)
}

// Debugln logs a message at level Debug on the standard logger.
func Debugln(args ...interface{}) {
    contextLog().Debugln(args...)
}

// Println logs a message at level Info on the standard logger.
func Println(args ...interface{}) {
    contextLog().Println(args...)
}

// Infoln logs a message at level Info on the standard logger.
func Infoln(args ...interface{}) {
    contextLog().Infoln(args...)
}

// Warnln logs a message at level Warn on the standard logger.
func Warnln(args ...interface{}) {
    contextLog().Warnln(args...)
}

// Warningln logs a message at level Warn on the standard logger.
func Warningln(args ...interface{}) {
    contextLog().Warningln(args...)
}

// Errorln logs a message at level Error on the standard logger.
func Errorln(args ...interface{}) {
    contextLog().Errorln(args...)
}

// Panicln logs a message at level Panic on the standard logger.
func Panicln(args ...interface{}) {
    contextLog().Panicln(args...)
}

// Fatalln logs a message at level Fatal on the standard logger then the process will exit with status set to 1.
func Fatalln(args ...interface{}) {
    contextLog().Fatalln(args...)
}
