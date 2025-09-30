package main

type Event struct {
	ID     int
	Format Format
}

type Racer struct {
	Event int
	Name  string
	Mmr   float64
	Score float64
}
