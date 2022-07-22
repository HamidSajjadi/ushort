package api

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"github.com/HamidSajjadi/ushort/internal/repositories"
	"github.com/gin-gonic/gin"
	"net/http"
	"net/url"
)

type Handler struct {
	urlRepo repositories.URLRepository
	gin     *gin.Engine
}

func New(stub *gin.Engine, urlRepo repositories.URLRepository) *Handler {
	handler := &Handler{
		urlRepo: urlRepo,
		gin:     stub,
	}
	handler.initRouter()
	return handler
}

func (h *Handler) Run(address string) {

	err := h.gin.Run(address)
	if err != nil {
		panic(err)
	}
	fmt.Printf("listening at %s", address)
}

func (h *Handler) initRouter() {
	h.gin.POST("/shorten", h.CreateShortURL)
	h.gin.GET("/:shortenedURL", h.Redirect)
}

func (h *Handler) Redirect(c *gin.Context) {
	shortenedURL := c.Param("shortenedURL")
	u, err := h.urlRepo.GetOne(shortenedURL)
	if err != nil {
		abort(c, err)
		return
	}

	c.Redirect(http.StatusMovedPermanently, u.Source)
}

func (h *Handler) CreateShortURL(c *gin.Context) {

	type ShortenURLReq struct {
		Source string
	}
	var req ShortenURLReq
	if err := c.BindJSON(&req); err != nil {
		abort(c, err)
		return
	}
	source, err := parseURL(req.Source)
	if err != nil {
		abort(c, err)
		return
	}
	shortUrl := shortenURL(source)
	_, err = h.urlRepo.Save(source, shortUrl)
	if err != nil {
		abort(c, err)
		return
	}
	c.JSON(http.StatusCreated, gin.H{"url": shortUrl})
}

func parseURL(inp string) (string, error) {
	u, err := url.Parse(inp)
	if err != nil {
		return "", err
	}
	if u.Scheme == "" {
		u.Scheme = "https"
	}
	return u.String(), nil
}

func shortenURL(url string) string {
	hash := md5.Sum([]byte(url))
	return hex.EncodeToString(hash[:])[:7]
}

func abort(c *gin.Context, err error) {
	c.Error(err)
	c.Abort()
}
