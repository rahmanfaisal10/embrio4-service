package util

import (
	"fmt"
	"sync"

	"github.com/jmoiron/sqlx"
	"github.com/labstack/gommon/log"
)

func DispatchWorkers(db *sqlx.DB, jobs <-chan []interface{}, wg *sync.WaitGroup) {
	for workerIndex := 0; workerIndex <= TOTALWORKER; workerIndex++ {
		go func(workerIndex int, db *sqlx.DB, jobs <-chan []interface{}, wg *sync.WaitGroup) {
			counter := 0

			for job := range jobs {
				DoTheJob(workerIndex, counter, db, job)
				wg.Done()
				counter++
			}
		}(workerIndex, db, jobs, wg)
	}
}

func DoTheJob(workerIndex, counter int, db *sqlx.DB, values []interface{}) {
	for {
		var outerError error
		func(outerError *error) {
			defer func() {
				if err := recover(); err != nil {
					*outerError = fmt.Errorf("%v", err)
				}
			}()

			query := fmt.Sprintf(`
				INSERT INTO embrio4.upload
				(
					Periode,
					Region,
					`+"`Main Branch`"+`, 
					Branch,
					Currency, 
					`+"`Nama AO`"+`, 
					`+"`LN Type`"+`, 
					`+"`Nomor rekening`"+`, 
					`+"`Nama Debitur`"+`, 
					`+"`Alamat Identitas`"+`, 
					`+"`Kode Pos Identitas`"+`, 
					`+"`Alamat Kantor`"+`, 
					`+"`Kode Pos Kantor`"+`, 
					Plafond, 
					`+"`Next Pmt Date`"+`, 
					`+"`Next Int Pmt Date`"+`, 
					Rate, 
					`+"`Tgl Menunggak`"+`,
					`+"`Tgl Realisasi`"+`, 
					`+"`Tgl Jatuh tempo`"+`, 
					`+"`Jangka Waktu`"+`, 
					`+"`Flag Restruk`"+`, 
					CIFNO, 
					`+"`Kolektibilitas Lancar`"+`, 
					`+"`Kolektibilitas DPK`"+`, 
					`+"`Kolektibilitas Kurang Lancar`"+`, 
					`+"`Kolektibilitas Diragukan`"+`, 
					`+"`Kolektibilitas Macet`"+`, 
					`+"`Tunggakan Pokok`"+`, 
					`+"`Tunggakan Bunga`"+`, 
					`+"`Tunggakan Pinalty`"+`, 
					`+"`PN   PENGELOLA`"+`, 
					`+"`NAMA  PENGELOLA`"+`
				)
				VALUES %s;
			`, `(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)`)

			_, err := db.Exec(query, values...)
			if err != nil {
				log.Error(err)
				return
			}

			err = db.Close()
			if err != nil {
				log.Error(err)
				return
			}
		}(&outerError)
		if outerError == nil {
			break
		}
	}

	if counter%100 == 0 {
		fmt.Println("=> worker", workerIndex, "inserted", counter, "data")
	}
}
