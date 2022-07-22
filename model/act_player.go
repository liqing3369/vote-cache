package model

import "time"

type Player struct {
	Id            int       `json:"id" gorm:"primary_key"`
	ActActivityId int       `json:"act_activity_id"` // 活动id
	PlayerCode    int       `json:"player_code"`     // 选手编号
	PlayerName    string    `json:"player_name"`     // 选手名称
	PhoneNumber   string    `json:"phone_number"`    // 电话号码
	Introduce     string    `json:"introduce"`       // 介绍
	VirtualVote   int       `json:"virtual_vote"`    // 虚拟票数
	Popularity    int       `json:"popularity"`      //  人气，访问次数
	AuditStatus   int       `json:"audit_status"`    // 审核状态，0-未审，1-审核通过，2-审核不通过
	VoteAction    int       `json:"vote_action"`     // 投票资格，1-正常，2-取消投票资格，3-禁止投钻石票，4-禁止投普通票
	Rank          int       `json:"rank"`            // 排名
	CreateTime    time.Time `json:"create_time"`
	UpdateTime    time.Time `json:"update_time"`
}

func (Player) TableName() string {
	return "act_player"
}

type PlayerImage struct {
	Id          int    `json:"id" gorm:"primary_key"`
	ActPlayerId int    `json:act_player_id`
	Image       string `json:image`
}

func (PlayerImage) TableName() string {
	return "act_player_image"
}
