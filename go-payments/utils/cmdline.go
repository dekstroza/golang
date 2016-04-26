package utils

import "flag"

//CmdLineFlags contains parsed flags
type CmdLineFlags struct {
	BindAddress        *string
	BindPort           *int
	DbHost             *string
	DbPort             *int
	DbUsername         *string
	DbPassword         *string
	DbName             *string
	DbSchema           *string
	MaxConnections     *int
	MaxIdleConnections *int
}

//ParseCmdArgs will parse command arguments and return them as CmdLineFlags struct
func ParseCmdArgs() CmdLineFlags {
	cmdLineFlags := CmdLineFlags{}
	cmdLineFlags.BindAddress = flag.String("b", "127.0.0.1", "Bind address for rest api.")
	cmdLineFlags.BindPort = flag.Int("p", 3000, "Port on which http will receive connections.")
	cmdLineFlags.DbHost = flag.String("dbHost", "localhost", "Address or name of the database host.")
	cmdLineFlags.DbPort = flag.Int("dbPort", 5432, "Port on which database server will accept connections.")
	cmdLineFlags.DbUsername = flag.String("user", "postgres", "Username for database connection")
	cmdLineFlags.DbPassword = flag.String("password", "", "Password for database connection")
	cmdLineFlags.DbName = flag.String("db", "postgres", "Name of the database to connect to.")
	cmdLineFlags.DbSchema = flag.String("schema", "public", "Name of database schema that will be used.")

	cmdLineFlags.MaxIdleConnections = flag.Int("maxIdle", 20, "Maximum number of idle connections in the pool.")
	cmdLineFlags.MaxConnections = flag.Int("maxConnections", 60, "Maximum number of database connections.")

	flag.Parse()
	return cmdLineFlags
}
