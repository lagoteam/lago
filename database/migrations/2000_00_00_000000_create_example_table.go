package migrations

import (
	"database/sql"
	"goravel/modules/lago/models"
	"goravel/packages/migrate"
	"gorm.io/gorm"
	"runtime"
)

func init() {
	type Example struct {
		models.BaseModel
		Name     string `gorm:"type:varchar(255);not null;index"`
		Email    string `gorm:"type:varchar(255);index;default:null"`
		Phone    string `gorm:"type:varchar(20);index;default:null"`
		Password string `gorm:"type:varchar(255)"`
	}

	up := func(migrator gorm.Migrator, DB *sql.DB) {
		migrator.AutoMigrate(&Example{})
	}

	down := func(migrator gorm.Migrator, DB *sql.DB) {
		migrator.DropTable(&Example{})
	}

	/*
		// 获取当前文件名称
		_, fullFilename, _, _ := runtime.Caller(0)
		fmt.Println(fullFilename)
		var filename string
		filename = path.Base(fullFilename)
		fmt.Println("filename=", filename)
	*/

	_, fullFilename, _, _ := runtime.Caller(0)
	migrate.AddMigration(migrate.GetMigrationFileName(fullFilename), up, down)
}
