package models

import (
	"time"
	"github.com/astaxie/beego/orm"
	"strconv"
	"github.com/astaxie/beego"
)

type UserInfo struct {
	Id        int
	UserName  string
	UserPwd   string
	Remark    string //备注
	AddDate   time.Time
	ModifDate time.Time
	DelFlag   int // 删除标记，软删除
	Roles     []*RoleInfo `orm:"rel(m2m)"`
	UserAction   []*UserAction `orm:"reverse(many)"`
}

func AddUserdb(userInfo *UserInfo, resp map[string]interface{}) map[string]interface{} {
	o := orm.NewOrm()
	_, err := o.Insert(userInfo)
	if err != nil {
		resp["flag"] = "no"
	} else {
		resp["flag"] = "ok"
		resp["rows"] = userInfo
	}
	return resp
}
func DeleteUserdb(userInfo *UserInfo, strId []string, resp map[string]interface{}) map[string]interface{} {
	o := orm.NewOrm()
	//转成int
	for i := 0; i < len(strId); i++ {
		id, _ := strconv.Atoi(strId[i])
		//查询操作
		userInfo.Id = id
		qb, _ := orm.NewQueryBuilder("mysql")
		qb.Select("*").From("user_info").Where("id=?")
		sql := qb.String()
		err := o.Raw(sql, id).QueryRow(userInfo)
		if err != nil {
			resp["flag"] = "no"
		}
		//删除操作
		//userInfo.DelFlag = 1
		//o.Update(userInfo)
		o.Delete(userInfo)
	}
	resp["flag"] = "ok"
	return resp
}
func ShowEditUserdb(userInfo *UserInfo, resp map[string]interface{}) map[string]interface{} {
	o := orm.NewOrm()
	err := o.Read(userInfo)
	if err != nil {
		resp["flag"] = "no"
	} else {
		resp["flag"] = "ok"
		resp["rows"] = userInfo
	}
	return resp
}
func EditUserdb(userInfo *UserInfo, resp map[string]interface{}) map[string]interface{} {
	o := orm.NewOrm()
	_, err := o.Update(userInfo)
	if err != nil {
		resp["flag"] = "no"
	} else {
		resp["flag"] = "ok"
		resp["rows"] = userInfo
	}
	return resp
}
func EditUserRoledb(userId int, idList []int, resp map[string]interface{}) map[string]interface{} {
	//查询到该用户
	o := orm.NewOrm()
	var userInfo = UserInfo{}
	err1 := o.QueryTable("user_info").Filter("id", userId).Filter("del_flag", 0).One(&userInfo)
         if err1!=nil{
         	resp["flag"]="no"
         	return resp
		 }
	//idList包含已有的角色信息，先把用户的角色全部清空
	var roleInfos = []RoleInfo{}
	n, err := o.QueryTable("role_info").Filter("Users__UserInfo__Id", userId).Filter(""+
		"del_flag", 0).All(&roleInfos)
	if err != nil || err1 != nil {
		if n == 0 {
			beego.Info("该用户没有任何角色")
		} else {
			resp["flag"] = "no"
			beego.Error("=========", err, err1)
			return resp
		}
	}
	//o.LoadRelated(&userInfo,"Roles")上面的代码可以用载入关系函数代替（锁定用户。载入关系）
	//获取m2m操作对象,移出用户中的所有角色信息
	//第一个参数为参数的对象，第二个为要操作（添加，删除）的字段
	m2m := o.QueryM2M(&userInfo, "Roles")
	//删除和添加为一致性操作，失败回滚，采用事务
	o.Begin()
	var acterr, acterr1, acterr2 error
	for i := 0; i < len(roleInfos); i++ {
		_, acterr = m2m.Remove(&roleInfos[i])
	}
	//给用户重新添加角色
	var roleInfo = RoleInfo{}
	for i := 0; i < len(idList); i++ {
		//查询到id为idList[i]的RoleInfo对象
		acterr1 = o.QueryTable("role_info").Filter("id", idList[i]).One(&roleInfo)
		_, acterr2 = m2m.Add(&roleInfo)
		if acterr != nil || acterr1 != nil || acterr2 != nil {
			o.Rollback()
			resp["flag"] = "no"
			beego.Info("查询用户角色失败")
			return resp
		}
	}
	// 提交事务
	o.Commit()
	resp["flag"] = "ok"
	return resp
}

