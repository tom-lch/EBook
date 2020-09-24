package database


func Setup(drive string) {
	dbType := drive
	if dbType == "mysql" {
		var db = new(MySQL)
		db.Setup()
	}
}
