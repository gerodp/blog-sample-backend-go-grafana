package backendapi

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"testing"

	"github.com/gavv/httpexpect/v2"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-testfixtures/testfixtures/v3"
)

const testUser string = "testint1"
const testUserPass string = "testint1"

func HelperBuildHttpexpect(t *testing.T) *httpexpect.Expect {
	return httpexpect.WithConfig(httpexpect.Config{
		TestName: t.Name(),

		// prepend this url to all requests
		BaseURL: os.Getenv("BACKEND_BASE_URL"),

		// use fatal failures
		Reporter: httpexpect.NewRequireReporter(t),

		// print all requests and responses
		Printers: []httpexpect.Printer{
			httpexpect.NewDebugPrinter(t, true),
		},
	})
}

func HelperLogin(e *httpexpect.Expect) string {
	loginParams := map[string]interface{}{
		"username": testUser,
		"password": testUserPass,
	}

	loginResponse := e.POST("/login").
		WithJSON(loginParams).
		Expect().
		Status(http.StatusOK).JSON().Object()

	return loginResponse.Value("token").String().Raw()
}

var (
	db       *sql.DB
	fixtures *testfixtures.Loader
)

func buildDBConnectionString(user string, pass string, address string, databaseName string) string {
	return fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, pass, address, databaseName)
}

func HelperSetup(m *testing.M) {

	fmt.Println("Helper Setup")
	var err error

	db, err = sql.Open("mysql",
		buildDBConnectionString(os.Getenv("DATABASE_USER"), os.Getenv("DATABASE_PASSWORD"), fmt.Sprintf("%s:%s", os.Getenv("DATABASE_HOST"), os.Getenv("DATABASE_PORT")), os.Getenv("DATABASE_NAME")),
	)

	if err != nil {
		log.Fatalf("Error connecting to DB. Error = %s", err)
	}

	fixtures, err = testfixtures.New(
		testfixtures.Database(db),                   // You database connection
		testfixtures.Dialect("mysql"),               // Available: "postgresql", "timescaledb", "mysql", "mariadb", "sqlite" and "sqlserver"
		testfixtures.Directory("testdata/fixtures"), // The directory containing the YAML files
	)
	if err != nil {
		log.Fatalf("Error connecting to DB. Error = %s", err)
	}

	if err1 := fixtures.Load(); err1 != nil {
		log.Fatalf("Error loading fixtures into DB. Error = %s", err1)
	}
}
