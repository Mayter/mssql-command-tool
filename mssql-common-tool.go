package main

import (
  "os"
  "fmt"
  "log"
  //"reflect"
  "github.com/urfave/cli"
  "database/sql"
  _ "github.com/denisenkom/go-mssqldb"
)

var (

	server string  = "127.0.0.1"
	user = "sa"
	password string
	port int
	query string
	cmd string
	cmd1 string
	debug bool
	enable bool
	ole bool
	connString string
	conn *sql.DB
	err error

)


func main() {
	app := cli.NewApp()
	app.Name = "Mssql Toolkit"
	app.Version = "1.0"
	
	app.Usage = "mssql command tool"
	app.Authors = []cli.Author{
		cli.Author{
		  Name:  "lostwolf",
		  Email: "linuxseclab@gmail.com",
		},
	  }

	app.Flags = []cli.Flag {
		cli.StringFlag {
			Name: "server,host,s",
			Value: "127.0.0.1",
			Usage: "The database server",
		},
		cli.StringFlag {
			Name: "user, u",
			Value: "sa",
			Usage: "The database user",
		},
		cli.StringFlag {
			Name: "password, p",
			Usage: "The database password",
		},
		cli.IntFlag {
			Name: "port, P",
			Usage: "The database port",
		},
		cli.StringFlag {
			Name: "query, sql,q",
			Value: "select @@version",
			Usage: "SQL query",
		},
	
		cli.StringFlag {
			Name: "exec,c,cmd",
			Value: "whoami",
			Usage: "Exec System Command",

		},
		cli.StringFlag {
			Name: "common,C,cmd1",
			Value: "c:\\windows\\system32\\cmd.exe /c whoami>whoami.txt",
			Usage: "Exec System Command",

		},
		cli.BoolFlag{
			Name: "debug,d",
			Usage: "Debug info",
		},
		cli.BoolFlag{
			Name: "enable,e",
			Usage: "Enabled xp_cmdshell",
		},
		cli.BoolFlag{
			Name: "ole,o",
			Usage: "Enabled sp_oacreate",
		},
	}

	app.Action = func(c *cli.Context) error {
		if c.IsSet("server"){
			server=c.String("server")
		}
		if c.IsSet("user"){
			user=c.String("user")
		}
		if c.IsSet("password"){
			password=c.String("password")
		}
		if c.IsSet("port"){
			port=c.Int("port")
		}
		if c.IsSet("query"){
			query=c.String("query")
		}
		if c.IsSet("cmd"){
			cmd=c.String("cmd")
		}
		if c.IsSet("cmd1"){
			cmd1=c.String("cmd1")
		}
		
		
		connString = fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;encrypt=disable", server, user, password ,port)
		conn,err = sql.Open("mssql", connString)
		defer conn.Close()
		//fmt.Println(reflect.TypeOf(conn))
		//fmt.Println(conn)
	
		if err != nil {
		log.Fatal("Open connection failed:", err.Error())
		}
		//fmt.Println(" Connect:",connString) 
		defer conn.Close()
    if c.IsSet("debug"){
		if c.Bool("debug"){
			log.Println("Debug info:")
			fmt.Printf(" server:%s\n", server)
			fmt.Printf(" user:%s\n", user)
			fmt.Printf(" password:%s\n", password)
			fmt.Printf(" port:%d\n", port)
			fmt.Printf(" Query:%s\n", query)
			fmt.Printf(" Cmd:%s\n", cmd)
			fmt.Printf(" Cmd1:%s\n", cmd1)
			fmt.Println(" Connect:",connString)
			}
		}
		if c.IsSet("enable"){
			if c.Bool("enable"){
				Open()
			}
		}
		if c.IsSet("ole"){
			if c.Bool("ole"){
				Openole()
			}
		}
	
		if c.IsSet("query"){		
				exec_sql()
		}
		if c.IsSet("cmd"){		
			os_shell()
		}
		if c.IsSet("cmd1"){		
			os_shellt()
		}
		
    return nil
	}

	if len(os.Args) <=1 {
	fmt.Printf("Try '%s  --help' for more options.\n",os.Args[0])
	}
	
	err :=app.Run(os.Args)
	if err !=nil {
		log.Fatal(err)
		
	}
  }

  func exec_sql(){
	rows, err := conn.Query(query)
	if err != nil {
		
		panic(err.Error()) 
		
	}
	defer rows.Close()

	columns, err := rows.Columns()
	if err !=nil{
		panic(err.Error())
	}
	values := make([]sql.RawBytes, len(columns))
	scanArgs := make([]interface{}, len(values))
	for i := range values {
		scanArgs[i] = &values[i]
	}
	
	for rows.Next(){
		err=rows.Scan(scanArgs...)
		if err !=nil{
			panic(err.Error())
		}
		var value string
		for _,col := range values{
			if col ==nil{
				value=""
			}else{
				value=string(col)
			}
			fmt.Println(value)

		}
		//fmt.Println("-----------------------------------")

	}
	if err = rows.Err(); err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
}

  func Open() {
	value, err :=conn.Prepare("select value_in_use from sys.configurations where name = 'xp_cmdshell'")
	if err != nil {
		log.Fatal("Prepare failed:", err.Error())
	}
	defer value.Close()

	row := value.QueryRow()
	//var somenumber int64
	var v int
	err = row.Scan( &v)
	if err != nil {
		log.Fatal("Query failed:", err.Error())
	}
	if v==1 {
		fmt.Printf("xp_cmdshell Enabled\n" )
		 
	}else{
		fmt.Printf("Open xp_cmdshell...\n")
	stmt, err := conn.Prepare("EXEC sp_configure 'show advanced options', 1;RECONFIGURE;EXEC sp_configure 'xp_cmdshell', 1;RECONFIGURE;")
	if err != nil {
		//fmt.Println("Query Error", err)
		return
	}
	
	defer stmt.Close()
	stmt.Query()
	

	}
	return
	
}

  func Openole() {
	value, err :=conn.Prepare("select value_in_use from sys.configurations where name = 'Ole Automation Procedures'")
	if err != nil {
		log.Fatal("Prepare failed:", err.Error())
	}
	defer value.Close()

	row := value.QueryRow()
	//var somenumber int64
	var v int
	err = row.Scan( &v)
	if err != nil {
		log.Fatal("Query failed:", err.Error())
	}
	if v==1 {
		fmt.Printf("sp_oacreate Enabled\n" )
		 
	}else{
		fmt.Printf("Open sp_oacreate...\n")
	stmt, err := conn.Prepare("EXEC sp_configure 'show advanced options', 1;RECONFIGURE WITH OVERRIDE;EXEC sp_configure 'Ole Automation Procedures',1;RECONFIGURE WITH OVERRIDE;")
	if err != nil {
		//fmt.Println("Query Error", err)
		return
	}
	
	defer stmt.Close()
	stmt.Query()
	

	}
	return
	
}


