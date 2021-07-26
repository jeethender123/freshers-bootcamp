package Models

type Student_table struct {
	Id      uint   `json:"id"`
	Name    string `json:"name"`
	LastName string `json:"last_name"`
	DOB     string  `json:"DOB"`
	Address string `json:"address"`
	Subject   string `json:"subject"`
	Marks    string `json:"marks"`
}
func (b *Student_table) TableName() string {
	return "student_table"
}
