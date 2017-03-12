package util

import "github.com/fatih/color"



func Print_title(message string){
	color.Blue(message)
}

func Print_subtitle(message string){
	color.Yellow("[ " + message + " ]")
}

func Print_message(message string, params ...interface{}){
	if(params == nil){
		color.White("\t" + message)
	} else {
		color.White("\t" + message, params[0])
	}
}

func Print_error(message string, params ...interface{}){
	if(params == nil){
		color.Red("[ ERROR ]\t" + message)
	} else {
		color.Red("[ ERROR ]\t" + message, params[0])
	}
}
func Print_warning(message string, params ...interface{}){
	if(params == nil){
		color.Magenta("[ WARNING ]\t" + message)
	}else {
		color.Magenta("[ WARNING ]\t" + message, params[0])
	}

}
