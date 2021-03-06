package db

import (
	"github.com/132yse/acgzone-server/api/def"
	"database/sql"
)

//func AddPageView(pid int) error {
//
//	stmtIns, err := dbConn.Prepare("INSERT INTO pv (pid,pv) VALUES (?,?)")
//
//	stmtUp, err := dbConn.Prepare("UPDATE pv SET pv=pv+1 WHERE pid =?")
//	if err != nil {
//		return err
//	}
//	_, err = stmtIns.Exec(content, ctime, pid, uid)
//	_, err = stmtUp.Exec(&pid)
//	if err != nil {
//		return err
//	}
//	defer stmtIns.Close()
//	defer stmtUp.Close()
//	return nil
//}

func GetCount(pid int) (*def.Count, error) {
	//stmtOut, err := dbConn.Prepare("SELECT pid,pv FROM pv WHERE pid = ?")
	stmtCount, err := dbConn.Prepare("SELECT COUNT(*) FROM comments WHERE pid = ?")
	if err != nil {
		return nil, err
	}

	var pv, cv int
	err = stmtCount.QueryRow(pid).Scan(&cv)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}

	res := &def.Count{Pid: pid, Pv: pv, Cv: cv}

	defer stmtCount.Close()

	return res, nil
}
