-- phpMyAdmin SQL Dump
-- version 5.1.1
-- https://www.phpmyadmin.net/
--
-- Host: localhost
-- Generation Time: Jul 24, 2021 at 01:15 PM
-- Server version: 10.4.20-MariaDB
-- PHP Version: 8.0.8

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `embrio4`
--

-- --------------------------------------------------------

--
-- Table structure for table `cabang`
--

CREATE TABLE `cabang` (
  `id` int(6) UNSIGNED NOT NULL,
  `id_cabang` int(10) NOT NULL,
  `nama_cabang` varchar(50) NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp(),
  `created_by` varchar(50) DEFAULT NULL,
  `updated_at` timestamp NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp(),
  `updated_by` varchar(50) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- --------------------------------------------------------

--
-- Table structure for table `cicilan`
--

CREATE TABLE `cicilan` (
  `id` int(6) UNSIGNED NOT NULL,
  `id_mantri` int(6) NOT NULL,
  `id_pinjaman` int(6) NOT NULL,
  `payment_date` date NOT NULL,
  `next_payment_date` date DEFAULT NULL,
  `jumlah_cicilan` double(15,2) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- --------------------------------------------------------

--
-- Table structure for table `jenis_aksi`
--

CREATE TABLE `jenis_aksi` (
  `id` int(6) UNSIGNED NOT NULL,
  `id_mantri` int(6) NOT NULL,
  `nama_aksi` varchar(50) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- --------------------------------------------------------

--
-- Table structure for table `jenis_kolektebilitas`
--

CREATE TABLE `jenis_kolektebilitas` (
  `id` int(6) UNSIGNED NOT NULL,
  `id_pinjaman` int(6) NOT NULL,
  `nama_jenis_kolektibilitas` varchar(50) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- --------------------------------------------------------

--
-- Table structure for table `jenis_tunggakan`
--

CREATE TABLE `jenis_tunggakan` (
  `id` int(6) UNSIGNED NOT NULL,
  `id_pinjaman` int(6) NOT NULL,
  `nama_tunggakan` varchar(50) DEFAULT NULL,
  `nominal_tunggakan` double(15,2) DEFAULT NULL,
  `tanggal_menunggak` date DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- --------------------------------------------------------

--
-- Table structure for table `mantri`
--

CREATE TABLE `mantri` (
  `id` int(6) UNSIGNED NOT NULL,
  `username_pn` varchar(50) NOT NULL,
  `password` varchar(255) NOT NULL,
  `nama_mantri` varchar(255) NOT NULL,
  `unit_kerja` varchar(255) DEFAULT NULL,
  `kode_branch` varchar(50) DEFAULT NULL,
  `jabatan` varchar(50) DEFAULT NULL,
  `id_cabang` int(50) DEFAULT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp(),
  `created_by` varchar(50) DEFAULT NULL,
  `updated_at` timestamp NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp(),
  `updated_by` varchar(50) DEFAULT NULL,
  `last_login` timestamp NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- --------------------------------------------------------

--
-- Table structure for table `nasabah`
--

CREATE TABLE `nasabah` (
  `id` int(6) UNSIGNED NOT NULL,
  `id_mantri` int(6) NOT NULL,
  `nama_debitur` varchar(60) NOT NULL,
  `nomor_rekening` varchar(60) NOT NULL,
  `cif_no` varchar(60) NOT NULL,
  `no_ktp` varchar(60) NOT NULL,
  `alamat` varchar(255) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- --------------------------------------------------------

--
-- Table structure for table `pinjaman`
--

CREATE TABLE `pinjaman` (
  `id` int(6) UNSIGNED NOT NULL,
  `id_nasabah` int(6) NOT NULL,
  `plafond` double(15,2) NOT NULL,
  `jenis_pinjaman` varchar(50) DEFAULT NULL,
  `jangka_waktu` date DEFAULT NULL,
  `tanggal_realisasi` date DEFAULT NULL,
  `tanggal_jatuh_tempo` date DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- --------------------------------------------------------

--
-- Table structure for table `users`
--

CREATE TABLE `users` (
  `id` int(6) UNSIGNED NOT NULL,
  `username_pn` varchar(255) NOT NULL,
  `password` varchar(255) DEFAULT NULL,
  `nama` varchar(255) DEFAULT NULL,
  `unit_kerja` varchar(255) DEFAULT NULL,
  `kode_branch` varchar(255) DEFAULT NULL,
  `jabatan` varchar(255) DEFAULT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp(),
  `created_by` varchar(255) DEFAULT NULL,
  `updated_at` timestamp NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp(),
  `updated_by` varchar(255) DEFAULT NULL,
  `last_login` timestamp NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp(),
  `groups` varchar(255) DEFAULT NULL,
  `phone_number` varchar(15) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `users`
--

INSERT INTO `users` (`id`, `username_pn`, `password`, `nama`, `unit_kerja`, `kode_branch`, `jabatan`, `created_at`, `created_by`, `updated_at`, `updated_by`, `last_login`, `groups`, `phone_number`) VALUES
(1, 'admin', '$2a$04$64WUUz3c2XBb/7jviayM0eFLmGHB4SJ2u6VEIAzxEvV9nWctCSWS2', 'admin', '', '', '', '2021-07-24 09:11:23', 'admin', '2021-07-24 09:11:23', 'admin', '2021-07-24 09:11:23', 'admin', '081393384561'),
(6, 'coba1', '$2a$04$.Dz2ovefz2m5a0YbDjVvCOaDx7oE8mt.x10ikvh68So6QCkp.UM3y', 'coba1', '', '', '', '2021-07-24 11:12:21', 'admin', '2021-07-24 11:12:21', 'admin', '2021-07-24 11:12:21', '', '');

--
-- Indexes for dumped tables
--

--
-- Indexes for table `cabang`
--
ALTER TABLE `cabang`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `cicilan`
--
ALTER TABLE `cicilan`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `jenis_aksi`
--
ALTER TABLE `jenis_aksi`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `jenis_kolektebilitas`
--
ALTER TABLE `jenis_kolektebilitas`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `jenis_tunggakan`
--
ALTER TABLE `jenis_tunggakan`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `mantri`
--
ALTER TABLE `mantri`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `nasabah`
--
ALTER TABLE `nasabah`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `pinjaman`
--
ALTER TABLE `pinjaman`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `users`
--
ALTER TABLE `users`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `username_pn` (`username_pn`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `cabang`
--
ALTER TABLE `cabang`
  MODIFY `id` int(6) UNSIGNED NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT for table `cicilan`
--
ALTER TABLE `cicilan`
  MODIFY `id` int(6) UNSIGNED NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT for table `jenis_aksi`
--
ALTER TABLE `jenis_aksi`
  MODIFY `id` int(6) UNSIGNED NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT for table `jenis_kolektebilitas`
--
ALTER TABLE `jenis_kolektebilitas`
  MODIFY `id` int(6) UNSIGNED NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT for table `jenis_tunggakan`
--
ALTER TABLE `jenis_tunggakan`
  MODIFY `id` int(6) UNSIGNED NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT for table `mantri`
--
ALTER TABLE `mantri`
  MODIFY `id` int(6) UNSIGNED NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT for table `nasabah`
--
ALTER TABLE `nasabah`
  MODIFY `id` int(6) UNSIGNED NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT for table `pinjaman`
--
ALTER TABLE `pinjaman`
  MODIFY `id` int(6) UNSIGNED NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT for table `users`
--
ALTER TABLE `users`
  MODIFY `id` int(6) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=7;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
