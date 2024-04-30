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

-- キー...一意に識別するためのカラム
-- 主キー...テーブル内の各レコードを一意に識別するためのカラム
-- 外部キー...外部テーブル内の各レコードを一意に識別するためのカラム
-- 　　(外部テーブルの各レコードを一意に識別できていればいいので外部キー自身は一意でなくてok)
-- 　　(主に外部テーブルの主キーが参照される)
-- 外部キー制約...外部キーにかけられる制約(ルール)のこと
-- 　　(外部テーブルの主キーを参照するためには外部テーブルのレコードが存在している必要がある)
-- 　　(そのため外部キーでは外部テーブルに存在しない主キーの値を追加することができない)
-- 　　(参照先の主キーが更新された場合に外部キーも更新したり削除したりすることがオプションを指定することでできる)
-- 　　(デフォルトではRESTRICTというオプションで主キーを削除しようとすると外部キーで使用されているので削除を中止するオプションが適用されている)