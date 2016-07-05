package middleware

import (
	"crypto/md5"
	"encoding/base64"
	"fmt"
	. "github.com/Dataman-Cloud/omega-napp/config"
	"github.com/Dataman-Cloud/omega-napp/store"
	"github.com/Dataman-Cloud/omega-napp/util"
	log "github.com/Sirupsen/logrus"
	redis "github.com/garyburd/redigo/redis"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func Authenticate(ctx *gin.Context) {
	authenticated := false
	token := util.Header(ctx, HeaderToken)
	sry_svc_token := util.Header(ctx, SRY_SVC_TOKEN)
	uid := util.Header(ctx, "Uid")

	// for EventSource does not header setting
	if len(token) == 0 {
		token = strings.Trim(ctx.Query(strings.ToLower(HeaderToken)), " ")
	}

	if len(token) > 0 {
		conn := store.Open()
		defer conn.Close()
		uid, err := redis.String(conn.Do("HGET", "s:"+token, "user_id"))
		log.Info("uid: ", uid)
		if err == nil {
			authenticated = true
			ctx.Set("uid", uid)
		} else if err != redis.ErrNil {
			log.Error("[app] got error1 ", err)
		}
	}
	if len(sry_svc_token) > 0 {
		if uid != "" {
			if sry_svc_token == CronTokenBuilder(uid) {
				authenticated = true
				ctx.Set("uid", uid)
			}
		}
	}

	if authenticated {
		ctx.Next()
	} else {
		ctx.JSON(http.StatusUnauthorized, gin.H{"code": 10002})
		ctx.Abort()
	}
}

func CronTokenBuilder(uid string) string {
	b64 := GetBaseEncoding()
	md5Token := fmt.Sprintf("%x", md5.Sum([]byte(uid)))
	b64Token := b64.EncodeToString([]byte(uid))
	token := b64.EncodeToString([]byte(fmt.Sprintf("%s:%s", md5Token, b64Token)))[:20]
	return strings.ToLower(token)
}

func GetBaseEncoding() *base64.Encoding {
	return base64.NewEncoding("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/")
}
