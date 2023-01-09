package middleware

import (
	"api-courseroom/entities"
	"api-courseroom/libraries"
	"api-courseroom/models"
	"bytes"
	"crypto/tls"
	"html/template"
	"net/http"
	"os"
	"path/filepath"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
	jsoniter "github.com/json-iterator/go"
	"github.com/k3a/html2text"
	"gopkg.in/gomail.v2"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

type Middleware struct {
	DB                    *gorm.DB
	Validator             *validator.Validate
	SECRET_TOKEN          string
	EMAIL_SERVER          string
	EMAIL_PORT            int
	EMAIL_ADDRESS         string
	EMAIL_CREDENTIALS     string
	EMAIL_VERIFICATOR_API string
	QR_SERVER_API         string
	COURSEROOM_CALCULATOR string
}

func NewMiddleware() *Middleware {

	// Cargar archivo .env
	err := godotenv.Load(".env")

	if err != nil {
		return nil
	}

	//Cargar variables:

	server := os.Getenv("SERVER")
	user := os.Getenv("USER")
	password := os.Getenv("PASSWORD")
	databaseName := os.Getenv("DATABASE")
	secretToken := os.Getenv("SECRET_TOKEN")
	emailServer := os.Getenv("EMAIL_SERVER")
	emailAddress := os.Getenv("EMAIL_ADDRESS")
	emailCredentials := os.Getenv("EMAIL_CREDENTIALS")
	emailAPI := os.Getenv("EmailVerificatorAPI")
	emailPort, _ := strconv.Atoi(os.Getenv("EMAIL_PORT"))
	qrAPI := os.Getenv("QRServerAPI")
	courseRoomCalculator := os.Getenv("COURSEROOM_CALCULATOR")

	dsn := "sqlserver://" + user + ":" + password + "@" + server + "?database=" + databaseName

	db, err := gorm.Open(sqlserver.Open(dsn), &gorm.Config{
		SkipDefaultTransaction: true,
		PrepareStmt:            true,
	})

	if err != nil {
		return nil
	}

	return &Middleware{
		DB:                    db,
		Validator:             validator.New(),
		SECRET_TOKEN:          secretToken,
		EMAIL_SERVER:          emailServer,
		EMAIL_ADDRESS:         emailAddress,
		EMAIL_CREDENTIALS:     emailCredentials,
		EMAIL_PORT:            emailPort,
		EMAIL_VERIFICATOR_API: emailAPI,
		QR_SERVER_API:         qrAPI,
		COURSEROOM_CALCULATOR: courseRoomCalculator}
}

func (middleware *Middleware) ValidateModel(data interface{}) error {
	return middleware.Validator.Struct(data)
}

/*Emails*/

func (middleware *Middleware) SendBienvenidaEmail(data *models.BienvenidaEmail) error {

	smtpPass := middleware.EMAIL_CREDENTIALS
	smtpUser := middleware.EMAIL_ADDRESS
	smtpHost := middleware.EMAIL_SERVER
	smtpPort := middleware.EMAIL_PORT

	var body bytes.Buffer

	template, err := ParseTemplateDir("app_data")
	if err != nil {
		return err
	}

	template.ExecuteTemplate(&body, "bienvenida.html", &data)

	m := gomail.NewMessage()

	m.SetHeader("From", smtpUser)
	m.SetHeader("To", data.CorreoElectronico)
	m.SetHeader("Subject", "Bienvenid@ a la comunidad de CourseRoom®")
	m.SetBody("text/html", body.String())
	m.AddAlternative("text/plain", html2text.HTML2Text(body.String()))

	d := gomail.NewDialer(smtpHost, smtpPort, smtpUser, smtpPass)
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	// Send Email
	if err := d.DialAndSend(m); err != nil {
		return err
	}

	return nil
}

func (middleware *Middleware) SendCalificacionEmail(data *models.CalificacionEmail) error {

	smtpPass := middleware.EMAIL_CREDENTIALS
	smtpUser := middleware.EMAIL_ADDRESS
	smtpHost := middleware.EMAIL_SERVER
	smtpPort := middleware.EMAIL_PORT

	var body bytes.Buffer

	template, err := ParseTemplateDir("app_data")
	if err != nil {
		return err
	}

	template.ExecuteTemplate(&body, "calificacion.html", &data)

	m := gomail.NewMessage()

	m.SetHeader("From", smtpUser)
	m.SetHeader("To", data.CorreoElectronico)
	m.SetHeader("Subject", "Nueva tarea calificada")
	m.SetBody("text/html", body.String())
	m.AddAlternative("text/plain", html2text.HTML2Text(body.String()))

	d := gomail.NewDialer(smtpHost, smtpPort, smtpUser, smtpPass)
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	// Send Email
	if err := d.DialAndSend(m); err != nil {
		return err
	}

	return nil
}

func (middleware *Middleware) SendCredencialesEmail(data *models.CredencialesEmail) error {

	smtpPass := middleware.EMAIL_CREDENTIALS
	smtpUser := middleware.EMAIL_ADDRESS
	smtpHost := middleware.EMAIL_SERVER
	smtpPort := middleware.EMAIL_PORT

	var body bytes.Buffer

	template, err := ParseTemplateDir("app_data")
	if err != nil {
		return err
	}

	template.ExecuteTemplate(&body, "credenciales.html", &data)

	m := gomail.NewMessage()

	m.SetHeader("From", smtpUser)
	m.SetHeader("To", data.CorreoElectronico)
	m.SetHeader("Subject", "Recuperación de credenciales")
	m.SetBody("text/html", body.String())
	m.AddAlternative("text/plain", html2text.HTML2Text(body.String()))

	d := gomail.NewDialer(smtpHost, smtpPort, smtpUser, smtpPass)
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	// Send Email
	if err := d.DialAndSend(m); err != nil {
		return err
	}

	return nil
}

// Email template parser
func ParseTemplateDir(dir string) (*template.Template, error) {
	var paths []string
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			paths = append(paths, path)
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	return template.ParseFiles(paths...)
}

/*Apis*/

func (middleware *Middleware) EmailVerificatorAPI(email *string) entities.AccionEntity {

	var response entities.AccionEntity

	jsonIter := jsoniter.ConfigCompatibleWithStandardLibrary

	query := libraries.FormatString(middleware.EMAIL_VERIFICATOR_API, *email)

	resp, err := http.Get(query)

	if err != nil {
		response = entities.AccionEntity{
			Codigo:  -1,
			Mensaje: err.Error()}
	} else {

		if resp.StatusCode == 200 {

			var modelo *models.EmailVerificatorAPISuccess

			err := jsonIter.NewDecoder(resp.Body).Decode(&modelo)

			if err != nil {
				response = entities.AccionEntity{
					Codigo:  -1,
					Mensaje: err.Error()}
			} else {

				success := modelo.Status

				if success {
					response = entities.AccionEntity{
						Codigo:  1,
						Mensaje: "Ok"}
				} else {
					response = entities.AccionEntity{
						Codigo:  -1,
						Mensaje: "El correo electrónico al que se hace referencia no existe"}
				}

			}

		} else {

			var modelo *models.EmailVerificatorAPIError

			err := jsonIter.NewDecoder(resp.Body).Decode(&modelo)

			if err != nil {
				response = entities.AccionEntity{
					Codigo:  -1,
					Mensaje: err.Error()}
			} else {
				response = entities.AccionEntity{
					Codigo:  -1,
					Mensaje: modelo.Error.Message}
			}
		}
	}

	return response
}
