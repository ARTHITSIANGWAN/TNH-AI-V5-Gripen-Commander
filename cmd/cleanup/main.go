package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	fmt.Println("==================================================")
	fmt.Println("🧹 TNH V84.9.2 - PUBLIC REPOSITORY CLEANUP UTILITY")
	fmt.Println("==================================================")

	cmdDir := "cmd"
	files, err := os.ReadDir(cmdDir)
	if err != nil {
		fmt.Printf("❌ Error reading directory: %v\n", err)
		return
	}

	for _, file := range files {
		if !file.IsDir() {
			fileName := file.Name()
			ext := filepath.Ext(fileName)
			if ext == ".go" {
				baseName := strings.TrimSuffix(fileName, ext)
				targetDir := filepath.Join(cmdDir, baseName)
				
				// สร้างโฟลเดอร์ย่อยรองรับโครงสร้างแบบ Modular
				os.MkdirAll(targetDir, 0755)

				oldPath := filepath.Join(cmdDir, fileName)
				newFileName := "main.go"
				if fileName == "main.go" {
					continue
				}
				newPath := filepath.Join(targetDir, newFileName)

				// ย้ายไฟล์เข้าโฟลเดอร์หลักประจำตัว
				err := os.Rename(oldPath, newPath)
				if err != nil {
					fmt.Printf("⚠️ Failed to move %s: %v\n", fileName, err)
				} else {
					fmt.Printf("🟢 Sanitized & Moved: /cmd/%s -> /cmd/%s/%s\n", fileName, baseName, newFileName)
				}
			}
		}
	}

	fmt.Println("--------------------------------------------------")
	fmt.Println("✅ Repository Cleanup Completed. Ready for Public Sync.")
	fmt.Println("==================================================")
}

