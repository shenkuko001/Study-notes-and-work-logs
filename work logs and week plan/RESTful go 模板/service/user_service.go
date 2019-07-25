

package service


func InserUser(user *model.User) error{
	session := xorm.Session{}
	defer session.Close() //确保session 关闭
	

	if err:=session.Begin();err!=nil{
		logger.Error(err)
		return err
	}
	if err:=dao.InsertUser(session,user);err!=nil{
		_ = session.Rollback()
		logger.Error()
	}
	if err:=session.Commit();err!=nil{
		logger.Error(error)
		return err
	}
	return nil
}