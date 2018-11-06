CREATE OR REPLACE FUNCTION create_if_not_exists ()
RETURNS void AS
$_$
BEGIN
  IF EXISTS (
    SELECT 1
      FROM   pg_catalog.pg_tables
      WHERE  tablename  = 'roles'
      ) THEN
      RAISE NOTICE 'Table roles already exists.';
  ELSE
    -- 角色表
    CREATE TABLE roles (
        id          INT NOT NULL UNIQUE,
        role_name   VARCHAR(16) NOT NULL
    );
    COMMENT ON COLUMN roles.id IS 'Role ID';
    COMMENT ON COLUMN roles.role_name IS 'Role name';
  END IF;

  IF EXISTS (
    SELECT 1
      FROM   pg_catalog.pg_tables
      WHERE  tablename  = 'users'
      ) THEN
      RAISE NOTICE 'Table users already exists.';
  ELSE
    -- 用户表
    CREATE TABLE users (
        id          SERIAL PRIMARY KEY,
        email       VARCHAR(64) NOT NULL UNIQUE,
        nickname    VARCHAR(16) NOT NULL,
        passwd      VARCHAR(32) NOT NULL,
        img         VARCHAR(64),
        me          TEXT,
        user_role   INT NOT NULL,
        create_time TIMESTAMP NOT NULL
    );
    COMMENT ON COLUMN users.id IS 'User ID';
    COMMENT ON COLUMN users.email IS 'User email as login name';
    COMMENT ON COLUMN users.nickname IS 'Nick name';
    COMMENT ON COLUMN users.passwd IS 'User password';
    COMMENT ON COLUMN users.img IS 'User head portrait image';
    COMMENT ON COLUMN users.me IS 'Its me';
    COMMENT ON COLUMN users.user_role IS 'User role';
    COMMENT ON COLUMN users.create_time IS 'User create time';
  END IF;

  IF EXISTS (
    SELECT 1
      FROM   pg_catalog.pg_tables
      WHERE  tablename  = 'jobs'
      ) THEN
      RAISE NOTICE 'Table jobs already exists.';
  ELSE
    -- 工作履历表
    CREATE TABLE jobs (
        id          SERIAL PRIMARY KEY,
        user_id     INT NOT NULL references users(id),
        company     VARCHAR(32) NOT NULL,
        department  VARCHAR(32) NOT NULL,
        post        VARCHAR(32) NOT NULL,
        job_level   VARCHAR(16) NOT NULL,
        job_desc    TEXT NOT NULL,
        start_time  TIMESTAMP NOT NULL,
        end_time    TIMESTAMP,
        is_current  BOOLEAN,
        create_time TIMESTAMP NOT NULL
    );
    COMMENT ON COLUMN jobs.id IS 'Job ID';
    COMMENT ON COLUMN jobs.user_id IS 'User ID';
    COMMENT ON COLUMN jobs.company IS 'Company name';
    COMMENT ON COLUMN jobs.department IS 'Department in company';
    COMMENT ON COLUMN jobs.post IS 'Job post';
    COMMENT ON COLUMN jobs.job_level IS 'Job level';
    COMMENT ON COLUMN jobs.job_desc IS 'Description of job';
    COMMENT ON COLUMN jobs.start_time IS 'Start time';
    COMMENT ON COLUMN jobs.end_time IS 'End time';
    COMMENT ON COLUMN jobs.is_current IS 'Is current job';
    COMMENT ON COLUMN jobs.create_time IS 'Job info create time';
  END IF;

  IF EXISTS (
    SELECT 1
      FROM   pg_catalog.pg_tables
      WHERE  tablename  = 'education'
      ) THEN
      RAISE NOTICE 'Table education already exists.';
  ELSE
    -- 教育表
    CREATE TABLE education (
        id          SERIAL PRIMARY KEY,
        user_id     INT NOT NULL references users(id),
        school      VARCHAR(32) NOT NULL,
        major       VARCHAR(32) NOT NULL,
        degree      VARCHAR(8) NOT NULL,
        start_time  TIMESTAMP NOT NULL,
        end_time    TIMESTAMP,
        is_current  BOOLEAN,
        create_time TIMESTAMP NOT NULL
    );
    COMMENT ON COLUMN education.id IS 'Education ID';
    COMMENT ON COLUMN education.user_id IS 'User ID';
    COMMENT ON COLUMN education.school IS 'School name';
    COMMENT ON COLUMN education.major IS 'Major';
    COMMENT ON COLUMN education.degree IS 'Degree';
    COMMENT ON COLUMN education.start_time IS 'Start time';
    COMMENT ON COLUMN education.end_time IS 'End time';
    COMMENT ON COLUMN education.is_current IS 'Is current school';
    COMMENT ON COLUMN education.create_time IS 'Job info create time';
  END IF;

  IF EXISTS (
    SELECT 1
      FROM   pg_catalog.pg_tables
      WHERE  tablename  = 'article'
      ) THEN
      RAISE NOTICE 'Table article already exists.';
  ELSE
    -- 文章表
    CREATE TABLE article (
        id          SERIAL PRIMARY KEY,
        user_id     INT NOT NULL references users(id),
        nickname    VARCHAR(16) NOT NULL,
        title       VARCHAR(64) NOT NULL,
        markdown    TEXT NOT NULL,
        html        TEXT NOT NULL,
        create_time TIMESTAMP NOT NULL
    );
    COMMENT ON COLUMN article.id IS 'Article ID';
    COMMENT ON COLUMN article.user_id IS 'User ID';
    COMMENT ON COLUMN article.nickname IS 'Nickname';
    COMMENT ON COLUMN article.title IS 'Article title';
    COMMENT ON COLUMN article.markdown IS 'Article markdown';
    COMMENT ON COLUMN article.html IS 'Article html';
    COMMENT ON COLUMN article.create_time IS 'Article create time';
  END IF;
END;
$_$ LANGUAGE plpgsql;

SELECT create_if_not_exists();
