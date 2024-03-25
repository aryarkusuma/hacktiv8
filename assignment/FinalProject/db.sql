CREATE TABLE shuttles (
    id UUID DEFAULT uuid_generate_v4() PRIMARY KEY,
    shuttle_type VARCHAR(100) NOT NULL,
    seats INT NOT NULL,
    start_date DATE NOT NULL,
    route_start VARCHAR(100) NOT NULL,
    route_end VARCHAR(100) NOT NULL,
    created_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE users (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    username VARCHAR(100) UNIQUE NOT NULL,
    password VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);


CREATE TABLE reservation (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    shuttle_id UUID NOT NULL,
    reserv_name VARCHAR(255) NOT NULL,
    seat_number VARCHAR(50) NOT NULL,
    user_id UUID NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (shuttle_id) REFERENCES shuttle(id),
    FOREIGN KEY (user_id) REFERENCES users(id)
);
