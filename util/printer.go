package util

import "github.com/fatih/color"



func Print_title(message string){
	color.Blue(message)
}

func Print_subtitle(message string){
	color.Magenta("[ " + message + " ]")
}

func Print_message(message string, params ...interface{}){
	if(params == nil){
		color.White("\t" + message)
	} else {
		color.White("\t" + message, params[0])
	}
}
func Print_success(message string, params ...interface{}){
	if(params == nil){
		color.Green("[ SUCCESS ]\t" + message)
	} else {
		color.Green("[ SUCCESS ]\t" + message, params[0])
	}
}

func Print_error(message string, params ...interface{}){
	if(params == nil){
		color.Red("[ ERROR ]\t" + message)
	} else {
		color.Red("[ ERROR ]\t" + message, params[0])
	}
}
