package usecase

import (
	"plainslate/backend/dto"
	"plainslate/backend/lib"

	"go.uber.org/zap"
)

func (u *Usecase) WriteLog(req dto.WriteLogReq) {
	var msg string
	fields := []zap.Field{}

	for key, v := range req.LogFields {
		if key == "msg" {
			msg = v.(string)
		} else {
			fields = append(fields, zap.Any(key, v))
		}
	}

	fields = append(fields, zap.String("component", "frontend"))
	lib.CommonLog(u.Ctx, req.Level, msg, fields...)
}
