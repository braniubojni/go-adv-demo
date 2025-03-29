package main

import (
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"testing"
	"time"
)

func getComposePath() string {
	wd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	return filepath.Join(wd, "docker-compose.test.yml")
}

func startTestDB() {
	log.Println("🟢 Starting test DB (postgres_test_go)...")

	composePath := getComposePath()

	cmd := exec.Command("docker-compose", "-f", composePath, "up", "-d", "postgres")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		log.Fatalf("❌ Failed to start test DB: %v", err)
	}

	log.Println("⏳ Waiting for DB to be ready...")
	time.Sleep(5 * time.Second) // Can be replaced with retry logic
}

func runMigrations() {
	log.Println("📦 Running DB migrations...")

	// Run `go run migrations/auto.go`
	cmd := exec.Command("go", "run", "../migrations/auto.go")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		log.Fatalf("❌ Failed to run migrations: %v", err)
	}
}

func stopTestDB() {
	log.Println("🔴 Stopping test DB (postgres_test_go)...")

	composePath := getComposePath()

	cmd := exec.Command("docker-compose", "-f", composePath, "down", "--volumes")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		log.Printf("⚠️ Failed to stop test DB: %v", err)
	}
}

func TestMain(m *testing.M) {
	startTestDB()
	runMigrations()
	code := m.Run()
	stopTestDB()
	os.Exit(code)
}
