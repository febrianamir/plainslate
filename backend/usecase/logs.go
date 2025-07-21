package usecase

import (
	"log"
)

func (u *Usecase) WriteLog(s string) {
	log.Println(s)
}
