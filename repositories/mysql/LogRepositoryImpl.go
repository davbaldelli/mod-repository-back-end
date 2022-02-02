package mysql

import (
	"github.com/davide/ModRepository/models/entities"
	"gorm.io/gorm"
)

type LogRepositoryImpl struct {
	Db *gorm.DB
}

func (l LogRepositoryImpl) SelectAllTrackLogs(_ bool) ([]entities.TrackLog, error) {
	var logs []entities.TrackLog

	if res := l.Db.Table("track_logs_view").Find(&logs); res.Error != nil {
		return nil, res.Error
	}

	return logs, nil
}

func (l LogRepositoryImpl) SelectAllCarLogs(_ bool) ([]entities.CarLog, error) {
	var logs []entities.CarLog

	if res := l.Db.Table("car_logs_view").Find(&logs); res.Error != nil {
		return nil, res.Error
	}

	return logs, nil
}
