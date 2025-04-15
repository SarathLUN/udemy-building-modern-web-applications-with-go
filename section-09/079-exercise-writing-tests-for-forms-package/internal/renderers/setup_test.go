package renderers

import (
	"encoding/gob"
	"github.com/SarathLUN/udemy-building-modern-web-applications-with-go/section-09/079-exercise-writing-tests-for-forms-package/internal/config"
	"github.com/SarathLUN/udemy-building-modern-web-applications-with-go/section-09/079-exercise-writing-tests-for-forms-package/internal/models"
	"github.com/alexedwards/scs/v2"
	"net/http"
	"os"
	"testing"
	"time"
)

var session *scs.SessionManager
var testApp config.AppConfig

func TestMain(m *testing.M) {

	// register none-primitive data type into session
	gob.Register(models.Reservation{})

	// change this to true when in production mode
	testApp.InProduction = false

	// working on session
	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = false

	testApp.Session = session
	app = &testApp

	os.Exit(m.Run())
}

type myWriter struct {
}

func (*myWriter) Header() http.Header {
	return http.Header{}
}

func (*myWriter) Write(b []byte) (int, error) {
	return len(b), nil
}

func (*myWriter) WriteHeader(statusCode int) {
}
