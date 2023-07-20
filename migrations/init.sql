CREATE TABLE organizations (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    bin TEXT,
    name TEXT,
    address TEXT,
    workers INTEGER
);

CREATE TABLE projects (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT,
    due_date TEXT,
    owner INTEGER,
    participants TEXT
);
