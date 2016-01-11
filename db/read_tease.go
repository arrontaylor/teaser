package db

func ReadTease(teaseid int, tousername string) {
	Connection.Exec("UPDATE teases SET read=1 WHERE teaseid=? AND tousername=?", teaseid, tousername)
}
