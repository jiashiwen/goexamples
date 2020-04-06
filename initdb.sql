--name:create-tasks-table
create table tasks(
    taskid text primary key not null,
    status int,
    tasktype text
);
