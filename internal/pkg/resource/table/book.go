package table

type Book struct {
	BookId   int    `xorm:"not null pk autoincr INT(10)"`
	BookName string `xorm:"not null VARCHAR(45)"`
	Author   string `xorm:"not null VARCHAR(45)"`
	Size     int    `xorm:"not null INT(11)"`
}
