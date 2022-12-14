create database desafio_final;

CREATE TABLE `dentistas` (
  `id` INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
  `nome` varchar(100) DEFAULT NULL,
  `sobrenome` varchar(200) DEFAULT NULL,
  `matricula` int not null
);
CREATE TABLE `pacientes` (
  `id` INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
  `nome` varchar(100) DEFAULT NULL,
  `sobrenome` varchar(200) DEFAULT NULL,
  `rg` int not null,
  `data_cadastro` date DEFAULT NULL
);
CREATE TABLE `consultas` (
  `id` INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
  `paciente_id` int NOT NULL,
  `dentista_id` int NOT NULL,
  `data` date DEFAULT NULL,
  `hora` time default null,
  `descricao` varchar(200) DEFAULT NULL,
  KEY `fk_paciente` (`paciente_id`),
  KEY `fk_dentista` (`dentista_id`),
CONSTRAINT `fk_paciente` FOREIGN KEY (`paciente_id`) REFERENCES `pacientes` (`id`) ON UPDATE CASCADE,
  CONSTRAINT `fk_dentista` FOREIGN KEY (`dentista_id`) REFERENCES `dentistas` (`id`) ON UPDATE CASCADE
);
INSERT INTO dentistas(nome,sobrenome, matricula) VALUES ("Maria","Helena",234563);

INSERT INTO dentistas(nome,sobrenome, matricula) VALUES ("Maria","Carla",13131);

INSERT INTO pacientes(nome,sobrenome,rg, data_cadastro) VALUES ("Carlos","Santos",986545678, "2000-01-23");
INSERT INTO pacientes(nome,sobrenome,rg, data_cadastro) VALUES ("Antônio","Alberto",098778, "2000-02-4");
INSERT INTO pacientes(nome,sobrenome,rg, data_cadastro) VALUES ("João","Marco",78866, "2001-06-20");
INSERT INTO pacientes(nome,sobrenome,rg, data_cadastro) VALUES ("Alberto","Pereira",909877, "2002-09-10");
INSERT INTO pacientes(nome,sobrenome,rg, data_cadastro) VALUES ("Marcio","Alencar",6788, "2005-02-09");

INSERT INTO consultas(paciente_id, dentista_id, data, hora, descricao) VALUES (2,2,"2005-02-09","08:33:20","não chegar atrasado");

INSERT INTO consultas(data, hora, descricao) VALUES ("2005-02-09","08:33:20","não chegar atrasado");

select * from dentistas;

select * from pacientes;

select * from consultas;

SELECT c.id, c.descricao, c.data, c.hora, c.paciente_id, c.dentista_id,p.id, p.nome, p.sobrenome, p.rg, p.data_cadastro,d.id, d.nome, d.sobrenome, d.matricula FROM consultas c INNER JOIN pacientes p on c.paciente_id = p.id INNER JOIN dentistas d on c.dentista_id = d.id WHERE p.rg = 909877 and d.matricula = 13131 ORDER BY p.data_cadastro;


START TRANSACTION;
  INSERT INTO dentistas(nome,sobrenome, matricula)
  VALUES("Alberto","Carla",13131);
  SELECT LAST_INSERT_ID() INTO @id;
  INSERT INTO pacientes(nome, sobrenome, rg, data_cadastro)
  VALUES("David", "Silva", 987652, "2000-01-23");
  SELECT LAST_INSERT_ID() INTO @id;
COMMIT;

select * from pacientes where rg = 909877;

SELECT c.id, 
c.descricao, 
c.data, 
c.hora, 
c.paciente_id, 
c.dentista_id,
p.id, 
p.nome, 
p.sobrenome, 
p.rg, 
p.data_cadastro,
d.id, 
d.nome, 
d.sobrenome, 
d.matricula FROM consultas c 
INNER JOIN pacientes p on c.paciente_id = p.id 
INNER JOIN dentistas d on c.dentista_id = d.id WHERE p.rg = 986545678 ORDER BY p.data_cadastro;

INSERT INTO consultas(data, hora,descricao) VALUES 
((SELECT id FROM pacientes WHERE rg = 98778), 
(SELECT id FROM dentistas WHERE matricula = 234563),
"2000-09-09", "09:09:09", "chegar atrasado");