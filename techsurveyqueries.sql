select * from techsurvey.points;
select * from techsurvey.faveprogramlang;
delete from techsurvey.users;
delete from techsurvey.faveprogramlang;
alter table techsurvey.faveprogramlang modify email varchar(255) unique;
delete from techsurvey.feedback;
alter table techsurvey.faveprogramlang modify number int;
create table techsurvey.updates(
    message varchar(255),
    time datetime
    );
    select * from techsurvey.updates;
    select * from techsurvey.updates order by time asc;
    delete from techsurvey.updates;
    
    
SELECT COUNT(DISTINCT table_name) AS tables_with_email
FROM (
    SELECT 'faveprogramlang' AS table_name FROM techsurvey.faveprogramlang WHERE email = 'parasprajapat8th@gmail.com'
    UNION ALL
    SELECT 'faveframe' AS table_name FROM techsurvey.faveframe WHERE email = 'parasprajapat8th@gmail.com'
    UNION ALL
    SELECT 'favdev' AS table_name FROM techsurvey.favdev WHERE email = 'parasprajapat8th@gmail.com'
) AS result;
select * from techsurvey.faveprogramlang;

create table techsurvey.faveframe(
     email varchar(255),
     vote varchar(255),
     date datetime,
     number int 
     );
     create table techsurvey.favdev(
     email varchar(255),
     vote varchar(255),
     date datetime,
     number int 
     );
     select * from techsurvey.faveframe;
     select * from techsurvey.favdev;
     delete from techsurvey.favdev;
     delete from techsurvey.faveframe;
     alter table techsurvey.users add column points int default 0;
     drop table techsurvey.points;
     select points from techsurvey.users where email = "parasprajapat8th@gmail.com";