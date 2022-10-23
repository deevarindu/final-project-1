CREATE TABLE IF NOT EXISTS todos (
  id VARCHAR(5) NOT NULL PRIMARY KEY,
  title VARCHAR(255) NOT NULL,
  status VARCHAR(255)
)

INSERT INTO todos (id,title,status) VALUES 
('1', 'do final project 1', 'completed'),
('2', 'do final project 2', 'completed'),
('3', 'do final project 3', 'in progress');