package model

import "time"

type ActActivity struct {
	Id                       uint      `json:"id" gorm:"primary_key"`
	Title                    string    `json:"title"`                         // 标题
	SignupStartTime          time.Time `json:"signup_start_time"`             // 报名开始时间
	SignupEndTime            time.Time `json:"signup_end_time"`               // 报名结束时间
	VoteStartTime            time.Time `json:"vote_start_time"`               // 投票开始时间
	VoteEndTime              time.Time `json:"vote_end_time"`                 // 投票结束时间
	MaxVoteNumPerPerson      int       `json:"max_vote_num_per_person"`       // 每人最多投票数量
	MaxVoteNumPerPsersonDays int       `json:"max_vote_num_per_pserson_days"` // 每人每天最多投票数量
	MaxVoteNumPerPlayer      int       `json:"max_vote_num_per_player"`       // 每人对一选手最多投票数量
	NeedAudit                int       `json:"need_audit"`                    // 报名是否需要审核，1-需要，2-不需要
	UserId                   int       `json:"user_id"`                       // 资料录入员id
	Notice                   string    `json:"notice"`                        // 公告内容
	Introduce                string    `json:"introduce"`                     // 投票介绍内容
	Prize                    string    `json:"prize"`                         // 奖品
	Remark                   string    `json:"remark"`                        // 备注
	Popularity               int       `json:"popularity"`                    // 人气，访问量
	CreateTime               time.Time `json:"create_time"`
	Status                   int       `json:"status"`
	RealVote                 int       `json:"real_vote"`    // 实票
	TodayVote                int       `json:"today_vote"`   // 今日普票
	TodayAmount              float32   `json:"today_amount"` // 今日金额
	AmountTotal              float32   `json:"amount_total"` // 总金额
	Isopen                   int       `json:"isopen"`
}
