package api

import (
	"bytes"
	"fmt"
	"github.com/dchest/captcha"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"time"
)
 
type CaptchaResponse struct {
	CaptchaId string `json:"captchaId"` //验证码Id
	ImageUrl  string `json:"imageUrl"`  //验证码图片url
}
 
func Captcha(c *gin.Context) {
	length := captcha.DefaultLen
	captchaId := captcha.NewLen(length)
	var captcha CaptchaResponse
	captcha.CaptchaId = captchaId
	captcha.ImageUrl = captchaId
	c.JSON(http.StatusOK, captcha)
}

func Verify(c *gin.Context) {
	captchaId := c.Query("captchaId")
	value := c.Query("value")
	if captchaId == "" || value == "" {
		c.String(http.StatusBadRequest, "参数错误")
	}
	if captcha.VerifyString(captchaId, value) {
		c.JSON(http.StatusOK, "验证成功")
	} else {
		c.JSON(http.StatusOK, "验证失败")
	}
}

func GetCaptcha(c *gin.Context) {
	captchaId := c.Query("captchaId")
	ServeHTTP(c.Writer, c.Request, captchaId)
}


func ServeHTTP(w http.ResponseWriter, r *http.Request , id string) {
	ext := ".png"
	fmt.Println("serr : " + id)

	if id == "" {
		http.NotFound(w, r)
		return
	}

	if r.FormValue("reload") != "" {
		captcha.Reload(id)
	}
	lang := strings.ToLower(r.FormValue("lang"))
	download := false
	if Serve(w, r, id, ext, lang, download, captcha.StdWidth, captcha.StdHeight) == captcha.ErrNotFound {
		http.NotFound(w, r)
	}
}
 
func Serve(w http.ResponseWriter, r *http.Request, id, ext, lang string, download bool, width, height int) error {
	w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
	w.Header().Set("Pragma", "no-cache")
	w.Header().Set("Expires", "0")
	
	var content bytes.Buffer
	switch ext {
	case ".png":
		w.Header().Set("Content-Type", "image/png")
		captcha.WriteImage(&content, id, width, height)
	case ".wav":
		w.Header().Set("Content-Type", "audio/x-wav")
		captcha.WriteAudio(&content, id, lang)
	default:
		return captcha.ErrNotFound
	}
 
	if download {
		w.Header().Set("Content-Type", "application/octet-stream")
	}
	http.ServeContent(w, r, id+ext, time.Time{}, bytes.NewReader(content.Bytes()))
	return nil
}
 
