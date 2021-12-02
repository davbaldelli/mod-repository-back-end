package mysql

import (
	"github.com/davide/ModRepository/models/entities"
	"gorm.io/gorm"
)

type LogRepositoryImpl struct {
	Db *gorm.DB
}

func (l LogRepositoryImpl) SelectAllTrackLogs(premium bool) ([]entities.TrackLog, error) {
	var logs []entities.TrackLog
	if !premium {
		if res := l.Db.Table("track_logs_view").Find(&logs, "premium = ?", premium); res.Error != nil{
			return nil, res.Error
		}
	} else {
		if res := l.Db.Table("track_logs_view").Find(&logs); res.Error != nil{
			return nil, res.Error
		}
	}

	return logs, nil
}

func (l LogRepositoryImpl) SelectAllCarLogs(premium bool) ([]entities.CarLog, error) {
	var logs []entities.CarLog
	if !premium {
		if res := l.Db.Table("car_logs_view").Find(&logs, "premium = ?", premium); res.Error != nil {
			return nil, res.Error
		}
	} else {
		if res := l.Db.Table("car_logs_view").Find(&logs); res.Error != nil {
			return nil, res.Error
		}
	}
	return logs, nil
}



