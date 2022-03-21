package model

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"myproject/global"
	"myproject/pkg/setting"
	"time"
)

const (
	STATE_OPEN = 1
	STATE_CLOSE = 0
)

type Model struct {
	ID uint32 `gorm:"primary_key" json:"id"`
	CreatedBy string `json:"created_by"`
	ModifiedBy string `json:"modified_by"`
	CreatedOn uint32 `json:"created_on"`
	ModifiedOn uint32 `json:"modified_on"`
	DeletedOn uint32 `json:"deleted_on"`
	IsDel uint8 `json:"is_del"`
}

func NewDBEngine(databaseSetting *setting.DatabaseSettingS) (*gorm.DB,error) {
	db,err := gorm.Open(databaseSetting.DBType,
		fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=%s&parseTime=%t&loc=Local",
			databaseSetting.Username,
			databaseSetting.Password,
			databaseSetting.Host,
			databaseSetting.DBName,
			databaseSetting.Charset,
			databaseSetting.ParseTime,
			))
	if err != nil{
		return nil,err
	}
	//开启详细日志
	if global.ServerSetting.RunMode == "debug"{
		db.LogMode(true)
	}
	//默认使用单数表
	db.SingularTable(true)
	//注册公共字段的回调行为
	db.Callback().Create().Replace("gorm:update_time_stamp", updateTimeStampForCreateCallback)
	db.Callback().Update().Replace("gorm:update_time_stamp", updateTimeStampForUpdateCallback)
	db.Callback().Delete().Replace("gorm:delete", deleteCallback)
	//设置最大连接数等
	db.DB().SetMaxIdleConns(databaseSetting.MaxIdleConns)
	db.DB().SetMaxOpenConns(databaseSetting.MaxOpenConns)
	return db,nil
}

//model callback回调处理公共字段

//新增行为的回调方法
func updateTimeStampForCreateCallback(scope *gorm.Scope)  {
	if !scope.HasError() {
		nowTime := time.Now().Unix()
		if createTimeField, ok := scope.FieldByName("CreatedOn"); ok {
			if createTimeField.IsBlank {
				_ = createTimeField.Set(nowTime)
			}
		}

		if modifyTimeField, ok := scope.FieldByName("ModifiedOn"); ok {
			if modifyTimeField.IsBlank {
				_ = modifyTimeField.Set(nowTime)
			}
		}
	}
}

//更新行为的回调方法
func updateTimeStampForUpdateCallback(scope *gorm.Scope)  {
	if _,ok := scope.Get("gorm:update_column"); !ok {
		_ = scope.SetColumn("ModifiedOn", time.Now().Unix())
	}
}

//删除行为的回调方法
func deleteCallback(scope *gorm.Scope)  {
	if !scope.HasError(){
		var extraOption string
		if str,ok := scope.Get("gorm:delete_option");ok{
			extraOption = fmt.Sprint(str)
		}
		deleteOnField,hasDeleteOnField := scope.FieldByName("DeleteOn")
		isDelField,hasIsDelField := scope.FieldByName("IsDel")
		if scope.Search.Unscoped && hasDeleteOnField &&hasIsDelField{
			now := time.Now().Unix()
			//调用scope.QuotedTableName方法获取当前引用的表名
			//在设置完必要参数后，调用scope.CombinedConditionSql方法完成SQL语句的拼接
			sqlStr := fmt.Sprintf("UPDATE %v SET %v=%v,%v=%v%v%v",
				scope.QuotedTableName(),
				scope.Quote(deleteOnField.DBName),
				scope.AddToVars(now),
				scope.Quote(isDelField.DBName),
				scope.AddToVars(1),
				addExtraSpaceIfExist(scope.CombinedConditionSql()),
				addExtraSpaceIfExist(extraOption))
			scope.Raw(sqlStr).Exec()
		}else {
			sqlStr := fmt.Sprintf("UPDATE FROM %v%v%v",
				scope.QuotedTableName(),
				addExtraSpaceIfExist(scope.CombinedConditionSql()),
				addExtraSpaceIfExist(extraOption))
			scope.Raw(sqlStr).Exec()
		}
	}
}

func addExtraSpaceIfExist(str string) string {
	if str != ""{
		return " "+str
	}
	return ""
}




