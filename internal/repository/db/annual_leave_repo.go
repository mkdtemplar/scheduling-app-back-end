package db

import "scheduling-app-back-end/internal/repository/interfaces"

func NewAnnualLeaveRepo() interfaces.IAnnualLeaveRepository {
	return &PostgresDB{DB: GetDb()}
}

func (p *PostgresDB) CreateAnnualLeave() {

}
