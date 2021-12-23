INSERT INTO usuarios(nome, nick, email, senha)
values
("Usr1", "usr1", "usr1@gmail.com","$2a$10$fgszYGZELvXFfc5kiioGeu3WFlSjir4F.8076WD/ZDqZRKr9jtYGm"),
("Usr2", "usr2", "usr2@gmail.com","$2a$10$fgszYGZELvXFfc5kiioGeu3WFlSjir4F.8076WD/ZDqZRKr9jtYGm"),
("Usr3", "usr3", "usr3@gmail.com","$2a$10$fgszYGZELvXFfc5kiioGeu3WFlSjir4F.8076WD/ZDqZRKr9jtYGm"),
("Usr4", "usr4", "usr4@gmail.com","$2a$10$fgszYGZELvXFfc5kiioGeu3WFlSjir4F.8076WD/ZDqZRKr9jtYGm");

INSERT INTO seguidores(usuario_id, seguidor_id)
values
(1,3),
(2,3),
(1,4),
(4,1);