func os_shell(){
	rows, err := conn.Query(`exec master..xp_cmdshell '` + cmd + `' `)
	if err != nil {
		
		panic(err.Error()) 
		
	}
	defer rows.Close()

	columns, err := rows.Columns()
	if err !=nil{
		panic(err.Error())
	}
	values := make([]sql.RawBytes, len(columns))
	scanArgs := make([]interface{}, len(values))
	for i := range values {
		scanArgs[i] = &values[i]
	}
	
	for rows.Next(){
		err=rows.Scan(scanArgs...)
		if err !=nil{
			panic(err.Error())
		}
		var value string
		for _,col := range values{
			if col ==nil{
				value=""
			}else{
				value=string(col)
			}
			fmt.Println(value)

		}

	}
	if err = rows.Err(); err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
}

func os_shellt(){
	rows, err := conn.Query(`declare @shell int exec sp_oacreate 'wscript.shell',@shell output exec sp_oamethod @shell,'run',null,'` + cmd1 + `' `)
	if err != nil {
		
		panic(err.Error()) 
		
	}
	defer rows.Close()

	columns, err := rows.Columns()
	if err !=nil{
		panic(err.Error())
	}
	values := make([]sql.RawBytes, len(columns))
	scanArgs := make([]interface{}, len(values))
	for i := range values {
		scanArgs[i] = &values[i]
	}
	
	for rows.Next(){
		err=rows.Scan(scanArgs...)
		if err !=nil{
			panic(err.Error())
		}
		var value string
		for _,col := range values{
			if col ==nil{
				value=""
			}else{
				value=string(col)
			}
			fmt.Println(value)

		}

	}
	if err = rows.Err(); err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
}