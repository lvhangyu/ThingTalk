package gen

import (
	"testing"

	"Kratos-demo/pkg/db/mysql"
	"gorm.io/gen"
)

func TestGen(t *testing.T) {
	g := gen.NewGenerator(gen.Config{
		OutPath:       "../query",
		ModelPkgPath:  "../model",
		Mode:          gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface, // generate mode
		FieldNullable: true,
	})

	db := mysql.Init("root:123456@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local")
	g.UseDB(db) // reuse your gorm db

	g.ApplyBasic(
		g.GenerateModel("products"),
	)

	// Generate the code
	g.Execute()
}
