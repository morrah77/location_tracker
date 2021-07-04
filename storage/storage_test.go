package storage

import (
	"morrah77.org/location_tracker/domain"
	"testing"
)
// TODO add more tests
func TestNewStorage(t *testing.T) {
	actualResult := NewStorage()
	if actualResult == nil {
		t.Fatalf(`NewStorage() = nil, expected %#v`, actualResult)
	}
}

func TestStore(t *testing.T) {
	storage := NewStorage()
	value := &domain.Location{11.22, 33.44}
	actualResult := storage.Store(`123abc`, value)
	if actualResult != nil {
		t.Fatalf(`Store("123abc", %#v) = %v, expected nil`, value, actualResult)
	}
}

func TestFetch(t *testing.T) {
	storage := NewStorage()
	value := &domain.Location{11.22, 33.44}
	err := storage.Store(`123abc`, value)
	actualResult, err := storage.Fetch(`123abc`, -1)
	if (actualResult.([]*domain.Location))[0] != value {
		t.Fatalf(`Fetch("123abc") = %#v, %v, expected %#v, nil`, actualResult, err, value)
	}
}
