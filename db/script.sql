--CREATE USER TABLE
-- SQLINES LICENSE FOR EVALUATION USE ONLY
CREATE TABLE USERS (
    id NUMBER(19) PRIMARY KEY NOT NULL,
    firstName NVARCHAR2(50) DEFAULT NULL,
    middleName NVARCHAR2(50) DEFAULT NULL,
    lastName NVARCHAR2(50) DEFAULT NULL,
    username VARCHAR2(50) NOT NULL,
    mobile VARCHAR2(15),
    email VARCHAR2(50),
    passwordHash VARCHAR2(32) NOT NULL
);
-- Generate ID using sequence and trigger
CREATE SEQUENCE user_seq START WITH 1 INCREMENT BY 1;
CREATE OR REPLACE TRIGGER user_seq_tr BEFORE
INSERT ON USERS FOR EACH ROW
    WHEN (NEW.id IS NULL) BEGIN
SELECT user_seq.NEXTVAL INTO :NEW.id
FROM DUAL;
END;
/ --CREATE TASK TABLE
-- SQLINES LICENSE FOR EVALUATION USE ONLY
CREATE TABLE tasks (
    id NUMBER(19) PRIMARY KEY NOT NULL,
    userId NUMBER(19) NOT NULL,
    title NVARCHAR2(512) NOT NULL,
    description NVARCHAR2(512) DEFAULT NULL,
    status NUMBER(5) DEFAULT 0 NOT NULL,
    hours BINARY_DOUBLE DEFAULT 0 NOT NULL,
    createdAt TIMESTAMP(0) NULL DEFAULT CURRENT_TIMESTAMP,
    updatedAt TIMESTAMP(0) DEFAULT NULL,
    plannedStartDate TIMESTAMP(0) DEFAULT NULL,
    plannedEndDate TIMESTAMP(0) DEFAULT NULL,
    actualStartDate TIMESTAMP(0) DEFAULT NULL,
    actualEndDate TIMESTAMP(0) DEFAULT NULL,
    content CLOB DEFAULT NULL
);
-- Generate ID using sequence and trigger
CREATE SEQUENCE task_seq START WITH 1 INCREMENT BY 1;
CREATE OR REPLACE TRIGGER task_seq_tr BEFORE
INSERT ON tasks FOR EACH ROW
    WHEN (NEW.id IS NULL) BEGIN
SELECT task_seq.NEXTVAL INTO :NEW.id
FROM DUAL;
END;
/ --CREATE TABLE TASK META
-- SQLINES LICENSE FOR EVALUATION USE ONLY
CREATE TABLE task_metas (
    id NUMBER(19) PRIMARY KEY NOT NULL,
    taskId NUMBER(19) NOT NULL,
    key VARCHAR2(50) NOT NULL,
    content CLOB DEFAULT NULL
);
-- Generate ID using sequence and trigger
CREATE SEQUENCE task_meta_seq START WITH 1 INCREMENT BY 1;
CREATE OR REPLACE TRIGGER task_meta_seq_tr BEFORE
INSERT ON task_metas FOR EACH ROW
    WHEN (NEW.id IS NULL) BEGIN
SELECT task_meta_seq.NEXTVAL INTO :NEW.id
FROM DUAL;
END;
/ --CREATE TABLE TAG
-- SQLINES LICENSE FOR EVALUATION USE ONLY
CREATE TABLE tags (
    id NUMBER(19) PRIMARY KEY NOT NULL,
    title NVARCHAR2(75) NOT NULL,
    slug VARCHAR2(100) NOT NULL
);
-- Generate ID using sequence and trigger
CREATE SEQUENCE tag_seq START WITH 1 INCREMENT BY 1;
CREATE OR REPLACE TRIGGER tag_seq_tr BEFORE
INSERT ON tags FOR EACH ROW
    WHEN (NEW.id IS NULL) BEGIN
