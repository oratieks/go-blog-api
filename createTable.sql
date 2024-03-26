create table if not exists articles (
	article_id integer unsigned auto_increment primary key, -- unsignedで非負を指定
	title varchar(100) not null,
	contents text not null,
	username varchar(100) not null,
	nice integer not null,
	created_at datetime
);

create table if not exists comments (
	comment_id integer unsigned auto_increment primary key,
	article_id integer unsigned not null,
	message text not null,
	created_at datetime,
	foreign key (article_id) references articles(article_id) -- 外部キー制約 
);

-- FOREIGN KEY (外部キーカラム名) REFERENCES 親テーブル名 (主キーカラム名)
-- 外部キーとは親テーブルの主キーを参照している子テーブルのカラムのこと
-- 主キーカラム名とは子テーブルに参照される親テーブルのカラム（主キー）のこと
-- 外部キー制約とは子テーブルのカラムが親テーブルのカラム（主キー）に制約されているということ