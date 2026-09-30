-- 1. Insert 1 Admin (Password: admin123)
INSERT INTO users (email, password, role) 
VALUES ('admin@siakad.com', '$2a$10$vI8aWBnW3fID.ZQ4/zo1G.q1lRhrmvlW16Gj9VzY.wzEx.RUyT4jm', 'admin');

-- 2. Insert 20 Akun User untuk Mahasiswa (Password: admin123)
INSERT INTO users (email, password, role) VALUES 
('mhs01@kampus.id', '$2a$10$vI8aWBnW3fID.ZQ4/zo1G.q1lRhrmvlW16Gj9VzY.wzEx.RUyT4jm', 'mahasiswa'),
('mhs02@kampus.id', '$2a$10$vI8aWBnW3fID.ZQ4/zo1G.q1lRhrmvlW16Gj9VzY.wzEx.RUyT4jm', 'mahasiswa'),
('mhs03@kampus.id', '$2a$10$vI8aWBnW3fID.ZQ4/zo1G.q1lRhrmvlW16Gj9VzY.wzEx.RUyT4jm', 'mahasiswa'),
('mhs04@kampus.id', '$2a$10$vI8aWBnW3fID.ZQ4/zo1G.q1lRhrmvlW16Gj9VzY.wzEx.RUyT4jm', 'mahasiswa'),
('mhs05@kampus.id', '$2a$10$vI8aWBnW3fID.ZQ4/zo1G.q1lRhrmvlW16Gj9VzY.wzEx.RUyT4jm', 'mahasiswa'),
('mhs06@kampus.id', '$2a$10$vI8aWBnW3fID.ZQ4/zo1G.q1lRhrmvlW16Gj9VzY.wzEx.RUyT4jm', 'mahasiswa'),
('mhs07@kampus.id', '$2a$10$vI8aWBnW3fID.ZQ4/zo1G.q1lRhrmvlW16Gj9VzY.wzEx.RUyT4jm', 'mahasiswa'),
('mhs08@kampus.id', '$2a$10$vI8aWBnW3fID.ZQ4/zo1G.q1lRhrmvlW16Gj9VzY.wzEx.RUyT4jm', 'mahasiswa'),
('mhs09@kampus.id', '$2a$10$vI8aWBnW3fID.ZQ4/zo1G.q1lRhrmvlW16Gj9VzY.wzEx.RUyT4jm', 'mahasiswa'),
('mhs10@kampus.id', '$2a$10$vI8aWBnW3fID.ZQ4/zo1G.q1lRhrmvlW16Gj9VzY.wzEx.RUyT4jm', 'mahasiswa'),
('mhs11@kampus.id', '$2a$10$vI8aWBnW3fID.ZQ4/zo1G.q1lRhrmvlW16Gj9VzY.wzEx.RUyT4jm', 'mahasiswa'),
('mhs12@kampus.id', '$2a$10$vI8aWBnW3fID.ZQ4/zo1G.q1lRhrmvlW16Gj9VzY.wzEx.RUyT4jm', 'mahasiswa'),
('mhs13@kampus.id', '$2a$10$vI8aWBnW3fID.ZQ4/zo1G.q1lRhrmvlW16Gj9VzY.wzEx.RUyT4jm', 'mahasiswa'),
('mhs14@kampus.id', '$2a$10$vI8aWBnW3fID.ZQ4/zo1G.q1lRhrmvlW16Gj9VzY.wzEx.RUyT4jm', 'mahasiswa'),
('mhs15@kampus.id', '$2a$10$vI8aWBnW3fID.ZQ4/zo1G.q1lRhrmvlW16Gj9VzY.wzEx.RUyT4jm', 'mahasiswa'),
('mhs16@kampus.id', '$2a$10$vI8aWBnW3fID.ZQ4/zo1G.q1lRhrmvlW16Gj9VzY.wzEx.RUyT4jm', 'mahasiswa'),
('mhs17@kampus.id', '$2a$10$vI8aWBnW3fID.ZQ4/zo1G.q1lRhrmvlW16Gj9VzY.wzEx.RUyT4jm', 'mahasiswa'),
('mhs18@kampus.id', '$2a$10$vI8aWBnW3fID.ZQ4/zo1G.q1lRhrmvlW16Gj9VzY.wzEx.RUyT4jm', 'mahasiswa'),
('mhs19@kampus.id', '$2a$10$vI8aWBnW3fID.ZQ4/zo1G.q1lRhrmvlW16Gj9VzY.wzEx.RUyT4jm', 'mahasiswa'),
('mhs20@kampus.id', '$2a$10$vI8aWBnW3fID.ZQ4/zo1G.q1lRhrmvlW16Gj9VzY.wzEx.RUyT4jm', 'mahasiswa');

