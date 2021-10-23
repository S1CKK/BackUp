package main

import (
	"github.com/S1CKK/med-record-system/controller"

	"github.com/S1CKK/med-record-system/entity"

	"github.com/gin-gonic/gin"
)

func main() {

	entity.SetupDatabase()
	/*gin.SetMode(gin.ReleaseMode)*/
	r := gin.Default()
	r.Use(CORSMiddleware())

	// Pharmacist Routes
	r.GET("/pharmacists", controller.ListPharmacists)
	r.GET("/pharmacist/:id", controller.GetPharmacist)
	r.POST("/pharmacists", controller.CreatePharmacist)
	r.PATCH("/pharmacists", controller.UpdatePharmacist)
	r.DELETE("/pharmacists/:id", controller.DeletePharmacist)

	// Medicine Routes
	r.GET("/medicines", controller.ListMedicine)
	r.GET("/medicine/:id", controller.GetMedicine)
	r.POST("/medicines", controller.CreateMedicine)
	r.PATCH("/medicines", controller.UpdateMedicine)
	r.DELETE("/medicines/:id", controller.DeleteMedicine)

	// Admission Routes
	r.GET("/admissions", controller.ListAdmission)
	r.GET("/admission/:id", controller.GetAdmission)
	//protected.GET("/admission/watched/user/:id", controller.GetPlaylistWatchedByUser)
	r.POST("/admissions", controller.CreateAdmission)
	r.PATCH("/admissions", controller.UpdateAdmission)
	r.DELETE("/admissions/:id", controller.DeleteAdmission)

	// Treatment Routes
	r.GET("/treatment_records", controller.ListTreatmentRecord)
	r.GET("/admission/treatments", controller.GetAdmissionByTreatment)
	r.GET("/treatment_record/:id", controller.GetTreatmentRecord)
	r.POST("/treatment_records", controller.CreateTreatmentRecord)
	r.PATCH("/treatment_records", controller.UpdateTreatmentRecord)
	r.DELETE("/treatments_records/:id", controller.DeleteTreatmentRecord)

	// MedRecord Routes
	r.GET("/medication_records", controller.ListMedicationRacord)
	r.GET("/medication_record/:id", controller.GetMedicationRacord)
	r.POST("/medication_records", controller.CreateMedicationRecord)
	r.PATCH("/medication_records", controller.UpdateMedicationRacord)
	r.DELETE("/medication_records/:id", controller.DeleteMedicationRacord)

	// User Routes
	//r.POST("/users", controller.CreateUser)

	// Authentication Routes
	//r.POST("/login", controller.Login)

	// Run the server
	r.Run()

}

func CORSMiddleware() gin.HandlerFunc {

	return func(c *gin.Context) {

		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()

	}

}
