package logger

import (
	"fmt"
	"os"
	"runtime"
)

type Logger struct{
	file *os.File
	outputInfo bool
	isDebug bool
}

func NewLogger(info ,debug bool,filePath string)*Logger{
	f := outfile(filePath)
	return &Logger{
		file:f,
		outputInfo: info,
		isDebug:debug,
	}
}

func (l *Logger)Close(){
	err := l.file.Close()
	if err !=nil{
		fmt.Fprintf(os.Stderr,"file close error : %#v",err)
	}
}

func outfile(filePath string)*os.File{
	if filePath=="" || filePath == "stderr"{
		return os.Stderr
	}
	if filePath == "stdout"{
		return os.Stdout
	}
	file, err := os.Open(filePath)
    if err != nil {
		fmt.Fprintf(os.Stderr,"faild to open file. output error message to stderr : %#v",err)
		return os.Stderr
	}
	return file
}

func (l *Logger)Fatal(format string,a ...interface{}){
	label := "[FATAL]"
	_, file, line, _ := runtime.Caller(1)
	preMsg := fmt.Sprintf("%s called from:%s:%d",label,file,line)
	fmt.Fprintf(l.file,"%s\n"+format,preMsg,a)
	panic(fmt.Errorf(format,a...))
}


func (l *Logger)Error(format string,a ...interface{}){
	label := "[ERROR]"
	fmt.Fprintf(l.file,"%s "+format,label,a)
}

func (l *Logger)Info(format string,a ...interface{}){
	if l.outputInfo{
		label := "[INFO] "
		fmt.Fprintf(l.file,"%s "+format,label,a)
	}
}

func (l *Logger)Debug(format string,a ...interface{}){
	if l.isDebug{
		label := "[DEBUG]"
		fmt.Fprintf(l.file,"%s "+format,label,a)
	}
}
