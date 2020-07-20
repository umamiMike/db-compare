# (Epic) - MVP

## Say Hi to Joe

Joe is a software is a software dev - valid credentials (Happy Path)

He has a production db and staging db

he is given a report, along with a bugfix ticket
because a customer is seeing incorrect data in their application interface.

## User Story #1 - Joe can query the [datasource](datasource)


Joe enters the url and is greeted with the page. 
there is an interface containing various forms.
He types in the [credentials](api-design.md#POST /datasource) needed for connecting to a postgres-db he has
identifier
name

access to and adds and hits add

he loads the report data in as another data source

then loads his db credentials in a form and hits submit
and now has populated a list of data sources he can select from

He fills in a query in a textarea and hits the query button
and sees the data from his selected data source in a table

<!-- The valuable final product -->

### Acceptance Criteria

#### Notes on Acceptence Criteria

The Acceptance Criteria is **what** things need doing to make the story successful
dont create criteria not mentioned in the story


- [api design](api-design.md#POST /datasource) for creating new datasource

- various forms
- there is a form on the page allowing you to upload csv data
    - returns TODO:
- there is a form for username, hostname, password, database name
-  [POST datasource-credentials](#POST /datasource-credentials)
    - returns id - there is a blank data source selection menu - there is a textarea with a submit button, which is disabled 



- when you hit the add button, the data is sent to the server via rest POST request
- the server responds with successfully added data source message in json 
- the data source shows up in the data source selection menu

- the system allows you to type in a VALID SQL QUERY to the textarea, 
- hitting submit will send a http POST (which breaks REST convention endpoints) request with a json object containing the selected data source and the query string

<!-- - What is the length of a max of a query param??? -->
- 
  <!-- NOTE: elasticsearch uses get WITH a req body (which is weird) -->

- after successful response from api, a table appears containing the rows and
columns from the sql query



## Comparing csv report data to sql query - (Happy Path)


## He creates a new instance of the interface he just used

Joe wants to compare the data on the staging server with the csv report

which has the same data sources populated for him to select
he can select on the staging db source
and make the same query to the staging db
which shows another table, containing the table data of the query for this other data source.


<!-- The valuable final product -->

Joe is able to visually compare the report with the staging db source and see the disparity between them.  He talks to his client and shows them the issue.



## Joe loads the report csv file
- loads 

you want storys having a valuable end

## joe uploads a jpeg to query

Cant query it, or kills the server attempting

## Joe queries the report



when I load the app
  - [ ] the page has initial data

when I click the query button
- [ ]  fetch the appropriate data from api
- [ ] on success  it should:
  - [ ]  populate instance table
- [ ]  columns and rows

## resources

### datasource

#### examples of a datasource

a db
a csv file
a text document


## db-connection

the data needed to establish a connection to a db you wish to connect with via
the application


### Notes about user stories in general

this is the **why** of the user story
be as specific as possible when describing the story context. If they can be
actual cases all the better
do as little in the story as possible still containing value


# tmp

I am keeping some old code in 
[scratch](scratch.tsx)
