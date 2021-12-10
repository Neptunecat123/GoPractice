package main

type TopicMysql struct {
	// topic_url,author_url,author_name,like_num,
	// collect_num,comment_num,create_time,latest_time,is_elite
	Id         int    `gorm:"column:id;not null;type:int(4) primary key auto_increment;comment:'ID'"`
	TopicName  string `gorm"column:topic_name;type:varchar(100)"`
	TopicUrl   string `gorm:"column:topic_url;type:varchar(50);index:idx_name"`
	AuthorUrl  string `gorm:"column:author_url;type:varchar(50)"`
	AuthorName string `gorm:"column:author_name;type:varchar(30)"`
	LikeNum    int    `gorm:"column:like_num;type:int(4)"`
	CollectNum int    `gorm:"collect_num;type:int(4)"`
	CommentNum int    `gorm:"comment_num;type:int(4)"`
	CreateTime string `gorm:"create_time;type:varchar(20)"`
	LatestTime string `gorm:"latest_time;type:varchar(20)"`
}

type CommentMysql struct {
	// topic_id, comment_content, comment_author_url, comment_author_name, comment_time,
	// comment_floor, comment_like_num, comment_for_comment(comment_for_comment_content)
	Id                int    `gorm:"column:id;not null;type:int(4) primary key auto_increment"`
	TopicUrl          string `gorm:"column:topic_url;not null;type:varchar(50)"`
	AuthorUrl         string `gorm:"column:author_url;type:varchar(50)"`
	AuthorName        string `gorm:"colunm:author_name;type:varchar(30)"`
	CreateTime        string `gorm:"colunm:comment_time;type:varchar(20)"`
	Floor             int    `gorm:"column:comment_floor;type:int(4)"`
	LikeNum           int    `gorm:"column:like_num;type:int(4)"`
	CommentForComment int    `gorm:"column:comment_for_comment;type:int(4)"`
}

type TopicMongo struct {
}

type CommentMongo struct {
}
