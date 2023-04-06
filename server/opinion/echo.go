package opinion

import (
	"encoding/base64"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/reearth/reearthx/log"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

const defaultToName = "PLATEAU VIEW ご意見ご要望"
const titlePrefix = "【PLATEAU VIEW ご意見ご要望】"
const defaultFromName = "PLATEAU CMS"

type Config struct {
	SendGridAPIKey string
	From           string
	To             string
	// optional
	FromName string
	// optional
	ToName string
}

type req struct {
	Title    string `json:"title" form:"title"`
	Name     string `json:"name" form:"name" validate:"required"`
	Email    string `json:"email" form:"email" validate:"required,email"`
	Content  string `json:"content" form:"content" validate:"required"`
	Category string `json:"category" form:"category"`
	Org      string `json:"org" form:"org"`
}

func Echo(g *echo.Group, conf Config) {
	fromName := conf.FromName
	if fromName == "" {
		fromName = defaultFromName
	}

	toName := conf.ToName
	if toName == "" {
		toName = defaultToName
	}
	client := sendgrid.NewSendClient(conf.SendGridAPIKey)

	g.POST("", func(c echo.Context) error {
		r := req{}
		if err := c.Bind(&r); err != nil {
			return err
		}
		if err := c.Validate(r); err != nil {
			return err
		}

		from := mail.NewEmail(fromName, conf.From)
		to := mail.NewEmail(toName, conf.To)
		title := fmt.Sprintf("%s%s", titlePrefix, r.Title)
		message := mail.NewSingleEmailPlainText(from, title, to, r.MessageContent())
		replyToEmail := mail.NewEmail(r.Name, r.Email)
		message.SetReplyTo(replyToEmail)

		// handle an image file
		if mfh, err := c.FormFile("file"); err == nil {
			mf, err := mfh.Open()
			if err != nil {
				return c.JSON(http.StatusUnprocessableEntity, "cannot open file")
			}

			defer func() { _ = mf.Close() }()
			data, err := io.ReadAll(mf)
			if err != nil {
				return c.JSON(http.StatusUnprocessableEntity, "cannot read file")
			}

			ty := http.DetectContentType(data)
			if !strings.HasPrefix(ty, "image/") {
				return c.JSON(http.StatusBadRequest, "invalid file")
			}

			a := mail.NewAttachment().
				SetContent(base64.StdEncoding.EncodeToString(data)).
				SetType(ty).
				SetFilename(mfh.Filename).
				SetDisposition("attachment")
			message.AddAttachment(a)
		}

		response, err := client.Send(message)
		if err != nil {
			e := ""
			if err != nil {
				e = err.Error()
			} else {
				e = fmt.Sprintf("code=%d,body=%s", response.StatusCode, response.Body)
			}

			log.Errorf("opinion: failed to send email: %s", e)
			return c.JSON(http.StatusBadGateway, "failed to send email")
		}

		return c.JSON(http.StatusOK, "ok")
	}, middleware.BodyLimit("10M"), middleware.CORS())
}

func (r req) MessageContent() string {
	content := ""
	if r.Category != "" {
		content += fmt.Sprintf("カテゴリ：%s\n", r.Category)
	}
	if r.Org != "" {
		content += fmt.Sprintf("所属組織：%s\n", r.Org)
	}
	if r.Category != "" || r.Org != "" {
		content += "\n"
	}
	content += r.Content
	return content
}
