package data

import (
	"fmt"
	"os"
	"testing"
)

var list = New()

func TestMain(m *testing.M) {
	game := InitiateGame()
	list.Add(game)
	os.Exit(m.Run())
}

func TestAdd(t *testing.T) {
	if len(list.Games) == 0 {
		t.Errorf("Add() failed. Expected: %d, Got: %d", len(list.Games), 0)
	}
}

func TestGetAll(t *testing.T) {
	result := list.GetAll()

	if len(result) == 0 {
		t.Errorf("GetAll() failed. Expected: %d, Got: %d", len(list.Games), 0)
	}
}

func TestInitiateGame(t *testing.T) {
	result := InitiateGame()
	fmt.Println(result.Board)
	if len(result.Board) != 8 {
		t.Errorf("InitiateGame() failed. Expected: %d, Got: %d", len(result.Board), 0)
	}
}
