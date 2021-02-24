package model

type Message struct {
	ID             uint    `gorm:"not null;autoIncrement;primaryKey"`
	IDInSMSService uint    `gorm:"not null"`
	DepartmentID   uint    `gorm:"not null"`
	ReceiverID     uint    `gorm:"not null"`
	Cost           float32 `gorm:"not null"`
}

type SendUserMessageRequest struct {
	DepartmentID      uint `json:"DepartmentID" validate:"required"`
	AnswerID          uint `json:"AnswerID" validate:"required"`
}

type SendInterviewMessageRequest struct {
	DepartmentID      uint `json:"DepartmentID" validate:"required"`
	AnswerID          uint `json:"AnswerID" validate:"required"`
	InterviewID       uint `json:"InterviewID" validate:"required"`
}

type MessageRequest struct {
	DepartmentID      uint `json:"DepartmentID" validate:"required"`
	AnswerID          uint `json:"AnswerID" validate:"required"`
	InterviewID       uint `json:"InterviewID"`
}

type MessageCostAPI struct {
	Cost    float32
	Balance float32
}

func CreateMessage(requestMessage *Message) error {
	result := gormDb.Create(requestMessage)
	return result.Error
}

var MessageTemplate = []string{
	`【求是潮纳新平台】亲爱的#name#，我们已经收到了您的报名表，请核对下列信息，重复提交可覆盖原表。学号：#stuid# 。部门志愿： #intent# 。我们将在之后通过短信通知您具体的初试信息，敬请留意。感谢您报名#event#，期待您的加入！`,
	`【求是潮纳新平台】#name#同学，我们已为你生成出了#depart##interview#的时间与地点，请点击以下链接进行选择与确认。（注意：在链接中我们不会要求你输入任何诸如密码等的敏感信息）感谢参与#event#。 #url#`,
	`【求是潮纳新平台】#name#同学您好，您刚刚提交的#depart##interview#已经确认。时间为#time#，地点为#location#。请提前10分钟到场准备。感谢报名#event#。`,
	`【求是潮纳新平台】#name#同学，我们很遗憾地通知你，你未能成功加入#association##depart#。感谢你的支持，欢迎继续关注#association#。`,
}
