CREATE KEYSPACE microgo
WITH REPLICATION = { 'class' : 'SimpleStrategy', 'replication_factor' : 1 };
  
SELECT keyspace_name FROM system_schema.keyspaces WHERE keyspace_name = 'microgo';

CREATE TABLE IF NOT EXISTS microgo.students (
  id TIMEUUID PRIMARY KEY,
  first_name TEXT,
  last_name TEXT,
  age INT,
  marks INT
);

SELECT table_name FROM system_schema.tables  WHERE keyspace_name = 'microgo';