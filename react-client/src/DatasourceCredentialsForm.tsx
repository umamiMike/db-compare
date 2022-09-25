import React, { useState } from 'react';
import { api } from './api-client'

type Props =  {
  creds: DatasourceCredentials;
  onSubmit: any;
}

const handleSubmit = (query: any) => {
  const formdata = query;

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

export default function DatasourceCredentialsForm(props: Props) {
  const { creds, onSubmit } = props;
  const submitFormData = onSubmit;
  const [formdata, setFormData] = useState(creds);
  console.log(formdata);

  return (
    <div className='ds-select'>
      <form>
        <label>
          db name
          <input
            type='text'
            placeholder='database name'
            onChange={e => {
              formdata.dbName = e.target.value;
              setFormData(formdata);
            }}
          />
        </label>
        <label>
          username
          <input
            type='text'
            placeholder='username'
            onChange={e => {
              formdata.username = e.target.value;
              setFormData(formdata);
            }}
          />
        </label>
        <label>
          host
          <input
            type='text'
            placeholder='host'
            onChange={e => {
              formdata.host = e.target.value;
              setFormData(formdata);
            }}
          />
        </label>
        <label>
          password{' '}
          <input
            type='text'
            placeholder='password'
            onChange={e => {
              formdata.password = e.target.value;
              setFormData(formdata);
            }}
          />
        </label>

        <button
          onClick={e => {
            e.preventDefault();
            submitFormData(formdata);
          }}
        >
          {' '}
          submit
        </button>
      </form>
    </div>
  );
}
