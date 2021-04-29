package squirrel

import "fmt"

type indexPart part

func newIndexPart(pred interface{}, args ...interface{}) Sqlizer {
	return &indexPart{pred: pred, args: args}
}

func (p indexPart) ToSql() (sql string, args []interface{}, err error) {
	switch pred := p.pred.(type) {
	case nil:
		// no-op
	case rawSqlizer:
		return pred.toSqlRaw()
	case Sqlizer:
		return pred.ToSql()
	case map[string]interface{}:
		return Eq(pred).ToSql()
	case string:
		sql = pred
		args = p.args
	default:
		err = fmt.Errorf("expected string-keyed map or string, not %T", pred)
	}
	return
}
