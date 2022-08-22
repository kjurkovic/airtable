package scheduler

import (
	"time"

	"github.com/kjurkovic/airtable/service/auth/datastore"
	"github.com/kjurkovic/airtable/service/auth/util"
)

func StartCleanScheduler(log *util.Logger) {
	ticker := time.NewTicker(5 * time.Hour)
	go func() {
		for range ticker.C {
			affected, err := datastore.RefreshTokenDao.DeleteOutdated()

			if err != nil {
				log.Error(err)
			} else {
				log.Info("Cleanup complete, deleted rows: %d", affected)
			}
		}
	}()
}
