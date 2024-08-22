package main

import "testing"

func TestSearch(t *testing.T) {
	t.Run("Search", func(t *testing.T) {
		dict := Dict{"test": "this is just a test"}
		got, e := dict.Search("test")
		want := "this is just a test"

		if e != nil {
			t.Errorf("Error returned, but should not")
		}
		if got != want {
			t.Errorf("got %s != want %s", got, want)
		}
	})
	t.Run("Search, word not found", func(t *testing.T) {
		dict := Dict{"test": "this is just a test"}
		_, e := dict.Search("test_")

		if e == nil {
			t.Errorf("Error not returned, but should")
		}
		if e.Error() != ErrorWordNotFound.Error() {
			t.Errorf("got %s != want %s", e.Error(), ErrorWordNotFound.Error())
		}
	})
	t.Run("Add", func(t *testing.T) {
		dict := Dict{"test": "this is just a test"}
		e := dict.Add("aa", "aaa")
		if e != nil {
			t.Errorf("Error returned, word exits, but should not")
		}
		got, e := dict.Search("aa")
		want := "aaa"

		if e != nil {
			t.Errorf("Error returned, but should not")
		}
		if got != want {
			t.Errorf("got %s != want %s", got, want)
		}
		e2 := dict.Add("aa", "aaa")
		if e2 == nil {
			t.Errorf("Error should be returned")
		}
		if e2.Error() != ErrorWordFound.Error() {
			t.Errorf("got %s != want %s", e2.Error(), ErrorWordFound.Error())
		}
	})
	t.Run("Update", func(t *testing.T) {
		dict := Dict{}
		want := "321"
		e := dict.Add("test", "123")
		if e != nil {
			t.Errorf("Error(%s) returned, but should not: Add", e.Error())
		}
		e2 := dict.Update("test", want)
		if e2 != nil {
			t.Errorf("Error(%s) returned, but should not: Update", e2.Error())
		}
		got, e3 := dict.Search("test")
		if e3 != nil {
			t.Errorf("Error(%s) returned, but should not: Search", e3.Error())
		}
		if got != want {
			t.Errorf("Got %s != want %s", got, want)
		}
	})
	t.Run("Update, word not found", func(t *testing.T) {
		dict := Dict{}
		e2 := dict.Update("test", "123")
		if e2 == nil {
			t.Errorf("Error not returned, but should : Update")
		}
		if e2.Error() != ErrorWordNotFound.Error() {
			t.Errorf("got %s != want %s", e2.Error(), ErrorWordNotFound.Error())
		}
	})
	t.Run("Delete", func(t *testing.T) {
		dict := Dict{}
		e := dict.Add("test", "123")
		if e != nil {
			t.Errorf("Error(%s) returned, but should not: Add", e.Error())
		}
		e2 := dict.Delete("test")
		if e2 != nil {
			t.Errorf("Error(%s) returned, but should not: Delete", e2.Error())
		}
		_, e3 := dict.Search("test")
		if e3 == nil {
			t.Errorf("Error not returned, but should: Search")
		}
	})
	t.Run("Delete, word not found", func(t *testing.T) {
		dict := Dict{}
		e2 := dict.Delete("test")
		if e2 == nil {
			t.Errorf("Error not returned, but should : Update")
		}
		if e2.Error() != ErrorWordNotFound.Error() {
			t.Errorf("got %s != want %s", e2.Error(), ErrorWordNotFound.Error())
		}
	})
}
