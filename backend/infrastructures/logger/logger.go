package logger

import (
	"fmt"
	"os"
)

type Logger struct{
	file *os.File
	outputInfo bool
	isDebug bool
}

func NewLogger(info ,isDebug bool)*Logger{
	f := os.Stdout
	return &Logger{
		file:f,
		outputInfo: info,
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
	if l.outputInfo{
		label := "[INFO]"
		fmt.Fprintf(l.file,"%s "+format,label,a)
	}
}

func (l *Logger)Debug(format string,a ...interface{}){
	if l.isDebug{
		label := "[DEBUG]"
		fmt.Fprintf(l.file,"%s "+format,label,a)
	}
}
