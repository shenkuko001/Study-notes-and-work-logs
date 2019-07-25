package dao

func InsertUser(session xorm.Session,user model.User) error{
	if _,err:=session.Table("tb_user").Insert(user);err!=nil{
		logger.Error(err)
		return err
	}
	return nil
}