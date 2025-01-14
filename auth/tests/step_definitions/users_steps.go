package step_definitions

import (
	"fmt"
	"strconv"

	"rekar.ir/v2/auth/config"
	"rekar.ir/v2/auth/internal/domain/models"

	"github.com/cucumber/godog"
)

var db = config.ConnectDatabase()

func theDatabaseIsReady() error {
	// تست اتصال به دیتابیس
	if db == nil {
		return fmt.Errorf("database connection is not initialized")
	}
	return nil
}

func iCreateAUserWithTheFollowingDetails(details *godog.Table) error {
	for _, row := range details.Rows[1:] { // Skip header row
		// تبدیل مقدار رشته به عدد صحیح
		roleID, err := strconv.Atoi(row.Cells[6].Value)
		if err != nil {
			return fmt.Errorf("invalid RoleID value: %v", err)
		}

		// ایجاد کاربر
		user := models.User{
			Name:     row.Cells[0].Value,
			Family:   row.Cells[1].Value,
			Mobile:   row.Cells[2].Value,
			Email:    row.Cells[3].Value,
			Username: row.Cells[4].Value,
			Password: row.Cells[5].Value,
			RoleID:   uint(roleID), // تبدیل به uint
		}
		if err := db.Create(&user).Error; err != nil {
			return fmt.Errorf("failed to create user: %v", err)
		}
	}
	return nil
}
func theUserShouldBeSavedSuccessfully() error {
	// بررسی اینکه کاربر در دیتابیس ذخیره شده است
	var user models.User
	if err := db.Last(&user).Error; err != nil {
		return fmt.Errorf("no user found: %v", err)
	}
	fmt.Println("User created successfully:", user)
	return nil
}

func iRetrieveTheUserWithUsername(username string) error {
	var user models.User
	if err := db.Where("username = ?", username).First(&user).Error; err != nil {
		return fmt.Errorf("failed to find user with username %s: %v", username, err)
	}
	return nil
}

func theUsersEmailShouldBe(expectedEmail string) error {
	var user models.User
	if err := db.Last(&user).Error; err != nil {
		return fmt.Errorf("failed to fetch user: %v", err)
	}
	if user.Email != expectedEmail {
		return fmt.Errorf("expected email %s, got %s", expectedEmail, user.Email)
	}
	return nil
}

func FeatureContext(s *godog.ScenarioContext) {
	s.Step(`^the database is ready$`, theDatabaseIsReady)
	s.Step(`^I create a user with the following details:$`, iCreateAUserWithTheFollowingDetails)
	s.Step(`^the user should be saved successfully$`, theUserShouldBeSavedSuccessfully)
	s.Step(`^I retrieve the user with username "([^"]*)"$`, iRetrieveTheUserWithUsername)
	s.Step(`^the user's email should be "([^"]*)"$`, theUsersEmailShouldBe)
}
