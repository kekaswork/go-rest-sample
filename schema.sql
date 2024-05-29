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

-- Inserting sample data rows
INSERT INTO Students (first_name, last_name) VALUES ('Almaz', 'Galimianov');
INSERT INTO Students (first_name, last_name) VALUES ('Kate', 'Galimianova');
INSERT INTO Students (first_name, last_name) VALUES ('Ivan', 'Pupkin');

INSERT INTO Subjects (name) VALUES ('Math');
INSERT INTO Subjects (name) VALUES ('Physics');
INSERT INTO Subjects (name) VALUES ('Programming');

INSERT INTO Marks (student_id, subject_id, mark, mark_date) VALUES (1, 1, 4, '2023-05-01');
INSERT INTO Marks (student_id, subject_id, mark, mark_date) VALUES (1, 2, 3, '2023-05-02');
INSERT INTO Marks (student_id, subject_id, mark, mark_date) VALUES (2, 2, 5, '2023-05-03');
INSERT INTO Marks (student_id, subject_id, mark, mark_date) VALUES (2, 3, 2, '2023-05-04');
INSERT INTO Marks (student_id, subject_id, mark, mark_date) VALUES (3, 1, 1, '2023-05-05');
INSERT INTO Marks (student_id, subject_id, mark, mark_date) VALUES (3, 3, 4, '2023-05-06');