Applications for automatically creating a schema for tables and columns in a Postgres database (ERD-diagram)

There are many different applications for automatically creating ERD diagrams,
however, they all display link arrows only to the table, and not to the desired column,
and do not have the ability to edit and update the edited diagram -
so I had to create my own application.

The application can automatically find in database and draw:
1. All tables
2. All table columns and types
3. Arrows of table relationships from column to column (foreign key)
4. Finds tables in the old existing .graphml file,
and places the tables at the same X,Y coordinates

The resulting .graphml file in the free yED editor can be:
1. Export as a .jpg drawing
2. Edit
3. Automatic placement of blocks

A sample implementation (drawings) can be found in the examples directory

Installation procedure:
1. Install the .graphml file editor yEd (free)
https://www.yworks.com/products/yed/download

2. Compile this repository
>make build
>
the image_database file will appear in the bin folder

3. fill parameters in "settings.txt" (or ".env") file:
```
FILENAME_GRAPHML=
INCLUDE_TABLES=
EXCLUDE_TABLES=

DB_HOST=
DB_NAME=
DB_SCHEME=
DB_PORT=
DB_USER=
DB_PASSWORD=
```
Run file:
>image_database

4. Open the resulting .graphml file in the yEd editor
(all elements will first be in the center of the screen)
and select from the menu:
Tools - Remove Node Overlaps
- the yEd editor will arrange all the elements of the diagram in the optimal form.

5. Export the diagram to an image.
Select from the menu:
File - Export


![database](https://github.com/ManyakRus/image_database/assets/30662875/6fbef84f-dfce-4e01-8e6b-fd535509bf5e)


Source code in Golang language.
Tested on Linux Ubuntu
Readme from 09/14/2023

Made by Alexander Nikitin
https://github.com/ManyakRus/image_database
