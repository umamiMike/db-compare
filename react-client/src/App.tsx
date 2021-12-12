import React, { useState } from 'react';

import './css/reset.css';
 /* import './App.css'; */
import DatasourceCredentialsForm from './DatasourceCredentialsForm';
import TextArea from './TextArea';

export default function App() {
  const api = 'http://localhost:9099';
  const [query, setQuery] = useState('');

  const initCreds = {
    type: '',
    username: '',
    host: '',
    password: '',
    dbName: '',
  };

  const handleSubmit = (formdata: any) => {
    fetch(api + '/datasources', {
      method: 'POST',
      body: JSON.stringify({
        data: {
          type: 'datasource',
          attributes: { ...formdata },
        },
      }),
    })
      .then(function (response) {
        if (!response.ok) {
          throw Error(response.statusText);
        }
        return response;
      })
      .then(function (response) {
        return response.json();
      })
      .then(data => {
        return data;
      })
      .catch(error => {
        console.log(error);
      });
  };

  const handleQuery = () => {
    console.log(query);

    fetch(api + '/queries', {
      method: 'POST',
      body: JSON.stringify({
        data: {
          type: 'query',
          attributes: {
            query_string: query,
          },
        },
      }),
    })
      .then(function (response) {
        if (!response.ok) {
          throw Error(response.statusText);
        }
        return response;
      })
      .then(function (response) {
        return response.json();
      })
      .then(data => {
        console.log(data);
        return data;
      })
      .catch(error => {
        console.log(error);
      });
  };

  return (
    <div className='main'>
      <DatasourceCredentialsForm creds={initCreds} onSubmit={handleSubmit} />
      <form>
      <TextArea name='query' onChange={e => setQuery(e.target.value)} />
      <button
        onClick={() => {
          handleQuery();
        }}
      >
        {' '}
        query the db
      </button>
      </form>
    </div>
  );
}
