CREATE TABLE IF NOT EXISTS rooms (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    hotel_id UUID NOT NULL REFERENCES hotels(id) ON DELETE CASCADE,
    number VARCHAR(50) NOT NULL,
    type VARCHAR(50) NOT NULL,
    capacity INT NOT NULL CHECK (capacity > 0),
    price_per_night DECIMAL(10,2) NOT NULL CHECK (price_per_night > 0),
    active BOOL NOT NULL DEFAULT TRUE,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT uq_rooms_hotel_number UNIQUE (hotel_id, number)
);

CREATE INDEX IF NOT EXISTS idx_rooms_hotel_id ON rooms (hotel_id);
