package repositories_test

import (
	"database/sql"
	"fmt"
	"os"
	"os/exec"
	"testing"
)

var testDB *sql.DB

var (
	dbUser     = "docker"
	dbPassword = "docker"
	dbDatabase = "sampledb"
	dbConn     = fmt.Sprintf("%s:%s@tcp(127.0.0.1:3300)/%s?parseTime=true", dbUser, dbPassword, dbDatabase)
)

// DBに接続する処理
func connectDB() error {
	var err error
	testDB, err = sql.Open("mysql", dbConn)
	if err != nil {
		return err
	}
	return nil
}

// テストの前処理
func setup() error {
	if err := connectDB(); err != nil {
		return err
	}
	if err := cleanupDB(); err != nil {
		return err
	}
	if err := setupTestData(); err != nil {
		return err
	}
	return nil
}

// テストの後処理
func teardown() {
	cleanupDB()
	testDB.Close()
}

// 前処理と後処理を実行する
func TestMain(m *testing.M) {
	err := setup() // テストの前処理
	if err != nil {
		os.Exit(1)
	}

	m.Run()    // テストの実行
	teardown() // テストの後処理
}

// mysqlクライアントからmysqlサーバーに接続してテストデータ作成のsqlを実行する
func setupTestData() error {
	cmd := exec.Command("mysql", "-h", "127.0.0.1", "-u", "docker", "sampledb", "--password=docker", "-P", "3300", "-e", "source ./testdata/setupDB.sql")
	err := cmd.Run()
	if err != nil {
		return err
	}
	return nil
}

// mysqlクライアントからmysqlサーバーに接続してテストデータ削除のsqlを実行する
func cleanupDB() error {
	cmd := exec.Command("mysql", "-h", "127.0.0.1", "-u", "docker", "sampledb", "--password=docker", "-P", "3300", "-e", "source ./testdata/setupDB.sql")
	err := cmd.Run()
	if err != nil {
		return err
	}
	return nil
}
