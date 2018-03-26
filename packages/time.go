// +build go1.6

package packages

import (
	"reflect"
	"time"
)

func init() {
	if _, ok := Packages["time"]; !ok {
		Packages["time"] = make(map[string]interface{})
	}
	if _, ok := PackageTypes["time"]; !ok {
		PackageTypes["time"] = make(map[string]interface{})
	}
	var duration time.Duration
	var location time.Location
	var month time.Month
	var parseerror time.ParseError
	var ticker time.Ticker
	var time_ time.Time
	var timer time.Timer
	var weekday time.Weekday
	Packages["time"]["ANSIC"] = time.ANSIC
	Packages["time"]["After"] = time.After
	Packages["time"]["AfterFunc"] = time.AfterFunc
	Packages["time"]["April"] = time.April
	Packages["time"]["August"] = time.August
	Packages["time"]["Date"] = time.Date
	Packages["time"]["December"] = time.December
	Packages["time"]["February"] = time.February
	Packages["time"]["FixedZone"] = time.FixedZone
	Packages["time"]["Friday"] = time.Friday
	Packages["time"]["Hour"] = time.Hour
	Packages["time"]["January"] = time.January
	Packages["time"]["July"] = time.July
	Packages["time"]["June"] = time.June
	Packages["time"]["Kitchen"] = time.Kitchen
	Packages["time"]["LoadLocation"] = time.LoadLocation
	Packages["time"]["Local"] = time.Local
	Packages["time"]["March"] = time.March
	Packages["time"]["May"] = time.May
	Packages["time"]["Microsecond"] = time.Microsecond
	Packages["time"]["Millisecond"] = time.Millisecond
	Packages["time"]["Minute"] = time.Minute
	Packages["time"]["Monday"] = time.Monday
	Packages["time"]["Nanosecond"] = time.Nanosecond
	Packages["time"]["NewTicker"] = time.NewTicker
	Packages["time"]["NewTimer"] = time.NewTimer
	Packages["time"]["November"] = time.November
	Packages["time"]["Now"] = time.Now
	Packages["time"]["October"] = time.October
	Packages["time"]["Parse"] = time.Parse
	Packages["time"]["ParseDuration"] = time.ParseDuration
	Packages["time"]["ParseInLocation"] = time.ParseInLocation
	Packages["time"]["RFC1123"] = time.RFC1123
	Packages["time"]["RFC1123Z"] = time.RFC1123Z
	Packages["time"]["RFC3339"] = time.RFC3339
	Packages["time"]["RFC3339Nano"] = time.RFC3339Nano
	Packages["time"]["RFC822"] = time.RFC822
	Packages["time"]["RFC822Z"] = time.RFC822Z
	Packages["time"]["RFC850"] = time.RFC850
	Packages["time"]["RubyDate"] = time.RubyDate
	Packages["time"]["Saturday"] = time.Saturday
	Packages["time"]["Second"] = time.Second
	Packages["time"]["September"] = time.September
	Packages["time"]["Since"] = time.Since
	Packages["time"]["Sleep"] = time.Sleep
	Packages["time"]["Stamp"] = time.Stamp
	Packages["time"]["StampMicro"] = time.StampMicro
	Packages["time"]["StampMilli"] = time.StampMilli
	Packages["time"]["StampNano"] = time.StampNano
	Packages["time"]["Sunday"] = time.Sunday
	Packages["time"]["Thursday"] = time.Thursday
	Packages["time"]["Tick"] = time.Tick
	Packages["time"]["Tuesday"] = time.Tuesday
	Packages["time"]["UTC"] = time.UTC
	Packages["time"]["Unix"] = time.Unix
	Packages["time"]["UnixDate"] = time.UnixDate
	Packages["time"]["Wednesday"] = time.Wednesday
	PackageTypes["time"]["Duration"] = reflect.TypeOf(&duration).Elem()
	PackageTypes["time"]["Location"] = reflect.TypeOf(&location).Elem()
	PackageTypes["time"]["Month"] = reflect.TypeOf(&month).Elem()
	PackageTypes["time"]["ParseError"] = reflect.TypeOf(&parseerror).Elem()
	PackageTypes["time"]["Ticker"] = reflect.TypeOf(&ticker).Elem()
	PackageTypes["time"]["Time"] = reflect.TypeOf(&time_).Elem()
	PackageTypes["time"]["Timer"] = reflect.TypeOf(&timer).Elem()
	PackageTypes["time"]["Weekday"] = reflect.TypeOf(&weekday).Elem()
}