SELECT tag_seq.NEXTVAL INTO :NEW.id
FROM DUAL;
END;
/ -- CREATE TABLE TASK TAG
-- SQLINES LICENSE FOR EVALUATION USE ONLY
CREATE TABLE task_tags (
    taskId NUMBER(19) NOT NULL,
    tagId NUMBER(19) NOT NULL,
    PRIMARY KEY (taskId, tagId)
);
-- CREATE TABLE ACTIVITY
-- SQLINES LICENSE FOR EVALUATION USE ONLY
CREATE TABLE activities (
    id NUMBER(19) PRIMARY KEY NOT NULL,
    userId NUMBER(19) NOT NULL,
    taskId NUMBER(19) NOT NULL,
    title NVARCHAR2(512) NOT NULL,
    description NVARCHAR2(512) DEFAULT NULL,
    status NUMBER(5) DEFAULT 0 NOT NULL,
    hours BINARY_DOUBLE DEFAULT 0 NOT NULL,
    createdAt TIMESTAMP(0) NULL DEFAULT CURRENT_TIMESTAMP,
    updatedAt TIMESTAMP(0) DEFAULT NULL,
    plannedStartDate TIMESTAMP(0) DEFAULT NULL,
    plannedEndDate TIMESTAMP(0) DEFAULT NULL,
    actualStartDate TIMESTAMP(0) DEFAULT NULL,
    actualEndDate TIMESTAMP(0) DEFAULT NULL,
    content CLOB DEFAULT NULL
);
-- Generate ID using sequence and trigger
CREATE SEQUENCE activity_seq START WITH 1 INCREMENT BY 1;
CREATE OR REPLACE TRIGGER activity_seq_tr BEFORE
INSERT ON activities FOR EACH ROW
    WHEN (NEW.id IS NULL) BEGIN
SELECT activity_seq.NEXTVAL INTO :NEW.id
FROM DUAL;
END;
/ -- CREATE TABLE COMMENT
-- SQLINES LICENSE FOR EVALUATION USE ONLY
CREATE TABLE comments (
    id NUMBER(19) PRIMARY KEY NOT NULL,
    taskId NUMBER(19) DEFAULT NULL,
    activityId NUMBER(19) DEFAULT NULL,
    title NVARCHAR2(100) NOT NULL,
    createdAt TIMESTAMP(0) NOT NULL CURRENT_TIMESTAMP,
    updatedAt TIMESTAMP(0) DEFAULT NULL,
    content CLOB DEFAULT NULL
);
-- Generate ID using sequence and trigger
CREATE SEQUENCE comment_seq START WITH 1 INCREMENT BY 1;
CREATE OR REPLACE TRIGGER comment_seq_tr BEFORE
INSERT ON comments FOR EACH ROW
    WHEN (NEW.id IS NULL) BEGIN
SELECT comment_seq.NEXTVAL INTO :NEW.id
FROM DUAL;
END;
/ -- SQLINES LICENSE FOR EVALUATION USE ONLY
CREATE UNIQUE INDEX uq_username ON "USERS" (username);
-- SQLINES LICENSE FOR EVALUATION USE ONLY
CREATE UNIQUE INDEX uq_mobile ON "USERS" (mobile);
-- SQLINES LICENSE FOR EVALUATION USE ONLY
CREATE UNIQUE INDEX uq_email ON "USERS" (email);
-- SQLINES LICENSE FOR EVALUATION USE ONLY
CREATE INDEX idx_task_user ON tasks (userId);
-- SQLINES LICENSE FOR EVALUATION USE ONLY
CREATE INDEX idx_meta_task ON task_metas (taskId);
-- SQLINES LICENSE FOR EVALUATION USE ONLY
CREATE UNIQUE INDEX uq_task_meta ON task_metas (taskId, key);
-- SQLINES LICENSE FOR EVALUATION USE ONLY
CREATE INDEX idx_tt_task ON task_tags (taskId);
-- SQLINES LICENSE FOR EVALUATION USE ONLY
CREATE INDEX idx_tt_tag ON task_tags (tagId);
-- SQLINES LICENSE FOR EVALUATION USE ONLY
CREATE INDEX idx_activity_user ON activities (userId);
-- SQLINES LICENSE FOR EVALUATION USE ONLY
CREATE INDEX idx_comment_task ON comments (taskId);
ALTER TABLE tasks
ADD CONSTRAINT fk_task_user FOREIGN KEY (userId) REFERENCES "USERS" (id);
ALTER TABLE task_metas
ADD CONSTRAINT fk_meta_task FOREIGN KEY (taskId) REFERENCES tasks (id);
ALTER TABLE task_tags
ADD CONSTRAINT fk_tt_task FOREIGN KEY (taskId) REFERENCES tasks (id);
ALTER TABLE task_tags
ADD CONSTRAINT fk_tt_tag FOREIGN KEY (tagId) REFERENCES tags (id);
ALTER TABLE activities
ADD CONSTRAINT fk_activity_user FOREIGN KEY (userId) REFERENCES "USERS" (id);
ALTER TABLE activities
ADD FOREIGN KEY (taskId) REFERENCES tasks (id);
ALTER TABLE comments
ADD CONSTRAINT fk_comment_task FOREIGN KEY (taskId) REFERENCES tasks (id);
ALTER TABLE comments
ADD FOREIGN KEY (activityId) REFERENCES activities (id);
-- PROCEDURE ADD TASK
create or replace NONEDITIONABLE PROCEDURE proc_addtask (
        userid IN NUMBER,
        title IN CLOB,
        description IN CLOB,
        hours IN NUMBER,
        plannedstartdate IN TIMESTAMP,
        plannedenddate IN TIMESTAMP,
        content IN CLOB,
        inserted_id OUT NUMBER
    ) IS BEGIN
