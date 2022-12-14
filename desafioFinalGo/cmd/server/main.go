package main

import (
	"desafioFinal/cmd/server/handler"
	"desafioFinal/docs"
	"desafioFinal/internal/consulta"
	"desafioFinal/internal/dentista"
	"desafioFinal/internal/paciente"
	"desafioFinal/pkg/store"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"os"
)

// @title Store API
// @version 1.0
// @description This API handle products from our store
// @termsOfService https://publicapis.ml/terms

// @contact.name Public APIs
// @contact.url https://dev.publicapis.ml

// @license.name Apache 2.0
// @license.url https://www.apache.org/licenses/LICENSE-2.0.html
func main() {
	//defer config.CloseDatabaseConnection(db)
	err := godotenv.Load(".env")
	if err != nil {
		panic(".Env cant be load")
	}

	sqlStorage := store.NewSQLStoreSql()

	repo := dentista.NewRepository(sqlStorage)
	service := dentista.NewService(repo)
	dentisHandler := handler.NewProductHandler(service)

	sqlStoragePaciente := store.NewSQLStoreSqlPaciente()

	repoPacient := paciente.NewRepository(sqlStoragePaciente)
	servicePacient := paciente.NewService(repoPacient)
	pacientHandler := handler.NewProductHandlerPaciente(servicePacient)

	sqlStorageConsulta := store.NewSQLStoreSqlConsulta()

	repoConsult := consulta.NewRepository(sqlStorageConsulta)
	serviceConsult := consulta.NewService(repoConsult)
	consultHandler := handler.NewConsultHandler(serviceConsult)

	r := gin.Default()
	docs.SwaggerInfo.Host = os.Getenv("HOST")
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	dentistas := r.Group("/dentistas")
	{
		dentistas.GET("", dentisHandler.GetAll())
		dentistas.GET(":id", dentisHandler.GetByID())
		dentistas.POST("", dentisHandler.Post())
		dentistas.DELETE(":id", dentisHandler.Delete())
		dentistas.PATCH(":id", dentisHandler.Patch())
		dentistas.PUT(":id", dentisHandler.Put())
	}
	pacientes := r.Group("/pacientes")
	{
		pacientes.GET("", pacientHandler.GetAll())
		pacientes.GET(":id", pacientHandler.GetByID())
		pacientes.POST("", pacientHandler.Post())
		pacientes.DELETE(":id", pacientHandler.Delete())
		pacientes.PATCH(":id", pacientHandler.Patch())
		pacientes.PUT(":id", pacientHandler.Put())
	}
	consultas := r.Group("/consultas")
	{
		consultas.GET("/consulta_paciente/:rg", consultHandler.GetByRg())
		consultas.GET(":id", consultHandler.GetByID())
		consultas.POST("", consultHandler.Post())
		consultas.DELETE(":id", consultHandler.Delete())
		consultas.PATCH(":id", consultHandler.Patch())
		consultas.PUT(":id", consultHandler.Put())
	}

	r.Run(":8080")
}
