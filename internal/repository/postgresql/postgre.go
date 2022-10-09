package repository

import (
	"fmt"
	"log"

	"github.com/AdiKhoironHasan/matkul/internal/models"
	"github.com/AdiKhoironHasan/matkul/internal/repository"
	matkulErrors "github.com/AdiKhoironHasan/matkul/pkg/errors"
	"github.com/jmoiron/sqlx"
)

const (
	SaveMahasiswa       = `INSERT INTO kampus.mahasiswas (nama, nim, created_at) VALUES ($1, $2, now()) RETURNING id`
	SaveMahasiswaAlamat = `INSERT INTO kampus.mahasiswa_alamats (jalan, no_rumah, created_at, id_mahasiswas) VALUES ($1,$2, now(), $3)`
	SaveMatkul          = `INSERT INTO kampus.mata_kuliah (id_dosen, nama, created_at) VALUES ($1, $2, now())`
)

var statement PreparedStatement

type PreparedStatement struct {
	saveMatkul *sqlx.Stmt //membungkus query untuk melindungi dari sql inject
}

type PostgreSQLRepo struct {
	Conn *sqlx.DB
}

func NewRepo(Conn *sqlx.DB) repository.Repository {

	repo := &PostgreSQLRepo{Conn}
	InitPreparedStatement(repo)
	return repo
}

func (p *PostgreSQLRepo) Preparex(query string) *sqlx.Stmt {
	statement, err := p.Conn.Preparex(query)
	if err != nil {
		log.Fatalf("Failed to preparex query: %s. Error: %s", query, err.Error())
	}

	return statement
}

func InitPreparedStatement(m *PostgreSQLRepo) {
	statement = PreparedStatement{
		saveMatkul: m.Preparex(SaveMatkul),
	}
}

func (p *PostgreSQLRepo) SaveMatkul(dataMahasiswa *models.MatkulModels) error {
	result, err := statement.saveMatkul.Exec(dataMahasiswa.ID, dataMahasiswa.Name)

	if err != nil {
		log.Println("Failed Query SaveMatkul : ", err.Error())
		return fmt.Errorf(matkulErrors.ErrorDB)
	}

	rows, err := result.RowsAffected()

	if err != nil {
		log.Println("Failed RowAffectd SaveMatkul : ", err.Error())
		return fmt.Errorf(matkulErrors.ErrorDB)
	}

	if rows < 1 {
		log.Println("SaveMatkul: No Data Changed")
		return fmt.Errorf(matkulErrors.ErrorNoDataChange)
	}

	return nil
}
