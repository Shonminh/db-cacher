CREATE DATABASE IF NOT EXISTS `travis_test_db`;

CREATE TABLE IF NOT EXISTS `travis_test_db`.`test_tab`
(
    id        bigint(21) unsigned         NOT NULL AUTO_INCREMENT,
    user_name varchar(64)      default '' NOT NULL,
    sex       tinyint(4)       default 0  NOT NULL,
    address   varchar(64)      default '' NOT NULL,
    email     varchar(64)      default '' NOT NULL,
    ctime     int(11) unsigned DEFAULT 0  NOT NULL,
    mtime     int(11) unsigned DEFAULT 0  NOT NULL,
    PRIMARY KEY (id)

) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_unicode_ci;

