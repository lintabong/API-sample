import sqlite3

conn = sqlite3.connect('mydatabase.db')

conn.execute('''CREATE TABLE IF NOT EXISTS USER
             (id        INTEGER PRIMARY KEY AUTOINCREMENT,
              username  VARCHAR(75)     NOT NULL,
              password  VARCHAR(75)     NOT NULL,
              level     VARCHAR(75)     NOT NULL);''')

conn.execute('''CREATE TABLE IF NOT EXISTS POST
             (id        INTEGER PRIMARY KEY AUTOINCREMENT,
              author    VARCHAR(75)     NOT NULL,
              title     TEXT            NOT NULL,
              createdAt DATETIME,
              content   TEXT);''')

conn.execute('''CREATE TABLE IF NOT EXISTS COMMENT
             (id        INTEGER PRIMARY KEY AUTOINCREMENT,
              postId    INTEGER         NOT NULL,
              author    VARCHAR(75)     NOT NULL,
              title     VARCHAR(75)     NOT NULL,
              createdAt DATETIME,
              content   TEXT);''')

conn.execute('''CREATE TABLE IF NOT EXISTS POST_COMMENT
            (post_id        INTEGER         NOT NULL,
             comment_id     INTEGER         NOT NULL);''')

conn.execute('''CREATE TABLE IF NOT EXISTS CATEGORY
             (id        INTEGER PRIMARY KEY AUTOINCREMENT,
              title     TEXT            NOT NULL);''')

conn.execute('''CREATE TABLE IF NOT EXISTS POST_CATEGORY
            (post_id        INTEGER         NOT NULL,
             category_id    INTEGER         NOT NULL);''')

conn.commit()
conn.close()
