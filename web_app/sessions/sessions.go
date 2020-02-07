package sessions

import(
	"github.com/gorilla/sessions"
)

var store = sessions.NewCookieStore([]byte("t0p-s3cr3t"))