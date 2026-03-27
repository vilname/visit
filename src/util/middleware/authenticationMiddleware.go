package middleware

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
	"regexp"
	"strings"

	"github.com/form3tech-oss/jwt-go"
	"github.com/gin-gonic/gin"
	"gopkg.in/yaml.v2"
)

type AccessRule struct {
	Path  string   `yaml:"path"`
	Roles []string `yaml:"roles"`
}

type FileRole struct {
	AccessControl []AccessRule `yaml:"access_control"`
}

type TokenRole struct {
	Roles map[string]string `json:"roles"`
}

type JWTClaims struct {
	Sub         string `json:"sub"`
	RealmAccess struct {
		Roles []string `json:"roles"`
	} `json:"realm_access"`
}

func AuthenticationMiddleware(c *gin.Context) {
	request := c.Request
	urlPath := request.URL.Path

	if strings.Contains(urlPath, "/swagger/") {
		c.Next()
		return
	}

	auth := request.Header.Get("Authorization")
	if auth == "" {
		c.Next()
		return
	}

	jwtString := strings.TrimPrefix(auth, "Bearer ")

	token, _ := jwt.Parse(jwtString, nil)

	if token == nil {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	tokenClaims, _ := tokenClaim(jwtString)
	isAccess, _ := checkPermission(tokenClaims.RealmAccess.Roles, urlPath)

	if !isAccess {
		c.AbortWithStatus(http.StatusForbidden)
		return
	}

	userUuid := tokenClaims.Sub

	c.Set("userUuid", userUuid)

	c.Next()
}

func tokenClaim(jwtString string) (JWTClaims, error) {
	var err error
	var claims JWTClaims

	parts := strings.Split(jwtString, ".")
	if len(parts) != 3 {
		return claims, errors.New("Неверный формат JWT")
	}

	// Декодируем payload
	payload, err := base64.RawURLEncoding.DecodeString(parts[1])
	if err != nil {
		return claims, errors.New(fmt.Sprintf("Ошибка: %v\n", err))
	}

	err = json.Unmarshal(payload, &claims)
	if err != nil {
		return claims, errors.New(fmt.Sprintf("Ошибка: %v\n", err))
	}
	return claims, nil
}

func checkPermission(tokenRoles []string, requestUrl string) (bool, error) {

	var err error
	isAccess := false

	userRoles := make([]string, 0)
	for _, role := range tokenRoles {
		if strings.Contains(role, "role_") {
			userRoles = append(userRoles, role)
		}
	}

	data, err := os.ReadFile("config/role.yaml")
	if err != nil {
		return isAccess, errors.New("неверный путь до файла с ролями")
	}

	var role FileRole
	err = yaml.Unmarshal(data, &role)

	rolesPath := ""
	for _, rule := range role.AccessControl {
		regexStr := fmt.Sprintf("%s", rule.Path)
		matched, _ := regexp.MatchString(regexStr, requestUrl)

		if matched {
			rolesPath = strings.Join(rule.Roles, ",")
			break
		}
	}

	for _, userRole := range userRoles {
		if strings.Contains(rolesPath, userRole) {
			isAccess = true
		}
	}

	return isAccess, nil
}
