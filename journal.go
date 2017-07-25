package main

import (
	"time"
)

type Journal struct {
	Name string
	Entries []*JournalEntry
}

type JournalEntry struct {
	Date time.Time
	Entries []*QuestionAnswerPair
}

type QuestionAnswerPair struct {
	Question string
	Answer string
}

func NewJournal(name string) *Journal {
	return &Journal{
		Name: name,
		Entries: []*JournalEntry{},
	}
}

func NewJournalEntry(date time.Time) *JournalEntry {
	return &JournalEntry {
		Date: date,
		Entries: []*QuestionAnswerPair{},
	}
}

func NewQuestionAnswerPair(question, answer string) *QuestionAnswerPair {
	return &QuestionAnswerPair{
		Question: question,
		Answer: answer,
	}
}

func (j *Journal) AddEntry(entry *JournalEntry) {
	j.Entries = append(j.Entries, entry)
}

func (j *JournalEntry) AddQuestionAnswer(qa *QuestionAnswerPair) {
	j.Entries = append(j.Entries, qa)
}