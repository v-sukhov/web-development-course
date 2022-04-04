--liquibase formatted sql

--changeset sukhov:1

create table T_USER
(
	user_id			SERIAL			PRIMARY KEY,
	login			VARCHAR(256)	NOT NULL,
    password_sha256	BYTEA,
	email			VARCHAR(256),
    cf_handle       VARCHAR(256),
	is_admin		SMALLINT        NOT NULL DEFAULT 0,
    is_teacher      SMALLINT        NOT NULL DEFAULT 0,
	surname			VARCHAR(256),
	first_name		VARCHAR(256),
	second_name		VARCHAR(256),
	display_name	VARCHAR(256)
);

COMMENT ON TABLE T_USER IS 'Все пользователи системы (включая регулярных пользователей и участников)';
COMMENT ON COLUMN T_USER.user_id IS 'Уникальный идентификатор';
COMMENT ON COLUMN T_USER.login IS 'Логин пользователя. Для регулярных пользователей должен совпадать с почтовым адресом (кроме специальных случаев)';
COMMENT ON COLUMN T_USER.password_sha256 IS 'Пароль, зашифрованный при помощи секретного salt-слова и хэш-алгоритма';
COMMENT ON COLUMN T_USER.email IS 'EMail пользователя. Для регулярных пользователей должен совпадать с логином (кроме специальных случаев)';
COMMENT ON COLUMN T_USER.cf_handle IS 'Хэндл пользователя на codeforce';
COMMENT ON COLUMN T_USER.is_admin IS '1 - админ, 0 - не админ';
COMMENT ON COLUMN T_USER.is_teacher IS '1 - учитель, 0 - не учитель';
COMMENT ON COLUMN T_USER.surname IS '';
COMMENT ON COLUMN T_USER.first_name IS '';
COMMENT ON COLUMN T_USER.second_name IS '';
COMMENT ON COLUMN T_USER.display_name IS '';

--rollback drop table T_USER;

--changeset sukhov:2

CREATE UNIQUE INDEX T_USER_LOGIN_INDEX ON T_USER(LOGIN);

--rollback drop index T_USER_LOGIN_INDEX;
