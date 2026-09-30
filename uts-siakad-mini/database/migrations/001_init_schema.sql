-- Tabel users (1-1 ke students)
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    email VARCHAR(255) UNIQUE NOT NULL,
    password VARCHAR(255) NOT NULL,
    role VARCHAR(20) NOT NULL CHECK (role IN ('admin', 'mahasiswa'))
);

-- Tabel students (punya deleted_at untuk soft delete)
CREATE TABLE students (
    id SERIAL PRIMARY KEY,
    user_id INT UNIQUE NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    nim VARCHAR(12) UNIQUE NOT NULL,
    nama VARCHAR(255) NOT NULL,
    prodi VARCHAR(100) NOT NULL,
    angkatan INT NOT NULL,
    ipk_terakhir DECIMAL(3,2) DEFAULT 0.00,
    deleted_at TIMESTAMPTZ DEFAULT NULL
);

-- Tabel courses
CREATE TABLE courses (
    id SERIAL PRIMARY KEY,
    kode_mk VARCHAR(20) UNIQUE NOT NULL,
    nama_mk VARCHAR(255) NOT NULL,
    sks INT NOT NULL,
    semester INT NOT NULL,
    kuota INT NOT NULL
);

-- Tabel enrollments (KRS)
CREATE TABLE enrollments (
    id SERIAL PRIMARY KEY,
    student_id INT NOT NULL REFERENCES students(id) ON DELETE CASCADE,
    course_id INT NOT NULL REFERENCES courses(id) ON DELETE CASCADE,
    tahun_akademik VARCHAR(20) NOT NULL,
    created_at TIMESTAMPTZ DEFAULT NOW(),
    -- Constraint Unik agar tidak ambil MK sama 2 kali di tahun yang sama
    CONSTRAINT unique_student_course_year UNIQUE (student_id, course_id, tahun_akademik)
);