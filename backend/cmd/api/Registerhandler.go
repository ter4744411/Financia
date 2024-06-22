package main

import (
	"encoding/csv"
	"fmt"
	"net/http"
	"os"
	"regexp"

	"github.com/labstack/echo"
)

type User struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func isValidEmail(email string) bool {
	// Format email
	pattern := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	match, _ := regexp.MatchString(pattern, email)
	return match
}

func isValidPassword(password string) bool {
	// password at least 8 characters
	pattern := `^.{8,}$`
	match, _ := regexp.MatchString(pattern, password)
	return match
}

func isValidUsername(username string) bool {
	// username can not over than 10 characters
	pattern := `^.{1,10}$`
	match, _ := regexp.MatchString(pattern, username)
	return match
}

func saveUserHandler(c echo.Context) error {

	user := new(User) //Bind() ใช้ในการ wrap request body ตามโครงสร้างของ User
	if err := c.Bind(user); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid request payload")
	}
	fmt.Printf("Received user data: %+v\n", user)

	// Check email missing
	if user.Email == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "Email is missing")
	}
	// Check email validation
	if !isValidEmail(user.Email) {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid email format")
	}

	// Check password missing
	if user.Password == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "Password is missing")
	}
	// Validate password
	if !isValidPassword(user.Password) {
		return echo.NewHTTPError(http.StatusBadRequest, "Password must contain at least 8 characters")
	}
	// Check username missing
	if user.Username == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "Username is missing")
	}
	// Validate username
	if !isValidUsername(user.Username) {
		return echo.NewHTTPError(http.StatusBadRequest, "Your Username is too long")
	}

	// Save user data to CSV
	if err := saveToCSV(user); err != nil {
		if err.Error() == "email already exists" {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
		return echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("Failed to save user: %s", err.Error()))
	}

	return c.String(http.StatusOK, "User data saved successfully!")
}

func saveToCSV(user *User) error {
	path := `C:\Users\User\OneDrive\Desktop\financia\frontend\users.csv`

	file, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		return fmt.Errorf("can't open csv file")
	}
	defer file.Close() //ปิดไฟล์ หลังจาก function returns

	info, err := file.Stat() //ดึงข้อมูลจากไฟล์
	if err != nil {
		return fmt.Errorf("file state error")
	}

	writer := csv.NewWriter(file) // สร้าง cs writer
	defer writer.Flush()          //ล้างบัฟเฟอร์ของผู้เขียนหลังจาก function returns

	var idnumber int
	//กรณีเป็นตารางว่างๆ
	if info.Size() == 0 {
		//กำหนดชื่อ attribute
		header := []string{"idnumber", "username", "email", "password", "role"}
		if err := writer.Write(header); err != nil {
			return fmt.Errorf("write header error")
		}
		idnumber = 1
	} else {
		//check email ซํ้าใน users.csv กับใน admins.csv

		file.Seek(0, 0)
		reader := csv.NewReader(file)
		records, err := reader.ReadAll()
		if err != nil {
			return fmt.Errorf("readall csv error")
		}

		for _, record := range records {
			if record[2] == user.Email {
				return fmt.Errorf("email already exists")
			}
		}
		adminpath := `C:\Users\User\OneDrive\Desktop\financia\frontend\admins.csv`
		adminrecords, err := readCSVFile(adminpath)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("Error opening or reading CSV file: %v", err))
		}

		for i, record := range adminrecords {
			if i == 0 {
				continue
			}
			if record[2] == user.Email {
				return fmt.Errorf("email already exists")
			}
		}

		idnumber = len(records)
	}

	//ถ้า email password username valid
	//ถ้า email ไม่มีซํ้ากับในฐานข้อมูลทั้งสอง
	//บันทึกข้อมูลผู้ใช้งานในฐานะ client
	record := []string{
		fmt.Sprintf("%d", idnumber),
		user.Username,
		user.Email,
		user.Password,
		"client",
	}
	if err := writer.Write(record); err != nil {
		return fmt.Errorf("write record error")
	}

	return nil
}
