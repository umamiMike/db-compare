import React, { useState } from 'react';
import { api } from './api-client'

import './css/reset.css';
/* import './App.css'; */
import DatasourceCredentialsForm from './DatasourceCredentialsForm';
import DatasourceList from './DatasourceList';
import TextArea from './TextArea';

export default function App() {
  const [query, setQuery] = useState('');

  const initCreds: DatasourceCredentials = {
    type: '',
    username: '',
    host: '',
    password: '',
    dbName: '',
  };
  const handleGetAll = () => {
    const formdata = query;
    fetch(api + '/datasources', {
      method: 'GET',
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

  const handleSubmit = () => {
    const formdata = query;
    console.log(query);
    fetch(api + '/datasources', {
      method: 'POST',
      body: JSON.stringify({
        data: {
          type: 'datasource',
          attributes: query,
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
      <DatasourceList creds={initCreds} onSubmit={handleGetAll} />
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
