/*
    This SQL code is for part of database that your codes are ready to 
    be sent to CodeX.

    Also when the result is received you have to save its result as well here!
    
    Each row is build with a (task_id, email, query_string, status=uploaded, result=null)

    *** Run This piece of code once in mysql as root!
*/

DROP DATABASE IF EXISTS Builder_Executer;
CREATE DATABASE IF NOT EXISTS Builder_Executer;
USE Builder_Executer;

-- Just in case :)
DROP TABLE IF EXISTS Tasks;
DROP TABLE IF EXISTS Results;

-- STORE RESULTS HERE
CREATE TABLE IF NOT EXISTS Results
(
    ID INT PRIMARY KEY AUTO_INCREMENT,
    TimeStamp INT,
    Stat INT,
    Output VARCHAR(4194304),
    Error VARCHAR(1048576),
    Language VARCHAR(256),
    Info VARCHAR(524288)
);

-- TOTAL QS (INPUR + CODE + LANG) = 4MB + 50B 
CREATE TABLE IF NOT EXISTS Tasks
(
    TASK_ID VARCHAR(512) NOT NULL UNIQUE,
    EMAIL VARCHAR(512) NOT NULL UNIQUE,
    STAT VARCHAR(10) CHECK(STAT IN {'UPLOAD', 'PROCESSING', 'FAILED', 'DONE'}) NOT NULL,
    QS VARCHAR(4194354) NOT NULL UNIQUE,
    RES INT FOREIGN KEY REFERENCES Results(ID),
    PRIMARY KEY (TASK_ID)
);

-- This user is for authx admin!
CREATE USER 'builder'@'localhost' IDENTIFIED BY 'password';
CREATE USER 'executer'@'localhost' IDENTIFIED BY 'password';

-- job_builder authorities
GRANT ALL ON Builder_Executer.Tasks TO 'builder'@'localhost';

-- job_executer authorities
GRANT SELECT, UPDATE ON Builder_Executer.Tasks TO 'executer'@'localhost';
GRANT ALL ON Builder_Executer.Results TO 'executer'@'localhost';
