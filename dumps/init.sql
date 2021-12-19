BEGIN;

CREATE TABLE `users` (
    `id` INT(10) NOT NULL AUTO_INCREMENT PRIMARY KEY,
    `username` VARCHAR(16) NOT NULL,
    `password` VARCHAR(256) NOT NULL,
    UNIQUE KEY (`username`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE `services` (
    `id` INT(10) NOT NULL AUTO_INCREMENT PRIMARY KEY,
    `ident` VARCHAR(32) NOT NULL,
    `type` ENUM('ticketing', 'vcs') NOT NULL,
    UNIQUE KEY (`ident`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE `users_services` (
    `id` INT(10) NOT NULL AUTO_INCREMENT PRIMARY KEY,
    `id_user` INT(10) NOT NULL,
    `id_service` INT(10) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE `credentials` (
    `id` INT(10) NOT NULL AUTO_INCREMENT PRIMARY KEY,
    `id_user_service` INT(10) NOT NULL,
    `key` VARCHAR(16) NOT NULL,
    `value` VARCHAR(256) NOT NULL,
    UNIQUE KEY (`id_user_service`, `key`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE `tickets` (
    `id` INT(10) NOT NULL AUTO_INCREMENT PRIMARY KEY,
    `id_user_service` INT(10) NOT NULL,
    `name` VARCHAR(16) NOT NULL,
    `type` VARCHAR(16) NOT NULL,
    `project` VARCHAR(64) NOT NULL,
    `caption` VARCHAR(256) NOT NULL,
    `status` VARCHAR(32) NOT NULL,
    `priority` VARCHAR(16) NOT NULL,
    `assignee` VARCHAR(256) NOT NULL,
    `creator` VARCHAR(256) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

INSERT INTO `services` (`ident`, `type`) VALUES
('jira', 'ticketing');

ALTER TABLE `users_services`
    ADD FOREIGN KEY (`id_user`) REFERENCES `users` (`id`) ON UPDATE CASCADE ON DELETE CASCADE,
    ADD FOREIGN KEY (`id_service`) REFERENCES `services` (`id`) ON UPDATE CASCADE ON DELETE CASCADE;

ALTER TABLE `credentials`
    ADD FOREIGN KEY (`id_user_service`) REFERENCES `users_services` (`id`) ON UPDATE CASCADE ON DELETE CASCADE;

ALTER TABLE `tickets`
    ADD FOREIGN KEY (`id_user_service`) REFERENCES `users_services` (`id`) ON UPDATE CASCADE ON DELETE CASCADE;

COMMIT;
