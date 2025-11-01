-- Grant all privileges to answer user from any host
-- This ensures the user can connect from the Docker network
-- The user is already created by MARIADB_USER, we just need to grant permissions
GRANT ALL PRIVILEGES ON answer.* TO 'answer'@'%';
FLUSH PRIVILEGES;

