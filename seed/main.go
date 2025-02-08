package main

import (
	"fmt"
	"log"

	"mnc-rest-api/database"
	"mnc-rest-api/internal/domain"
	"mnc-rest-api/pkg/inits"

	"gorm.io/gorm"
)

func main() {
	cfg := inits.InitializeConfig()
	database.ConnectDB()
	db := inits.InitializeDatabase(cfg.Env.Db, cfg.Env.Tz)

	err := db.AutoMigrate(&domain.User{}, &domain.TopUp{}, &domain.Transaction{})
	if err != nil {
		log.Fatalf("failed to migrate schema: %v", err)
	}

	insertSeedData(db)

	fmt.Println("Seed data inserted successfully.")
}

func insertSeedData(db *gorm.DB) {
	userQuery := `
		INSERT INTO users (user_id, first_name, last_name, phone_number, address, pin, balance, created_date)
		VALUES 
			('bc1c823e-b0fb-4b20-88c0-dff25e283252', 'Tom', 'Araya', '08123456789', 'Jl. Diponegoro No. 215', '1234', 50000.00, NOW()),
			('bd2d723e-c0fa-4c10-88f0-eef45e293353', 'Alice', 'Johnson', '08129876543', 'Jl. Merdeka No. 10', '5678', 30000.00, NOW());
	`

	err := db.Exec(userQuery)
	if err != nil {
		log.Fatalf("Failed to insert users seed data: %v", err)
	}

	// Seed data for topups
	topUpQuery := `
		INSERT INTO topups (top_up_id, user_id, amount, balance_before, balance_after, created_date)
		VALUES 
			('201ddde1-f797-484b-b1a0-07d1190e790a', 'bc1c823e-b0fb-4b20-88c0-dff25e283252', 500000.00, 0.00, 500000.00, NOW()),
			('302ddde1-f897-485b-b1b0-07d1290e890a', 'bd2d723e-c0fa-4c10-88f0-eef45e293353', 300000.00, 0.00, 300000.00, NOW());
	`

	err = db.Exec(topUpQuery)
	if err != nil {
		log.Fatalf("Failed to insert topups seed data: %v", err)
	}

	// Seed data for transactions
	transactionQuery := `
		INSERT INTO transactions (transaction_id, user_id, top_up_id, transaction_type, amount, balance_before, balance_after, remarks, created_date)
		VALUES 
			('a7d39cf6-44b6-41fc-b3e9-7b16df5321c5', 'bc1c823e-b0fb-4b20-88c0-dff25e283252', '201ddde1-f797-484b-b1a0-07d1190e790a', 'CREDIT', 500000.00, 0.00, 500000.00, 'Initial Top Up', NOW()),
			('b8d49df7-55c7-52ed-b4e8-8c27de5432d6', 'bd2d723e-c0fa-4c10-88f0-eef45e293353', '302ddde1-f897-485b-b1b0-07d1290e890a', 'CREDIT', 300000.00, 0.00, 300000.00, 'Initial Top Up', NOW());
	`

	err = db.Exec(transactionQuery)
	if err != nil {
		log.Fatalf("Failed to insert transactions seed data: %v", err)
	}
}
