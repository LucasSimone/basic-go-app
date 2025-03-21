-- Creates the table schema for the routes table
CREATE TABLE climbs(
    id serial PRIMARY KEY,
    title varchar(255),
    category varchar(255),
    grade varchar(5),
    setter varchar(255),
    time_created timestamp
);

-- Inserts an new route(row) into the routes table
INSERT INTO climbs (
    title, 
    category,
    grade,
    setter,
    time_created
) VALUES (
    'Underswing',
    'boulder',
    'V3',
    'John Doe',
    '2025-02-28 15:30:45'
);

-- Insert dummy data to populate the routes table
INSERT INTO climbs (title, category, grade, setter, time_created) VALUES
('Overhang Fury', 'boulder', 'V5', 'Alice Smith', '2024-08-15 14:20:00'),
('Crux Delight', 'sport', '5.10b', 'Bob Johnson', '2024-09-03 16:30:15'),
('Pocket Rocket', 'boulder', 'V6', 'Emma Brown', '2024-09-25 17:00:30'),
('Sloper Madness', 'sport', '5.11a', 'Charlie White', '2024-10-10 10:45:20'),
('Crimp City', 'boulder', 'V2', 'David Lee', '2024-10-28 18:10:05'),
('Dyno King', 'sport', '5.12c', 'Sophia Green', '2024-11-12 08:30:45'),
('Slab Wizard', 'boulder', 'V1', 'Ethan Black', '2024-11-27 19:00:00'),
('Campus Crusher', 'sport', '5.10d', 'Mia Harris', '2024-12-05 09:15:25'),
('Heel Hook Heaven', 'boulder', 'V3', 'Olivia Scott', '2024-12-12 19:45:50'),
('Edge of Glory', 'sport', '5.11c', 'Noah Adams', '2024-12-20 13:10:15'),
('Gaston Gambit', 'boulder', 'V4', 'Liam Wilson', '2025-01-04 20:45:30'),
('High Step Hero', 'sport', '5.9', 'Ava Martinez', '2025-01-15 07:00:40'),
('Compression King', 'boulder', 'V7', 'Lucas Hall', '2025-01-20 21:30:10'),
('Toe Hook Trick', 'sport', '5.12a', 'Harper Young', '2025-01-28 12:00:05'),
('Mantle Mania', 'boulder', 'V8', 'Evelyn Carter', '2025-02-02 22:15:20'),
('Balance Beam', 'sport', '5.11b', 'James Walker', '2025-02-10 22:45:35'),
('Tension Tango', 'boulder', 'V3', 'Charlotte King', '2025-02-18 23:00:50'),
('Flash Pump', 'sport', '5.10a', 'Benjamin Turner', '2025-02-24 23:30:00'),
('Heel Hook Hurdle', 'boulder', 'V4', 'Isabella Wright', '2025-02-26 23:45:25'),
('Endurance Test', 'sport', '5.13a', 'Mason Rivera', '2025-02-28 00:15:10');