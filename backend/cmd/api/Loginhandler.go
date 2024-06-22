package main

import (
	"encoding/csv"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo"
)

// jwtCustomClaims are custom claims extending default ones.
type jwtCustomClaims struct {
	Email    string `json:"username"`
	Role     string `json:"role"`
	Idnumber string `json:"idnumber"`
	jwt.StandardClaims
}

type UserLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

var secretKey = []byte("secret")

func jwtMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		//echo.Context เก็บทั้ง request response,path parameter,query parameter,data ต่างๆ
		authHeader := c.Request().Header.Get("Authorization")
		//ใน frontend ที่ใส่ header เราจะ check ก่อนว่ามี authheader มั้ย
		if authHeader == "" {
			return echo.NewHTTPError(http.StatusUnauthorized, "Missing or invalid token")
		}

		//เอา prefix bearer ออกเอาแค่ token
		tokenString := authHeader[len("Bearer "):]
		claims := &jwtCustomClaims{} //ใช้ในการ store data จาก token
		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return []byte(secretKey), nil
		}) // Provide the secret key used for signing the token

		if err != nil || !token.Valid {
			return echo.NewHTTPError(http.StatusUnauthorized, "Invalid token jwtparse")
		}

		if claims, ok := token.Claims.(*jwtCustomClaims); ok && token.Valid {
			// Check token expiration
			if time.Now().Unix() > claims.ExpiresAt {
				return echo.NewHTTPError(http.StatusUnauthorized, "Token expired")
			}

			// Set user claims context
			c.Set("user", claims)

			// Token is valid, proceed to the next middleware or handler
			return next(c)
		}
		return echo.NewHTTPError(http.StatusUnauthorized, "Invalid token jwtCustomClaims")
	}
}

func adminMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		user, ok := c.Get("user").(*jwtCustomClaims) //จาก jwt c.Set("user", claims)
		if !ok {
			return echo.NewHTTPError(http.StatusUnauthorized, "Invalid token from adminMiddleware")
		}
		userRole := user.Role

		if userRole != "Admin" {
			return echo.NewHTTPError(http.StatusForbidden, "You do not have access to this dashboard")
		}

		return next(c)
	}
}

func userMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		user, ok := c.Get("user").(*jwtCustomClaims) //จาก jwt c.Set("user", claims)
		if !ok {
			return echo.NewHTTPError(http.StatusUnauthorized, "Invalid token from userMiddleware")
		}
		requestedIdnumber := c.Param("idnumber")
		userRole := user.Role
		idnumber := user.Idnumber
		if userRole != "Admin" && idnumber != requestedIdnumber {
			return echo.NewHTTPError(http.StatusForbidden, "You do not have access to this dashboard")
		}

		return next(c)
	}
}

func readCSVFile(filePath string) ([][]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.Comma = ','
	return reader.ReadAll()
}

func loginUserHandler(c echo.Context) error {
	userlogin := new(UserLogin)
	if err := c.Bind(userlogin); err != nil { //Bind() ใช้ในการ wrap request body ตามโครงสร้างของ userlogin
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid request payload"})
	}
	if userlogin.Email == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Email is Missing"})
	}
	if userlogin.Password == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Password is Missing"})
	}

	userCSVPath := `C:\Users\User\OneDrive\Desktop\financia\frontend\users.csv`
	adminCSVPath := `C:\Users\User\OneDrive\Desktop\financia\frontend\admins.csv`

	records, err := readCSVFile(userCSVPath)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": fmt.Sprintf("Error opening or reading CSV file: %v", err)})
	}

	for i, record := range records {
		if i == 0 {
			//ข้ามแถวแรก เพราะเป็นชื่อ attribute
			continue
		}
		recordEmail := record[2]    // column 3  เป็น (email)
		recordPassword := record[3] // column 4 เป็น (password)
		recordRole := record[4]     // column 5 เป็น (role)
		recordIdnumber := record[0]
		//fmt.Printf("recordEmail:%s , recordPassword:%s , inputemail:%s , inputpassword:%s \n", recordEmail, recordPassword, userlogin.Email, userlogin.Password)

		// ถ้า email กับ password ตรง กับบัญชีใน users.csv
		if recordEmail == userlogin.Email && recordPassword == userlogin.Password {

			claims := &jwtCustomClaims{
				Email:    userlogin.Email,
				Role:     recordRole,
				Idnumber: recordIdnumber,
				StandardClaims: jwt.StandardClaims{
					ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
				},
			}

			//สร้าง token
			token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
			t, err := token.SignedString(secretKey)
			if err != nil {
				return c.JSON(http.StatusBadRequest, map[string]string{"message": "generate token error"})
			}

			//ถ้าสร้างสำเร็จ return json status กับ token
			return c.JSON(http.StatusOK, echo.Map{
				"token": t,
			})
		}
	}

	//ถ้าไม่เจอใน user.csv
	// ทำแบบเดียวกับ client แต่ทำใน admins.csv
	records, err = readCSVFile(adminCSVPath)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": fmt.Sprintf("Error opening or reading CSV file: %v", err)})
	}

	for i, record := range records {
		if i == 0 {
			continue
		}
		recordEmail := record[2]
		recordPassword := record[3]
		recordRole := record[4]
		//fmt.Printf("recordEmail:%s , recordPassword:%s , inputemail:%s , inputpassword:%s \n", recordEmail, recordPassword, userlogin.Email, userlogin.Password)

		if recordEmail == userlogin.Email && recordPassword == userlogin.Password {

			claims := &jwtCustomClaims{
				Email: userlogin.Email,
				Role:  recordRole,
				StandardClaims: jwt.StandardClaims{
					ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
				},
			}

			token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
			t, err := token.SignedString(secretKey)
			if err != nil {
				return c.JSON(http.StatusBadRequest, map[string]string{"message": "generate token error"})
			}

			return c.JSON(http.StatusOK, echo.Map{
				"token": t,
			})
		}
	}

	return c.JSON(http.StatusBadRequest, map[string]string{"message": "Incorrect Email or Password!"})
}
