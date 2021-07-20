BEGIN;

CREATE TABLE `users` (
    `id` BIGINT AUTO_INCREMENT NOT NULL PRIMARY KEY,
    `username` VARCHAR(64) NOT NULL,
    `password` VARCHAR(256) NOT NULL,
    UNIQUE KEY (`username`)
);

CREATE TABLE `services` (
    `ident` VARCHAR(32) NOT NULL PRIMARY KEY,
    `type` VARCHAR(32) NOT NULL
);

CREATE TABLE `service_types` (
    `ident` VARCHAR(32) NOT NULL PRIMARY KEY
);

CREATE TABLE `settings` (
    `id` BIGINT AUTO_INCREMENT NOT NULL PRIMARY KEY,
    `id_user` BIGINT NOT NULL,
    `service` VARCHAR(32) NOT NULL
);

CREATE TABLE `credentials` (
    `id` BIGINT AUTO_INCREMENT NOT NULL PRIMARY KEY,
    `id_setting` BIGINT NOT NULL,
    `key` VARCHAR(64) NOT NULL,
    `value` TEXT NOT NULL
);

CREATE TABLE `tickets` (
    `id` BIGINT AUTO_INCREMENT NOT NULL PRIMARY KEY,
    `id_user` BIGINT NOT NULL,
    `name` VARCHAR(16) NOT NULL,
    `type` VARCHAR(16) NOT NULL,
    `project` VARCHAR(64) NOT NULL,
    `caption` VARCHAR(256) NOT NULL,
    `status` VARCHAR(32) NOT NULL,
    `priority` VARCHAR(16) NOT NULL,
    `assignee` VARCHAR(256) NOT NULL,
    `creator` VARCHAR(256) NOT NULL
);

INSERT INTO `service_types` (`ident`) VALUES
('Ticketing');

INSERT INTO `services` (`ident`, `type`) VALUES
('Jira', 'Ticketing');

ALTER TABLE `settings`
    ADD FOREIGN KEY (`id_user`) REFERENCES `users`(`id`)
        ON UPDATE CASCADE
        ON DELETE CASCADE,
    ADD FOREIGN KEY (`service`) REFERENCES `services`(`ident`)
        ON UPDATE CASCADE
        ON DELETE CASCADE;

ALTER TABLE `credentials`
    ADD FOREIGN KEY (`id_setting`) REFERENCES `settings`(`id`)
        ON UPDATE CASCADE
        ON DELETE CASCADE;

ALTER TABLE `tickets`
    ADD FOREIGN KEY (`id_user`) REFERENCES `users`(`id`)
        ON UPDATE CASCADE
        ON DELETE CASCADE;

ALTER TABLE `services`
    ADD FOREIGN KEY (`type`) REFERENCES `service_types`(`ident`)
        ON UPDATE CASCADE
        ON DELETE RESTRICT;

COMMIT;
