package conv

import "time"

//dont know whether it just used in customer contract or in any other
//so lets stay it here for now

func ParseDate(s string) time.Time {
	t, err := time.Parse("01-2006", s) // "01-2006" — month-year template
	if err != nil {
		return time.Time{}
	} // or handle error
	return t
}

func ParseDatePtr(s *string) *time.Time {
	if s == nil {
		return nil
	}
	t, err := time.Parse("01-2006", *s)
	if err != nil {
		return nil
	}
	return &t
}

func FormatPtrDate(t *time.Time, template string) *string {
	if t == nil {
		return nil
	}
	//dateStr := t.Format("01-2006")
	dateStr := t.Format(template)
	return &dateStr
}
