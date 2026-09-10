CREATE TABLE IF NOT EXISTS room_reservations (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    room_id UUID NOT NULL REFERENCES rooms(id) ON DELETE CASCADE,
    booking_id BIGINT NOT NULL,
    check_in_date DATE NOT NULL,
    check_out_date DATE NOT NULL,
    status VARCHAR(20) NOT NULL DEFAULT 'ACTIVE',
    expires_at TIMESTAMP WITH TIME ZONE,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT chk_room_reservations_dates CHECK (check_out_date > check_in_date),
    CONSTRAINT chk_room_reservations_status CHECK (status IN ('ACTIVE', 'RELEASED', 'EXPIRED'))
);

CREATE INDEX IF NOT EXISTS idx_room_reservations_room_dates
    ON room_reservations (room_id, check_in_date, check_out_date);

CREATE UNIQUE INDEX IF NOT EXISTS uq_room_reservations_active_booking
    ON room_reservations (booking_id)
    WHERE status = 'ACTIVE';
