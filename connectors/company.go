package connectors

type Company struct {
	ExtId         string `bson:"extId" json:"extId"`
	CompanyName   string `bson:"companyName" json:"companyName"`
	CompanyType   string `bson:"companyType" json:"companyType"`
	INN           string `bson:"inn" json:"inn"`
	Website       string `bson:"website" json:"website"`
	People        string `bson:"people" json:"people"`
	Phones        string `bson:"phones" json:"phones"`
	Emails        string `bson:"emails" json:"emails"`
	Address       string `bson:"address" json:"address"`
	ITEquipment   string `bson:"ITEquipment" json:"ITEquipment"`
	SoftwareName  string `bson:"softwareName" json:"softwareName"`
	IsMinPromTorg bool   `bson:"isMinPromTorg" json:"isMinPromTorg"`
	IsMincifr     bool   `bson:"isMincifr" json:"isMincifr"`
	Description   string `bson:"description" json:"description"`
	Status        string `bson:"status" json:"status"`
	Approbation   string `bson:"approbation" json:"approbation"`
	Feedback      string `bson:"feedback" json:"feedback"`
	Comments      string `bson:"comments" json:"comments"`
}
