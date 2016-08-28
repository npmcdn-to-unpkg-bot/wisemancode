package model

import (
	"time"
	"wisemancode/log"
	"wisemancode/utils"
)

// Head 表结构必须字段
type Head struct {
	ID         string    //主键id
	CTime      time.Time //创建时间
	UpdateTime time.Time //更新时间
	Version    string    //版本号
}

//Header 表头数据需要进行对必要字段进行默认设置
type Header interface {
	GetID() string
	CreateTime() time.Time
}

//GetID 创建表id string 类型
func (head *Head) GetID() string {
	log.Logger.Info("获取表ID")
	id := utils.Rand()
	log.Logger.Info("获取表ID %+s", id)
	return id
}

//CreateTime 表生成时间
func (head *Head) CreateTime() time.Time {
	return time.Now()
}

//NewHead 创建头信息
func NewHead() *Head {
	h := new(Head)
	h.ID = h.GetID()
	h.CTime = h.CreateTime()
	return h
}