-- 3. Insert 20 Profil Mahasiswa (user_id 2 sampai 21)
INSERT INTO students (user_id, nim, nama, prodi, angkatan, ipk_terakhir) VALUES 
(2, '187221000001', 'Ali Mahfud', 'Teknik Informatika', 2024, 3.85),
(3, '187221000002', 'Budi Santoso', 'Teknik Informatika', 2024, 3.10),
(4, '187221000003', 'Citra Kirana', 'Sistem Informasi', 2023, 2.95),
(5, '187221000004', 'Dewi Lestari', 'Sistem Informasi', 2023, 3.50),
(6, '187221000005', 'Eko Prasetyo', 'Teknik Komputer', 2022, 2.45),
(7, '187221000006', 'Fajar Ramadhan', 'Teknik Informatika', 2024, 3.20),
(8, '187221000007', 'Gita Gutawa', 'Sistem Informasi', 2025, 0.00),
(9, '187221000008', 'Hadi Sucipto', 'Teknik Informatika', 2023, 3.65),
(10, '187221000009', 'Indah Permata', 'Sistem Informasi', 2022, 3.90),
(11, '187221000010', 'Joko Widodo', 'Teknik Komputer', 2024, 2.75),
(12, '187221000011', 'Kurniawan Dwi', 'Teknik Informatika', 2025, 0.00),
(13, '187221000012', 'Lestari Ayu', 'Sistem Informasi', 2023, 3.15),
(14, '187221000013', 'Muhammad Ilham', 'Teknik Informatika', 2024, 2.80),
(15, '187221000014', 'Nina Marlina', 'Sistem Informasi', 2022, 3.40),
(16, '187221000015', 'Oka Antara', 'Teknik Komputer', 2024, 3.00),
(17, '187221000016', 'Putri Ayu', 'Teknik Informatika', 2023, 2.60),
(18, '187221000017', 'Qori Ahmad', 'Sistem Informasi', 2025, 0.00),
(19, '187221000018', 'Rizky Febian', 'Teknik Informatika', 2022, 3.75),
(20, '187221000019', 'Siti Nurhaliza', 'Sistem Informasi', 2024, 3.33),
(21, '187221000020', 'Taufik Hidayat', 'Teknik Komputer', 2023, 2.10);

-- 4. Insert 10 Mata Kuliah
INSERT INTO courses (kode_mk, nama_mk, sks, semester, kuota) VALUES 
('IF101', 'Algoritma dan Pemrograman', 3, 1, 40),
('IF102', 'Struktur Data', 3, 2, 40),
('IF201', 'Basis Data', 3, 3, 30),
('IF202', 'Pemrograman Web', 3, 3, 35),
('IF301', 'Pemrograman Backend Lanjut', 3, 5, 20),
('IF302', 'Kecerdasan Buatan', 3, 5, 25),
('SI101', 'Pengantar Sistem Informasi', 2, 1, 40),
('SI201', 'Analisis dan Desain Sistem', 3, 3, 30),
('TK101', 'Arsitektur Komputer', 3, 2, 30),
('TK201', 'Jaringan Komputer', 3, 4, 30);