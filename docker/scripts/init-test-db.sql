DO
$do$
BEGIN
   IF NOT EXISTS (SELECT FROM pg_database WHERE datname = 'visit_test') THEN
      CREATE DATABASE visit_test;
END IF;
END
$do$;