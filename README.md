# mnc-rest-api


Prerequisites
Make sure the following prerequisites are met:

Go: Ensure Go is installed on your system.
Database: Set up your database and update the configuration accordingly in the local.yaml file.
Steps to Run the Project
1. Update Configuration
Update the local.yaml configuration file with the correct values for the database connection, timezone, and server settings.

2. Insert Seed Data
Run the seed data script to populate the database with initial values for users, topups, and transactions.

go run seed/main.go

3. Run the Application
Once the seed data has been inserted, start the main application:

go run main.go