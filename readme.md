# SalaryWithdrawal
---
SalaryWithdrawal is a management feature designed to handle financial and personnel data within an organization efficiently. It includes the following capabilities:

Features:
- *Job Position Management*: Perform CRUD (Create, Read, Update, Delete) operations to manage various job positions, including their respective salary scales.
- *Employee Management*: Manage employee data, enabling admins to add, update, or remove employee records using CRUD operations.
- *Company Balance Management*: Administrators have the authority to top up the company’s balance, allowing for continuous operational funding.
- *Salary Withdrawal*: Enables salary withdrawals, requiring an Employee ID and a secure Secret ID to verify identity. The withdrawal amount is based on the employee’s current position and respective salary scale.
- *Transaction History*: Maintain a clear record of all balance top-ups and salary deductions, offering transparency in financial transactions.

# How to Run Locally

To run the application, execute the following commands:

1. `cp .env.example .env` and fill it in according to the environment settings on your PC.
2. `go mod tidy && go mod vendor`
3. `go run main.go`

---
# Project Database Structure

This document outlines the database structure for the project. The project's database consists of four collections: "users", "companies," "positions," and "transactions."

## Users Collection

This collection stores information about registered users.

| Column       | Data Type | Description                                                           |
|--------------|-----------|-----------------------------------------------------------------------|
| id           | INT       | Primary key, unique identifier for each user.                         |
| secret_id    | VARCHAR   | Unique secret identifier for secure actions.                          |
| name         | VARCHAR   | User's full name.                                                     |
| email        | VARCHAR   | User's email address.                                                 |
| phone        | VARCHAR   | User's phone number.                                                  |
| address      | TEXT      | User's residential address.                                           |
| position_id  | INT       | Foreign key referencing the `Position` table, represents the user's job position. |
| created_at   | TIMESTAMP | Timestamp for when the user record was created.                       |
| updated_at   | TIMESTAMP | Timestamp for the latest update to the user record.  |

### Database Structure for Company Model

| Column       | Data Type | Description                                       |
|--------------|-----------|---------------------------------------------------|
| id           | INT       | Primary key, unique identifier for each company.  |
| name         | VARCHAR   | Name of the company.                              |
| address      | TEXT      | Address of the company.                           |
| balance      | INT       | Current balance of the company.                   |
| created_at   | TIMESTAMP | Timestamp for when the company record was created.|
| updated_at   | TIMESTAMP | Timestamp for the latest update to the company record.|

---

### Database Structure for Position Model

| Column       | Data Type | Description                                        |
|--------------|-----------|----------------------------------------------------|
| id           | INT       | Primary key, unique identifier for each position.  |
| name         | VARCHAR   | Name of the job position.                          |
| salary       | INT       | Salary associated with this job position.          |
| created_at   | TIMESTAMP | Timestamp for when the position record was created.|
| updated_at   | TIMESTAMP | Timestamp for the latest update to the position record.|

---

### Database Structure for Transaction Model

| Column       | Data Type | Description                                        |
|--------------|-----------|----------------------------------------------------|
| id           | INT       | Primary key, unique identifier for each transaction. |
| amount       | INT       | Amount of the transaction.                         |
| note         | TEXT      | Note or description of the transaction.            |
| type         | VARCHAR   | Type of transaction (e.g., "credit", "debit").|
| created_at   | TIMESTAMP | Timestamp for when the transaction record was created.|
| updated_at   | TIMESTAMP | Timestamp for the latest update to the transaction record.|

