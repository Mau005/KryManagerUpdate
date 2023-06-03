package schema

import (
	"fmt"
)

type querys struct {
	mysql mysql
}

/* generate new query struct, for manipuled mysql */
func NewQuerys(path string) *querys {
	mysql := getDataMysql(path)
	var q querys
	q.mysql = *mysql
	return &q
}

/*Preparaing change position for playing in update */
func (q *querys) ChangePosition(pos_x, pos_y, pos_z int32) error {
	err := q.mysql.Execute(fmt.Sprintf("UPDATE players SET posx = %v, posy = %v, posz = %v;", pos_x, pos_y, pos_z))
	if err != nil {
		return err
	}
	return nil
}
