create table user_info (
    name  varchar(64) not null check(length(name) > 5),
    email varchar(256) not null unique primary key check(length(email) > 5) check(email LIKE '%@%'),
    age integer not null check(age >= 18) 
);

create table bank_account (
    bank_uuid varchar(256) not null unique,
    money real default(0),
    user_email varchar(256) unique not null, 
    foreign key (user_email) references user_info(email) 
);