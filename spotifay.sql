-- phpMyAdmin SQL Dump
-- version 5.0.1
-- https://www.phpmyadmin.net/
--
-- Host: localhost
-- Generation Time: Jan 27, 2020 at 11:34 PM
-- Server version: 10.3.20-MariaDB-1
-- PHP Version: 7.3.12-1

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
SET AUTOCOMMIT = 0;
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `spotifay`
--

-- --------------------------------------------------------

--
-- Table structure for table `artist`
--

CREATE TABLE `artist` (
  `id` int(11) NOT NULL,
  `name` varchar(30) NOT NULL,
  `debut` date NOT NULL,
  `genre_id` int(11) NOT NULL,
  `image_url` varchar(255) NOT NULL,
  `category` varchar(10) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `artist`
--

INSERT INTO `artist` (`id`, `name`, `debut`, `genre_id`, `image_url`, `category`) VALUES
(1, 'Armada', '2014-03-08', 3, 'http://3.bp.blogspot.com/_gvYQg7_TOU0/S3msYfqSL_I/AAAAAAAAByg/q5_-s1kkrvw/s320/ARMADA.jpg', 'Group'),
(2, 'Judika', '2012-04-28', 3, 'http://1.bp.blogspot.com/-ETZUK8yptE0/VhMlXlXspFI/AAAAAAAAAR0/Q3Pt2B3JnU0/s1600/jud.jpg', 'Solo'),
(3, 'tipe-x', '2008-09-16', 5, 'https://i.ytimg.com/vi/wRZzI6bxAb0/hqdefault.jpg', 'Group'),
(23, 'Dhyo Haw', '2013-06-04', 1, 'https://4.bp.blogspot.com/-OhQ1EQKzh-U/VI6Xb0Q4i-I/AAAAAAAAAFA/Qis05gd1qLk/s1600/dhyo-haw-music-lyric-1-0-s-307x512.jpg', 'Solo');

-- --------------------------------------------------------

--
-- Table structure for table `genre`
--

CREATE TABLE `genre` (
  `id` int(11) NOT NULL,
  `name` varchar(10) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `genre`
--

INSERT INTO `genre` (`id`, `name`) VALUES
(1, 'Reggae'),
(2, 'Jazz'),
(3, 'Pop'),
(4, 'Rock'),
(5, 'Ska');

-- --------------------------------------------------------

--
-- Table structure for table `song`
--

CREATE TABLE `song` (
  `id` int(11) NOT NULL,
  `artist_id` int(11) NOT NULL,
  `song_name` varchar(25) NOT NULL,
  `imageUrl` varchar(255) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `song`
--

INSERT INTO `song` (`id`, `artist_id`, `song_name`, `imageUrl`) VALUES
(1, 3, 'selamat jalan', 'https://i.ytimg.com/vi/YwvJJBEHUGc/maxresdefault.jpg'),
(2, 3, 'mawar hitam', 'https://i.ytimg.com/vi/M-ENv1kSkWM/maxresdefault.jpg'),
(3, 3, 'salam rindu', 'http://i.ytimg.com/vi/mYk-ElV6vnA/maxresdefault.jpg'),
(4, 23, 'kecewa', 'https://i.ytimg.com/vi/qNLlwAco44U/maxresdefault.jpg'),
(5, 23, 'tetap tersenyum kawan', 'https://i.ytimg.com/vi/a1o3k-lRHfo/maxresdefault.jpg');

--
-- Indexes for dumped tables
--

--
-- Indexes for table `artist`
--
ALTER TABLE `artist`
  ADD PRIMARY KEY (`id`),
  ADD KEY `genre_id` (`genre_id`);

--
-- Indexes for table `genre`
--
ALTER TABLE `genre`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `song`
--
ALTER TABLE `song`
  ADD PRIMARY KEY (`id`),
  ADD KEY `artist_id` (`artist_id`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `artist`
--
ALTER TABLE `artist`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=25;

--
-- AUTO_INCREMENT for table `genre`
--
ALTER TABLE `genre`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=6;

--
-- AUTO_INCREMENT for table `song`
--
ALTER TABLE `song`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=24;

--
-- Constraints for dumped tables
--

--
-- Constraints for table `artist`
--
ALTER TABLE `artist`
  ADD CONSTRAINT `artist_ibfk_1` FOREIGN KEY (`genre_id`) REFERENCES `genre` (`id`);

--
-- Constraints for table `song`
--
ALTER TABLE `song`
  ADD CONSTRAINT `song_ibfk_1` FOREIGN KEY (`artist_id`) REFERENCES `artist` (`id`) ON DELETE CASCADE;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
