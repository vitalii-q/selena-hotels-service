package services

import (
	"errors"
	"testing"

	"github.com/gofrs/uuid"
	"github.com/shopspring/decimal"
	"github.com/vitali-q/selena-hotels-service/internal/models"
)

type roomRepositoryStub struct {
	created *models.Room
}

func (r *roomRepositoryStub) CreateRoom(room *models.Room) error {
	r.created = room
	return nil
}

func (r *roomRepositoryStub) GetRoomByID(id uuid.UUID) (*models.Room, error) {
	return nil, errors.New("not implemented")
}

func (r *roomRepositoryStub) GetRoomsByHotelID(hotelID uuid.UUID) ([]models.Room, error) {
	return nil, errors.New("not implemented")
}

func (r *roomRepositoryStub) UpdateRoom(room *models.Room) error {
	return errors.New("not implemented")
}

func (r *roomRepositoryStub) DeleteRoom(id uuid.UUID) error {
	return errors.New("not implemented")
}

func TestCreateRoomRejectsInvalidPrice(t *testing.T) {
	service := NewRoomService(&roomRepositoryStub{})
	room := validRoom()
	room.PricePerNight = decimal.Zero

	_, err := service.CreateRoom(room)
	if !errors.Is(err, ErrRoomPricePerNightInvalid) {
		t.Fatalf("expected ErrRoomPricePerNightInvalid, got %v", err)
	}
}

func TestCreateRoomRejectsInvalidCapacity(t *testing.T) {
	service := NewRoomService(&roomRepositoryStub{})
	room := validRoom()
	room.Capacity = 0

	_, err := service.CreateRoom(room)
	if !errors.Is(err, ErrRoomCapacityInvalid) {
		t.Fatalf("expected ErrRoomCapacityInvalid, got %v", err)
	}
}

func TestCreateRoomPersistsValidRoom(t *testing.T) {
	repository := &roomRepositoryStub{}
	service := NewRoomService(repository)
	room := validRoom()

	created, err := service.CreateRoom(room)
	if err != nil {
		t.Fatalf("CreateRoom returned an error: %v", err)
	}
	if created != room || repository.created != room {
		t.Fatal("CreateRoom did not pass the valid room to the repository")
	}
}

func validRoom() *models.Room {
	return &models.Room{
		HotelID:       uuid.Must(uuid.NewV4()),
		Number:        "101",
		Type:          "STANDARD",
		Capacity:      2,
		PricePerNight: decimal.RequireFromString("99.99"),
		Active:        true,
	}
}
