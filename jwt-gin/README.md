list of APIs

public API:
http://127.0.0.1:8080/api/register
http://127.0.0.1:8080/api/login

Protected API (requires bearer token):
http://127.0.0.1:8080/api/wallet/user
http://127.0.0.1:8080/api/wallet/balance
http://127.0.0.1:8080/api/wallet/top-up
http://127.0.0.1:8080/api/wallet/transactions
http://127.0.0.1:8080/api/wallet/confirm-biller-transaction
http://127.0.0.1:8080/api/wallet/biller-transactions
http://127.0.0.1:8080/api/wallet/delete-user
http://127.0.0.1:8080/api/wallet/reset-password

how to test:
update .env file with proper variables
create db according to the variable name
use 'go run main.go' to run script
on startup, the script will create the required tables and columns (but not DB)
you can register then login to obtain bearer token
the token is then used for all protected 

Refer to Postman API folder for more info