package usecase

import (
	"log"
)

func (u *Usecase) WriteLog(v any) {
	log.Println(v)
}
