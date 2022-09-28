import React, { useState } from 'react';
import { handleQuery } from './handlers'

interface Props {
  creds: DatasourceCredentials;
  onSubmit: any;
}

export default function DatasourceList(props: Props) {
  const { creds, onSubmit } = props;
  const submitFormData = handleQuery;
  const [formdata, setFormData] = useState(creds);

  return (
    <div className='ds-list'>
      <form>
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
