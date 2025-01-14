 CREATE TABLE "role" (
	"id" INT PRIMARY KEY,
	"name" VARCHAR(20) NOT NULL UNIQUE
);

INSERT INTO "role"
VALUES
	(1, 'normal'),
	(2, 'admin'),
	(3, 'moderator');

CREATE TABLE "appuser" (
	"id" SERIAL PRIMARY KEY,
	"username" VARCHAR(80) UNIQUE,
	"name" VARCHAR(80) NOT NULL,
	"role" INT NOT NULL REFERENCES "role" ("id"),
	"created_on" TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);

CREATE TABLE "thread_type" ("id" INT PRIMARY KEY, "title" VARCHAR);

INSERT INTO
	"thread_type"
VALUES
	(1, 'Question'), 
	(2, 'Post'), 
	(3, 'Reply');

CREATE TABLE "thread" (
	"id" SERIAL PRIMARY KEY,
	"title" VARCHAR,
	"content" VARCHAR,
	"type" INT NOT NULL REFERENCES "thread_type" ("id"),
	"created_by" INT REFERENCES "appuser"("id"),
	"created_on" TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
	"edited_on" TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);

CREATE FUNCTION update_edited_time() RETURNS TRIGGER AS $update_edited_time$
    BEGIN 
        UPDATE TG_TABLE_NAME
        SET "edited_on" = now()
        WHERE OLD.id=NEW.id;
    END;
$update_edited_time$ LANGUAGE plpgsql;

CREATE TRIGGER updating_edited_on_thread
AFTER
UPDATE OF "title",
"content" ON "thread" FOR EACH ROW
EXECUTE FUNCTION update_edited_time();

CREATE TABLE "thread_jt" (
	"parent_id" INT NOT NULL REFERENCES "thread" ("id"),
	"child_id" INT NOT NULL REFERENCES "thread" ("id")
);

CREATE TABLE "category" (
	"id" SERIAL PRIMARY KEY,
	"title" VARCHAR NOT NULL,
	"description" VARCHAR
);

CREATE TABLE "thread_category" (
	"thread_id" INT NOT NULL REFERENCES "thread" ("id"),
	"category_id" INT NOT NULL REFERENCES "category" ("id")
);

