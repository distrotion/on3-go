package main

import (
	_ "encoding/json"
	"fmt"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	// cors "github.com/itsjamie/gin-cors"
)

type Tokenschema struct {
	Token string `json:"token"`
}

type Gqlquery struct {
	Query string `json:"query"`
}

func main() {
	r := gin.Default()
	// r := gin.New()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"PUT", "POST"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return origin == "https://github.com"
		},
		MaxAge: 12 * time.Hour,
	}))

	// r.Use(cors.Middleware(cors.Config{
	// 	Origins:         "*",
	// 	Methods:         "GET, PUT, POST, DELETE",
	// 	RequestHeaders:  "Origin, Authorization, Content-Type",
	// 	ExposedHeaders:  "",
	// 	MaxAge:          50 * time.Second,
	// 	Credentials:     true,
	// 	ValidateHeaders: false,
	// }))

	r.GET("/", func(c *gin.Context) {
		fmt.Println("graphqlResponse")

		c.JSON(200, "Cloud Test Service on Amazon Elastic Beanstalk")
	})

	// r.POST("/getpatientlist", func(c *gin.Context) {
	// 	var input model.GetListData
	// 	c.BindJSON(&input)
	// 	// // c.ShouldBind(&input)
	// 	// fmt.Println(input)

	// 	reg := getlisdata.PatientListData(c, input, staffDB, staffDB_clinic_profile, staffDB_clinic_profile_model, clinicDB, clinicDB_clinic_profile, clinicDB_clinic_profile_model, authDB, authDB_profile, authDB_profile_model)

	// 	//=========================================
	// 	c.JSON(200, reg)
	// })

	r.Run(":9760")
	// r.RunTLS(":9105", "./testdata/server.pem", "./testdata/server.key")
}

var (
	clinicDB                      = `clinic`
	clinicDB_clinic_profile_model = `clinic_Profile_model`
	clinicDB_clinic_profile       = `clinic_Profile_model`
	staffDB                       = `auth_staff`
	staffDB_clinic_profile_model  = `staff_Profile_model`
	staffDB_clinic_profile        = `staff_Profile_model`
	authDB                        = `auth_main`
	authDB_profile_model          = `patient_Profile_model`
	authDB_profile                = `patient_Profile_model`
	caregiverDB                   = `auth_caregiver`
	caregiverDB_profile_model     = `caregiver_Profile_model`
	caregiverDB_profile           = `caregiver_Profile_model`
)
