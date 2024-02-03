package database

var Databases *BD

type BD struct {
	DBPostgresql *Postgres
}

func Connect(dbType string) {

	Databases = &BD{}

	switch dbType {
	case "postgres":

		Databases.DBPostgresql = new(Postgres).Connect()

	default:
		panic("No dbType found for connection")
	}

}
