package mysql

import (
	klog "goklmmx/lib/log"
)

func SelectAccount(deviceId string)  int {
	rows, err := MysqlClient.Query("select accountId from account where deviceId=?;", deviceId)
	if err != nil {
		klog.Klog.Println(err)
	}
	if rows.Next(){
		var id int
		if err := rows.Scan(&id); err != nil {
			klog.Klog.Println(err)
			return 0
		}
		return id
	}
	return 0
}
func CreateAccount(accountId int, deviceId string)  error {
	stmt , err := MysqlClient.Prepare("insert into account (accountId, deviceId) values (?,?);")
	if err != nil {
		klog.Klog.Println(err)
		return err
	}
	_, err = stmt.Exec(accountId, deviceId)
	if err != nil {
		klog.Klog.Println(err)
		return err
	}
	return nil
}

