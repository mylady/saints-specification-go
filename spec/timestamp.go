package spec

import (
	"database/sql/driver"
	"fmt"
	"strconv"
	"time"
)

//Timestamp :custom format type for time.Time,it accepts epoch from json and output time with timezone.
type Timestamp struct {
	value time.Time
}

//String :to string
func (ts *Timestamp) String() string {
	return fmt.Sprintf("\"%s\"", ts.value.Format("2006-01-02 15:04:05"))
}

//MarshalJSON :marshal to json
func (ts *Timestamp) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("%s", ts.value.Format("2006-01-02 15:04:05"))), nil
}

//UnmarshalJSON :unmarshal from json
func (ts *Timestamp) UnmarshalJSON(data []byte) (err error) {
	s := string(data)
	var i int
	if i, err = strconv.Atoi(s); err == nil {
		*ts = Timestamp{
			value: time.Unix(int64(i), 0),
		}
	} else {
		var t time.Time
		if t, err = time.ParseInLocation("2006-01-02 15:04:05", s, time.Local); err == nil {
			*ts = Timestamp{
				value: t,
			}
		}
	}
	return err
}

//MarshalText :marshal to text
func (ts *Timestamp) MarshalText() ([]byte, error) {
	return []byte(fmt.Sprintf("%s", ts.value.Format("2006-01-02 15:04:05"))), nil
}

// Value insert timestamp into db need this function.
func (ts Timestamp) Value() (driver.Value, error) {
	var zeroTime time.Time
	if ts.value.UnixNano() == zeroTime.UnixNano() {
		return nil, nil
	}
	return ts.value, nil
}

// Scan valueof time.Time
func (ts *Timestamp) Scan(v interface{}) error {
	tValue, ok := v.(time.Time)
	if ok {
		*ts = Timestamp{value: tValue}
		return nil
	}
	return fmt.Errorf("can not convert %v to timestamp", v)
}
