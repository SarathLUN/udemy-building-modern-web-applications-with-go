package renderers

import (
	"encoding/gob"
	"github.com/SarathLUN/udemy-building-modern-web-applications-with-go/section-09/076-writing-tests-for-our-Render-package/internal/config"
	"github.com/SarathLUN/udemy-building-modern-web-applications-with-go/section-09/076-writing-tests-for-our-Render-package/internal/models"
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
