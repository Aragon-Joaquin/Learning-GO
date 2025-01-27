package helloExport

import (
	"errors"
	"fmt"
	"math/rand"
)

//uppercase functions are exported while lowercase ones are not
func Hello(name string) (string, error) {
	if(name == "") {return "", errors.New("name is empty")}
	var message string = fmt.Sprintf(randomFormat(), name)
	return message, nil
}

func HelloMultiple(names []string) (map[string]string, error) {
	messages := make(map[string]string) // a map consist of key value pairs like {"name1": "message1", "name2": "message2"}

	for _, nameValue := range names { // i suppose the first value is the key
        message, err := Hello(nameValue)
        if err != nil {
            return nil, err
        }
		messages[nameValue] = message
	}
	return messages, nil
}

func randomFormat() string {
	formats := []string{ //! slice of strings
		"Hi, %v. Welcome!",
		"Hello, %v. Good to see you!",
		"Hi, %v. Glad to see you!",
	}

	return formats[rand.Intn(len(formats))]
}