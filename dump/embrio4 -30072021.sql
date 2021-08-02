-- phpMyAdmin SQL Dump
-- version 5.1.1
-- https://www.phpmyadmin.net/
--
-- Host: localhost
-- Generation Time: Jul 30, 2021 at 02:37 AM
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
-- Table structure for table `aksi_misi`
--

CREATE TABLE `aksi_misi` (
  `id` int(6) NOT NULL,
  `id_aksi` int(6) DEFAULT NULL,
  `id_misi` int(6) DEFAULT NULL,
  `description` varchar(255) DEFAULT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp(),
  `updated_at` timestamp NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- --------------------------------------------------------

--
-- Table structure for table `cabang`
--

CREATE TABLE `cabang` (
  `id` int(6) NOT NULL,
  `kode` varchar(255) DEFAULT NULL,
  `nama` varchar(255) DEFAULT NULL,
  `alamat` varchar(255) DEFAULT NULL,
  `description` varchar(255) DEFAULT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp(),
  `updated_at` timestamp NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- --------------------------------------------------------

--
-- Table structure for table `cicilan`
--

CREATE TABLE `cicilan` (
  `id` int(6) NOT NULL,
  `id_pinjaman` int(6) DEFAULT NULL,
  `nominal` double(20,2) DEFAULT NULL,
  `next_payment_date` date DEFAULT NULL,
  `status` int(1) DEFAULT NULL,
  `description` varchar(255) DEFAULT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp(),
  `updated_at` timestamp NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- --------------------------------------------------------

--
-- Table structure for table `dashboard`
--

CREATE TABLE `dashboard` (
  `id` int(6) NOT NULL,
  `os_total` varchar(255) DEFAULT NULL,
  `os_kupedes` varchar(255) DEFAULT NULL,
  `os_kur` varchar(255) DEFAULT NULL,
  `os_kur_supermi` varchar(255) DEFAULT NULL,
  `os_gbt` varchar(255) DEFAULT NULL,
  `dpk_total` varchar(255) DEFAULT NULL,
  `dpk_kecil` varchar(255) DEFAULT NULL,
  `dpk_baru` varchar(255) DEFAULT NULL,
  `dpk_1` varchar(255) DEFAULT NULL,
  `dpk_2` varchar(255) DEFAULT NULL,
  `dpk_3` varchar(255) DEFAULT NULL,
  `dpk_cnpl` varchar(255) DEFAULT NULL,
  `dpk_blm_restruk` varchar(255) DEFAULT NULL,
  `dpk_masih_GP` varchar(255) DEFAULT NULL,
  `dpk_dalam_GP` varchar(255) DEFAULT NULL,
  `lancar_blm_jth_tempo` varchar(255) DEFAULT NULL,
  `npl_total` varchar(255) DEFAULT NULL,
  `npl_kl_kecil` varchar(255) DEFAULT NULL,
  `npl_kl_total` varchar(255) DEFAULT NULL,
  `npl_diragukan_total` varchar(255) DEFAULT NULL,
  `npl_macet` varchar(255) DEFAULT NULL,
  `npl_belum_restruk` varchar(255) DEFAULT NULL,
  `dh_total` varchar(255) DEFAULT NULL,
  `ph_total` varchar(255) DEFAULT NULL,
  `dh_pemasukan_tahunberjalan` varchar(255) DEFAULT NULL,
  `dh_bersaldo_simpanan` varchar(255) DEFAULT NULL,
  `dh_yg_mengangsur` varchar(255) DEFAULT NULL,
  `simpanan_total` varchar(255) DEFAULT NULL,
  `simpanan_topup_besar` varchar(255) DEFAULT NULL,
  `simpanan_pengambilan_besar` varchar(255) DEFAULT NULL,
  `simpanan_besar_tgl_lahir_bln_ini` varchar(255) DEFAULT NULL,
  `restruk_os` varchar(255) DEFAULT NULL,
  `restruk_dalam_grace_periode` varchar(255) DEFAULT NULL,
  `Id_mantri` varchar(50) DEFAULT NULL,
  `description` varchar(255) DEFAULT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp(),
  `updated_at` timestamp NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- --------------------------------------------------------

--
-- Table structure for table `detail_cicilan`
--

CREATE TABLE `detail_cicilan` (
  `id` int(6) NOT NULL,
  `id_cicilan` int(6) DEFAULT NULL,
  `id_kolektebilitas` int(6) DEFAULT NULL,
  `id_aksi_misi` int(6) DEFAULT NULL,
  `tanggal_bayar` date DEFAULT NULL,
  `tgl_janji_bayar` date DEFAULT NULL,
  `keterangan` varchar(255) DEFAULT NULL,
  `description` varchar(255) DEFAULT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp(),
  `updated_at` timestamp NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- --------------------------------------------------------

--
-- Table structure for table `detail_tunggakan`
--

CREATE TABLE `detail_tunggakan` (
  `id` int(6) NOT NULL,
  `id_detail_cicilan` int(6) DEFAULT NULL,
  `id_jenis_tunggakan` int(6) DEFAULT NULL,
  `nominal` double(20,0) DEFAULT NULL,
  `description` varchar(255) DEFAULT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp(),
  `updated_at` timestamp NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- --------------------------------------------------------

--
-- Table structure for table `jenis_aksi`
--

CREATE TABLE `jenis_aksi` (
  `id` int(6) NOT NULL,
  `kode` varchar(20) DEFAULT NULL,
  `nama` int(255) DEFAULT NULL,
  `description` varchar(255) DEFAULT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp(),
  `updated_at` timestamp NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- --------------------------------------------------------

--
-- Table structure for table `jenis_pinjaman`
--

CREATE TABLE `jenis_pinjaman` (
  `id` int(6) NOT NULL,
  `kode` varchar(255) DEFAULT NULL,
  `nama` varchar(255) DEFAULT NULL,
  `description` varchar(255) DEFAULT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp(),
  `updated_at` timestamp NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- --------------------------------------------------------

--
-- Table structure for table `kolektebilitas`
--

CREATE TABLE `kolektebilitas` (
  `id` int(6) NOT NULL,
  `kode` varchar(20) DEFAULT NULL,
  `nama` varchar(255) DEFAULT NULL,
  `description` varchar(255) DEFAULT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp(),
  `updated_at` timestamp NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- --------------------------------------------------------

--
-- Table structure for table `log_activity`
--

CREATE TABLE `log_activity` (
  `id` int(6) NOT NULL,
  `id_mantri` int(6) DEFAULT NULL,
  `id_aksi_misi` int(6) DEFAULT NULL,
  `nama_log` varchar(255) DEFAULT NULL,
  `tgl_activity` timestamp NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp(),
  `waktu_aktifitas` timestamp NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp(),
  `description` varchar(255) DEFAULT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp(),
  `updated_at` timestamp NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- --------------------------------------------------------

--
-- Table structure for table `mantri`
--

CREATE TABLE `mantri` (
  `id` int(6) NOT NULL,
  `id_unit` int(6) DEFAULT NULL,
  `kode` varchar(20) DEFAULT NULL,
  `nama` varchar(255) DEFAULT NULL,
  `alamat` varchar(255) DEFAULT NULL,
  `description` varchar(255) DEFAULT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp(),
  `updated_at` timestamp NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- --------------------------------------------------------

--
-- Table structure for table `misi`
--

CREATE TABLE `misi` (
  `id` int(6) NOT NULL,
  `id_mantri` int(6) DEFAULT NULL,
  `kode` varchar(20) DEFAULT NULL,
  `nama` int(255) DEFAULT NULL,
  `description` varchar(255) DEFAULT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp(),
  `updated_at` timestamp NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- --------------------------------------------------------

--
-- Table structure for table `nasabah`
--

CREATE TABLE `nasabah` (
  `id` int(6) NOT NULL,
  `nama` varchar(255) DEFAULT NULL,
  `no_ktp` varchar(255) DEFAULT NULL,
  `nomor_rekening` varchar(255) DEFAULT NULL,
  `saldo` double(255,2) DEFAULT NULL,
  `no_telepon` varchar(15) DEFAULT NULL,
  `nama_ibu_kandung` varchar(255) DEFAULT NULL,
  `kode_pos_tempat_tinggal` varchar(255) DEFAULT NULL,
  `kelurahan_tempat_tinggal` varchar(255) DEFAULT NULL,
  `kecamatan_tempat_tinggal` varchar(255) DEFAULT NULL,
  `kode_pos_tempat_usaha` varchar(255) DEFAULT NULL,
  `kelurahan_tempat_usaha` varchar(255) DEFAULT NULL,
  `kecamatan_tempat_usaha` varchar(255) DEFAULT NULL,
  `description` varchar(255) DEFAULT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp(),
  `updated_at` timestamp NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- --------------------------------------------------------

--
-- Table structure for table `pinjaman`
--

CREATE TABLE `pinjaman` (
  `id` int(6) NOT NULL,
  `id_mantri` int(6) DEFAULT NULL,
  `id_jenis_pinjaman` int(6) DEFAULT NULL,
  `id_nasabah` int(6) DEFAULT NULL,
  `kode` varchar(20) DEFAULT NULL,
  `jumlah` double(20,2) DEFAULT NULL,
  `jangka_waktu` int(11) DEFAULT NULL,
  `tanggal_pencairan` date DEFAULT NULL,
  `description` varchar(255) DEFAULT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp(),
  `updated_at` timestamp NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- --------------------------------------------------------

--
-- Table structure for table `tunggakan`
--

CREATE TABLE `tunggakan` (
  `id` int(6) NOT NULL,
  `kode` varchar(20) DEFAULT NULL,
  `nama` varchar(255) DEFAULT NULL,
  `description` varchar(255) DEFAULT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp(),
  `updated_at` timestamp NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- --------------------------------------------------------

--
-- Table structure for table `unit`
--

CREATE TABLE `unit` (
  `id` int(6) NOT NULL,
  `id_cabang` int(6) DEFAULT NULL,
  `kode` varchar(255) DEFAULT NULL,
  `nama` varchar(255) DEFAULT NULL,
  `alamat` varchar(255) DEFAULT NULL,
  `description` varchar(255) DEFAULT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp(),
  `updated_at` timestamp NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp()
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
(1, 'admin', '$2a$04$64WUUz3c2XBb/7jviayM0eFLmGHB4SJ2u6VEIAzxEvV9nWctCSWS2', 'admin', '', '', '', '2021-07-26 22:12:40', 'admin', '2021-07-26 22:12:40', 'admin', '2021-07-26 22:12:40', 'admin', '081393384561'),
(6, 'coba1', '$2a$04$.Dz2ovefz2m5a0YbDjVvCOaDx7oE8mt.x10ikvh68So6QCkp.UM3y', 'coba1', '', '', '', '2021-07-24 11:12:21', 'admin', '2021-07-24 11:12:21', 'admin', '2021-07-24 11:12:21', '', '');

--
-- Indexes for dumped tables
--

--
-- Indexes for table `aksi_misi`
--
ALTER TABLE `aksi_misi`
  ADD PRIMARY KEY (`id`),
  ADD KEY `aksi_misi_FK` (`id_misi`),
  ADD KEY `aksi_misi_FK_1` (`id_aksi`);

--
-- Indexes for table `cabang`
--
ALTER TABLE `cabang`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `cicilan`
--
ALTER TABLE `cicilan`
  ADD PRIMARY KEY (`id`),
  ADD KEY `cicilan_FK` (`id_pinjaman`);

--
-- Indexes for table `dashboard`
--
ALTER TABLE `dashboard`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `detail_cicilan`
--
ALTER TABLE `detail_cicilan`
  ADD PRIMARY KEY (`id`),
  ADD KEY `detail_cicilan_FK` (`id_cicilan`),
  ADD KEY `detail_cicilan_FK_1` (`id_kolektebilitas`),
  ADD KEY `detail_cicilan_FK_2` (`id_aksi_misi`);

--
-- Indexes for table `detail_tunggakan`
--
ALTER TABLE `detail_tunggakan`
  ADD PRIMARY KEY (`id`),
  ADD KEY `detail_tunggakan_FK` (`id_detail_cicilan`),
  ADD KEY `detail_tunggakan_FK_1` (`id_jenis_tunggakan`);

--
-- Indexes for table `jenis_aksi`
--
ALTER TABLE `jenis_aksi`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `jenis_pinjaman`
--
ALTER TABLE `jenis_pinjaman`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `kolektebilitas`
--
ALTER TABLE `kolektebilitas`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `log_activity`
--
ALTER TABLE `log_activity`
  ADD PRIMARY KEY (`id`),
  ADD KEY `log_activity_FK` (`id_mantri`),
  ADD KEY `log_activity_FK_1` (`id_aksi_misi`);

--
-- Indexes for table `mantri`
--
ALTER TABLE `mantri`
  ADD PRIMARY KEY (`id`),
  ADD KEY `mantri_FK` (`id_unit`);

--
-- Indexes for table `misi`
--
ALTER TABLE `misi`
  ADD PRIMARY KEY (`id`),
  ADD KEY `misi_FK` (`id_mantri`);

--
-- Indexes for table `nasabah`
--
ALTER TABLE `nasabah`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `nomor_rekening` (`nomor_rekening`);

--
-- Indexes for table `pinjaman`
--
ALTER TABLE `pinjaman`
  ADD PRIMARY KEY (`id`),
  ADD KEY `pinjaman_FK` (`id_mantri`),
  ADD KEY `pinjaman_FK_1` (`id_jenis_pinjaman`),
  ADD KEY `pinjaman_FK_2` (`id_nasabah`);

--
-- Indexes for table `tunggakan`
--
ALTER TABLE `tunggakan`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `unit`
--
ALTER TABLE `unit`
  ADD PRIMARY KEY (`id`),
  ADD KEY `unit_FK` (`id_cabang`);

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
-- AUTO_INCREMENT for table `aksi_misi`
--
ALTER TABLE `aksi_misi`
  MODIFY `id` int(6) NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT for table `cabang`
--
ALTER TABLE `cabang`
  MODIFY `id` int(6) NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT for table `cicilan`
--
ALTER TABLE `cicilan`
  MODIFY `id` int(6) NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT for table `dashboard`
--
ALTER TABLE `dashboard`
  MODIFY `id` int(6) NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT for table `detail_cicilan`
--
ALTER TABLE `detail_cicilan`
  MODIFY `id` int(6) NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT for table `detail_tunggakan`
--
ALTER TABLE `detail_tunggakan`
  MODIFY `id` int(6) NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT for table `jenis_aksi`
--
ALTER TABLE `jenis_aksi`
  MODIFY `id` int(6) NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT for table `jenis_pinjaman`
--
ALTER TABLE `jenis_pinjaman`
  MODIFY `id` int(6) NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT for table `kolektebilitas`
--
ALTER TABLE `kolektebilitas`
  MODIFY `id` int(6) NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT for table `log_activity`
--
ALTER TABLE `log_activity`
  MODIFY `id` int(6) NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT for table `mantri`
--
ALTER TABLE `mantri`
  MODIFY `id` int(6) NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT for table `misi`
--
ALTER TABLE `misi`
  MODIFY `id` int(6) NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT for table `nasabah`
--
ALTER TABLE `nasabah`
  MODIFY `id` int(6) NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT for table `pinjaman`
--
ALTER TABLE `pinjaman`
  MODIFY `id` int(6) NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT for table `tunggakan`
--
ALTER TABLE `tunggakan`
  MODIFY `id` int(6) NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT for table `unit`
--
ALTER TABLE `unit`
  MODIFY `id` int(6) NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT for table `users`
--
ALTER TABLE `users`
  MODIFY `id` int(6) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=7;

--
-- Constraints for dumped tables
--

--
-- Constraints for table `aksi_misi`
--
ALTER TABLE `aksi_misi`
  ADD CONSTRAINT `aksi_misi_FK` FOREIGN KEY (`id_misi`) REFERENCES `misi` (`id`),
  ADD CONSTRAINT `aksi_misi_FK_1` FOREIGN KEY (`id_aksi`) REFERENCES `jenis_aksi` (`id`);

--
-- Constraints for table `cicilan`
--
ALTER TABLE `cicilan`
  ADD CONSTRAINT `cicilan_FK` FOREIGN KEY (`id_pinjaman`) REFERENCES `pinjaman` (`id`);

--
-- Constraints for table `detail_cicilan`
--
ALTER TABLE `detail_cicilan`
  ADD CONSTRAINT `detail_cicilan_FK` FOREIGN KEY (`id_cicilan`) REFERENCES `cicilan` (`id`),
  ADD CONSTRAINT `detail_cicilan_FK_1` FOREIGN KEY (`id_kolektebilitas`) REFERENCES `kolektebilitas` (`id`),
  ADD CONSTRAINT `detail_cicilan_FK_2` FOREIGN KEY (`id_aksi_misi`) REFERENCES `aksi_misi` (`id`);

--
-- Constraints for table `detail_tunggakan`
--
ALTER TABLE `detail_tunggakan`
  ADD CONSTRAINT `detail_tunggakan_FK` FOREIGN KEY (`id_detail_cicilan`) REFERENCES `detail_cicilan` (`id`),
  ADD CONSTRAINT `detail_tunggakan_FK_1` FOREIGN KEY (`id_jenis_tunggakan`) REFERENCES `tunggakan` (`id`);

--
-- Constraints for table `log_activity`
--
ALTER TABLE `log_activity`
  ADD CONSTRAINT `log_activity_FK` FOREIGN KEY (`id_mantri`) REFERENCES `mantri` (`id`),
  ADD CONSTRAINT `log_activity_FK_1` FOREIGN KEY (`id_aksi_misi`) REFERENCES `aksi_misi` (`id`);

--
-- Constraints for table `mantri`
--
ALTER TABLE `mantri`
  ADD CONSTRAINT `mantri_FK` FOREIGN KEY (`id_unit`) REFERENCES `unit` (`id`);

--
-- Constraints for table `misi`
--
ALTER TABLE `misi`
  ADD CONSTRAINT `misi_FK` FOREIGN KEY (`id_mantri`) REFERENCES `mantri` (`id`);

--
-- Constraints for table `pinjaman`
--
ALTER TABLE `pinjaman`
  ADD CONSTRAINT `pinjaman_FK` FOREIGN KEY (`id_mantri`) REFERENCES `mantri` (`id`),
  ADD CONSTRAINT `pinjaman_FK_1` FOREIGN KEY (`id_jenis_pinjaman`) REFERENCES `jenis_pinjaman` (`id`),
  ADD CONSTRAINT `pinjaman_FK_2` FOREIGN KEY (`id_nasabah`) REFERENCES `nasabah` (`id`);

--
-- Constraints for table `unit`
--
ALTER TABLE `unit`
  ADD CONSTRAINT `unit_FK` FOREIGN KEY (`id_cabang`) REFERENCES `cabang` (`id`);
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
