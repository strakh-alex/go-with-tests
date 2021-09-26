package maps

type Dictionary map[string]string

const (
	ErrNotFount         = DictionaryErr("there is no such word")
	ErrWordExists       = DictionaryErr("the word is exist")
	ErrWordDoesNotExist = DictionaryErr("word does not exist")
)

type DictionaryErr string

func (e DictionaryErr) Error() string {
	return string(e)
}

func (dictionary Dictionary) Search(word string) (string, error) {

	definition, ok := dictionary[word]

	if !ok {
		return "", ErrNotFount
	}

	return definition, nil
}

func (dictionary Dictionary) Add(word, definition string) error {
	_, err := dictionary.Search(word)

	switch err {
	case ErrNotFount:
		dictionary[word] = definition
	case nil:
		return ErrWordExists
	default:
		return err
	}

	return nil
}

func (dictionary Dictionary) Update(word, definition string) error {
	_, err := dictionary.Search(word)

	switch err {
	case ErrNotFount:
		return ErrWordDoesNotExist
	case nil:
		dictionary[word] = definition
	default:
		return err
	}

	return nil
}

func (dictionary Dictionary) Delete(word string) {
	delete(dictionary, word)
}
