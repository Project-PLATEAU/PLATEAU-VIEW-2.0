package opinion

import (
	"bytes"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/jarcoal/httpmock"
	"github.com/labstack/echo/v4"
	"github.com/samber/lo"
	"github.com/stretchr/testify/assert"
)

func TestEcho(t *testing.T) {
	httpmock.Activate()
	defer httpmock.Deactivate()
	httpmock.RegisterResponder("POST", "https://api.sendgrid.com/v3/mail/send", httpmock.NewJsonResponderOrPanic(http.StatusOK, `{}`))

	e := echo.New()
	e.Validator = &customValidator{validator: validator.New()}
	g := e.Group("")
	Echo(g, Config{
		SendGridAPIKey: "xxx",
		From:           "hoge@example.com",
		To:             "hoge@example.com",
	})

	// bad request
	rb := `{"title":"aaa","email":"from@examle.com","content":"","name":"name"}`
	r := httptest.NewRequest("POST", "/", strings.NewReader(rb))
	r.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	e.ServeHTTP(w, r)
	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.NotEmpty(t, strings.TrimSpace(w.Body.String()))

	// application/json
	rb = `{"email":"from@examle.com","content":"aaaa","name":"name","category":"カテゴリ","org":"所属組織"}`
	r = httptest.NewRequest("POST", "/", strings.NewReader(rb))
	r.Header.Set("Content-Type", "application/json")
	w = httptest.NewRecorder()
	e.ServeHTTP(w, r)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, `"ok"`, strings.TrimSpace(w.Body.String()))

	// application/x-www-form-urlencoded
	form := url.Values{}
	form.Add("title", "aaa")
	form.Add("email", "from@examle.com")
	form.Add("content", "aaaa")
	form.Add("name", "name")
	r = httptest.NewRequest("POST", "/", strings.NewReader(form.Encode()))
	r.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	w = httptest.NewRecorder()
	e.ServeHTTP(w, r)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, `"ok"`, strings.TrimSpace(w.Body.String()))

	// multipart/form-data
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	_ = writer.WriteField("name", "NAME")
	_ = writer.WriteField("title", "TITLE")
	_ = writer.WriteField("email", "from@example.com")
	_ = writer.WriteField("content", "CONTENT")
	file := lo.Must(os.Open("testdata/test.jpg"))
	defer func() { _ = file.Close() }()
	part := lo.Must(writer.CreateFormFile("file", filepath.Base(file.Name())))
	_ = lo.Must(io.Copy(part, file))
	lo.Must0(writer.Close())

	r = httptest.NewRequest("POST", "/", body)
	r.Header.Add("Content-Type", writer.FormDataContentType())
	w = httptest.NewRecorder()
	e.ServeHTTP(w, r)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, `"ok"`, strings.TrimSpace(w.Body.String()))

	// multipart/form-data with invalid file
	body = &bytes.Buffer{}
	writer = multipart.NewWriter(body)
	_ = writer.WriteField("name", "NAME")
	_ = writer.WriteField("title", "TITLE")
	_ = writer.WriteField("email", "from@example.com")
	_ = writer.WriteField("content", "CONTENT")
	part = lo.Must(writer.CreateFormFile("file", "aaa.txt"))
	_, _ = part.Write([]byte("foobar"))
	lo.Must0(writer.Close())

	r = httptest.NewRequest("POST", "/", body)
	r.Header.Add("Content-Type", writer.FormDataContentType())
	w = httptest.NewRecorder()
	e.ServeHTTP(w, r)
	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Equal(t, `"invalid file"`, strings.TrimSpace(w.Body.String()))
}

type customValidator struct {
	validator *validator.Validate
}

func (cv *customValidator) Validate(i any) error {
	if err := cv.validator.Struct(i); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return nil
}

func TestReq_MessageContent(t *testing.T) {
	assert.Equal(t, "カテゴリ：cate\n所属組織：org\n\naaa", req{
		Content:  "aaa",
		Category: "cate",
		Org:      "org",
	}.MessageContent())
}
