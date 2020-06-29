# Story (Epic) - MVP

Joe is a software is a software dev

He has a production db and staging db

he is given a report, along with a bugfix ticket
because a customer is seeing incorrect data in their interface

Joe fires up the tool
there is an interface with various forms
and loads the report data in as a data source
then loads his db credentials in a form and hits submit
and now has populated a list of data sources he can select from

He fills in a query in a textarea and hits the query button
and sees the data from his selected data source in a table

## He creates a new instance of the interface he just used

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
  - [ ] load the save

when I click the query button
- [ ]  fetch the appropriate data from api
- [ ] on success  it should:
  - [ ]  populate instance table
- [ ]  columns and rows

- query should be read only
```js

  let querystring = " ";
  let querypos = 0;
  let dbname = "new";
  let columns = [];
  let rowdata = [];
  // list of dbs, of type string, when sent to the api will dictate the db to read from
  let dbs = ["ae1", "New", "Test"];
  let previousqueries = [];



  const saveCursor = (event: any) => {
    event.preventDefault();
    let querypos = event.target.selectionStart;
    let se = event.target.selectionEnd;
    return "";
  };

const insertIntoQuery = (newVal: any, event: any) => {
  let astr = querystring.substr(0, querypos);
  let cstr = querystring.substr(querypos);
  querystring = astr + " " + newVal + cstr;
};
  const replaceQueryString = (newVal: any, event: Event) => {
    let querystring = newVal;
    return querystring;
  };
```

