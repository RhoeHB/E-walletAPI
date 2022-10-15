list of APIs<br /><br />

public API:<br />
http://127.0.0.1:8080/api/register<br />
http://127.0.0.1:8080/api/login<br />

Protected API (requires bearer token):<br />
http://127.0.0.1:8080/api/wallet/user<br />
http://127.0.0.1:8080/api/wallet/balance<br />
http://127.0.0.1:8080/api/wallet/top-up<br />
http://127.0.0.1:8080/api/wallet/transactions<br />
http://127.0.0.1:8080/api/wallet/confirm-biller-transaction<br />
http://127.0.0.1:8080/api/wallet/biller-transactions<br />
http://127.0.0.1:8080/api/wallet/delete-user<br />
http://127.0.0.1:8080/api/wallet/reset-password<br />

how to test:<br />
update .env file with proper variables<br />
create db according to the variable name<br />
use 'go run main.go' to run script<br />
on startup, the script will create the required tables and columns (but not DB)<br />
you can register then login to obtain bearer token
the token is then used for all protected<br /> 

Refer to Postman API folder for more info<br />
