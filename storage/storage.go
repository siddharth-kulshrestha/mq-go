package storage

import "fmt"

type Storage interface {
	Get(offset int) (Data, error)
	Set(Data) (int, error)
}

type inMemory struct {
	data []Data
}

func Init() Storage {
	return initInMem()
}

func initInMem() Storage {
	return &inMemory{
		data: InitDataSlice(),
	}
}

func (im *inMemory) Get(offset int) (Data, error) {
	if len(im.data) > offset {
		return nil, fmt.Errorf("offset %d cannot be greater than length of data %d", offset, len(im.data))
	}
	return im.data[offset], nil
}

func (im *inMemory) Set(val Data) (int, error) {
	im.data = append(im.data, val)
	return len(im.data) - 1, nil
}
