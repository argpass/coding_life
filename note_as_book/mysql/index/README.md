# MySQL 索引笔记

## Notice
- If you index more than one column, the column order is very important, because MySQL can only search efficiently on a leftmost prefix of the index.
- ORM vs. index
    >ORMs produce logically and syntactically correct queries (most of the time), but they rarely produce index-friendly queries unless you use them for only the basic types of queries,such as primary key lookups.
### typical index types:(the below are not the all index types)
-  B-Tree 
-  Full text indexes
- Spatial indexes(空间索引，支持不够好不如PostGIS)
- Hash indexes
    - In MySQL, only the Memory storage engine supports explicit hash indexes. They are the default index type for Memory tables,thouth Memory tables can have B-Tree indexes too.
    - You can emulate hash index yourself.example
        >增加一个column作为hash column 使用适合的hash函数来计算hash，使用trigger来更新hash（当row被insert or update时）.hash 函数要返回整型.

        >mysql>SELECT id FROM url WHERE url="http://www.mysql.com”;

        >mysql> SELECT id FROM url WHERE url="http://www.mysql.com" AND url_crc=CRC32("http://www.mysql.com”);

## Indexing Strategies for High Performance
### Isolating the column
“Isolating”means it should not be part of an expression or be inside a function in the query.Here’s an example of a common mistake:

>mysql> SELECT ... WHERE TO_DAYS(CURRENT_DATE) - TO_DAYS(date_col) <= 10;

### Prefix Indexes and index selectivity
Sometimes you need to index very long character columns, which makes your index large and slow. One strategy is to emulate Hash type.But sometimes that isn’t good enough.You can often save space and get good performance by indexing the first few characters instead of the whole value.This makes your indexes use less space, but it also makes them less selective.A prefix of the column is often selective enough to give good performance.The trick is to choose a prefix that’s long enough to give good selectivity, but short enough to save space.
MySQL can not use prefix indexes for ORDER BY or GROUP BY queries, nor can it use them as covering indexes.
how to create a prefix index on the column: 
>mysql> ALTER TABLE sakila.city_demo ADD KEY (city(7));

### Multicolumn indexes
The index merge strategy sometimes works very well, but it’s more common for it to actually be an indication of a poorly indexed table:
- When the server intersects indexes(usually for AND conditions), it usually means that you need a single index
- The server unions indexes (usually for OR condition)
