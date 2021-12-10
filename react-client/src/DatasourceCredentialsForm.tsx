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
    <div className="ds-select">
      <form>
        <label>database</label>
        <input
          type='text'
          placeholder='database name'
          onChange={e => {
            formdata.dbName = e.target.value;
            setFormData(formdata);
          }}
        />
        <label>username</label>
        <input
          type='text'
          placeholder='username'
          onChange={e => {
            formdata.username = e.target.value;
            setFormData(formdata);
          }}
        />
        <label>host</label>
        <input
          type='text'
          placeholder='host'
          onChange={e => {
            formdata.host = e.target.value;
            setFormData(formdata);
          }}
        />
        <label>password</label>
        <input
          type='text'
          placeholder='password'
          onChange={e => {
            formdata.password = e.target.value;
            setFormData(formdata);
          }}
        />
      <button
        onClick={() => {
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
