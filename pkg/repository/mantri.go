package repository

import (
	"fmt"
	"rahmanfaisal10/embrio4-service/pkg/model"
	"strings"

	"github.com/labstack/gommon/log"
)

var (
	AMOUNTFIELDMANTRI = 7
)

func (repo *repository) BUlkUpsertMantri(request []*model.Mantri) ([]*model.Mantri, error) {
	addsQuery := make([]string, 0, len(request))

	for _, mantri := range request {
		addsQuery = append(addsQuery, fmt.Sprintf(`(SELECT * FROM mantri m WHERE m.username_pn = %s)`, mantri.UsernamePN))
	}

	if len(request) > 0 {
		query := `INSERT INTO embrio4.mantri
		(id_cabang, username_pn, nama_mantri, unit_kerja, kode_branch, jabatan, description, created_at, updated_at)` + strings.Join(addsQuery, ",") +
			`ON DUPLICATE KEY UPDATE SET 
			id_cabang = :id_cabang, 
			username_pn = :username_pn, 
			nama_mantri = :nama_mantri, 
			unit_kerja = :unit_kerja, 
			kode_branch = :kode_branch, 
			jabatan = :jabatan, 
			description = :description, 
			created_at = NOW(), 
			updated_at = NOW();`

		_, err := repo.db.NamedExec(query, request)
		if err != nil {
			log.Error(err)
			return nil, err
		}

	}

	return request, nil
}
