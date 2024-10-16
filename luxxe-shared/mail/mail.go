package mail

type MailInfoStruct struct {
	From         string            `json:"from" bson:"from"`
	To           string            `json:"to" bson:"to"`
	Subject      string            `json:"subject" bson:"subject"`
	Text         string            `json:"text" bson:"text"`
	// TemplateID   string            `json:"templateID" bson:"templateID"`
	TemplateData map[string]string `json:"templateData" bson:"templateData"`
}
