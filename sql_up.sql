create table if not exists Visitors (
	id serial primary key,
	name varchar(30) not null,
	email varchar(50) not null,
	password varchar(280) not null
);

create table if not exists PhysicalInfo (
	id serial primary key,
	visitor_id int references Visitors(id) on delete cascade not null unique,
	weight numeric(5,2) null,
	height numeric(5,2) null
);

create table if not exists Clubs (
	id serial primary key,
	address varchar(50) null,
	club_name varchar(50) not null
);

create table if not exists Coaches (
	id serial primary key,
	name varchar(30) not null,
	email varchar(50) not null,
	password varchar(280) not null
);

create table if not exists Activities (
	id serial primary key,
	activity_name varchar(30) not null unique,
	activity_desc text null,
	club_id int references Clubs(id) on delete set null
);

create table if not exists StatesTypes (
	id serial primary key,
	type_name varchar(30) not null,
	type_desc text null,
	type_unit varchar(10) not null
); 

create table if not exists Trainings (
	id serial primary key,
	start_datetime TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
	end_datetime TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
	coach_id int references Coaches(id) on delete set null,
	club_id int not null references Clubs(id) on delete cascade
);

create table if not exists Attendance (
	id serial primary key,
	visitor_id int not null references Visitors(id) on delete cascade,
	training_id int not null references Trainings(id) on delete cascade
);

create table if not exists ActivityUsage (
	id serial primary key,
	visitor_id int not null references Visitors(id) on delete cascade,
	activity_name varchar(30) not null references Activities(activity_name) on delete cascade,
	usage_start_datetime TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
	usage_end_datetime TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
	training_id int null references Trainings(id) on delete set null
);

create table if not exists PhysicalStates (
	id serial primary key,
	state_type_id int not null references StatesTypes(id) on delete cascade,
	unit_amount numeric(5,2) not null,
	at_datetime TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
	activity_usage_id int null references ActivityUsage(id) on delete cascade,
	duration_secs numeric(5,2) null
);

create table if not exists ActivityStates (
	id serial primary key,
	state_type_id int not null references StatesTypes(id) on delete cascade,
	unit_amount numeric(5,2) not null,
	at_datetime TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
	activity_id int not null references Activities(id) on delete cascade,
	duration_secs numeric(5,2) null
);


drop table if exists ActivityStates;
drop table if exists PhysicalStates;
drop table if exists ActivityUsage;
drop table if exists Attendance;
drop table if exists Trainings;
drop table if exists StatesTypes;
drop table if exists Activities;
drop table if exists Clubs;
drop table if exists Coaches;
drop table if exists PhysicalInfo;
drop table if exists Visitors;





use SportClubDB;

select ast.* from activitystates ast
	join activityusage aus on ast.activity_name = aus.activity_name AND
		ast.at_datetime between aus.usage_start_datetime and aus.usage_end_datetime
	where aus.visitor_id = 3 AND aus.activity_name = 'veloerg#12345679';




insert into activityusage(visitor_id, activity_name,usage_start_datetime,usage_end_datetime)
values
(3, 'veloerg#12345679', '2023-08-07 00:30:00', '2023-08-07 00:34:00');


select * from activitystates;
select * from visitors;
select * from activityusage;


--delete from activitystates;

select * from activities;









