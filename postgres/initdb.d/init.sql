DROP DATABASE IF EXISTS ctf_db;
CREATE DATABASE ctf_db;
\c ctf_db;
DROP TABLE IF EXISTS users;

CREATE TABLE users (
    id SERIAL NOT NULL PRIMARY KEY,
    first_name VARCHAR(30) NOT NULL,
    last_name VARCHAR(30) NOT NULL,
    job VARCHAR(30) NOT NULL,
    delete_flag BOOLEAN NOT NULL DEFAULT FALSE
);

INSERT INTO users (first_name, last_name, job) 
    VALUES 
    ('太郎', '山田', 'サーバーサイドエンジニア'),
    ('次郎', '鈴木', 'フロントエンドエンジニア'),
    ('三郎', '田中', 'インフラエンジニア'),
    ('花子', '佐藤', 'デザイナー');
INSERT INTO users (first_name, last_name, job, delete_flag)
    VALUES ('一郎', '渡辺', 'myctf{scf_sql_injection_flag}', TRUE); /* フラグ1つ目(同じテーブル) */

DROP TABLE IF EXISTS flag;

CREATE TABLE flag (
    id SERIAL NOT NULL PRIMARY KEY,
    flag VARCHAR(60) NOT NULL,
    create_date TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    update_date TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);

INSERT INTO flag (flag)
    VALUES ('myctf{next_flag_[/var/ctf/flag.md]}'); /* フラグ2つ目(別のテーブル) */