package logger

import (
	"fmt"
	"os"
)

type Logger struct{
	file *os.File
	isDebug bool
}

func NewLogger(isDebug bool)*Logger{
	return &Logger{
		file:nil,
		isDebug:isDebug,
	}
}

func outfile(isDebug bool)(*os.File,error){

	return nil,nil
}

 func (l *Logger)Fatal(format string,a ...interface{}){
	label := "[FATAL]"
	fmt.Fprintf(l.file,"%s "+format,label,a)
	panic(fmt.Errorf(format,a...))
}

func (l *Logger)Error(format string,a ...interface{}){
	label := "[ERROR]"
	fmt.Fprintf(l.file,"%s "+format,label,a)
}

func (l *Logger)Info(format string,a ...interface{}){
	label := "[INFO]"
	fmt.Fprintf(l.file,"%s "+format,label,a)
}
