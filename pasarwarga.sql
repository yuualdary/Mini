-- phpMyAdmin SQL Dump
-- version 4.9.0.1
-- https://www.phpmyadmin.net/
--
-- Host: 127.0.0.1
-- Generation Time: Aug 04, 2021 at 11:32 PM
-- Server version: 10.4.6-MariaDB
-- PHP Version: 7.3.9

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
SET AUTOCOMMIT = 0;
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `pasarwarga`
--

-- --------------------------------------------------------

--
-- Table structure for table `articles`
--

CREATE TABLE `articles` (
  `id` bigint(20) UNSIGNED NOT NULL,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `category_id` bigint(20) UNSIGNED DEFAULT NULL,
  `content` longtext DEFAULT NULL,
  `title` longtext DEFAULT NULL,
  `slug` longtext DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

--
-- Dumping data for table `articles`
--

INSERT INTO `articles` (`id`, `created_at`, `updated_at`, `deleted_at`, `category_id`, `content`, `title`, `slug`) VALUES
(1, '2021-08-05 02:14:44.321', '2021-08-05 02:14:44.321', '2021-08-05 02:42:45.719', 5, 'sangat horor', 'Scream 2', 'scream-2-3'),
(2, '2021-08-05 02:40:29.582', '2021-08-05 02:40:29.582', '2021-08-05 02:42:45.719', 5, 'sangat horor bagian 2', 'Scream 3', 'scream-3-5'),
(3, '2021-08-05 02:41:32.609', '2021-08-05 02:41:32.609', NULL, 1, 'Mindblow', 'Durarara', 'durarara-1'),
(4, '2021-08-05 02:41:38.321', '2021-08-05 02:41:38.321', NULL, 1, 'Mindblow 2', 'Durarara!!', 'durarara-1'),
(5, '2021-08-05 03:35:05.560', '2021-08-05 03:35:05.560', NULL, 2, 'Cerita tentang detektif', 'The Detective Is Already Dead', 'the-detective-is-already-dead-2'),
(6, '2021-08-05 03:36:26.789', '2021-08-05 04:20:02.093', '2021-08-05 04:31:42.543', 8, 'Time Leaper', 'Tokyo Revenger', 'tokyo-revenger-8');

-- --------------------------------------------------------

--
-- Table structure for table `categories`
--

CREATE TABLE `categories` (
  `id` bigint(20) UNSIGNED NOT NULL,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `category_name` longtext DEFAULT NULL,
  `category_slug` longtext DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

--
-- Dumping data for table `categories`
--

INSERT INTO `categories` (`id`, `created_at`, `updated_at`, `deleted_at`, `category_name`, `category_slug`) VALUES
(1, '2021-08-04 20:15:58.657', '2021-08-04 22:31:04.141', NULL, 'Fiksi Dan Sains', 'fiksi-dan-sains'),
(2, '2021-08-04 20:16:54.503', '2021-08-04 22:38:40.018', NULL, 'Light Novel', 'light-novel'),
(3, '2021-08-04 22:38:56.951', '2021-08-04 22:38:56.951', '2021-08-05 02:07:45.793', 'Horor', 'horor'),
(4, '2021-08-05 02:10:40.367', '2021-08-05 02:10:40.367', '2021-08-05 02:11:02.809', 'Horor', 'horor'),
(5, '2021-08-05 02:15:15.764', '2021-08-05 02:15:15.764', '2021-08-05 02:42:45.726', 'Horor', 'horor'),
(6, '2021-08-05 03:35:37.606', '2021-08-05 03:37:54.772', NULL, 'Misteri', 'misteri'),
(7, '2021-08-05 03:36:57.180', '2021-08-05 03:36:57.180', '2021-08-05 03:37:29.839', 'Misteri', 'misteri'),
(8, '2021-08-05 03:55:31.387', '2021-08-05 03:55:31.387', NULL, 'Action', 'action');

--
-- Indexes for dumped tables
--

--
-- Indexes for table `articles`
--
ALTER TABLE `articles`
  ADD PRIMARY KEY (`id`),
  ADD KEY `idx_articles_deleted_at` (`deleted_at`),
  ADD KEY `fk_articles_categories` (`category_id`);

--
-- Indexes for table `categories`
--
ALTER TABLE `categories`
  ADD PRIMARY KEY (`id`),
  ADD KEY `idx_categories_deleted_at` (`deleted_at`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `articles`
--
ALTER TABLE `articles`
  MODIFY `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=7;

--
-- AUTO_INCREMENT for table `categories`
--
ALTER TABLE `categories`
  MODIFY `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=9;

--
-- Constraints for dumped tables
--

--
-- Constraints for table `articles`
--
ALTER TABLE `articles`
  ADD CONSTRAINT `fk_articles_categories` FOREIGN KEY (`category_id`) REFERENCES `categories` (`id`);
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
