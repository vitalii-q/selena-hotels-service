package models

import (
	"testing"

	"github.com/gofrs/uuid"
	"github.com/shopspring/decimal"
)

func TestRoomBeforeCreateAssignsUUID(t *testing.T) {
	room := Room{
		HotelID:       uuid.Must(uuid.NewV4()),
		Number:        "101",
		Type:          "STANDARD",
		Capacity:      2,
		PricePerNight: decimal.RequireFromString("99.99"),
		Active:        true,
	}

	if err := room.BeforeCreate(nil); err != nil {
		t.Fatalf("BeforeCreate returned an error: %v", err)
	}
	if room.ID == uuid.Nil {
		t.Fatal("BeforeCreate did not assign a room UUID")
	}
}

func TestRoomBeforeCreatePreservesExistingUUID(t *testing.T) {
	existingID := uuid.Must(uuid.NewV4())
	room := Room{ID: existingID}

	if err := room.BeforeCreate(nil); err != nil {
		t.Fatalf("BeforeCreate returned an error: %v", err)
	}
	if room.ID != existingID {
		t.Fatal("BeforeCreate replaced an existing room UUID")
	}
}
