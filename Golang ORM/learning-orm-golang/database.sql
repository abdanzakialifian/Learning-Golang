CREATE TABLE sample(
    id VARCHAR(100) PRIMARY KEY NOT NULL,
    name VARCHAR(100) NOT NULL
);

SELECT * FROM sample;

CREATE TABLE user(
    id VARCHAR(100) PRIMARY KEY NOT NULL,
    password VARCHAR(100) NOT NULL,
    name VARCHAR(100) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

ALTER TABLE user rename COLUMN name to first_name;

ALTER TABLE user add COLUMN middle_name VARCHAR(100) NULL AFTER first_name;

ALTER TABLE user add COLUMN last_name VARCHAR(100) NULL AFTER middle_name;

SELECT * FROM user;

DELETE FROM `user` where id = "50";

CREATE TABLE user_logs(
    id INT PRIMARY KEY AUTO_INCREMENT,
    user_id VARCHAR(100) NOT NULL,
    action VARCHAR(100) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

SELECT * FROM user_logs;

DELETE FROM user_logs;

ALTER TABLE user_logs MODIFY created_at BIGINT NOT NULL;

ALTER TABLE user_logs MODIFY updated_at BIGINT NOT NULL;

DESCRIBE user_logs;

CREATE TABLE todo(
    id BIGINT PRIMARY KEY NOT NULL AUTO_INCREMENT,
    user_id VARCHAR(100) NOT NULL,
    title VARCHAR(100) NOT NULL,
    description TEXT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL
);

SELECT * FROM todo;

DELETE FROM todo WHERE id = 2;

CREATE TABLE wallet(
    id VARCHAR(100) PRIMARY KEY NOT NULL,
    user_id VARCHAR(100) NOT NULL,
    balance BIGINT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES user (id)
);

DESCRIBE wallet;

SHOW CREATE TABLE wallet;

SELECT * FROM wallet;

ALTER TABLE wallet DROP FOREIGN KEY wallet_ibfk_1;

ALTER TABLE wallet ADD CONSTRAINT fk_wallet_user_id FOREIGN KEY (user_id) REFERENCES user (id) ON DELETE CASCADE;

CREATE TABLE address(
    id BIGINT PRIMARY KEY AUTO_INCREMENT NOT NULL,
    user_id VARCHAR(100) NOT NULL,
    address VARCHAR(100) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES user (id)
);

DESC address;

SHOW CREATE TABLE address;

SELECT * FROM address;

ALTER TABLE address DROP FOREIGN KEY address_ibfk_1;

ALTER TABLE address ADD CONSTRAINT fk_address_user_id FOREIGN KEY (user_id) REFERENCES user (id) ON DELETE CASCADE;

CREATE TABLE product(
    id VARCHAR(100) PRIMARY KEY NOT NULL,
    name VARCHAR(100) NOT NULL,
    price BIGINT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

DESC product;

SELECT * FROM product;

DELETE FROM product WHERE id = "P001";

CREATE TABLE user_like_product(
    user_id VARCHAR(100) NOT NULL,
    product_id VARCHAR(100) NOT NULL,
    PRIMARY KEY (user_id, product_id),
    FOREIGN KEY (user_id) REFERENCES user(id),
    FOREIGN KEY (product_id) REFERENCES product(id)
);

DESC user_like_product;

SELECT * FROM user_like_product;

SELECT COUNT(id) from wallet;


SELECT COUNT(id) from user;

SELECT `user`.`id`,`user`.`first_name`,`user`.`middle_name`,`user`.`last_name`,`user`.`password`,`user`.`created_at`,`user`.`updated_at`,`Wallet`.`id` AS `Wallet__id`,`Wallet`.`user_id` AS `Wallet__user_id`,`Wallet`.`balance` AS `Wallet__balance`,`Wallet`.`created_at` AS `Wallet__created_at`,`Wallet`.`updated_at` AS `Wallet__updated_at` FROM `user` LEFT JOIN `wallet` `Wallet` ON `user`.`id` = `Wallet`.`user_id`

SELECT `user`.`id`,`user`.`first_name`,`user`.`middle_name`,`user`.`last_name`,`user`.`password`,`user`.`created_at`,`user`.`updated_at` FROM `user` join wallet on wallet.user_id = user.id

SELECT * FROM guest_book;