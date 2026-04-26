/*юзеры*/
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    email TEXT UNIQUE NOT NULL,
    password TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT NOW()
);


/*расписание*/
CREATE TABLE shedules (
    id SERIAL PRIMARY KEY,
    group_name TEXT NOT NULL,
    data TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT NOW()
);

CREATE TABLE substitutions (
    id SERIAL PRIMARY KEY,
    shedule_id INT REFERENCES shedules(id),
    date DATE NOT NULL,
    lessons_old TEXT,
    lessons_new TEXT,
    created_at TIMESTAMP DEFAULT NOW()
);