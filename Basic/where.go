package Basic

import "fmt"

func (d *Db)Where(pattern string, value string) (*Db) {
	d.where = fmt.Sprintf(pattern, value)
	return d
}
