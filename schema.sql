CREATE TABLE IF NOT EXISTS Students (
    id SERIAL PRIMARY KEY,
    first_name VARCHAR(50) NOT NULL,
    last_name VARCHAR(50) NOT NULL
);

CREATE TABLE IF NOT EXISTS Subjects (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL
);

CREATE TABLE IF NOT EXISTS Marks (
    id SERIAL PRIMARY KEY,
    student_id INT,
    subject_id INT,
    mark SMALLINT CHECK (mark >= 1 AND mark <= 5),
    created DATE,
    FOREIGN KEY (student_id) REFERENCES Students(id),
    FOREIGN KEY (subject_id) REFERENCES Subjects(id)
);

