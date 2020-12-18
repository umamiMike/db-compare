import React, { useState } from 'react';
import { DsCreds } from './credentials';
import './App.css';

interface Props {
  creds: DsCreds;
  onSubmit: any;
}

export default function DatasourceCredentialsForm(props: Props) {
  const { creds, onSubmit } = props;
  const submitFormData = onSubmit;
  const [formdata, setFormData] = useState(creds);

  return (
    <div className='db-creds'>
      <h1>database credentials</h1>
      <form>
        <input
          type='text'
          placeholder='database name'
          onChange={e => {
            formdata.dbName = e.target.value;
            setFormData(formdata);
          }}
        />
        <br />
        <input
          type='text'
          placeholder='username'
          onChange={e => {
            formdata.username = e.target.value;
            setFormData(formdata);
          }}
        />
        <br />
        <input
          type='text'
          placeholder='host'
          onChange={e => {
            formdata.host = e.target.value;
            setFormData(formdata);
          }}
        />
        <input
          type='text'
          placeholder='password'
          onChange={e => {
            formdata.password = e.target.value;
            setFormData(formdata);
          }}
        />
      </form>
      <button
        onClick={() => {
          submitFormData(formdata);
        }}
      >
        {' '}
        submit
      </button>
    </div>
  );
}
