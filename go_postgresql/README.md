0. You must create a postgreSQL server first
1. create a table named "mst_person_info" using scripts below
    create table mst_person_info (
    person_cd character varying(10) not null
    , person_name character varying(100)
    , person_name_kana character varying(100)
    , birthday date
    , sex character(1)
    , zip character varying(11)
    , address1 character varying(100)
    , address2 character varying(100)
    , address3 character varying(100)
    , address4 character varying(100)
    , tel character varying(30)
    , mobile character varying(30)
    , mail_address character varying(255)
    , authentication_date date
    , delete_flag character(1) default '0'
    , create_user character varying(20)
    , create_date timestamp(6) without time zone
    , update_user character varying(20)
    , update_date timestamp(6) without time zone
    , primary key (person_cd)
    );

    And insert some data like query below
    [Insert into "mst_person_info" values ('4','孫　三郎','ｿﾝ ｻﾌﾞﾛｳ','2015-06-01','1','2060021','東京都','多摩市','XXXX１－２５－35','','0123456789','08012345678','12335@google.com','0001-01-01','0','tester',to_timestamp('2020-04-16 17:54:29.751','null'),'tester',to_timestamp('2020-04-16 17:54:29.751','null'));]
    [Insert into "mst_person_info" values ('2','孫　二郎','ｿﾝ ジロウ','2001-06-21','1','2060002','東京都','多摩市','一ノXX宮１－２４－１','メゾンドノア１２３室','0123456789','08012345678','123@google.com','0001-01-01','0','tester',to_timestamp('2020-04-16 17:31:29.371','null'),'tester',to_timestamp('2020-04-16 17:31:29.371','null'));]
    [Insert into "mst_person_info" values ('1','孫　一郎','ｿﾝ イチロウ','2001-06-21','1','2060021','東京都','多摩市','XXXX１－２５－２６','','','','','0001-01-01','0','tester',to_timestamp('2020-04-14 21:23:33.125','null'),'tester',to_timestamp('2020-04-16 17:57:48.736','null'));]

2. Open workspace folder with vs Code, and input [go run main.go]

3. Add some others action to finish CRUD

4. In PostgreSQL create a sequence for getting a new personcd using query below
   [create sequence SEQ_PERSON_CD INCREMENT BY 1 maxvalue 99999999 start with 1 cycle]

   So wo can see the column[person_cd] in table MST_PERSON_INFO is varchar(10), and the maxlength of sequence is 8, 
   because i want to add YY at front of the value like [to_char(CURRENT_DATE,'YY') || LPAD(nextval('SEQ_PERSON_CD')::varchar, 8, '0')]

5. So we create some method
         getPerson() to get person data by person code, 
         createNewPerson() create a new person data and print the new personcd to console, 
         updatePerson() update person data, 
         deletePerson() delete person data

Done.
