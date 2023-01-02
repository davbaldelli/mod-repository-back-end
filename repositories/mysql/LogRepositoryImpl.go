package mysql

import (
	"github.com/davide/ModRepository/models"
	"gorm.io/gorm"
)

type LogRepositoryImpl struct {
	Db *gorm.DB
}

func (l LogRepositoryImpl) SelectAllTrackLogs() ([]models.TrackLog, error) {
	var logs []models.TrackLog

	if res := l.Db.Table("track_logs_view").Find(&logs); res.Error != nil {
		return nil, res.Error
	}

	return logs, nil
}

func (l LogRepositoryImpl) SelectAllCarLogs() ([]models.CarLog, error) {
	var logs []models.CarLog

	if res := l.Db.Table("car_logs_view").Find(&logs); res.Error != nil {
		return nil, res.Error
	}

	return logs, nil
}
