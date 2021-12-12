import React, { useState } from 'react';
import { DsCreds } from './credentials';

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
        <label>username
        <input
          type='text'
          placeholder='username'
          onChange={e => {
            formdata.username = e.target.value;
            setFormData(formdata);
          }}
          />
          </label>
        <label>host
        <input
          type='text'
          placeholder='host'
          onChange={e => {
            formdata.host = e.target.value;
            setFormData(formdata);
          }}
        /></label>
        <label>password        <input
          type='text'
          placeholder='password'
          onChange={e => {
            formdata.password = e.target.value;
            setFormData(formdata);
          }}
        /></label>

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
