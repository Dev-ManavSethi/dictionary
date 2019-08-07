package utils

import "log"

func HandleError(err error, Errmsg, Smsg string) {
	if err != nil {
		log.Println(Errmsg)
		log.Fatalln(err)
	} else if err == nil && Smsg != "" {
		log.Println(Smsg)
	}
}
