// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package entity

import "log"

const TableNameTSiteRatePlan = "t_site_rate_plan"

// TSiteRatePlan mapped from table <t_site_rate_plan>
type TSiteRatePlan struct {
	ID         int    `gorm:"column:id;type:bigint;primaryKey;autoIncrement:true" json:"id"`
	SiteID     string `gorm:"column:site_id;type:varchar(64);not null" json:"site_id"`       // 站点ID
	StartTime  string `gorm:"column:start_time;type:varchar(64);not null" json:"start_time"` // 开始
	EndTime    string `gorm:"column:end_time;type:varchar(64);not null" json:"end_time"`     // 结束
	Type       int   `gorm:"column:type;type:tinyint;not null;default:1" json:"type"`       // 类型:1-波峰,2-波谷
	Status     int   `gorm:"column:status;type:tinyint;not null;default:1" json:"status"`   // 状态:0-不使用,1-未删除,2-删除
	CreateTime int    `gorm:"column:create_time;type:int unsigned;autoCreateTime" json:"create_time"`
	UpdateTime int    `gorm:"column:update_time;type:int unsigned;autoUpdateTime" json:"update_time"`
}

// TableName TSiteRatePlan's table name
func (*TSiteRatePlan) TableName() string {
	return TableNameTSiteRatePlan
}

func (m *TSiteRatePlan)IsValid() bool {
	//传入时间格式验证
	start := m.StartTime
	end := m.EndTime
	if len(start) != 5 || start[2:3] != ":" {
		log.Println("长度不正确")
		return false
	}
	if len(end) != 5 || end[2:3] != ":" {
		log.Println("长度不正确")
		return false
	}
	shour := start[:2]
	sminu := start[3:]
	ehour := end[:2]
	eminu := end[3:]

	if shour < ehour {
		log.Println("格式正确", shour, ehour, sminu, eminu)
		return true
	} else if shour == ehour {
		if sminu <= eminu {
			log.Println("格式正确", shour, ehour, sminu, eminu)
			return true
		}
	}

	log.Println("格式错误", shour, ehour, sminu, eminu)
	return false
}
