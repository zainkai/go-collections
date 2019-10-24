package set

import "testing"

func TestAddToSet(t *testing.T) {
	s := New()

	s.Add(1)
	if _, ok := s.Data[1]; !ok {
		t.Errorf("set did not add a new key")
	}
	s.Add("test")
	if _, ok := s.Data["test"]; !ok {
		t.Errorf("set did not add a new key")
	}
}

func TestHasInSet(t *testing.T) {
	s := New()
	s.Add(1)
	s.Add("test")

	if !s.Has(1) {
		t.Errorf("set should contain key 1")
	}
	if !s.Has("test") {
		t.Errorf("set should contain key \"test\"")
	}
}

func TestDeleteInSet(t *testing.T) {
	s := New()

	s.Add(1)
	s.Delete(1)
	if _, ok := s.Data[1]; ok {
		t.Errorf("set did not delete key")
	}

	s.Add("test")
	s.Delete("test")
	if _, ok := s.Data["test"]; ok {
		t.Errorf("set did not delete key")
	}
}

func TestLenSet(t *testing.T) {
	s := New()

	s.Add(1)
	if s.Len() != 1 {
		t.Errorf("Length was not correct")
	}

	s.Add("test")
	if s.Len() != 2 {
		t.Errorf("Length was not correct")
	}
}

func TestKeysInSet(t *testing.T) {
	s := New()
	s.Add(1)
	s.Add("test")

	keys := s.Keys()
	if v, ok := keys[0].(string); ok && v != "test" {
		t.Errorf("key \"test\" was not saved to array")
	}
	if v, ok := keys[0].(int); ok && v != 1 {
		t.Errorf("key 1 was not saved to array")
	}
	if v, ok := keys[1].(string); ok && v != "test" {
		t.Errorf("key \"test\" was not saved to array")
	}
	if v, ok := keys[1].(int); ok && v != 1 {
		t.Errorf("key 1 was not saved to array")
	}
}

func TestForEachInSet(t *testing.T) {
	s := New()
	s.Add(1)
	s.Add("test")

	callCount := 0

	callback := func(key interface{}) {
		if v, ok := key.(string); ok && v != "test" {
			t.Errorf("key \"test\" was not saved to array")
		}
		if v, ok := key.(int); ok && v != 1 {
			t.Errorf("key 1 was not saved to array")
		}

		callCount++
	}

	s.ForEach(callback)

	if callCount != 2 {
		t.Errorf("callback function was not called in forEach loop")
	}
}
