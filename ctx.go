package fiberRoute

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"time"

	"github.com/gofiber/fiber/v2"
	contracthttp "github.com/sujit-baniya/framework/contracts/http"
)

type Context struct {
	instance *fiber.Ctx
}

func (c *Context) Origin() *http.Request {
	headers := make(map[string][]string)
	for header, value := range c.instance.GetReqHeaders() {
		headers[header] = []string{value}
	}
	parsedUrl, _ := url.Parse(c.instance.OriginalURL())
	return &http.Request{
		Method: c.instance.Method(),
		URL:    parsedUrl,
		Proto:  c.instance.Protocol(),
		Header: headers,
		Body:   io.NopCloser(bytes.NewReader(c.instance.Body())),
		GetBody: func() (io.ReadCloser, error) {
			return io.NopCloser(bytes.NewReader(c.instance.Body())), nil
		},
		Host:       c.instance.Hostname(),
		RemoteAddr: c.instance.IP(),
		RequestURI: c.instance.Request().URI().String(),
	}
}

func NewContext(ctx *fiber.Ctx) contracthttp.Context {
	return &Context{ctx}
}

func (c *Context) WithValue(key string, value any) {
	c.instance.Context().SetUserValue(key, value)
}

func (c *Context) Deadline() (deadline time.Time, ok bool) {
	return c.instance.Context().Deadline()
}

func (c *Context) Done() <-chan struct{} {
	return c.instance.Context().Done()
}

func (c *Context) Err() error {
	return c.instance.Context().Err()
}

func (c *Context) Value(key any) any {
	return c.instance.Context().Value(key)
}

func (c *Context) Params(key string) string {
	return c.instance.Params(key)
}

func (c *Context) Query(key, defaultValue string) string {
	return c.instance.Query(key, defaultValue)
}

func (c *Context) Form(key, defaultValue string) string {
	return c.instance.FormValue(key, defaultValue)
}

func (c *Context) Bind(obj any) error {
	return nil
}

func (c *Context) SaveFile(name string, dst string) error {
	file, err := c.File(name)
	if err != nil {
		return err
	}
	return c.instance.SaveFile(file, dst)
}

func (c *Context) File(name string) (*multipart.FileHeader, error) {
	return c.instance.FormFile(name)
}

func (c *Context) Header(key, defaultValue string) string {
	header := c.instance.Get(key)
	if header != "" {
		return header
	}

	return defaultValue
}

func (c *Context) Headers() http.Header {
	mp := make(map[string][]string)
	headers := c.instance.GetReqHeaders()
	for key, header := range headers {
		mp[key] = []string{header}
	}
	return mp
}

func (c *Context) Method() string {
	return c.instance.Method()
}

func (c *Context) Status(code int) contracthttp.Context {
	c.instance.Status(code)
	return c
}

func (c *Context) Url() string {
	return c.instance.OriginalURL()
}

func (c *Context) FullUrl() string {
	prefix := "https://"
	if !c.instance.Secure() {
		prefix = "http://"
	}

	if c.instance.Hostname() == "" {
		return ""
	}

	return prefix + string(c.instance.Request().Host()) + string(c.instance.Request().RequestURI())
}

func (c *Context) AbortWithStatus(code int) {
	c.instance.Status(code)
}

func (c *Context) Next() error {
	return c.instance.Next()
}

func (c *Context) Cookies(key string, defaultValue ...string) string {
	return c.instance.Cookies(key, defaultValue...)
}

func (c *Context) Cookie(co *contracthttp.Cookie) {
	c.instance.Cookie(&fiber.Cookie{
		Name:        co.Name,
		Value:       co.Value,
		Path:        co.Path,
		Domain:      co.Domain,
		MaxAge:      co.MaxAge,
		Expires:     co.Expires,
		Secure:      co.Secure,
		HTTPOnly:    co.HTTPOnly,
		SameSite:    co.SameSite,
		SessionOnly: co.SessionOnly,
	})
}

func (c *Context) Path() string {
	return string(c.instance.Request().URI().Path())
}

func (c *Context) EngineContext() any {
	return c.instance
}

func (c *Context) Secure() bool {
	return c.instance.Secure()
}

func (c *Context) Ip() string {
	return c.instance.IP()
}

func (c *Context) String(format string, values ...any) error {
	return c.instance.SendString(fmt.Sprintf(format, values...))
}

func (c *Context) Json(obj any) error {
	return c.instance.JSON(obj)
}

func (c *Context) SendFile(filepath string, compress ...bool) error {
	return c.instance.SendFile(filepath, compress...)
}

func (c *Context) Download(filepath, filename string) error {
	return c.instance.Download(filepath, filename)
}

func (c *Context) StatusCode() int {
	return c.instance.Response().StatusCode()
}

func (c *Context) Render(name string, bind any, layouts ...string) error {
	return c.instance.Render(name, bind, layouts...)
}

func (c *Context) SetHeader(key, value string) contracthttp.Context {
	c.instance.Set(key, value)
	return c
}

func (c *Context) Vary(key string, value ...string) {
	c.instance.Vary(key)
}
