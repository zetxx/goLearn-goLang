package main

type Dict map[string]string

const (
	ErrorWordNotFound = DictError("WordNotFound")
	ErrorWordFound    = DictError("WordFound")
)

type DictError string

func (e DictError) Error() string {
	return string(e)
}

func (dict Dict) Search(word string) (string, error) {
	w, f := dict[word]
	if f == false {
		return "", ErrorWordNotFound
	}
	return w, nil
}

func (dict Dict) Add(key, value string) error {
	_, err := dict.Search(key)

	if err == nil {
		return ErrorWordFound
	}
	dict[key] = value
	return nil
}

func (dict Dict) Update(key, newVal string) error {
	_, e := dict.Search(key)
	if e != nil {
		return e
	}
	dict[key] = newVal
	return nil
}

func (dict Dict) Delete(key string) error {
	_, e := dict.Search(key)
	if e != nil {
		return e
	}
	delete(dict, key)
	return nil
}
