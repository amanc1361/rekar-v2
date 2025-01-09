package main

import "rekar.ir/v2/auth/app"

func main() {
	// ایجاد نمونه جدید از App
	a := app.App{}

	// Initialization of app settings
	a.Initialize()

	// اجرای سرور HTTP
	a.Run(":8081")
}
