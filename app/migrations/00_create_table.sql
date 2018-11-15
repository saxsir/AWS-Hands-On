CREATE TABLE `eventlog` (
  `at` datetime NOT NULL DEFAULT NOW() COMMENT 'イベント発生時刻',
  `name`  varchar(255) NOT NULL COMMENT 'イベント名',
  `value` varchar(255) NOT NULL COMMENT 'イベントの値'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='イベントログ';
