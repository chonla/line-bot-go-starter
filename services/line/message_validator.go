package line

import (
	"bytes"
	"io"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (l *Line) MessageValidator(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var body []byte
		if c.Request().Body != nil {
			// Read request body, this will close io reader automatically and the body will be gone.
			body, _ = io.ReadAll(c.Request().Body)
			// So we need to put the body data back to Request
			c.Request().Body = io.NopCloser(bytes.NewBuffer(body))
		}

		signature := []byte(c.Request().Header.Get("x-line-signature"))
		if len(signature) == 0 {
			return c.NoContent(http.StatusForbidden)
		}

		if l.VerifyMessage(body, signature) {
			return next(c)
		}
		return c.NoContent(http.StatusForbidden)
	}
}
