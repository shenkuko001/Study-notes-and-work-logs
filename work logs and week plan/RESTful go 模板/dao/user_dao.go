package dao

func InsertUser(session *xorm.Session,user *model.User) error{
	if _,err:=session.Table("tb_user").Insert(user);err!=nil{
		logger.Error(err)
		return err
	}
	return nil
}

func UpdateSysUser(session *xorm.Session, user *model.User, columns ...string) error{
	session.Table("tb_user")
	if columns != nil{
		session.Cols(columns...)
	}
	if _,err:=session.Where().Update(user);err!=nil{
		logger.Error(err)
		return err
	}
	return nil
}

func FindSysUserById(session *xorm.Session, id int64)(*model.SysUser,error){
	u := new(model.User)
	if ext,err:=session.Table("tb_user").Where("tb_user.id = ?",id).Get(user);err!=nil{
		return nil,err
	}
	return user,nil
}

func FindSysUserCount(session *xorm.session, user *model.User) (int64, error){
	session.Table("tb_user")
	if user.id!=0{
		session.Where("tb_user.id = ? ", user.id)
	}

	if user.name!=""{
		session.Where("tb_user.name = ?", user.name)
	}

	return session.Count(new(model.User))
}

func FindSysUserList(session *xorm.session, user *model.User,size int, offset int) ([]*model.User, error]){
	session.Table("tb_user")
	if user.id!=0{
		session.Where("tb_user.id = ? ", user.id)
	}

	if user.name!=""{
		session.Where("tb_user.name = ?", user.name)
	}
	
	session.OrderBy("tb_user.id desc").Limit(size,offset)
	var userlist []*model.User
	if err:=session.Find(userlist);err!=nil{
		logger.Error(err)
		return nil,err
	}
	if userlist == nil{
		return make([]*model.User,0),nil
	}else{
	return userlist,nil
}

func FindSysUserPager(session *xrom.Session, user *model.User, pager *model.Pager)(*Pager, error){
	count,err := FindSysUserCount(session,user) // 获取总记录数
	if err!=nil{
		logger.Error(err)
		return nil,err
	}

	pager.SetTotal(int(count)) //设置pager对象的总数
	userlist,err := FindsysUserList(session,user,pager.size,pager.offset)
}


type Pager struct{
	Page int `json:"page"`
	Size int `json:"page_size"`
	offset int `json:"-"`
	Total int `json:"total"`
	List interface{} `json:"list"`
	Next bool `json:"next"`
	Prev bool `json:"prev"`
}
