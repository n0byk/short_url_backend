package filestorage

import (
	"bufio"
	"encoding/json"
	"errors"
	"os"

	entities "github.com/n0byk/short_url_backend/dataservice/entities"
)

type storageSet struct {
	file    *os.File
	encoder *json.Encoder
}

func NewStorageSet(fileName string) (*storageSet, error) {
	file, err := os.OpenFile(fileName, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0777)
	if err != nil {
		return nil, err
	}
	return &storageSet{
		file:    file,
		encoder: json.NewEncoder(file),
	}, nil
}

func (p *storageSet) WriteURL(url *entities.URLCatalog) error {
	return p.encoder.Encode(&url)
}
func (p *storageSet) CloseURLCatalog() error {
	return p.file.Close()
}

type storageGet struct {
	file    *os.File
	decoder *json.Decoder
}

func NewStorageGet(fileName string) (*storageGet, error) {
	file, err := os.OpenFile(fileName, os.O_RDONLY|os.O_CREATE, 0777)
	if err != nil {
		return nil, err
	}
	return &storageGet{
		file:    file,
		decoder: json.NewDecoder(file),
	}, nil
}
func (c *storageGet) GetURL(ShortURL string) (string, error) {

	urls := bufio.NewScanner(c.file)

	for urls.Scan() {
		bytes := []byte(urls.Text())
		var item entities.URLCatalog
		err := json.Unmarshal(bytes, &item)
		if err == nil && item.ShortURL == ShortURL {
			return item.FullURL, nil
		}

	}
	return "", errors.New("Cant_get_URL")
}

func (c *storageGet) CloseURLCatalog() error {
	return c.file.Close()
}

// var events = []*Event{
// 	{
// 		ID:       1,
// 		CarModel: "Lada",
// 		Price:    400000,
// 	},
// 	{
// 		ID:       2,
// 		CarModel: "Mitsubishi",
// 		Price:    650000,
// 	},
// 	{
// 		ID:       3,
// 		CarModel: "Toyota",
// 		Price:    800000,
// 	},
// 	{
// 		ID:       4,
// 		CarModel: "BMW",
// 		Price:    875000,
// 	},
// 	{
// 		ID:       5,
// 		CarModel: "Mercedes",
// 		Price:    999999,
// 	},
// }

// func main() {
// 	fileName := "events.log"
// 	defer os.Remove(fileName)
// 	producer, err := NewProducer(fileName)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer producer.Close()
// 	consumer, err := NewConsumer(fileName)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer consumer.Close()
// 	for _, event := range events {
// 		if err := producer.WriteEvent(event); err != nil {
// 			log.Fatal(err)
// 		}
// 		readedEvent, err := consumer.ReadEvent()
// 		if err != nil {
// 			log.Fatal(err)
// 		}
// 		fmt.Println(readedEvent)
// 	}
// }
