package viewmodels

import (
	"ItcastCms/models"
	"github.com/astaxie/beego/orm"
)

//定义搜索对象
type UserData struct {
	Name string
	Remark string
	PageSize int
	PageIndex int
	TotalCount int64
	Start int
}
func (this *UserData)SearchUserData()([]models.UserInfo){
	//模糊查询实现搜索功能
	o := orm.NewOrm()
	temp := o.QueryTable("user_info")
	//判断名字和备注是否为空
	if this.Name != "" {
		temp = temp.Filter("user_name__contains", this.Name)
	} else if this.Remark != "" {
		temp = temp.Filter("remark__contains", this.Remark)
	}
	//过滤未被删除的数据，都为空都显示
	temp = temp.Filter("del_flag", 0)
	//实现分页
	this.TotalCount, _ = temp.Count()
	//查询
	var userInfos=[]models.UserInfo{}
	temp.OrderBy("id").Limit(this.PageSize, this.Start).All(&userInfos)
	///resp["data"] = &userInfos
	//return resp
	return userInfos
}
