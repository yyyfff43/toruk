package table

type Hachi struct {
	Id    int    `xorm:"not null pk autoincr unique INT(10)"`
	TestA string `xorm:"not null VARCHAR(45)"`
	TestB string `xorm:"VARCHAR(45)"`
}
