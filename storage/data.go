package storage

type Data interface {
	Raw() []byte
}

func InitDataSlice() []Data {
	return make([]Data, 0)
}

func CreateData(val []byte) Data {
	return &RawData{
		data: val,
	}
}

type RawData struct {
	data []byte
}

func (rd *RawData) Raw() []byte {
	return rd.data
}
