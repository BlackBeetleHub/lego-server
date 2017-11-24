package json

type Dictionary interface {
	GetWords() []Word
	GetCountWords() int
}

type Word interface {
	GetWord() string
	GetTranslates() []Translate
	HasLearned() bool
}

type Translate interface {
	GetValue() string
}
