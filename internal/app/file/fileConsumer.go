package file

import (
	"bufio"
	"encoding/json"
	"os"
)

type Consumer interface {
	ReadEvent() (*Event, error) // для чтения события
	Close() error               // для закрытия ресурса (файла)
}

type consumer struct {
	file    *os.File
	scanner *bufio.Scanner
}

func NewConsumer(filename string) (*consumer, error) {
	file, err := os.OpenFile(filename, os.O_RDONLY|os.O_CREATE, 0777)
	if err != nil {
		return nil, err
	}

	return &consumer{
		file: file,
		// создаём новый scanner
		scanner: bufio.NewScanner(file),
	}, nil
}

func (c *consumer) ReadEvent() (*Event, error) {

	if !c.scanner.Scan() {
		return nil, c.scanner.Err()
	}

	data := c.scanner.Bytes()

	event := Event{}
	err := json.Unmarshal(data, &event)
	if err != nil {
		return nil, err
	}

	return &event, nil
}

func (c *consumer) Close() error {
	return c.file.Close()
}
