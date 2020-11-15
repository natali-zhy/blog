# blog

Clone this repository into your go home path or a separate directory.

> Note: If you are making clone in a separate directory then you have to export your environments.

Go to pkg/config/config.json directory and change the data to yours

> {
    "AddrHttp": ":3006",
    "Templates": "/",
    "db": {
      "AddrMysql": "localhost",
      "UserMysql": "root",
      "PasswordMysql": "",
      "Database": "article"
    }
  }

Go to your main directory and enter the following command

> go run main.go

This command will run the project successfully.
