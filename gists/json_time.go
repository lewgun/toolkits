type JSONTime time.Time

func (t JSONTime) MarshalJSON() ([]byte, error) {

	const format = "2006-01-02 15:04:05"

	b := make([]byte, 0, len(format)+2)
	b = append(b, '"')
	b = time.Time(t).AppendFormat(b, format)
	b = append(b, '"')
	return b, nil

}