INSERT INTO tasks (
        userid,
        title,
        description,
        hours,
        plannedstartdate,
        plannedenddate,
        content
    )
VALUES (
        userid,
        title,
        description,
        hours,
        TO_TIMESTAMP(plannedstartdate),
        TO_TIMESTAMP(plannedenddate),
        content
    )
RETURNING id INTO inserted_id;
END proc_addtask;
-- PROCEDURE ADD USER
create or replace NONEDITIONABLE PROCEDURE proc_adduser (
        firstname IN NVARCHAR2,
        middlename IN NVARCHAR2,
        lastname IN NVARCHAR2,
        username IN VARCHAR2,
        mobile IN VARCHAR2,
        email IN VARCHAR2,
        passwordhash IN VARCHAR2
    ) AS BEGIN
INSERT INTO users (
        firstname,
        middlename,
        lastname,
        username,
        mobile,
        email,
        passwordhash
    )
VALUES (
        firstname,
        middlename,
        lastname,
        username,
        mobile,
        email,
        passwordhash
    );
END proc_adduser;
-- PROCEDURE ADD ACTIVITY
create or replace NONEDITIONABLE PROCEDURE proc_addactivity (
        userId IN NUMBER,
        taskId IN NUMBER,
        title IN CLOB,
        description IN CLOB,
        hours IN NUMBER,
        plannedStartDate IN TIMESTAMP,
        plannedEndDate IN TIMESTAMP,
        content IN CLOB
    ) AS BEGIN
INSERT INTO activities (
        userId,
        taskId,
        title,
        description,
        hours,
        plannedStartDate,
        plannedEndDate,
        content
    )
VALUES (
        userId,
        taskId,
        title,
        description,
        hours,
        plannedStartDate,
        plannedEndDate,
        content
    );
END proc_addactivity;
-- PROCEDURE ADD COMMENT
create or replace NONEDITIONABLE PROCEDURE proc_addcomment (
        taskId IN NUMBER,
        activityId IN NUMBER,
        title IN NVARCHAR2,
        content IN CLOB,
        inserted_id OUT NUMBER
    ) IS BEGIN
INSERT INTO comments (
        taskId,
        activityId,
        title,
        content
    )
VALUES (
        taskId,
        activityId,
        title,
        content
    )
RETURNING id INTO inserted_id;
END proc_addcomment;
-- PROCEDURE ADD TAG
create or replace NONEDITIONABLE PROCEDURE proc_addtag (
        title IN VARCHAR2,
        slug IN VARCHAR2,
        inserted_id OUT NUMBER
    ) IS BEGIN
INSERT INTO tags (title, slug)
VALUES (title, slug)
RETURNING id INTO inserted_id;
END proc_addtag;
-- PROCEDURE ADD TASK TO TAG
create or replace NONEDITIONABLE procedure proc_addtasktotag (taskid in number, tagid in number) as begin
INSERT INTO task_tags(taskid, tagid)
VALUES (taskid, tagid);
end proc_addtasktotag;
-- PROCEDURE UPDATE STATUS OF ACITIVITY
CREATE OR REPLACE PROCEDURE proc_updatestatusactivity (
        idActivity IN NUMBER,
        newStatus IN NUMBER
    ) AS BEGIN
UPDATE activities
SET status = newStatus
WHERE id = idActivity;
END proc_updatestatusactivity;